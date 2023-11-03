package routes

import (
	"github.com/gin-gonic/gin"
	clientV1 "goapi/app/controllers/email/v1"
	middlewares "goapi/app/middlewares/v3"
)

// RegisterEmailRoutes 注册 Email 路由
func RegisterEmailRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AddressLimit()) // 地区限制
	//router.Use(middlewares.Client())
	// 路由分组 邮件 模块
	EmailRoute := router.Group("/email")
	{
		{ // V1 版本
			EmailV1Group := new(clientV1.GroupMail)
			V1Route := EmailRoute.Group("/")
			{
				// 139.com
				// qq.com
				// gmail.com
				// 163.com

				// 1、【API】查询剩余发送条数接口
				V1Route.POST("/QueryMargin.json", EmailV1Group.MailController.QueryMargin)
				// 2、【API】单发邮件接口
				V1Route.POST("/SendOneEmail.json", EmailV1Group.MailController.SendOneEmail)
				// 3、【API】群发邮件接口
				V1Route.POST("/SendBatchEmail.json", EmailV1Group.MailController.SendBatchEmail)
				// 4、【定时任务】重置邮箱剩余条数：
				// 执行时间/周期：
				//    - 每天的凌晨12:02分
				V1Route.POST("/ResetMargin.json", EmailV1Group.MailController.ResetMargin)
			}
		}

		{ // V2 版本

		}

		{ // V3 版本

		}

	}
}
