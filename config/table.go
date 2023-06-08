package config

import (
	"goapi/pkg/config"
)

// 数据库连接参数
func init() {
	config.Add("table", config.StrMap{
		"kline_log": config.Env("KLINE_TABLE", "kline_log"),
	})
}
