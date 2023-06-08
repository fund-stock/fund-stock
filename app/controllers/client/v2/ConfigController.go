package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/redis"
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
