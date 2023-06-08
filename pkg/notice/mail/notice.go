package NoticeMail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	conf "goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/notice/telegram"
	"goapi/pkg/redis"
	"net/url"
	"runtime/debug"
)

func getNoticeKey(address, t string) string {
	return fmt.Sprintf("notice:mail:%v:%v", t, address)
}

// 检测通知次数，超过次数不进行通知
func check(address, t string) bool {
	var pass = true
	key := getNoticeKey(address, t)
	switch t {
	case "msgHeartbeatAnomaly": // 心跳通知
		{
			if !helpers.InArray(address, []string{AdminMail, "yezhiming@yuctime.com"}) {
				return false
			}
			ok, _ := redis.Client.Get(key)
			// 初始化
			if len(ok) > 0 {
				return false
			}
			// 15分钟通知一次
			_, _ = redis.Client.Add(key, "ok", 15*60)
		}
		break
	case "msgSocketNetworkAnomaly": // 检测累计通知次数：网络异常：一分钟内超过15次进行通知
		pass = checkGrandTotal(address, t, 0, 1*60, 15)
		break
	case "msgAssetSleep": // 开市休市提醒
		pass = true
	case "msgCheckSwitchSource": // 自动切换数据源
	case "msgNoSwitchSource": // 没有可用数据源切换
		if !helpers.InArray(address, []string{AdminMail, "yezhiming@yuctime.com"}) {
			return false
		}
		pass = true
		break
	default:
		break
	}
	return pass
}

// 检查累计次数
func checkGrandTotal(address, t string, InitNum, GrandTotalExTime, GrandTotal int) bool {
	key := getNoticeKey("grandTotal:"+address, t) // key=======>notice:mail:msgSocketNetworkAnomaly:grandTotal:mail@54zm.com
	value, _ := redis.Client.Get(key)
	// 初始化
	if len(value) <= 0 {
		// 四分钟累计
		_, err := redis.Client.Add(key, InitNum, GrandTotalExTime)
		if err != nil {
			return false
		}
	} else {
		num := helpers.StringToInt(value)
		// 最大通知次数
		if num < GrandTotal {
			return false
		}
		if redis.Client.Incr(key) > int64(GrandTotal) {
			return true
		} else {
			return false
		}
	}
	return false
}

// 重启 socket 通知
func msgRebootSocket(service, symbol string, num interface{}, lastRuntime int64) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf(`
断线重连==
【环境：%v】
【服务：%v】
【币种：%v】
【上次socket运行：%v 秒 】
`, conf.GetString("app.env"), service, symbol, lastRuntime/1000)
	info["content"] = fmt.Sprintf("这是第【%v】次socket断线重连", num)
	return info
}

// 赔率变动
func msgRateChange(service, tips string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】，服务【%v】赔率变动", conf.GetString("app.env"), service)
	info["content"] = tips
	return info
}

// 数据异常
func msgDataAbnormal(c *gin.Context, tips string) map[string]interface{} {
	info := cmap.New().Items()
	info["title"] = fmt.Sprintf("环境【%v】==K线图历史数据异常通知，当前查询历史数据小于150条：", conf.GetString("app.env"))
	info["content"] = fmt.Sprintf(
		`请求 reqId 为%v<br>
						用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
						请求头信息：%v<br>
						查询范围：%v<br>
						请求的url为：%v<br>`,
		logger.RequestId,
		c.ClientIP(), c.ClientIP(),
		c.Request.Header,
		tips,
		c.Request.URL,
	)
	return info
}

// 堆栈异常信息
func msgStack(c *gin.Context, r interface{}, description, errInfo string) map[string]interface{} {
	Info := cmap.New().Items()
	Info["title"] = description + "当前环境：" + conf.GetString("app.env") + fmt.Sprintf("Recover: %v", r)
	if c != nil {
		method := c.Request.Method
		Info["content"] = fmt.Sprintf(`
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
							堆栈信息 【%v】<br>
						`, c.ClientIP(), c.ClientIP(), c.Request.URL, method, r, errInfo)
	} else {
		Info["content"] = errInfo
	}
	return Info
}

// 网络异常
func msgNetworkAnomaly(title interface{}, c *gin.Context, r, description, errInfo interface{}) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	method := c.Request.Method
	brokenInfo["title"] = fmt.Sprintf("环境【%v】 %v", conf.GetString("app.env"), title)
	brokenInfo["content"] =
		fmt.Sprintf(
			`
							【%v】<br>
							用户IP【<a href="https://www.ip138.com/iplookup.asp?ip=%v&action=2">%v</a>】，<br>
							RequestId 【%v】<br>
							请求地址 【%v】<br>
							请求方法 【%v】<br>
							错误信息 【%v】<br>
					`, description, c.ClientIP(), c.ClientIP(), logger.RequestId, c.Request.URL, method, r)
	if errInfo != nil {
		brokenInfo["content"] = fmt.Sprintf(`%v具体堆栈信息： <br>%v`, brokenInfo["content"], errInfo)
	}
	return brokenInfo
}

// Socket网络异常
func msgSocketNetworkAnomaly(name string, proxy *url.URL) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf(`
环境【%v】
名称【%v】
socket采集连接异常，请查看详细信息
`, conf.GetString("app.env"), name)
	if proxy != nil {
		brokenInfo["content"] = fmt.Sprintf(`当前使用的代理ip【%v】链接失败`, proxy.String())
	} else {
		brokenInfo["content"] = fmt.Sprintf(`正常链接失败，未使用代理`)
	}
	return brokenInfo
}

// Socket关闭
func msgSocketClose(name, description string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf(`
环境【%v】
服务【%v】
socket已关闭，请查看详细信息
`, conf.GetString("app.env"), name)
	brokenInfo["content"] = description
	return brokenInfo
}

// 服务堆栈异常
func msgServiceAnomaly(name string, r interface{}) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】【%v】服务异常，请查看详细信息", conf.GetString("app.env"), name)
	brokenInfo["content"] = fmt.Sprintf(`
【%v】服务异常【%v】
<br>详细堆栈错误信息 :
<br>%v
<br>`,
		name, r, string(debug.Stack()))
	return brokenInfo
}

// 心跳检测，服务健康通知
func msgHeartbeatAnomaly(name, msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】【%v】服务异常，请查看详细信息", conf.GetString("app.env"), name)
	brokenInfo["content"] = msg
	return brokenInfo
}

// 心跳检测，服务健康通知
func msgAssetSleep(AssetsCode, msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】资产(%v)变动，请查看详细信息", conf.GetString("app.env"), AssetsCode)
	brokenInfo["content"] = msg
	return brokenInfo
}

// 检测切换数据源
func msgCheckSwitchSource(msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】自动切换数据源", conf.GetString("app.env"))
	brokenInfo["content"] = msg
	return brokenInfo
}

// 没有可以切换的数据源
func msgNoSwitchSource(msg string) map[string]interface{} {
	brokenInfo := cmap.New().Items()
	brokenInfo["title"] = fmt.Sprintf("环境【%v】无数据源", conf.GetString("app.env"))
	brokenInfo["content"] = msg
	return brokenInfo
}

// 发送通知

func Notice(t string, data interface{}) {
	var Info map[string]interface{}
	switch t {
	case "msgRebootSocket": // 重启 socket 通知
		arr := data.(map[string]interface{})
		Info = msgRebootSocket(arr["service"].(string), arr["symbol"].(string), arr["num"].(int64), arr["lastRuntime"].(int64))
		break
	case "msgRateChange": // 赔率变动通知
		arr := data.(map[string]interface{})
		Info = msgRateChange(arr["service"].(string), arr["tips"].(string))
		break
	case "msgDataAbnormal":
		arr := data.(map[string]interface{})
		Info = msgDataAbnormal(arr["c"].(*gin.Context), arr["tips"].(string))
		break
	case "msgStack":
		arr := data.(map[string]interface{})
		Info = msgStack(arr["c"].(*gin.Context), arr["r"], arr["description"].(string), arr["errInfo"].(string))
		break
	case "msgNetworkAnomaly":
		arr := data.(map[string]interface{})
		Info = msgNetworkAnomaly(arr["title"], arr["c"].(*gin.Context), arr["r"], arr["description"], arr["errInfo"])
		break
	case "msgSocketNetworkAnomaly":
		arr := data.(map[string]interface{})
		Info = msgSocketNetworkAnomaly(arr["name"].(string), arr["proxy"].(*url.URL))
		break
	case "msgSocketClose":
		arr := data.(map[string]interface{})
		Info = msgSocketClose(arr["name"].(string), arr["description"].(string))
		break
	case "msgServiceAnomaly":
		arr := data.(map[string]interface{})
		Info = msgServiceAnomaly(arr["name"].(string), arr["r"])
		break
	case "msgHeartbeatAnomaly":
		arr := data.(map[string]interface{})
		Info = msgHeartbeatAnomaly(arr["name"].(string), arr["msg"].(string))
		break
	case "msgAssetSleep":
		arr := data.(map[string]interface{})
		Info = msgAssetSleep(arr["assets_code"].(string), arr["msg"].(string))
		break
	case "msgCheckSwitchSource":
		arr := data.(map[string]interface{})
		Info = msgCheckSwitchSource(arr["msg"].(string))
		break
	case "msgNoSwitchSource":
		arr := data.(map[string]interface{})
		Info = msgNoSwitchSource(arr["msg"].(string))
		break
	default:
		return
	}
	// 时差不进行通知，避免通知泛滥d
	if helpers.InArray(t, []string{"msgSocketClose", "msgRebootSocket"}) {
		return
	}
	// telegram 通知
	go telegram.Notice(Info)
	//if !helpers.InArray(t, []string{"msgHeartbeatAnomaly", "msgAssetSleep", "msgCheckSwitchSource", "msgNoSwitchSource"}) { // 这四个才进行邮件通知
	//	return
	//}
	//// 邮件通知
	//for _, address := range conf.GetStringSlice("notify.email") {
	//	if (len(address) > 0) && check(address, t) {
	//		go func(address string) {
	//			// 随机十秒以内进行休息，然后发送，避免发送过于频繁发送失败
	//			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	//			err := email.SendEmail(Info, address)
	//			if err != nil {
	//				logger.Error(err)
	//			}
	//		}(address)
	//	}
	//}
}
