package client

import (
	"github.com/xxjwxc/public/tools"
	"goapi/pkg/helpers"
	"strings"
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
