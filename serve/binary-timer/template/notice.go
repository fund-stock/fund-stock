package template

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/serve/binary-stock/client"
)

func TplUp(Stock *client.SharesInfoDetails, diffPrice, zfPercent, currentPrice, PrePrice interface{}) map[string]interface{} {
	msg := cmap.New().Items()
	// 上涨通知
	msg["title"] = fmt.Sprintf("%v-%v-上涨", client.ClassifyBoard(Stock.Code), Stock.Name)
	msg["body"] = fmt.Sprintf(""+
		"⬆️涨：%v ¥\n"+
		"⬆️涨幅：%v %%\n"+
		"当前价位：%v\n"+
		"昨日收盘价：%v\n",
		diffPrice, zfPercent, currentPrice, PrePrice,
	)
	return msg
}

func TplDown(Stock *client.SharesInfoDetails, diffPrice, zfPercent, currentPrice, PrePrice interface{}) map[string]interface{} {
	msg := cmap.New().Items()
	// 下跌通知
	msg["title"] = fmt.Sprintf("%v-%v-下跌", client.ClassifyBoard(Stock.Code), Stock.Name)
	msg["body"] = fmt.Sprintf(""+
		"⬇️跌：%v ¥\n"+
		"⬇️跌幅：%v %%\n"+
		"当前价位：%v\n"+
		"昨日收盘价：%v\n",
		diffPrice, zfPercent, currentPrice, PrePrice,
	)
	return msg
}
