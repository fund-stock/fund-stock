package routes

import (
	"github.com/gin-gonic/gin"
	WebCtl "goapi/app/controllers/web/v1"
	WebMw "goapi/app/middlewares/v1"
)

var WebV1Group = new(WebCtl.Group)

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(ApiRoute *gin.RouterGroup) {
	// 路由分组 Web 接口
	ApiRoute.Use(WebMw.Web())
	WebV1 := ApiRoute.Group("/web/v1")
	// 看板模块
	Kanban := WebV1.Group("/kanban")
	{
		// 【E32】 历史行情-mysql读取
		Kanban.Any("/getHisKlineList.json", WebV1Group.KlineController.HistoryHandler)
	}

	// 设置模块
	Config := WebV1.Group("/config")
	{
		// 【E36】 所有 redis 的Key
		Config.Any("/getRedisKey.json", WebV1Group.KlineController.GetRedisKeyHandler)
		// 【E37】 设置token
		Config.Any("/setKlineToken.json", WebV1Group.KlineController.SetKlineToken)
	}

}
