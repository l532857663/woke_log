package main

import (
	"backend-gateway/conf"
	"backend-gateway/model"
	"backend-gateway/service"
	"backend-gateway/utils"
	errMsg "backend-gateway/utils/error_message"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	routerSvr *RouterImpl
)

// @Description 路由服务对象结构
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
type RouterImpl struct {
	BackendSvc *service.Service
}

// @Description 创建一个 rouder 服务实例
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func newRouter(c *conf.Config, appletName string) *RouterImpl {
	if routerSvr == nil {
		routerSvr = &RouterImpl{
			BackendSvc: service.New(c, appletName),
		}
	}

	return routerSvr
}

// @Description 路由服务ping测试
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func (s *RouterImpl) ping(c *gin.Context) {
	lang := c.DefaultQuery(model.LANGUAGE_KEY_NAME, model.LANGUAGE_DEFAULT)

	c.JSON(http.StatusOK, utils.SuccessResponse(utils.GetI18nLocaleMessage(errMsg.PongMessage, lang)))
	return
}
