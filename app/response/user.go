package response

type AppUser struct {
	ItemCode          string  `json:"item_code"`                  //item_code
	AppUserId         string  `json:"app_user_id"`                //app用户表
	Balance           float64 `json:"balance"`                    //余额
	VirtualBalance    float64 `json:"virtual_balance"`            //虚拟余额
	BonusBalance      float64 `json:"bonus_balance"`              //奖励余额
	CurrencyType      string  `json:"currency_type"`              //货币类型 IDR USD CNY
	Token             string  `json:"token"`                      //登录凭证
	LastToken         string  `json:"last_token"`                 //最后的登录凭证
	IsNew             int64   `json:"is_new"`                     //是否新用户 1是 0否 首充后变成0
	Nation            string  `json:"nation"`                     //是否新用户 1是 0否 首充后变成0
	HisTradingCount   int64   `gorm:"-" json:"hisTradingCount"`   //历史交易次数
	HisTradingAmounts float64 `gorm:"-" json:"hisTradingAmounts"` //历史交易金额
	FlowBindId        string  `json:"flow_bind_id"`               //整合流量方公司绑定ID
}

type HistoryTrading struct {
	HisTradingCount   int64   `json:"hisTradingCount"`   //历史交易次数
	HisTradingAmounts float64 `json:"hisTradingAmounts"` //历史交易金额
}

type NoticeInfo struct {
	Total int64 `json:"total"` //小红点数量
}

/** 用户登录数据 **/

type ClientUserBeans struct {
	AccountEmail            string `json:"accountEmail"`
	AppUserId               string `json:"appUserId"`
	BankCredit              int    `json:"bankCredit"`
	BonusBalance            string `json:"bonusBalance"`
	CurrencyType            string `json:"currencyType"`
	EmailCredit             int    `json:"emailCredit"`
	HadRechargeAmounts      string `json:"hadRechargeAmounts"`
	HadWithdrawAmounts      string `json:"hadWithdrawAmounts"`
	IdentityCredit          int    `json:"identityCredit"`
	IsFinishCredit          int    `json:"isFinishCredit"`
	IsNew                   int    `json:"isNew"`
	MobileCredit            int    `json:"mobileCredit"`
	Nation                  string `json:"nation"`
	RealBalance             string `json:"realBalance"`
	RechargeAmounts         string `json:"rechargeAmounts"`
	TotalHisBonusAmounts    string `json:"totalHisBonusAmounts"`
	TotalHisRechargeAmounts string `json:"totalHisRechargeAmounts"`
	TotalHisWithdrawAmounts string `json:"totalHisWithdrawAmounts"`
	TradingAmounts          string `json:"tradingAmounts"`
	UserLevel               int    `json:"userLevel"`
	VirtualBalance          string `json:"virtualBalance"`
	WithdrawAmounts         string `json:"withdrawAmounts"`
}
