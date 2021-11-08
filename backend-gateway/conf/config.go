package conf

import (
	"log"

	"backend-gateway/library/config"

	logger "github.com/cihub/seelog"
)

// 程序相关配置
type Config struct {
	ServerGeneral *config.ServerGeneralConfig
	Cache         *config.CacheConfig
	DBBackend     *config.DatabaseConfig
}

// 获取业务配置信息
func GetServiceCfg(appletName string, runEnv string) *Config {
	conf := &Config{}

	myViper := config.InitConfig(appletName)

	err := myViper.Unmarshal(conf)

	if err != nil {
		log.Panic("Error loading configuration: ", err)
	}

	logger.Infof("load configuration sucessed")

	// 设置服务运行环境
	conf.ServerGeneral.RunEnv = runEnv

	// 设置数据库DSN
	config.InitDSN(conf.DBBackend)

	return conf
}
