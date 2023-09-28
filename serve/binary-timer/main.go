package main

import (
	"github.com/robfig/cron/v3"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/pprof"
	"goapi/serve/binary-timer/task_route"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东七
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-timer")
}

func main() {
	// pprof 分析
	pprof.Debug(":6068")
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	// 新建一个定时任务对象 根据 cron 表达式进行时间调度，cron 可以精确到秒，大部分表达式格式也是从秒开始。
	crontab := cron.New() // 默认从分开始进行时间调度，精确到秒
	task_route.RegisterTaskRoutes(crontab)
	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	select {}
	// 根据实际情况进行控制
}
