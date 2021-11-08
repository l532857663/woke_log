package http

import (
	"backend-gateway/library/config"
	"fmt"
	"log"
	"net/http"
	"runtime"

	logger "github.com/cihub/seelog"
)

// 创建一个 pprof 服务实例
func InitPprofService(c *config.ServerGeneralConfig) error {
	if c.PProfConf.Enabled {
		// 服务 ip:port
		address := fmt.Sprintf("%s:%d", c.PProfConf.Address, c.PProfConf.Port)

		go func() {
			// runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
			runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
			runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

			log.Panic("Go pprof service failed:", http.ListenAndServe(address, nil))
		}()

		logger.Info("Starting Go pprof service on:", address)
	}

	return nil
}
