package main

import (
	"backend-gateway/cmd"
	"backend-gateway/conf"
	"backend-gateway/library/log"
	_ "net/http/pprof"

	logger "github.com/cihub/seelog"
)

// @Description ldflags编译传导的参数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
var (
	appletVersion   string
	appletBuild     string
	appletName      string
	appletBuildTime string
	appletRunEnv    string
)

// @Description 初始化命令行参数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func init() {
	// 命令行参数解析初始化
	cmdObject := cmd.New(appletVersion, appletBuild, appletName, appletBuildTime, appletRunEnv)

	cmdObject.CmdInit()
}

// @Description 本服务为 backend gateway 接口服务
// @Author Wangch
// @Version 1.0
// @Host 127.0.0.1:4396
// @BasePath /
func main() {
	// 日志初始化
	log.InitLog()

	// 应用配置文件初始化
	conf := conf.GetServiceCfg(appletName, appletRunEnv)
	logger.Info(conf.ServerGeneral.Name, "configuration initialization is complete!")

	// 启动服务
	backendServer := newServer(conf, appletName)
	backendServer.start()
	logger.Infof("server %s start success!", appletName)

	// 监听并捕获消息
	signalHandler(appletName)

	return
}
