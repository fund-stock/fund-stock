package config

import "goapi/pkg/config"

// email 参数
func init() {
	config.Add("elastic", config.StrMap{
		"sniff": config.Env("ELASTIC_SNIFF", true),                    // 启用或禁用嗅探器（默认启用）。
		"url":   config.Env("ELASTIC_URL", "http://10.10.10.10:9200"), // elasticsearch 服务地址，多个服务地址使用逗号分隔
		"gzip":  config.Env("ELASTIC_GZIP", true),                     // 启用gzip压缩
	})
}
