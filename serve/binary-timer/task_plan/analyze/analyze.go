package analyze

import (
	"fmt"
	"goapi/app/models"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/client"
	"time"
)

// 分析市场行情

func Market() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	DB := mysql.DB
	var codes []string
	models.GoStockMgr(DB.Where("status = ?", 1).Group("code")).Debug().Select("code").Find(&codes)
	var outs []*client.SharesInfoDetails
	num := 800
	for {
		if len(codes) > num {
			searchCodes := codes[:num]
			SearchDetailsOuts := client.SearchDetails(searchCodes)
			outs = append(outs, SearchDetailsOuts...)
			codes = codes[num:]
		} else {
			SearchDetailsOuts := client.SearchDetails(codes)
			outs = append(outs, SearchDetailsOuts...)
			break
		}
	}
	for _, item := range outs {
		Analyze(item)
	}
}

func Analyze(Stock *client.SharesInfoDetails) {
	// 心里预期价位检测
	go MonitorExpectedPrice(Stock)
	// 与昨日收盘价进行对比，涨幅超过预期就报警
	go MonitorPercentageChange(Stock)
}
