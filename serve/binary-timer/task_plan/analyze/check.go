package analyze

import (
	"fmt"
	"github.com/shopspring/decimal"
	"goapi/app/models"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/notice"
	"goapi/serve/binary-stock/client"
	"goapi/serve/binary-timer/template"
)

func MonitorExpectedPrice(Stock *client.SharesInfoDetails) {
	// 最新的价位
	currentPrice := Stock.Price
	DB := models.GoStockMgr(mysql.DB)
	DbStock, err := DB.Debug().GetByOption(DB.WithStatus(1), DB.WithCode(Stock.Code))
	if err != nil {
		logger.Error(err)
		return
	}
	if DbStock.Nav <= 0 {
		return
	}
	// 股价降到心里预期价位，赶快预警
	if decimal.NewFromFloat(currentPrice).
		Sub(decimal.NewFromFloat(DbStock.Nav)).
		GreaterThan(decimal.NewFromFloat(0)) {
		return
	}
	notice.Notice(fmt.Sprintf("达到心里预期价格:%v", currentPrice), fmt.Sprintf(`请留意[%v][%v],该价位是买入的好时机`, Stock.Name, Stock.Code))
}

// 监控涨跌幅变化

func MonitorPercentageChange(Stock *client.SharesInfoDetails) {
	// 最新的价位
	currentPrice := Stock.Price
	// 昨日收盘价
	PrePrice := Stock.YesterdayClosePrice
	// 差价
	diffPrice, _ := decimal.NewFromFloat(currentPrice).Sub(decimal.NewFromFloat(PrePrice)).Float64()
	if diffPrice == 0 {
		return
	}
	zf, _ := decimal.NewFromFloat(diffPrice).Div(decimal.NewFromFloat(PrePrice)).Float64()
	zfPercent := decimal.NewFromFloat(zf * 100).Round(2)
	// 涨跌幅大于 10 % 进行报警（考虑加入到监控列表）
	if zfPercent.Abs().LessThan(decimal.NewFromFloat(10)) {
		return
	}
	if diffPrice > 0 {
		msg := template.TplUp(Stock, diffPrice, zfPercent, currentPrice, PrePrice)
		notice.Notice(msg["title"].(string), msg["body"].(string))
	} else {
		msg := template.TplDown(Stock, diffPrice, zfPercent, currentPrice, PrePrice)
		notice.Notice(msg["title"].(string), msg["body"].(string))
	}
}
