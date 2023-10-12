package utils

import (
	"strings"
)

func CompleteStockCode(stockCode string) string {
	stockCode = strings.ToLower(stockCode)
	// 去除可能存在的前缀
	stockCode = strings.TrimPrefix(stockCode, "sh")
	stockCode = strings.TrimPrefix(stockCode, "sz")

	// 判断应该添加哪个前缀
	if strings.HasPrefix(stockCode, "6") {
		return "sh" + stockCode
	} else if strings.HasPrefix(stockCode, "0") || strings.HasPrefix(stockCode, "3") {
		return "sz" + stockCode
	} else {
		// 返回默认的sz前缀
		return "sz" + stockCode
	}
}
