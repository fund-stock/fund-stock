package bootstrap

import (
	"fmt"
	"goapi/app/middlewares/common"
	"goapi/pkg/redis"
	"goapi/routes"
	"goapi/routes/client_v1"
	"goapi/routes/client_v2"
	"goapi/routes/client_v3"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(common.TraceLogger()) // 日志追踪
	router.Use(common.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		num := redis.Client.SCard("onlineUsers")
		context.String(200, fmt.Sprintf("Hello World!：%v--%v\n\n\n", requestId, num))
		context.String(200, "下面是所有接口服务：\n\n\n\t")
		routers := router.Routes()
		for _, v := range routers {
			context.String(200, fmt.Sprintf("Method：%v\tURL：%v  \n\tHandler: %v\n\t=====================\n\t\n\t", v.Method, v.Path, v.Handler))
		}
	})
	router.GET("/shell", func(c *gin.Context) {
		cmd := exec.Command("sh", "bin/Synchronize")
		output, err := cmd.Output()
		if err != nil {
			c.String(200, err.Error()+"\n\n\n")
			return
		}
		c.String(200, string(output))
	})
	// 二元 client 接口
	client_v1.RegisterClientRoutes(router.Group("/"))
	client_v2.RegisterClientRoutes(router.Group("/"))
	client_v3.RegisterClientRoutes(router.Group("/"))
	return router
}

// SetupRouteWeb 路由初始化
func SetupRouteWeb() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(common.TraceLogger()) // 日志追踪
	router.Use(common.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello Web!："+requestId.(string)+"\n\n\n")
	})
	// 二元 web 接口
	routes.RegisterWebRoutes(router.Group("/"))
	return router
}

// SetupRouteEmail 路由初始化
func SetupRouteEmail() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(common.TraceLogger()) // 日志追踪
	router.Use(common.Cors())        // 跨域
	router.NoRoute(NoResponse)
	router.GET("/", func(context *gin.Context) {
		requestId, _ := context.Get("Tracking-Id")
		context.String(200, "Hello Email!："+requestId.(string)+"\n\n\n")
	})
	// 二元 Email 接口
	routes.RegisterEmailRoutes(router.Group("/"))
	return router
}

func NoResponse(c *gin.Context) {
	//返回 404 状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}
