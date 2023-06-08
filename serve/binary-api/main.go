package main

import (
	"flag"
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"time"

	"github.com/gin-contrib/pprof"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-client")
}

// @title 用户端接口服务
// @version 3.0
// @description 3.0版本，基于之前的2.0改造的
// @termsOfService http://127.0.0.1/docs/index.html

// @contact.name 追梦小窝
// @contact.url http://github.com/iszmxw
// @contact.email mail@54zm.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath
func main() {
	AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	// 初始化路由绑定
	logger.Info("加载 client 路由")
	router := bootstrap.SetupRoute()
	pprof.Register(router) // 开启 pprof
	// 启动路由
	logger.Info("启动路由")
	logger.Info(fmt.Sprintf("当前环境:%v", *AppPort))
	_ = router.Run(fmt.Sprintf(":%v", *AppPort))
}
