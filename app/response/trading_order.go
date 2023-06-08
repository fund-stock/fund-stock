package response

import "time"

type TradingOrder struct {
	TradeOrderID        string    `json:"trade_order_id" gorm:"column:trade_order_id"`               // 交易表主键自增ID
	AppUserID           string    `json:"app_user_id" gorm:"column:app_user_id"`                     // 客户ID
	Amounts             float64   `json:"amounts" gorm:"column:amounts"`                             // 交易金额
	PlanGainAmounts     float64   `json:"plan_gain_amounts" gorm:"column:plan_gain_amounts"`         // 预估收益金额
	RealGainAmounts     float64   `json:"real_gain_amounts" gorm:"column:real_gain_amounts"`         // 真实收益金额
	ItemCode            string    `json:"item_code" gorm:"column:item_code"`                         // 关联项目编码
	AssetsID            int64     `json:"assets_id" gorm:"column:assets_id"`                         // 资产ID
	AssetsName          string    `json:"assets_name" gorm:"column:assets_name"`                     // 资产名称
	AssetsLogoUrl       string    `json:"assets_logo_url" gorm:"column:assets_logo_url"`             // 资产logo的URL
	BeginAssetsPrice    float64   `json:"begin_assets_price" gorm:"column:begin_assets_price"`       // 起始资产指数值
	EndAssetsPrice      float64   `json:"end_assets_price" gorm:"column:end_assets_price"`           // 截止资产指数值
	TradeBeginDate      time.Time `json:"trade_begin_date" gorm:"column:trade_begin_date"`           // 交易起始日期
	TradeEndDate        time.Time `json:"trade_end_date" gorm:"column:trade_end_date"`               // 交易截止日期
	QuizResult          int       `json:"quiz_result" gorm:"column:quiz_result"`                     // 竞猜结果 1对 0错 -1未检出结果 2结果检出中 3平
	IsValid             int       `json:"is_valid" gorm:"column:is_valid"`                           // 是否有效 1是 0否
	CreateDate          time.Time `json:"create_date" gorm:"column:create_date"`                     // 添加日期
	UpdateDate          time.Time `json:"update_date" gorm:"column:update_date"`                     // 更新日期
	CreateTs            int64     `json:"create_ts" gorm:"column:create_ts"`                         // 创建时间戳
	UpdateTs            int64     `json:"update_ts" gorm:"column:update_ts"`                         // 更新时间戳
	AmountType          int       `json:"amount_type" gorm:"column:amount_type"`                     // 交易金额类型 1余额 2bonus 3虚拟货币
	UpDownType          string    `json:"up_down_type" gorm:"column:up_down_type"`                   // 买涨买跌类型
	RealCostAmounts     float64   `json:"real_cost_amounts" gorm:"column:real_cost_amounts"`         // 实际消费金额
	KlineTime           int64     `json:"kline_time" gorm:"column:kline_time"`                       // k线交易时间点
	TradeBeginTimestamp int64     `json:"trade_begin_timestamp" gorm:"column:trade_begin_timestamp"` // 交易起始日期时间戳
	TradeEndTimestamp   int64     `json:"trade_end_timestamp" gorm:"column:trade_end_timestamp"`     // 交易截止日期时间戳
	UserGainAmounts     float64   `json:"user_gain_amounts" gorm:"column:user_gain_amounts"`         // 用户真实亏损，大于0为客户盈利反之则亏
	LockingKey          string    `json:"locking_key" gorm:"column:locking_key"`
	AssetsCode          string    `json:"assets_code" gorm:"column:assets_code"`
	ReturnRate          float64   `json:"return_rate" gorm:"column:return_rate"`                   // 收益率
	CurrencyType        string    `json:"currency_type" gorm:"column:currency_type"`               // 货币类型 IDR USD CNY
	CurrencyRate        float64   `json:"currency_rate" gorm:"column:currency_rate"`               // 货币费率对标美元
	Balance             float64   `json:"balance" gorm:"column:balance"`                           // 交易时余额
	Nation              string    `json:"nation" gorm:"column:nation"`                             // 绑定国家
	EqualReturnAmounts  float64   `json:"equal_return_amounts" gorm:"column:equal_return_amounts"` // 平返还金额
}
