package test

import (
	"goapi/config"
	"goapi/pkg/logger"
	"goapi/pkg/notice/bark"
	"testing"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func TestBark(t *testing.T) {
	bark.Notice("测试通知消息", "恭喜您获取成功，该次上涨成功，请您留意当前的股票走势，看好买点时间，当前已经达到最好的入场时间，估计比较低错过了就要等待下一波了，当前赢利点达到您的预期，是否前往查看并且卖出适当的仓位，减少持仓")
	time.Sleep(time.Second * 5)
}
