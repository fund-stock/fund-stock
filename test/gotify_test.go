package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestGotify(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			form, err := http.PostForm("https://gotify.ethanyep.cn/message?token=AWxSCZ.uG4nV4zV",
				url.Values{
					"title":   {"恭喜您获利成功"},
					"message": {"恭喜您获取成功，该次上涨成功，请您留意当前的股票走势，看好买点时间，当前已经达到最好的入场时间，估计比较低错过了就要等待下一波了，当前赢利点达到您的预期，是否前往查看并且卖出适当的仓位，减少持仓"},
				})
			if err != nil {
				fmt.Println(err)
				return
			}
			defer form.Body.Close() // 关闭响应体

			body, err := ioutil.ReadAll(form.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Response Body:", string(body))
		}(i)
		time.Sleep(time.Millisecond * 100)
	}
	time.Sleep(time.Second) // 等待goroutines完成
}
