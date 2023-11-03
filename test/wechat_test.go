package test

import (
	"context"
	"fmt"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"goapi/bootstrap"
	"goapi/config"
	conf "goapi/pkg/config"
	"goapi/pkg/logger"
	"testing"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	// 初始化配置信息
	config.Initialize()
	// 定义日志目录
	logger.Init("binary-wechat-test")
}

func TestWechat(t *testing.T) {
	// 初始化 SQL
	logger.Info("初始化 SQL")
	bootstrap.SetupDB()
	// 初始化 Redis
	logger.Info("初始化 Redis")
	db := conf.GetInt("redis.db")
	bootstrap.SetupRedis(db)
	defer bootstrap.RedisClose()
	wechatServe()
	fmt.Println(1212)
}

func wechatServe() {
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
	data := make(map[string]*message.TemplateDataItem)
	// 项目名称
	// {{thing2.DATA}}
	data["thing2"] = &message.TemplateDataItem{Value: "测试"}
	// 账目金额
	// {{amount4.DATA}}
	data["amount4"] = &message.TemplateDataItem{Value: "1000000.00"}
	// 操作人
	// {{thing9.DATA}}
	data["thing9"] = &message.TemplateDataItem{Value: "测试"}
	// 账目日期
	// {{time8.DATA}}
	data["time8"] = &message.TemplateDataItem{Value: time.Now().Format(time.DateTime)}
	// 账目类型
	// {{thing3.DATA}}
	data["thing3"] = &message.TemplateDataItem{Value: "出账"}
	send, err := officialAccount.GetTemplate().Send(&message.TemplateMessage{
		ToUser: "oWtmj6kMuFODx0bqPuwKkyZE-MF8",
		//ToUser:     "oWtmj6j6GQ3p12syNCT14V48mY-4",
		TemplateID: "q6cozXUE5t842msi1QQan89RBVfUA-Nd6f1ax4M-pBo",
		Data:       data,
		URL:        "https://baidu.com",
		Color:      "",
	})
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(send)
}
