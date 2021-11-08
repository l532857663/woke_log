package basic

import (
	"backend-gateway/conf"
	"backend-gateway/dao"
	"regexp"
	"sync"
)

var (
	SC *ServiceContext
)

// @Description 服务管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type ServiceContext struct {
	Conf           *conf.Config
	Dao            dao.Storage
	DbConf         *DbServiceConfig
	AddressRegexer map[uint64]*regexp.Regexp
}

// @Description 服务相关业务配置
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type DbServiceConfig struct {
	sync.Mutex
}

// @Description 新建一个服务并初始化资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(c *conf.Config) *ServiceContext {
	if SC == nil {
		SC = &ServiceContext{
			Conf: c,
		}
	}

	return SC
}
