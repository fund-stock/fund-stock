package test

import (
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/yiyan"
	"testing"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-wechat-test")
}

func TestChatGpt(t *testing.T) {
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	yiyan.Chat()
}
