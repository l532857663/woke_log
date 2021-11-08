package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	logger "github.com/cihub/seelog"
	"github.com/spf13/viper"
)

const (
	OFFICIAL_PATH = "./conf"
)

// 初始化配置信息
func InitConfig(configName string) *viper.Viper {
	config := viper.New()
	initViper(config, configName)

	// 获取配置文件路径前缀信息
	prefix := strings.ToUpper(configName)

	config.SetEnvPrefix(prefix)
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	err := config.ReadInConfig()
	if err != nil {
		log.Panicf("Error reading configuration: %s, make sure your config file exists or format is correct.", err.Error())
	} else {
		logger.Debugf("Using config file: %s", config.ConfigFileUsed())
	}

	return config
}

// Viper 配置层初始化。
// 主要的目标是建立路径，参考这些路径来找到我们需要的配置
// 如果 viper 实例为空，将会初始化全局变量
func initViper(v *viper.Viper, configName string) error {
	prefix := strings.ToUpper(configName)
	logger.Debugf("Init component configuration with prefix: %s", prefix)

	// 获取用户设置的环境变量
	var altPath = os.Getenv(fmt.Sprintf("%s_CFG_PATH", prefix))
	if altPath != "" {
		logger.Debugf("adding altPath [%s] for viper configuration", altPath)
		addConfigPath(v, altPath)
	} else {
		// 添加当前工作路径CWD作为处理路径
		addConfigPath(v, "./")

		// 添加配置文件处理路径
		if dirExists(OFFICIAL_PATH) {
			logger.Debugf("adding [%s] to config path", OFFICIAL_PATH)
			addConfigPath(v, OFFICIAL_PATH)
		}
	}

	// 设置配置文件
	if v != nil {
		v.SetConfigName(configName)
	} else {
		viper.SetConfigName(configName)
	}

	return nil
}

func addConfigPath(v *viper.Viper, p string) {
	if v != nil {
		v.AddConfigPath(p)
	} else {
		viper.AddConfigPath(p)
	}
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func InitDSN(dbConf *DatabaseConfig) {
	// 设置数据库DSN
	dbConf.Dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Address,
		dbConf.DBname,
		dbConf.Charset,
		dbConf.ParseTime,
	)

	return
}
