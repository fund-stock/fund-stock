package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/serve/binary-stock/response/k780"
	"goapi/serve/binary-stock/response/qtimg"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 获取分时原始数据

func GetMinute(code string) (*qtimg.Minute, error) {
	// 开始获取
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/minute/query?code=%v", code), "")
	out = strings.Replace(out, code, "info", -1)
	var info qtimg.MinuteWWW
	err := json.Unmarshal([]byte(out), &info)
	if err != nil {
		logger.Error(err)
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
			if len(out) >= 3 {
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
	} else {
		return nil, errors.New(fmt.Sprintf("%v", info.Code))
	}
	return &tmp, nil
}

// 获取历史数据

func GetHistoryData(code string) qtimg.Resp {
	var Resp qtimg.Resp
	out := SendGet(fmt.Sprintf("https://web.ifzq.gtimg.cn/appstock/app/fqkline/get?param=%v,day,,,320,qfq", code), "")
	out = strings.ReplaceAll(out, code, "info")
	out = strings.ReplaceAll(out, "qfqday", "day")
	err := json.Unmarshal([]byte(out), &Resp)
	if err != nil {
		logger.Error(err)
		return Resp
	}
	return Resp
}

func StockList() {
	urls := "https://sapi.k780.com/?app=finance.stock_list&category=hs&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"118\", \"Google Chrome\";v=\"118\", \"Not=A?Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("Host", "sapi.k780.com")
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
	var Resp k780.List
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		logger.Error(err)
		return
	}
	if Resp.Success != "1" {
		logger.Info(Resp.Success)
		return
	}
	DB := models.GoStockMgr(mysql.DB)
	for _, item := range Resp.Result.Lists {
		if strings.Contains(item.Symbol, "bj") {
			// 北交所暂时不需要
			continue
		}
		option, err := DB.GetByOption(DB.WithStatus(1), DB.WithCode(item.Symbol))
		if err != nil {
			logger.Error(err)
			return
		}
		if option.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoStock{
			Code:     item.Symbol,
			Name:     item.Sname,
			Amount:   100.00,
			Nav:      0,
			Status:   1,
			CreateAt: time.Now().UnixMilli(),
			UpdateAt: time.Now().UnixMilli(),
		})
	}
	fmt.Println("Success")

}

// 创业板和中小板、主板就股票代码上有什么区别
// 创业板股票代码区间为：300000-399999，一般按发行顺序依次编码，如：300001、300002等。
// 沪市代码600、601、603开头，
// 深市股票代码，000、001、002、300开头。
// 创业板是深圳交易所的一个子版块，是300开头的股票。目前共计356只创业板个股
//
// 作者：生而为人2019
// 链接：https://xueqiu.com/7851772761/134168423
// 来源：雪球

func ClassifyBoard(stockCode string) string {
	stockCode = stockCode[2:]
	code := helpers.StringToInt64(stockCode)
	if code >= 300000 && code <= 399999 {
		return "创业板"
	}
	return "主板"
}
