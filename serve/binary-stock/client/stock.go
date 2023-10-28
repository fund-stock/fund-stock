package client

import (
	"fmt"
	"github.com/xxjwxc/public/tools"
	"goapi/pkg/helpers"
	"strings"
	"time"
)

var _extMp = map[string]string{
	"200": "us", // 美股
	"100": "hk", // 港股
	"51":  "sz", // 深圳
	"1":   "sh", // 上海
}

// SharesInfo 股票信息
type SharesInfo struct {
	Code       string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`             // 股票代码
	SimpleCode string  `protobuf:"bytes,2,opt,name=simpleCode,proto3" json:"simpleCode"` // 股票代码简写
	Ext        string  `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext"`               // 后缀
	Name       string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`             // 股票名字
	Price      float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price"`         // 当前价格
	Percent    float64 `protobuf:"fixed64,6,opt,name=percent,proto3" json:"percent"`     // 百分比
}

// Searches 确定的搜索
func Searches(codes []string) (resp []*SharesInfo) {
	if len(codes) == 0 {
		return nil
	}
	query := strings.Join(codes, ",s_")
	url := "https://qt.gtimg.cn/q=s_" + query
	out := SendGet(url, "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	if len(out) == 0 {
		return nil
	}
	out = strings.Replace(out, "\n", "", -1)

	// todo 分析结果
	list := strings.Split(out, ";")
	for _, v := range list {
		if len(v) < 4 {
			continue
		}
		list1 := strings.Split(v, "=")
		if len(list1) == 2 {
			if strings.HasPrefix(list1[0], "v_s_") {
				tmp := &SharesInfo{
					Code: list1[0][4:],
				}
				list1[1] = strings.Trim(list1[1], "\"")
				list2 := strings.Split(list1[1], "~")
				if len(list2) > 9 {
					tmp.SimpleCode = list2[2]
					tmp.Ext = _extMp[list2[0]]
					tmp.Name = strings.Replace(list2[1], " ", "", -1)
					tmp.Price = helpers.StrToFloat64(list2[3])
					tmp.Percent = helpers.StrToFloat64(list2[5])
				}
				resp = append(resp, tmp)
			}
		}
	}
	return resp
}

// SharesInfoDetails 股票详细信息
type SharesInfoDetails struct {
	Code         string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`                    // 股票代码
	SimpleCode   string  `protobuf:"bytes,2,opt,name=simpleCode,proto3" json:"simpleCode"`        // 股票代码简写
	Ext          string  `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext"`                      // 后缀
	Name         string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                    // 股票名字
	Price        float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price"`                // 当前价格
	Percent      float64 `protobuf:"fixed64,6,opt,name=percent,proto3" json:"percent"`            // 百分比
	Color        string  `protobuf:"bytes,7,opt,name=color,proto3" json:"color"`                  // 颜色
	Img          string  `protobuf:"bytes,8,opt,name=img,proto3" json:"img"`                      // 图片地址
	Volume       float64 `protobuf:"fixed64,9,opt,name=volume,proto3" json:"volume"`              // 成交量（手）
	Turnover     float64 `protobuf:"fixed64,10,opt,name=turnover,proto3" json:"turnover"`         // 成交额（万）
	TurnoverRate float64 `protobuf:"fixed64,11,opt,name=turnoverRate,proto3" json:"turnoverRate"` // 换手率
	Pe           float64 `protobuf:"fixed64,12,opt,name=pe,proto3" json:"pe"`                     // 市盈率
	Pb           float64 `protobuf:"fixed64,13,opt,name=pb,proto3" json:"pb"`                     // 市净率
	Total        float64 `protobuf:"fixed64,15,opt,name=total,proto3" json:"total"`               // 流通市值
	Cap          float64 `protobuf:"fixed64,16,opt,name=cap,proto3" json:"cap"`                   // 总市值
	Zljlr        float64 `protobuf:"fixed64,17,opt,name=zljlr,proto3" json:"zljlr"`               // 主力资金净流入
	OpenPrice    float64 `protobuf:"fixed64,18,opt,name=openPrice,proto3" json:"openPrice"`       // 开盘价
	ClosePrice   float64 `protobuf:"fixed64,19,opt,name=closePrice,proto3" json:"closePrice"`     // 收盘价
	Macd         float64 `protobuf:"fixed64,20,opt,name=macd,proto3" json:"macd"`                 // macd
	Dif          float64 `protobuf:"fixed64,21,opt,name=dif,proto3" json:"dif"`                   // 快线
	Dea          float64 `protobuf:"fixed64,22,opt,name=dea,proto3" json:"dea"`                   // 慢线
	Max          float64 `protobuf:"fixed64,23,opt,name=max,proto3" json:"max"`                   // 当日最高点
	Min          float64 `protobuf:"fixed64,24,opt,name=min,proto3" json:"min"`                   // 当日最低
}

func SplitCode(code string) (ext, simplecode string) {
	if len(code) < 2 {
		return
	}
	ext = code[:2]
	simplecode = code[2:]
	return
}

type MacdInfo struct {
	Day0 time.Time
	Dif  float64
	Dea  float64
	Macd float64
}

type Macd struct {
	DayStr string
	Dif    float64
	Dea    float64
	Macd   float64
}

func GetMacd(code string) ([]Macd, error) {
	//kdata, err := GetDayliy(code)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

// GetDayMACD 当日macd
func GetDayMACD(code string) (resp *MacdInfo) {
	list, err := GetMacd(code)
	if err == nil {
		index := len(list) - 1
		if index >= 0 {
			return &MacdInfo{
				Day0: tools.StrToTime(list[index].DayStr, "2006-01-02", time.Local),
				Macd: list[index].Macd,
				Dif:  list[index].Dif,
				Dea:  list[index].Dea,
			}
		}
	}

	return nil
}

// SearchDetails 确定的搜索(全量搜索)
func SearchDetails(codes []string) (resp []*SharesInfoDetails) {
	if len(codes) == 0 {
		return nil
	}
	parm := strings.Join(codes, ",")
	url := "http://qt.gtimg.cn/q=" + parm
	out := SendGet(url, "")
	out = tools.ConvertToString(out, "gbk", "utf8")
	if len(out) == 0 {
		return nil
	}
	out = strings.Replace(out, "\n", "", -1)

	// todo 分析结果
	list := strings.Split(out, ";")
	for _, v := range list {
		if len(v) < 4 {
			continue
		}
		list1 := strings.Split(v, "=")
		if len(list1) == 2 {
			if strings.HasPrefix(list1[0], "v_") {
				tmp := &SharesInfoDetails{
					Code: list1[0][2:],
				}
				list1[1] = strings.Trim(list1[1], "\"")
				list2 := strings.Split(list1[1], "~")
				if len(list2) > 45 {
					tmp.Ext = _extMp[list2[0]]
					tmp.Name = strings.Replace(list2[1], " ", "", -1)
					tmp.SimpleCode = list2[2]
					tmp.Price = helpers.StrToFloat64(list2[3])
					tmp.OpenPrice = helpers.StrToFloat64(list2[5])
					tmp.ClosePrice = tmp.Price
					tmp.Percent = helpers.StrToFloat64(list2[32])
					tmp.Volume = helpers.StrToFloat64(list2[36])
					tmp.Turnover = helpers.StrToFloat64(list2[37])
					tmp.TurnoverRate = helpers.StrToFloat64(list2[38])
					tmp.Pe = helpers.StrToFloat64(list2[39])
					tmp.Max = helpers.StrToFloat64(list2[41])
					tmp.Min = helpers.StrToFloat64(list2[42])
					tmp.Pb = helpers.StrToFloat64(list2[46])
					tmp.Total = helpers.StrToFloat64(list2[44])
					tmp.Cap = helpers.StrToFloat64(list2[45])
				}
				// 主力资金
				_tmp := getDayZLJLR(SplitCode(tmp.Code))
				if _tmp != nil {
					tmp.Zljlr = _tmp.Price
				}
				// -----
				// macd
				_tmp1 := GetDayMACD(tmp.Code)
				if _tmp1 != nil {
					tmp.Macd = _tmp1.Macd
					tmp.Dif = _tmp1.Dif
					tmp.Dea = _tmp1.Dea
				}
				//
				resp = append(resp, tmp)
			}
		}
	}
	return resp
}

type HyInfo struct {
	Day0  time.Time
	Price float64
}

type HyResp struct {
	Rc   int        `json:"rc"`
	Data HyRespData `json:"data"`
}

type HyRespData struct {
	Total int    `json:"total"`
	Diffs []Diff `json:"diff"`
}

type Diff struct {
	F12 string  `json:"f12"`
	F13 int     `json:"f13"`
	F14 string  `json:"f14"`
	F62 float64 `json:"f62"`
}

// getDayZLJLR 当日主力资金净流入
func getDayZLJLR(ext, simplecode string) (resp *HyInfo) {
	switch ext {
	case "sh":
		ext = "1"
	case "sz":
		ext = "0"
	}
	url := fmt.Sprintf("https://push2.eastmoney.com/api/qt/ulist.np/get?secids=%v.%v&fields=f12,f13,f14,f62", ext, simplecode)
	out := SendGet(url, "")
	if len(out) == 0 {
		return nil
	}

	var tmp HyResp
	tools.JSONEncode(out, &tmp)
	if tmp.Rc == 0 && len(tmp.Data.Diffs) > 0 {
		return &HyInfo{
			Day0:  time.Now(),
			Price: tmp.Data.Diffs[0].F62 * 0.0001,
		}
	}

	return
}
