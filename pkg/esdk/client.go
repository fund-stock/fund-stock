package esdk

import (
	"github.com/olivere/elastic/v7"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"log"
	"os"
)

var EsClient *elastic.Client

func Client() (*elastic.Client, error) {
	var err error
	// 创建client连接ES
	EsClient, err = elastic.NewClient(
		// 启用或禁用嗅探器（默认启用）。
		elastic.SetSniff(config.GetBool("elastic.sniff")),
		// elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(config.GetString("elastic.url")),
		//// 设置基于http base auth验证的账号和密码
		//elastic.SetBasicAuth("user", "secret"),
		//// 启用gzip压缩
		//elastic.SetGzip(config.GetBool("elastic.gzip")),
		//// 设置监控检查时间间隔
		//elastic.SetHealthcheckInterval(10*time.Second),
		//// 设置请求失败最大重试次数
		//elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ERROR：", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "SUCCESS：", log.LstdFlags)),
	)
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return EsClient, nil
}
