package service

import (
	"backend-gateway/conf"
	"backend-gateway/dao"
	"backend-gateway/library/config"
	"backend-gateway/service/backend"
	"os"
	"regexp"

	logger "github.com/cihub/seelog"
)

var (
	s *Service
)

// @Description 服务管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type Service struct {
	*backend.BackendService
	AddressRegexer map[uint64]*regexp.Regexp
	terminate      func(status int)
}

// @Description 新建一个服务并初始化资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(c *conf.Config, appletName string) *Service {
	if s == nil {
		s = &Service{
			terminate: os.Exit,
		}
		bs := backend.New(c)
		s.BackendService = bs

		if c.ServerGeneral.RunEnv == config.RUN_ENV_JOIN_DEBUG {
			s.BackendService.Dao = dao.NewMock()
			dao.CloseMockExpect(s.BackendService.Dao)
		} else {
			s.BackendService.Dao = dao.New(c, appletName)
		}

		// 加载需要缓存的数据
		if s.BackendService.Conf.Cache.Enabled && s.BackendService.Conf.Cache.Redis.Enable {
			// TODO: preload ...
			s.LoadCache()
		}
	}

	return s
}

// @Description 关闭服务运行的资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (s *Service) Close() {
	s.BackendService.Dao.Close()
}

// @Description 装载backend钱包token表数据缓存
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-03-09 init
func (s *Service) LoadCache() {
	// 预热缓存
	logger.Info("Load cache success.")
}
