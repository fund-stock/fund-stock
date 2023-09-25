package test

import (
	"encoding/json"
	"fmt"
	"goapi/app/response/qtimg"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestStock(t *testing.T) {
	fmt.Println("测试")
	TxClient("sz002261")
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
	fmt.Println(Resp.Data.StockInfo.Day)
	fmt.Println(Resp.Data.StockInfo.Day)
	for _, item := range DayData {
		fmt.Println(item[0], item[1], item[2], item[3], item[4], item[5])
	}
}
