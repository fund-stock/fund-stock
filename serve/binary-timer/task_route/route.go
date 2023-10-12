package task_route

import (
	"github.com/robfig/cron/v3"
	"goapi/serve/binary-timer/task_plan/analyze"
)

// RegisterTaskRoutes 注册定时任务路由

func RegisterTaskRoutes(crontab *cron.Cron) {
	// 分析市场行情 每隔1分钟
	_, _ = crontab.AddFunc("*/2 * * * * ?", analyze.Market)
}
