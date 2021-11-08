package http

import (
	"backend-gateway/conf"
	"backend-gateway/model"
	"backend-gateway/utils"
	"bytes"
	"encoding/json"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	logger "github.com/cihub/seelog"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	WRITE_WAIT = 10 * time.Second

	// NOTE: ws服务端设置的超时时间是60秒, 这里最好不要超过或者接近这个值, 会有些问题 const defaultTimeout = 60 * time.Second
	// WRITE_WAIT 和 PONG_WAIT 最好设置一样的值, 不然的话会有些超时断开的问题?
	// Time allowed to read the next pong message from the peer.
	PONG_WAIT = 10 * time.Second

	// Send pings to peer with this period. Must be less than PONG_WAIT.
	PING_PERIOD = (PONG_WAIT * 9) / 10

	// Maximum message size allowed from peer.
	MAX_MESSAGE_SIZE = 102400

	// Maximum message id.
	MAX_MESSAGE_ID = 99999999

	// Send queue size
	SEND_QUEUE_SIZE = 2000

	// Reconnect count
	RECONNECT_COUNT = 100

	// Scheme
	TRANSPORT_SCHEME_WEBSOCKET      = "ws"
	TRANSPORT_SCHEME_WEBSOCKET_PATH = "/websocket"
)

type (
	reqClients struct {
		*sync.Map
	}

	BlockbookWsRequest struct {
		ID     string      `json:"id"`
		Method string      `json:"method"`
		Params interface{} `json:"params"`
	}

	Info struct {
		Name       string `json:"name"`
		Shortcut   string `json:"shortcut"`
		Decimals   int    `json:"decimals"`
		Version    string `json:"version"`
		BestHeight int    `json:"bestHeight"`
		BestHash   string `json:"bestHash"`
		Block0Hash string `json:"block0Hash"`
		Testnet    bool   `json:"testnet"`
	}

	// 用于 sync.map, 与请求的 goroutines 进行交互
	ChanData struct {
		Data chan interface{}
	}

	BlockbookWsResponse struct {
		ID   string      `json:"id"`
		Data interface{} `json:"data"`
	}

	wsClient struct {
		*websocket.Conn // websocket客户端连接
		*time.Ticker    // 心跳保持定时器
	}

	// 消息集合器
	WsMsgHub struct {
		Conf *conf.Config
		url.URL
		isRunning    bool             // 服务运行状态
		platformName string           // 主链名称
		wsCtx        wsClient         // websocket客户端上下文
		clients      reqClients       // 注册的请求客户端
		sendQueue    chan interface{} // 待发送消息缓冲队列
		id           uint64           // 消息请求id
		exitChan     chan struct{}    // 发送消息模块退出控制
	}
)

// @Description 初始化 websocket 消息集合器
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func NewWsMsgHub(host, platformName string, conf *conf.Config) *WsMsgHub {
	u := url.URL{Scheme: TRANSPORT_SCHEME_WEBSOCKET, Host: host, Path: TRANSPORT_SCHEME_WEBSOCKET_PATH}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.Error("Blockbook websocket dial error: ", err)
		time.Sleep(100 * time.Millisecond)
		os.Exit(1)
	}

	return &WsMsgHub{
		Conf:         conf,
		URL:          u,
		platformName: platformName,
		wsCtx: wsClient{
			Conn:   ws,
			Ticker: time.NewTicker(PING_PERIOD),
		},
		clients:   reqClients{Map: new(sync.Map)},
		sendQueue: make(chan interface{}, SEND_QUEUE_SIZE),
		id:        0,
		exitChan:  make(chan struct{}),
	}
}

// @Description 初始化 websocket 消息集合器收发消息模块
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) Init() {
	go h.writePump()
	go h.readPump()

	// 确保writePump和readPump都已执行成功
	time.Sleep(10 * time.Millisecond)
	h.isRunning = true
	return
}

// @Description websocket 消息发送模块
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) writePump() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, model.PANIC_STACK_SIZE)
			buf = buf[:runtime.Stack(buf, false)]

			logger.Errorf("Internal panic occurred: %s, stack: %s", r, string(buf))
		}
	}()

	for {
		select {
		// 通过ws发送消息
		case msg := <-h.sendQueue:
			err := h.wsCtx.WriteJSON(msg)
			if err != nil {
				logger.Errorf("Send message: %+v to blockbook ws error: %+v, message: %+v", msg, err)
				// 清空阻塞队列
				h.cleanRespData(msg.(BlockbookWsRequest).ID)
			}
		case <-h.wsCtx.C:
			// 保持心跳
			h.wsCtx.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
			if err := h.wsCtx.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Error("Send ping message to blockbook ws error: ", err)
				logger.Debugf("[%s]平台心跳包发送异常, 尝试重连", h.platformName)
				h.reconnectWs()
			}
			// logger.Debugf("[%s]平台心跳包发送正常", h.platformName)
		case <-h.exitChan:
			// Cleanly shutdown: 向ws服务端发送关闭消息后再关闭连接, 保证链接正常友好关闭
			err := h.wsCtx.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Error("Write close message to blockbook ws error: ", err)
				return
			}
			h.wsCtx.Close()

			logger.Infof("Gracefully exiting blockbook %s websocket write pump module ...", h.platformName)
			return
		}
	}
}

// @Description websocket 消息接收模块
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) readPump() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, model.PANIC_STACK_SIZE)
			buf = buf[:runtime.Stack(buf, false)]

			logger.Errorf("Internal panic occurred: %s, stack: %s", r, string(buf))
		}
	}()

	// 设置读参数
	h.wsCtx.SetReadLimit(MAX_MESSAGE_SIZE)
	h.wsCtx.SetReadDeadline(time.Now().Add(PONG_WAIT))
	h.wsCtx.SetPongHandler(func(string) error { h.wsCtx.SetReadDeadline(time.Now().Add(PONG_WAIT)); return nil })

	for {
		resp := BlockbookWsResponse{}
		err := h.wsCtx.ReadJSON(&resp)
		// NOTE:
		// Applications must break out of the application's read loop when this method
		// returns a non-nil error value. Errors returned from this method are
		// permanent. Once this method returns a non-nil error, all subsequent calls to
		// this method return the same error.
		//
		if err != nil {
			// 读取消息超出阈值
			if strings.Contains(err.Error(), websocket.ErrReadLimit.Error()) {
				logger.Warn("websocket: read limit exceeded!")
				continue
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("The blockbook ws connection had going out without CloseGoingAway or CloseAbnormalClosure, error: ", err)
			}

			if h.isRunning {
				logger.Warnf("The blockbook ws connection had going out, error: %+v, readPump exiting ...", err)
			} else {
				logger.Infof("Gracefully exiting blockbook %s websocket read pump module ...", h.platformName)
			}

			// FIXME: 连接中断后读模块会第一时间感应
			// blockbook服务停止时, 会报如下错误
			// websocket: close 1006 (abnormal closure): unexpected EOF
			h.isRunning = false
			h.wsCtx.Close()
			// 多测试几种情况, 决定是否需要在此关闭连接, 关闭后写模块发送ping会失败触发重连机制

			return
		}
		// logger.Debug("Read data from blockbook ws: ", resp)

		// 将应答消息发送至已注册的clients
		v, ok := h.clients.Load(resp.ID)
		if !ok {
			logger.Error("Get data from client sync.Map error, message id: ", resp.ID)
		}
		v.(ChanData).Data <- resp.Data
	}
}

// @Description websocket JSON消息（未序列化结构体或map）发送请求
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) WriteJSON(req BlockbookWsRequest) error {
	if h.isRunning {
		// 注册消息请求ID对应的客户端应答结构
		h.clients.registerMsgId(req.ID)

		// 发送消息至缓冲队列
		h.sendQueue <- req
	} else {
		err := logger.Error("WsMsgHub had already exit, message id: ", req.ID)
		return err
	}

	return nil
}

// @Description websocket JSON消息（未序列化结构体或map）读取请求
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) GetRespData(id string, result interface{}) error {
	respData, ok := h.clients.Load(id)
	if !ok {
		err := logger.Error("Receive data from client sync.Map error, message id: ", id)
		return err
	}
	unknownTypeData := <-respData.(ChanData).Data

	var (
		data []byte
		err  error
	)

	switch unknownTypeData.(type) {
	case map[string]interface{}:
		data, err = json.Marshal(unknownTypeData)
		if err != nil {
			logger.Errorf("Platform unmarshal error: %+v, response body: %+v", err, unknownTypeData)
			h.clients.unregisterMsgId(id)
			return utils.ErrTypePlatformUnmarshal
		}
	case string:
		data = []byte(unknownTypeData.(string))
	case []byte:
		data = unknownTypeData.([]byte)
	case []interface{}:
		// 重新解析列表"[...map[string]interface{}]"结构, 再拼装结构体
		list := unknownTypeData.([]interface{})
		var buffer bytes.Buffer
		buffer.Write([]byte("["))
		for index, v := range list {
			vByte, err := json.Marshal(v)
			if err != nil {
				logger.Errorf("Platform marshal error: %+v, response body: %+v", err, v)
				h.clients.unregisterMsgId(id)
				return utils.ErrTypePlatformUnmarshal
			}
			buffer.Write(vByte)
			if index+1 == len(list) {
				break
			}
			buffer.Write([]byte(","))
		}
		buffer.Write([]byte("]"))
		data = buffer.Bytes()
	case nil: // 因网络问题而清空的消息队列
		err = logger.Error("Nil message data")
		return err
	default:
		err = logger.Errorf("Unknown data type: %+v", unknownTypeData)
		h.clients.unregisterMsgId(id)
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		logger.Errorf("Platform unmarshal error: %+v, response body: %+v", err, data)
		h.clients.unregisterMsgId(id)
		return utils.ErrTypePlatformUnmarshal
	}

	// logger.Debugf("Blockbook ws response data json unmarshal result: %+v\n ", result)

	// 从clients中删除该请求的注册信息
	h.clients.unregisterMsgId(id)

	return nil
}

// @Description 清空阻塞等待请求队列
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) cleanRespData(id string) {
	respData, ok := h.clients.Load(id)
	if !ok {
		logger.Error("Receive data from client sync.Map error, message id: ", id)
	}
	respData.(ChanData).Data <- nil

	return
}

// @Description 获取 websocket 请求消息ID
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) GetMsgId() uint64 {
	// 原子操作增加
	id := atomic.AddUint64(&h.id, 1)

	// 到达最大值之后重新从0开始
	if id == MAX_MESSAGE_ID {
		// 重置id值
		atomic.StoreUint64(&h.id, 0)
		id = 0
	}

	// 这个函数的作用是让当前 goroutine 让出CPU，好让其它的 goroutine 获得执行的机会。同时，当前的 goroutine 也会在未来的某个时间点继续运行
	runtime.Gosched()

	return id
}

// @Description 根据请求消息ID 注册请求客户端（应答channel）
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (c *reqClients) registerMsgId(id string) {
	data := ChanData{
		Data: make(chan interface{}),
	}
	c.Store(id, data)

	return
}

// @Description 根据请求消息ID 注销请求客户端（应答channel）
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (c *reqClients) unregisterMsgId(id string) {
	c.Delete(id)

	return
}

// @Description 退出 websocket 消息集合器
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-19 init
func (h *WsMsgHub) Exit() {
	// 关闭心跳保持
	h.wsCtx.Stop()

	h.isRunning = false

	// 退出消息发送模块
	close(h.exitChan)

	logger.Infof("Exiting blockbook %s websocket connection ...", h.platformName)

	return
}

// @Description 重新建立 websocket 连接
// @Author Ryen
// @Version 1.0
// @Update Ryen 2020-12-22 init
func (h *WsMsgHub) reconnectWs() error {
	// FIXME:
	// 读模块先感应到连接已中断, 并主动关闭连接描述符
	// h.isRunning = false

	var (
		ws  *websocket.Conn
		err error
	)

	// TODO: 重试次数、机制及每次间隔时间
	for i := 1; i <= RECONNECT_COUNT; i++ {
		if i == RECONNECT_COUNT {
			err := logger.Errorf("Reconnect blockbook websocket reconnect error: %+v, retry count: %d", err, i)
			return err
		}

		ws, _, err = websocket.DefaultDialer.Dial(h.URL.String(), nil)
		if err != nil {
			logger.Error("Reconnect blockbook websocket error: ", err)
			time.Sleep(time.Second)
			continue
		}

		break
	}

	h.wsCtx.Conn = ws

	go h.readPump()

	// 确保 readPump 已执行成功
	time.Sleep(10 * time.Millisecond)
	h.isRunning = true

	logger.Infof("Reconnect blockbook %s websocket connection success ...", h.platformName)

	return nil
}
