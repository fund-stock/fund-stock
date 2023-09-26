package test

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/app/response/qtimg"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("test")
}

func TestStock(t *testing.T) {
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	fmt.Println("测试")
	TxClient("sz002194")
	//TxClient("sz002261")
}

func TxClient(code string) {
	url := fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%v,day,,,320,qfq", code)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	str := strings.ReplaceAll(string(body), code, "info")
	str = strings.ReplaceAll(str, "qfqday", "day")
	var Resp qtimg.Resp
	err = json.Unmarshal([]byte(str), &Resp)
	if err != nil {
		logger.Error(err)
		return
	}
	if Resp.Code != 0 {
		logger.Info(Resp.Msg)
		return
	}
	DayData := Resp.Data.StockInfo.Day
	Info := Resp.Data.StockInfo.Qt.Info
	fmt.Println(Info)
	fmt.Println(DayData)
	fmt.Println(Resp.Data.StockInfo.Day)
	for _, item := range DayData {
		fmt.Println(item[0], item[1], item[2], item[3], item[4], item[5])
		DB := models.GoStockDayMgr(mysql.DB)
		//
		t, _ := time.Parse("2006-01-02", item[0].(string))
		logger.Info(t.Format("2006-01-02 15:04:05.999"), "===", t.UnixMilli(), "===", item[1], "===", item[2], "===", item[3], "===", item[4], "===", item[5])
		options, err := DB.Debug().GetByOption(DB.WithCode(code), DB.WithDayAt(t))
		if err != nil && !models.IsNotFound(err) {
			logger.Error(err)
			return
		}
		if options.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoFundDay{
			Code:     code,
			Name:     Info[1],
			Amount:   0,
			Nav:      item[2].(float64),
			DayTs:    t.UnixMilli(),
			DayAt:    t,
			CreateAt: time.Now().UnixMilli(),
			UpdateAt: time.Now().UnixMilli(),
		})
	}
}
