package task_plan

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/shopspring/decimal"
	"goapi/pkg/logger"
	"goapi/pkg/notice"
	"goapi/serve/binary-stock/client"
	"time"
)

// 分析市场行情

func AnalyzeMarket() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	code := "sz002261"
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
			"⬆️涨幅：%v %%\n"+
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
