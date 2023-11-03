package v3

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"goapi/pkg/logger"
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
