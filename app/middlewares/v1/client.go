package v1

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"goapi/app/response"
	"goapi/pkg/address"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"goapi/pkg/request"
)

// AddressLimit 地区限制
func AddressLimit() gin.HandlerFunc {
	/** 您所在的地区限制使用 */
	return func(c *gin.Context) {
		logger.Info(c.ClientIP())
		if c.FullPath() == "/app/v2/config/setKlineToken.json" {
			c.Next()
		} else {
			// 正式环境检测
			if config.GetBool("app.address_limit") && address.LimitAddress(c.ClientIP()) {
				echo.Error(c, "CountryLimitUse", "")
				c.Abort()
				return
			} else {
				c.Next()
			}
		}
	}
}

// Client 定义客户端中间件
func Client() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置语言，默认英文
		SetLang(c, "en")
		path := c.FullPath()
		switch path {
		// websocket
		case "/v1/api/kline/ws_index":
		case "/v1/api/kline/history":
			// 继续往下面执行
			c.Next()
			break
		default:
			CheckLogin(c)
			c.Next()
			break
		}
	}
}

// CheckLogin 检测登录
func CheckLogin(c *gin.Context) {
	var (
		strInfo string
		err     error
	)
	// 获取 "token"
	tokenString := request.GetParam(c, "token")
	if len(tokenString) <= 0 || tokenString == "<nil>" {
		logger.Error(errors.New("未检测到token"))
		echo.Error(c, "LoginInvalid", "")
		c.Abort()
		return
	}
	strInfo, err = redis.Client.Get(tokenString)
	if err != nil {
		echo.Error(c, "LoginInvalid", "")
		c.Abort()
		return
	}
	// json字符串数组,转换成切片
	var user response.ClientUserBeans
	multiErr := json.Unmarshal([]byte(strInfo), &user)
	if multiErr != nil {
		logger.Error(errors.New("转换出错"))
		logger.Error(multiErr)
		return
	}
	logger.Info(user)
	// 保存用户到 上下文
	c.Set("user", user)
	// 再次确认设置语言
	SetLang(c, "en")
	// 继续往下面执行
	c.Next()
}

// SetLang 设置语言
func SetLang(c *gin.Context, lang string) {
	language := c.Request.Header.Get("Language")
	if len(language) <= 0 {
		// 默认中文
		c.Request.Header.Set("Language", lang)
	}
	l, _ := c.Get("language")
	switch l {
	case 1:
		// 中文
		c.Request.Header.Set("Language", "zh")
		break
	case 2:
		// 英文
		c.Request.Header.Set("Language", "en")
		break
	case 3:
		// 日语
		c.Request.Header.Set("Language", "jp")
		break
	}
}
