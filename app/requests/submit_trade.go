package requests

type SubmitTradeParams struct {
	AssetsId         string  `json:"assetsId" validate:"required"`         //资产
	Amounts          float64 `json:"amounts" validate:"required"`          // 交易金额
	TradingBeginDate string  `json:"tradingBeginDate" validate:"required"` // 交易起始日期
	TradeEndDate     string  `json:"tradeEndDate" validate:"required"`     // 交易截止日期
	BeginAssetsPrice float64 `json:"beginAssetsPrice" validate:"required"` // 起始交易资产价格
	AmountType       int64   `json:"amountType" validate:"required"`       // 交易金额类型 1 余额 2 bonus 3 虚拟货币
	UpDownType       string  `json:"upDownType" validate:"required"`       // 买涨买跌类型 up涨 down跌
	TradeOrderId     string  `json:"tradeOrderId"`                         // 交易订单ID
	KlineTime        int64   `json:"klineTime" validate:"required"`        // k线交易时间点
}
