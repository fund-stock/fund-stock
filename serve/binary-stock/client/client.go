package client

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/serve/binary-stock/response/qtimg"
	"io/ioutil"
	"net/http"
	"strings"
)

// 获取分时原始数据

func GetMinute(code string) (*qtimg.Minute, error) {
	// 开始获取
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/minute/query?code=%v", code), "")
	out = strings.Replace(out, code, "info", -1)
	var info qtimg.MinuteWWW
	err := json.Unmarshal([]byte(out), &info)
	if err != nil {
		return nil, err
	}
	var sumPrice float64
	var sumVol int64
	var tmp qtimg.Minute
	if info.Code == 0 {
		if len(info.Data.Info.QT.Info) > 4 {
			tmp.PrePrice = helpers.StrToFloat64(info.Data.Info.QT.Info[4])
		}
		// index := len(info.Data.Info.Data.Data) - 1
		for _, v := range info.Data.Info.Data.Data {
			out := strings.Split(v, " ")
			if len(out) == 4 {
				minfo := qtimg.MinuteInfo{
					Time:  out[0],
					Price: helpers.StrToFloat64(out[1]),
					Vol:   int64(helpers.StringToInt(out[2])),
				}
				minfo.Vol = minfo.Vol - sumVol
				sumVol += minfo.Vol // 总成交量
				sumPrice += minfo.Price * float64(minfo.Vol)
				if sumVol > 0 {
					minfo.Ave = helpers.Decimal(sumPrice/float64(sumVol), 2)
				}
				tmp.List = append(tmp.List, minfo)
			}
		}
	}
	return &tmp, nil
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
