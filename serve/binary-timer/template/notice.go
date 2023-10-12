package template

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
)

func TplUp(name string, diffPrice, zfPercent, currentPrice, PrePrice interface{}) map[string]interface{} {
	msg := cmap.New().Items()
	// 上涨通知
	msg["title"] = fmt.Sprintf("%v-上涨", name)
	msg["body"] = fmt.Sprintf(""+
		"⬆️涨：%v ¥\n"+
		"⬆️涨幅：%v %%\n"+
		"当前价位：%v\n"+
		"昨日收盘价：%v\n",
		diffPrice, zfPercent, currentPrice, PrePrice,
	)
	return msg
}

func TplDown(name string, diffPrice, zfPercent, currentPrice, PrePrice interface{}) map[string]interface{} {
	msg := cmap.New().Items()
	// 下跌通知
	msg["title"] = fmt.Sprintf("%v-下跌", name)
	msg["body"] = fmt.Sprintf(""+
		"⬇️跌：%v ¥\n"+
		"⬇️跌幅：%v %%\n"+
		"当前价位：%v\n"+
		"昨日收盘价：%v\n",
		diffPrice, zfPercent, currentPrice, PrePrice,
	)
	return msg
}
