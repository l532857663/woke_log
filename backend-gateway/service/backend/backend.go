package backend

import (
	"backend-gateway/conf"
	"backend-gateway/service/basic"
)

var (
	S               *BackendService
	serviceDisabled bool
)

// @Description 服务管理资源结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type BackendService struct {
	*basic.ServiceContext
}

// @Description 新建一个服务并初始化资源
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func New(c *conf.Config) *BackendService {
	if S == nil {
		sc := basic.New(c)
		S = &BackendService{sc}
	}

	return S
}
