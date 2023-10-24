package test

import (
	"fmt"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
	"goapi/pkg/helpers"
	"goapi/serve/binary-stock/client"
	"strings"
	"testing"
	"time"
)

const (
	CacheCode        = "shares_code_cache"
	CacheCodeTimeout = 2 * time.Hour
	_timeoutTicker   = 3
	_maxSpan         = 3 * (60 / _timeoutTicker) // 三分钟内涨跌幅
	_maxUpPercent    = 3                         // 快速涨幅 3%
	_maxDownPercent  = 4                         // 快速跌幅 4%
)

var _analy Analy

type priceInfo struct {
	now   int64
	price float64
}

type Analy struct {
	mp   map[string][]priceInfo
	span int //  跨度
	cap  int // 报警容量
}

func (a *Analy) Init() { //
	a.mp = make(map[string][]priceInfo)
	a.span = _maxSpan
	a.cap = a.span * 2
}

func (a *Analy) Add(info *client.SharesInfo, now int64) {
	a.mp[info.Code] = append(a.mp[info.Code], priceInfo{
		now:   now,
		price: info.Price,
	})

	if len(a.mp[info.Code]) > a.cap { // 去掉不用的
		a.mp[info.Code] = a.mp[info.Code][(len(a.mp[info.Code]) - a.span):]
	}
}

func TestSearch(t *testing.T) {
	_analy.Init()
	for {
		now := time.Now()
		day0 := tools.GetUtcDay0(now)
		offset := now.Unix() - day0
		//if offset >= (15*60*60 + 1*60) { // 3点之后是空闲时间
		//	break
		//}
		ticker := time.NewTicker(3 * time.Second)
		<-ticker.C
		if offset > (11*60+30)*60 && offset < (13*60*60+60) { // 盘中休息
			mylog.Infof("timer 盘中休息.....:%v", tools.GetTimeStr(time.Now()))
		} else {
			mylog.Infof("timer 盘中执行.....:%v", tools.GetTimeStr(time.Now()))
			outs := client.Searches([]string{"sz002261"})
			now := time.Now().Unix()
			for _, v := range outs { // 保存

				_analy.Add(v, now) // 添加
			}
			fmt.Println(_analy.mp)
			fmt.Println(_analy.mp)
		}
	}
}

func Search(codes []string) string {
	parm := strings.Join(codes, ",s_")
	url := "http://qt.gtimg.cn/q=s_" + parm
	out := client.SendGet(url, "")
	out = helpers.ConvertToString(out, "gbk", "utf8")
	if len(out) == 0 {
		return ""
	}
	out = strings.Replace(out, "\n", "", -1)
	// todo 分析结果
	list := strings.Split(out, ";")
	fmt.Println(out)
	fmt.Println(list)
	fmt.Println(list)
	return ""
}
