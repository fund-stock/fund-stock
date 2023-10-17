package collect

import (
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/response/qtimg"
	"time"
)

// 采集每天的数据

func saveData(stockCode string, HistoryData qtimg.Resp) {
	DayData := HistoryData.Data.StockInfo.Day
	Info := HistoryData.Data.StockInfo.Qt.Info
	for _, item := range DayData {
		DB := models.GoStockDayMgr(mysql.DB)
		//
		t, _ := time.Parse("2006-01-02", item[0].(string))
		options, err := DB.Debug().GetByOption(DB.WithCode(stockCode), DB.WithDayAt(t))
		if err != nil && !models.IsNotFound(err) {
			logger.Error(err)
			return
		}
		if options.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoStockDay{
			Code:     stockCode,
			Name:     Info[1],
			Amount:   0,
			Nav:      helpers.StrToFloat64(item[2].(string)),
			DayTs:    t.UnixMilli(),
			DayAt:    t,
			CreateAt: time.Now().UnixMilli(),
			UpdateAt: time.Now().UnixMilli(),
		})
	}
}
