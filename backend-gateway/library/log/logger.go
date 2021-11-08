package log

import (
	"os"

	logger "github.com/cihub/seelog"
)

const (
	LOG_CONFIG_FILE = "./conf/conf-logger.xml"
)

func InitLog() {
	// 日志文件初始化
	loggerObj, err := logger.LoggerFromConfigAsFile(LOG_CONFIG_FILE)
	if err != nil {
		logger.Critical("seelog configuration initialization is error：", err)
		os.Exit(-1)
	}

	logger.ReplaceLogger(loggerObj)
	logger.Flush()

	return
}
