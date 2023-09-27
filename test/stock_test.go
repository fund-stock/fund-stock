package test

import (
	"fmt"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/serve/binary-stock/client"
	"testing"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-stock-test")
}

func TestGetHistoryData(t *testing.T) {
	fmt.Println(1212)
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	fmt.Println("测试")
	//client.GetHistoryData("sz002194")
	//client.GetHistoryData("sz002261")
	minute, err := client.GetMinute("sz002261")
	if err != nil {
		return
	}
	fmt.Println(minute)
	fmt.Println(minute)
	time.Sleep(time.Second * 20)
	//TxClient("sz002261")
}
