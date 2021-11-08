package http

import (
	"strconv"
	"testing"
	"time"
)

// 测试使用
func TestBlockbookWs(t *testing.T) {
	// 收发消息集合器
	wsMsgHub := NewWsMsgHub("localhost:9130", "bitcoin", nil)

	// 启动收发器
	wsMsgHub.Init()

	for i := 0; i < 1000; i++ {
		go func() {
			req := BlockbookWsRequest{
				ID:     strconv.FormatUint(wsMsgHub.GetMsgId(), 10),
				Method: "getInfo",
			}

			// 将消息传送至消息集合器
			wsMsgHub.WriteJSON(req)

			respData, ok := wsMsgHub.clients.Load(req.ID)
			if !ok {
				t.Log("Receive message data from sync.Map error, id: ", req.ID)
			}
			data := <-respData.(ChanData).Data
			// 从clients中删除该请求的注册信息
			wsMsgHub.clients.unregisterMsgId(req.ID)

			// 仅用于测试是否还存在该ID的信息
			/*
				_, ok = wsMsgHub.clients.Load(req.ID)
				if !ok {
					t.Logf("Id: %s had already deleted", req.ID)
				}
			*/

			t.Log("Receive data : ", data)
		}()
	}

	time.Sleep(time.Second)
	wsMsgHub.Exit()

	return
}
