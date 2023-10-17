package task_route

import (
	"github.com/robfig/cron/v3"
	"goapi/serve/binary-timer/task_plan/analyze"
	"goapi/serve/binary-timer/task_plan/collect"
)

// RegisterTaskRoutes 注册定时任务路由

func RegisterTaskRoutes(crontab *cron.Cron) {
	// 分析市场行情 每隔1分钟
	_, _ = crontab.AddFunc("*/60 * * * * *", analyze.Market)
	// 采集数据，每天下午 15:05 分执行一次
	_, _ = crontab.AddFunc("0 5 15 * * *", collect.StartStock)
}
