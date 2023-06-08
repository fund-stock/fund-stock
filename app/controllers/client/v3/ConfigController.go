package v3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/redis"
	"golang.org/x/text/unicode/norm"
	"unicode/utf8"
)

// k线图服务

type ConfigController struct {
	BaseController
}

// 获取 15 库的所有键值对

func (h *ConfigController) SetKlineToken(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	exTime := helpers.StringToInt(c.Query("exTime"))
	if len(value) > 0 {
		ok, err := redis.Client.SelectDbAdd(15, fmt.Sprintf("kline:token:%s", key), value, exTime)
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
		echo.Success(c, gin.H{"result": ok}, "")
		return
	}
	echo.Error(c, "Failed", "")
}

func isAllChinese(s string) bool {
	// 将字符串规范化为NFC格式
	s = norm.NFC.String(s)

	// 遍历字符串的每个字符
	for len(s) > 0 {
		// 读取第一个字符
		r, size := utf8.DecodeRuneInString(s)
		// 如果字符不是中文字符，则返回false
		if r < 0x4e00 || r > 0x9fff {
			return false
		}
		// 去除已经读取的字符
		s = s[size:]
	}
	// 所有字符都是中文字符，则返回true
	return true
}

// 检查多语言是否包含中文，然后发送警告

func (h *ConfigController) CheckLanguage(c *gin.Context) {
	echo.Success(c, nil, "ok")
}
