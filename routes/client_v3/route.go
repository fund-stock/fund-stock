package client_v3

import (
	clientV3 "goapi/app/controllers/client/v3"
	middlewaresV3 "goapi/app/middlewares/v3"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	router.Use(middlewaresV3.AddressLimit()) // 地区限制
	router.Use(middlewaresV3.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V3 版本
			ClientV3Group := new(clientV3.Group)
			V3Route := AppRoute.Group("/v3")

			// 设置token接口
			config := V3Route.Group("/config")
			{
				// 设置token
				config.POST("/setKlineToken.json", ClientV3Group.ConfigController.SetKlineToken)
				// 检查多语言是否包含中文，然后发送警告
				config.POST("/checkLanguage.json", ClientV3Group.ConfigController.CheckLanguage)
			}
		}

	}
}
