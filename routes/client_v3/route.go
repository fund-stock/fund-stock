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

			// 微信接口
			wechat := V3Route.Group("/wechat")
			{
				wechat.Any("/init.json", ClientV3Group.WechatController.ServeWechat)
				wechat.Any("/userSynchronize.json", ClientV3Group.WechatController.Synchronize)
				// 发送消息接口
				wechat.Any("/Send.json", ClientV3Group.WechatController.Send)
			}
		}

	}
}
