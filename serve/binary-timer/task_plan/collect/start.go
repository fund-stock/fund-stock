package collect

import (
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/client"
)

func StartStock() {
	DB := models.GoStockMgr(mysql.DB)
	options, err := DB.Debug().GetByOptions(DB.WithStatus(1))
	if err != nil {
		return
	}
	logger.Info(options)
	for _, item := range options {
		StartCollect(item.Code)
	}
}

// 采集

func StartCollect(stockCode string) {
	// 获取每天的历史数据
	HistoryData := client.GetHistoryData(stockCode)
	if HistoryData.Code == 0 && len(HistoryData.Data.StockInfo.Day) > 0 {
		go saveData(stockCode, HistoryData)
	}
}
