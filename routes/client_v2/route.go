package client_v2

import (
	clientV2 "goapi/app/controllers/client/v2"
	middlewaresV2 "goapi/app/middlewares/v2"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	router.Use(middlewaresV2.AddressLimit()) // 地区限制
	//router.Use(middlewares.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V2 版本
			ClientV2Group := new(clientV2.Group)
			V2Route := AppRoute.Group("/v2")

			// 代理落地页接口
			config := V2Route.Group("/config")
			{
				// 设置token
				config.POST("/setKlineToken.json", ClientV2Group.ConfigController.SetKlineToken)
			}
		}

	}
}
