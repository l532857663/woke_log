package main

import (
	"backend-gateway/library/config"
	httpLib "backend-gateway/library/net/http"

	"github.com/gin-gonic/gin"
)

var (
	// 路由鉴权账户组
	accounts = map[string]string{
		"user1": "ceshi",
	}
	// 数据验签路由
	verifySignatureRouteMap = map[string]string{}
)

// @Description 绑定所有PATH并设置回调
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
func setupRouter(g *gin.Engine, runEnv string) {
	backendRouter := g.Group("/api/")
	{
		// 设置日志模式
		if runEnv != config.RUN_ENV_PROD {
			// TODO: backendRouter.Use(httpLib.LogHttpWrapper())
		}
		// 跨域资源处理
		backendRouter.Use(httpLib.CORS())
		// Token验证
		// backendRouter.Use(httpLib.TokenHandler(verifyXAuthToken))

		// 测试服务通讯状况专用
		g.GET("ping", routerSvr.ping)

		backendRouter := backendRouter.Group("backend/")
		{
			SetupBackendRouter(backendRouter)
		}
	}

}

// @Description 绑定Backend处理函数
// @Author Wangch
// @Version 1.0
// @Update Wangch 2021-09-26 init
// @Update Wangch 2021-10-08 添加商品列表路由
func SetupBackendRouter(router gin.IRouter) {
	// 后台管理 加入鉴权功能
	backstageRouter := router.Group("backstage/", gin.BasicAuth(accounts))
	{
		// adminRouter := backstageRouter.Group("admin/") {}
		backstageRouter.Group("admin/")
	}
	// 商品列表路由
	goodsRouter := router.Group("goods/")
	{
		goodsRouterV1 := goodsRouter.Group("v1/")
		{
			goodsRouterV1.GET("list", routerSvr.getGoodsListV1)
		}
	}
}
