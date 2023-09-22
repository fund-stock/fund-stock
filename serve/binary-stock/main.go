package main

import (
	"flag"
	"goapi/app/models"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/client"
	"goapi/serve/binary-stock/params"
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
	// stockCode := "005819"
	collect := map[int]map[string]string{
		0: {"code": "002261"},
	}
	for _, item := range collect {
		go StartCollect(item["code"])
	}
	for true {
		time.Sleep(time.Second)
	}
}

// 采集

func StartCollect(fundCode string) {
	data, err := client.GetStockDetail(fundCode)
	if err != nil {
		logger.Error(err)
		return
	}
	DB := models.GoFundMgr(mysql.DB)
	options, err := DB.Debug().GetByOption(DB.WithCode(fundCode))
	if err != nil && !models.IsNotFound(err) {
		logger.Error(err)
		return
	}
	if options.ID == 0 {
		DB.Debug().Create(&models.GoFund{
			Code:      data.MaterialInfo.FundCode,
			ProductID: data.MaterialInfo.ProductId,
			Name:      data.MaterialInfo.FundBrief.FundNameAbbr,
			Amount:    0,
			Nav:       0,
			Status:    0,
			CreateAt:  time.Now().UnixMilli(),
			UpdateAt:  time.Now().UnixMilli(),
		})
	}
	// 获取每天的历史数据
	currentPage := 1
	HistoryData := client.GetHistoryData(data.MaterialInfo.ProductId, currentPage)
	if HistoryData.Success {
		go saveData(fundCode, data.MaterialInfo.FundBrief.FundNameAbbr, HistoryData)
		//for currentPage < HistoryData.TotalPages {
		//	time.Sleep(time.Second * 1)
		//	currentPage = currentPage + 1
		//	res := client.GetHistoryData(data.MaterialInfo.ProductId, currentPage)
		//	if res.Success {
		//		saveData(fundCode, data.MaterialInfo.FundBrief.FundNameAbbr, res)
		//	}
		//}
	}
}

// 采集每天的数据

func saveData(fundCode, name string, HistoryData params.HistoryData) {
	for _, item := range HistoryData.List {
		DB := models.GoFundDayMgr(mysql.DB)
		t, _ := time.Parse("2006-01-02", item.NetValueDate)
		logger.Info(t.Format("2006-01-02 15:04:05.999"), "===", t.UnixMilli(), "===", item.NetValueDate, "===", item.NetValue, "===", item.TotalNetValue, "===", item.DayOfGrowth)
		options, err := DB.Debug().GetByOption(DB.WithCode(fundCode), DB.WithDayAt(t))
		if err != nil && !models.IsNotFound(err) {
			logger.Error(err)
			return
		}
		if options.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoFundDay{
			Code:     fundCode,
			Name:     name,
			Amount:   0,
			Nav:      helpers.StrToFloat64(item.TotalNetValue),
			DayTs:    t.UnixMilli(),
			DayAt:    t,
			CreateAt: time.Now().UnixMilli(),
			UpdateAt: time.Now().UnixMilli(),
		})
	}
}
