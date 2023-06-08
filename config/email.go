package config

import "goapi/pkg/config"

// email 参数
func init() {
	config.Add("email", config.StrMap{
		"type": config.Env("EMAIL_TYPE", ""),
		// qq邮箱
		"qq": map[string]interface{}{
			"addr":     config.Env("QQ_EMAIL_ADDR", ""),
			"host":     config.Env("QQ_EMAIL_HOST", ""),
			"user":     config.Env("QQ_EMAIL_USER", ""),
			"password": config.Env("QQ_EMAIL_PASSWORD", ""),
		},
		// 谷歌邮箱
		"google": map[string]interface{}{
			"addr":     config.Env("GOOGLE_EMAIL_ADDR", ""),
			"host":     config.Env("GOOGLE_EMAIL_HOST", ""),
			"user":     config.Env("GOOGLE_EMAIL_USER", ""),
			"password": config.Env("GOOGLE_EMAIL_PASSWORD", ""),
		},
	})
}
