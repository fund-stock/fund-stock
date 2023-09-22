package client

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goapi/pkg/logger"
	"goapi/serve/binary-stock/params"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func GetHistoryData(productId string, currencyPage int) params.HistoryData {
	var HistoryData params.HistoryData
	url := "https://www.fund123.cn/api/fund/queryFundHistoryNetValueList?_csrf=zxnNDjyl-YPCPJfBajwuDAurCNpEzDBFoJeI"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
		"productId": "%s",
		"startDate": "20220525",
		"endDate": "%s",
		"pageNum": %d,
		"pageSize": 10
	}`, productId, time.Now().Format("20060102"), currencyPage))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return HistoryData
	}
	req.Header.Add("authority", "www.fund123.cn")
	req.Header.Add("accept", "json")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "ctoken=v1twu0ri-RFc-QWIbXitWSgl; ALIPAYJSESSIONID=uK0yhf1j7hLqSfEDuGDeCoDDweEJTGOCfinfundpcbff; spanner=Piek2n3Wl3P+cBm7toZJZqAbL3jSujDGXt2T4qEYgj0=; spanner=Xt9ppIMQST4Oy1dCyZ/e4NMJMnliwLyz")
	req.Header.Add("origin", "https://www.fund123.cn")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://www.fund123.cn/matiaria?fundCode=012863")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return HistoryData
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logger.Error(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return HistoryData
	}

	err = json.Unmarshal(body, &HistoryData)
	if err != nil {
		logger.Error(err)
		return HistoryData
	}
	return HistoryData
}
