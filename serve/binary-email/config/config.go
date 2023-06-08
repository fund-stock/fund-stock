package config

import (
	"errors"
	"goapi/app/models"
	"goapi/app/response"
	"goapi/pkg/mysql"
)

// 读取一条配置信息

func GetOne(email string) (error, response.SystemEmailRoutingNode) {
	var result response.SystemEmailRoutingNode
	mysql.DB.Debug().Model(models.SystemEmailRoutingNode{}).Where(map[string]interface{}{
		"valid": 1,
	}).Order("rate_fail asc").Limit(1).Find(&result)
	if result.Id == 0 {
		return errors.New("，数据库未找到相关渠道配置"), result
	}
	return nil, result
}
