package main

import (
	"backend-gateway/conf"
	httpLib "backend-gateway/library/net/http"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	logger "github.com/cihub/seelog"
)

// gracefulTimeout controls how long we wait before forcefully terminating
const (
	gracefulTimeout   = 3 * time.Second
	asyncManagerCount = 1
)

var (
	server *ServerImpl
	wg     sync.WaitGroup
)

// @Description 服务管理对象结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type ServerImpl struct {
	httpRouterSvr *RouterImpl
	httpSvr       *http.Server
	shutdownCh    chan struct{}
}

// @Description 创建一个 server 服务实例
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func newServer(c *conf.Config, appletName string) *ServerImpl {
	if server == nil {
		server = &ServerImpl{
			httpRouterSvr: newRouter(c, appletName),
			shutdownCh:    make(chan struct{}),
		}
	}
	// service.Init(c.BingooPlatform, &wg)
	server.init(c)

	return server
}

// @Description 初始化服务
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (s *ServerImpl) init(c *conf.Config) {
	// 初始化 pprof 分析工具
	// httpLib.InitPprofService(c.ServerGeneral)

	// 初始化服务以及路由
	s.httpSvr = httpLib.Init(c.ServerGeneral, setupRouter)

	return
}

// @Description 启动服务
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (s *ServerImpl) start() {
	// 启动 http 服务
	go s.httpSvr.ListenAndServe()

	return
}

// @Description 关闭服务
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (s *ServerImpl) shutdown() {
	// 使用 google 官方建议的方式关闭 server 服务
	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()

	err := s.httpSvr.Shutdown(ctx)
	if err != nil {
		logger.Error("HTTP server shutdown failure: ", err.Error())
	}

	// 关闭业务层处理
	s.httpSvr.Close()

	// 关闭 dao 层资源
	s.httpRouterSvr.BackendSvc.Close()

	return
}

// @Description 系统信号捕获
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func signalHandler(appletName string) {
	var (
		ch = make(chan os.Signal)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

WAIT_SIGNAL:
	for {
		sig := <-ch
		logger.Infof("get a signal \"%s\", stop the %s process", sig.String(), appletName)
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP:
			// 安全关闭服务
			go func() {
				server.shutdown()
				close(server.shutdownCh)
				return
			}()

			break WAIT_SIGNAL

		default:
			return
		}
	}

	select {
	case <-server.shutdownCh:
		fmt.Println("Gracefully shutting down server ...")
	}

	return
}
