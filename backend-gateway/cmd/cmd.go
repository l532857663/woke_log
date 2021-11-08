package cmd

import (
	"backend-gateway/conf"
	"backend-gateway/dao"
	"backend-gateway/library/config"
	"backend-gateway/model"
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cmd *CmdObject

	start  = kingpin.Command("start", "Start backend gateway server.").Default()
	syncdb = kingpin.Command("syncdb", "Sync database to latest version.")
	dbName = syncdb.Arg("dbname", "backend db sync").Required().String()
)

// @Description Command Line对象结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type CmdObject struct {
	AppletVersion   string
	AppletName      string
	AppletBuildTime string
	AppletRunEnv    string
	terminate       func(status int)
}

// @Description 初始化命令行参数: 设置 "-h" 参数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func init() {
	// -h 代表 help 参数
	kingpin.HelpFlag.Short('h')
}

// @Description 创建命令行参数对象
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(appletVersion, appletBuild, appletName, appletBuildTime, appletRunEnv string) *CmdObject {
	if cmd == nil {
		cmd = &CmdObject{
			AppletVersion:   appletVersion + "-build-" + appletBuild,
			AppletName:      appletName,
			AppletBuildTime: appletBuildTime,
			AppletRunEnv:    appletRunEnv,
			terminate:       os.Exit,
		}
	}

	return cmd
}

// @Description Command Line参数初始化
// @Author Wangch
// @Version 1.0
// @param CmdObject
// @return void
// @Update Wangch 2021-09-26 init
func (cmd *CmdObject) CmdInit() {
	// 设置版本号
	var versionStr string
	if cmd.AppletRunEnv == config.RUN_ENV_PROD {
		versionStr = fmt.Sprintf("Version: \t%s\r\nBuildTime: \t%s", cmd.AppletVersion, cmd.AppletBuildTime)
	} else {
		versionStr = fmt.Sprintf("Version: \t%s\r\nBuildTime: \t%s\r\nRun Env: \t%s", cmd.AppletVersion, cmd.AppletBuildTime, cmd.AppletRunEnv)
	}
	kingpin.Version(versionStr)
	// -v 代表 version 参数
	kingpin.CommandLine.VersionFlag.Short('v')

	// 命令行参数处理
	switch kingpin.MustParse(kingpin.Parse(), nil) {
	// "start" command 执行 main 函数
	case start.FullCommand():
		break

		// "syncdb" command 同步最新的 model 至指定数据库
	case syncdb.FullCommand():
		syncDbHandler(nil)
	}

	return
}

// @Description 使用 syncdb 参数, 根据 model 同步对应的数据库
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func syncDbHandler(*kingpin.ParseContext) error {
	conf := conf.GetServiceCfg(cmd.AppletName, cmd.AppletRunEnv)

	conf.Cache.Enabled = false // 数据库同步模式下不需要连接缓存
	db := dao.New(conf, model.APPLET_NAME_BACKEND)

	// 同步指定数据库
	db.SyncDB(*dbName)
	fmt.Printf("Sync db %s complete.\n", *dbName)

	db.Close()

	cmd.terminate(0)

	return nil
}
