package client

import (
	"github.com/xxjwxc/public/mylog"
	"io/ioutil"
	"net/http"
)

// SendGet 发送get 请求 返回对象
func SendGet(url, params string) string {
	//解析这个 URL 并确保解析没有出错。
	var urls = url
	if len(params) > 0 {
		urls += "?" + params
	}
	resp, err := http.Get(urls)
	if err != nil {
		mylog.Error(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		mylog.Error(err)
		return ""
	}

	return string(body)
}
