package client

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goapi/pkg/logger"
	"goapi/serve/binary-stock/params"
	"goapi/serve/binary-stock/response/qtimg"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetStockDetail(id string) (params.Data, error) {
	var data params.Data
	url := "https://www.fund123.cn/matiaria?fundCode=" + id
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error(err)
		return data, err
	}
	doc.Find("script").Each(func(index int, s *goquery.Selection) {
		scriptContent := s.Text()
		if strings.Contains(scriptContent, "window.context =") {
			context := strings.Replace(scriptContent, "window.context =", "", 1)
			context = strings.TrimRight(strings.TrimSpace(context), ";")
			fmt.Println(context)
			err = json.Unmarshal([]byte(context), &data)
			if err != nil {
				return
			}
		}
	})
	return data, nil
}

// 获取历史数据

func GetHistoryData(code string) qtimg.Resp {
	var Resp qtimg.Resp
	url := fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%v,day,,,320,qfq", code)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error(err)
		return Resp
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return Resp
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return Resp
	}
	str := strings.ReplaceAll(string(body), code, "info")
	str = strings.ReplaceAll(str, "qfqday", "day")
	err = json.Unmarshal([]byte(str), &Resp)
	if err != nil {
		logger.Error(err)
		return Resp
	}
	return Resp
}
