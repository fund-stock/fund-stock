package test

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/shopspring/decimal"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/notice"
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
	code := "sz002528"
	code = "sz002261"
	AnalyzeMarket(code)
	time.Sleep(time.Second * 60)
}

// 分析市场行情
func AnalyzeMarket(code string) {
	minute, err := client.GetMinute(code)
	if err != nil {
		logger.Error(err)
		return
	}
	for _, item := range minute.List {
		// ToDo 分析数据
		fmt.Println(item)
	}
	// 最新的价位
	currentPrice := minute.List[len(minute.List)-1].Price
	// 昨日收盘价
	PrePrice := minute.PrePrice
	// 差价
	diffPrice, _ := decimal.NewFromFloat(currentPrice).Sub(decimal.NewFromFloat(PrePrice)).Float64()
	msg := cmap.New().Items()
	if diffPrice > 0 {
		zf, _ := decimal.NewFromFloat(diffPrice).Div(decimal.NewFromFloat(PrePrice)).Float64()
		// 上涨通知
		msg["title"] = fmt.Sprintf("%v-上涨通知", code)
		msg["body"] = fmt.Sprintf(""+
			"⬆️涨：%v ¥\n"+
			"涨幅：%v %%\n"+
			"当前价位：%v\n"+
			"昨日收盘价：%v\n",
			diffPrice, decimal.NewFromFloat(zf*100).Round(2), currentPrice, PrePrice,
		)
	} else {
		zf, _ := decimal.NewFromFloat(diffPrice).Div(decimal.NewFromFloat(PrePrice)).Float64()
		// 下跌通知
		msg["title"] = fmt.Sprintf("%v-下跌通知", code)
		msg["body"] = fmt.Sprintf(""+
			"⬇️跌：%v ¥\n"+
			"⬇️跌幅：%v %%\n"+
			"当前价位：%v\n"+
			"昨日收盘价：%v\n",
			diffPrice, decimal.NewFromFloat(zf*100).Round(2), currentPrice, -PrePrice,
		)
	}
	go notice.Notice(msg["title"].(string), msg["body"].(string))
}
