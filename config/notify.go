package config

import "goapi/pkg/config"

// email 参数
func init() {
	config.Add("notify", config.StrMap{
		"email": config.Env("NOTIFY_EMAIL", "543619552@qq.com"),
	})
}
