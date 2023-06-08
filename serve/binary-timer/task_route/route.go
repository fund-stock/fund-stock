package task_route

import (
	"github.com/robfig/cron/v3"
	"goapi/serve/binary-timer/task_plan"
)

// RegisterTaskRoutes 注册定时任务路由

func RegisterTaskRoutes(crontab *cron.Cron) {
	// 5、同步分时间段统计注册、充值、交易 定时任务 每隔1分钟
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.ClientDemoTradeResult)
	// 7、每分钟执行一次，自动刷新充值超时状态定时任务
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.WebStatisticsYesterday)
	// 8、每隔2秒自动刷新真实交易结果定时任务
	_, _ = crontab.AddFunc("*/5 * * * * ?", task_plan.WebCheckNewUserData)

}
