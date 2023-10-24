package analyze

import (
	"fmt"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/client"
	"time"
)

// 分析市场行情

func Market() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	DB := models.GoStockMgr(mysql.DB)
	list, err := DB.Debug().GetByOptions(DB.WithStatus(1))
	if err != nil {
		logger.Error(err)
		return
	}
	for _, item := range list {
		Analyze(item)
	}
}

func Analyze(Stock *models.GoStock) {
	code := Stock.Code
	minute, err := client.GetMinute(code)
	if err != nil {
		logger.Error(err)
		return
	}
	if len(minute.List) <= 0 {
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
	// 心里预期价位检测
	go MonitorExpectedPrice(currentPrice, Stock)
	// 与昨日收盘价进行对比，涨幅超过预期就报警
	go MonitorPercentageChange(currentPrice, PrePrice, Stock)
}
