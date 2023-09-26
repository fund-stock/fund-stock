package main

import (
	"flag"
	"fmt"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/client"
	"goapi/serve/binary-stock/response/qtimg"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-stock")
	//AppPort := flag.Int64("APP_PORT", conf.GetInt64("app.port"), "服务端口")
	flag.Parse()
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
}

func main() {
	// 声明一个切片来存储 map
	var collect []map[string]string
	// 添加 map 到切片
	collect = append(collect,
		map[string]string{"code": "sz002261"},
		map[string]string{"code": "sz002194"},
	)
	for _, item := range collect {
		go StartCollect(item["code"])
	}
	for true {
		time.Sleep(time.Second)
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

// 采集每天的数据

func saveData(stockCode string, HistoryData qtimg.Resp) {
	DayData := HistoryData.Data.StockInfo.Day
	Info := HistoryData.Data.StockInfo.Qt.Info
	fmt.Println(Info)
	fmt.Println(DayData)
	fmt.Println(HistoryData.Data.StockInfo.Day)
	for _, item := range DayData {
		fmt.Println(item[0], item[1], item[2], item[3], item[4], item[5])
		DB := models.GoStockDayMgr(mysql.DB)
		//
		t, _ := time.Parse("2006-01-02", item[0].(string))
		logger.Info(t.Format("2006-01-02 15:04:05.999"), "===", t.UnixMilli(), "===", item[1], "===", item[2], "===", item[3], "===", item[4], "===", item[5])
		options, err := DB.Debug().GetByOption(DB.WithCode(stockCode), DB.WithDayAt(t))
		if err != nil && !models.IsNotFound(err) {
			logger.Error(err)
			return
		}
		if options.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoFundDay{
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
