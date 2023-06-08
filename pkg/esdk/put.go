package esdk

import (
	"context"
	"goapi/pkg/logger"
	"strings"
)

func IndexPut(index, id string, body interface{}) bool {
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	// 使用client创建一个新的文档
	esClient, err := Client()
	if err != nil {
		logger.Error(err)
		return false
	}
	_, err = esClient.Index().
		Index(strings.ToLower(index)). // 设置索引名称
		//Type("doc").
		Id(id).          // 设置文档id
		BodyJson(body).  // 指定前面声明struct对象
		Refresh("true"). // 执行操作后刷新索引。
		Do(ctx)          // 执行请求，需要传入一个上下文对象
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}
