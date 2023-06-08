package client_v1

import (
	clientV1 "goapi/app/controllers/client/v1"
	middlewaresV1 "goapi/app/middlewares/v1"

	"github.com/gin-gonic/gin"
)

// RegisterClientRoutes 注册Client路由
func RegisterClientRoutes(router *gin.RouterGroup) {
	router.Use(middlewaresV1.AddressLimit()) // 地区限制
	//router.Use(middlewares.Client())
	// 路由分组 客户端 模块
	AppRoute := router.Group("/app")
	{
		{ // V1 版本
			ClientV1Group := new(clientV1.Group)
			// 【A】默认模块
			{
				// 【A02】获取系统多语言包【POST】
				AppRoute.Group("/v1/config").POST("/getLanguagePage.json", ClientV1Group.ConfigController.GetLanguagePageHandler)
				// 【A03】校验app版本【POST】
				AppRoute.Group("/config/v1").POST("/checkAppVersion.json", ClientV1Group.ConfigController.CheckAppVersionHandler)
				// 【A04】获取引导教程【POST】
				AppRoute.Group("/v1/config").POST("/getTutorialByType.json", ClientV1Group.ConfigController.GetTutorialByTypeHandler)
				// 【A05】获取测试账号发送验证码【POST】
				AppRoute.Group("/v1/config").POST("/getLoginCodeByTest.json", ClientV1Group.ConfigController.GetLoginCodeByTestHandler)
				// 【A06】清除测试账号【POST】
				AppRoute.Group("/v1/config").POST("/clearTestAccount.json", ClientV1Group.ConfigController.ClearTestAccountHandler)
				// 【A09】获取国家信息【POST】
				AppRoute.Group("/v1/config").POST("/getExistingCountry.json", ClientV1Group.ConfigController.GetExistingCountryHandler)
				// 【A10】获取系统多语言包V2【POST】
				AppRoute.Group("/v1/config").POST("/getLanguagePageV2.json", ClientV1Group.ConfigController.GetLanguagePageV2Handler)
				// 【A11】获取系统展示图片【POST】
				AppRoute.Group("/v1/config").POST("/getImages.json", ClientV1Group.ConfigController.GetImagesHandler)
			}

		}

	}
}
