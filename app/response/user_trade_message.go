package response

type UserTradeMessage struct {
	AppUserId            string `json:"app_user_id"`
	IsNew                string `json:"is_new"`
	BlackList            string `json:"black_list"`
	TradingCount         int    `json:"trading_count"`
	TradingAmounts       int    `json:"trading_amounts"`
	RechargeAmounts      string `json:"recharge_amounts"`
	RechargeCount        string `json:"recharge_count"`
	RechargeTodayAmounts string `json:"recharge_today_amounts"`
	RechargeTodayCount   string `json:"recharge_today_count"`
	WithdrawAmounts      string `json:"withdraw_amounts"`
	WithdrawCount        string `json:"withdraw_count"`
	WithdrawTodayAmounts string `json:"withdraw_today_amounts"`
	WithdrawTodayCount   string `json:"withdraw_today_count"`
	HisTradingCount      string `json:"his_trading_count"`
	HisTradingAmounts    string `json:"his_trading_amounts"`
}
