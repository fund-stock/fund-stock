package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/config"
	"goapi/pkg/echo"
	"goapi/pkg/helpers"
	"goapi/pkg/output"
	"goapi/pkg/redis"
	"strings"
)

// k线图服务

type KlineController struct {
	BaseController
}

// HistoryHandler
// 历史行情 mysql 查询
func (h *KlineController) HistoryHandler(c *gin.Context) {
	// 监控K线数据时间段的数据波动，实时调整赔率
	output.Rjson(c, gin.H{
		"gzip":          1,
		"HistKlineLogs": 2,
		"totals":        3,
		"contractCode":  4,
	}, "success", 1)
}

// 获取 15 库的所有键值对

func (h *KlineController) GetRedisKeyHandler(c *gin.Context) {
	keys := c.Query("keys")
	arr, cursor, err := redis.Client.GetKeys(15, keys)
	if err != nil {
		echo.Error(c, "Failed", "")
	} else {
		result := cmap.New().Items()
		for _, item := range arr {
			value, _ := redis.Client.SelectDbGet(15, strings.ReplaceAll(item, config.GetString("app.name")+":", ""))
			result[item] = value
		}
		echo.Success(c, gin.H{"result": result, "cursor": cursor}, "")
	}
	return
}

// 获取 15 库的所有键值对

func (h *KlineController) SetKlineToken(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	exTime := helpers.StringToInt(c.Query("exTime"))
	if len(value) > 0 {
		ok, err := redis.Client.SelectDbAdd(15, fmt.Sprintf("kline:token:%s", key), value, exTime)
		if err != nil {
			echo.Success(c, gin.H{"result": ok}, err.Error())
			return
		}
		echo.Success(c, gin.H{"result": ok}, "")
	}
	echo.Error(c, "Failed", "")
	return
}
