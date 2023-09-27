package test

import (
	"fmt"
	"goapi/pkg/helpers"
	"goapi/serve/binary-stock/client"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	Search([]string{"sz002519"})
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
