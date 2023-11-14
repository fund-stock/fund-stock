package v3

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 微信服务

type WechatController struct {
	BaseController
}

func (h *WechatController) ServeWechat(c *gin.Context) {
	wc := wechat.NewWechat()
	// 这里本地内存保存 access_token，也可选择 redis，memcache 或者自定cache
	redisCache := cache.NewRedis(context.Background(), &cache.RedisOpts{
		Host:        "127.0.0.1:6379",
		Database:    0,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60, // second
	})
	cfg := &offConfig.Config{
		AppID:          "wx23d857c691e1c681a",
		AppSecret:      "m1470bbbe0423835940e252a52b03bfb5d",
		Token:          "xmg",
		EncodingAESKey: "YaGm3bNwd3U13Jd0OJ9O4Pu3ZGCeSYLN5ZgSAnodmyo",
		Cache:          redisCache,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	// 传入request和responseWriter
	server := officialAccount.GetServer(c.Request, c.Writer)
	// 设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		// TODO
		// 回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
	// 处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		logger.Error(err)
		return
	}
	// 发送回复的消息
	err = server.Send()
	if err != nil {
		logger.Error(err)
		return
	}
}

func (h *WechatController) Synchronize(c *gin.Context) {
	url := "http://106.52.198.173:8000/contact/get_contacts"
	method := "POST"
	BelongWx := "xiaomg_zs"
	payload := strings.NewReader(`{"guid": "48b678e5-a377-3067-b651-26ac580c87df"}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
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

	var wxResp response.RespWx
	err = json.Unmarshal(body, &wxResp)
	if err != nil {
		logger.Error(err)
		return
	}
	if wxResp.Status != 1 {
		logger.Error(errors.New(wxResp.Msg))
		return
	}
	DB := models.GoWechatMgr(mysql.DB)
	for _, data := range wxResp.Data {
		option, err := DB.GetByOption(DB.WithBelongWx(BelongWx), DB.WithWxid(data.Wxid))
		if err != nil {
			logger.Error(err)
			continue
		}
		if option.ID > 0 {
			continue
		}
		DB.Debug().Create(&models.GoWechat{
			BelongWx:    BelongWx,
			Wxid:        data.Wxid,
			Account:     data.Account,
			Sex:         data.Sex,
			Avatar:      data.Avatar,
			City:        data.City,
			Country:     data.Country,
			LabelidList: data.LabelidList,
			Nickname:    data.Nickname,
			Province:    data.Province,
			Remark:      data.Remark,
			CreateTs:    time.Now().UnixMilli(),
			UpdateTs:    time.Now().UnixMilli(),
			DeleteTs:    time.Now().UnixMilli(),
		})
	}
}
