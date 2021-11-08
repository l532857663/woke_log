package http

import (
	"backend-gateway/library/config"
	"fmt"
	"net/http"

	logger "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

type routerHandler func(*gin.Engine, string)

// Init 创建一个http服务实例
func Init(c *config.ServerGeneralConfig, rh routerHandler) *http.Server {
	// 设置日志模式
	switch c.RunEnv {
	case config.RUN_ENV_DEV, config.RUN_ENV_JOIN_DEBUG, config.RUN_ENV_LOCAL:
		gin.SetMode(gin.DebugMode)
	case config.RUN_ENV_BETA:
		gin.SetMode(gin.TestMode)
	case config.RUN_ENV_PROD:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	httpHandler := gin.Default()
	rh(httpHandler, c.RunEnv)

	logger.Info("Server initialized successfully.")

	// 监听 ip:port
	address := fmt.Sprintf("%s:%d", c.ListenAddress, c.ListenPort)

	httpSvc := &http.Server{
		Addr:    address,
		Handler: httpHandler,
	}

	// 启动服务
	logger.Infof("starting server on %s", address)

	return httpSvc
}
