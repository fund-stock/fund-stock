package bark

import (
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
)

func icon() string {
	images := config.GetString("bark.logo")
	return images
}

func Notice(title, bodys string) {
	name := config.GetString("app.name")
	title = fmt.Sprintf("【%v】---- %v", name, title)
	go send(config.GetString("bark.key.mac"), title, bodys)    // 发往 mac
	go send(config.GetString("bark.key.iphone"), title, bodys) // 发往 iPhone
}

func send(appKey, title, bodys string) {
	url := config.GetString("bark.url") + appKey + fmt.Sprintf("/%v/%v?icon=%v", title, bodys, icon())
	res, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return
	}
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
	logger.Info(string(body))
}
