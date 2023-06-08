package requests

type Sub struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

type SubHistory struct {
	Symbol    string `json:"symbol"`
	BeginDate string `json:"beginDate"`
	EndDate   string `json:"endDate"`
	Limit     string `json:"limit"`
	Gzip      string `json:"gzip"`
}

type SubUserInfo struct {
	AppUserId string `json:"appUserId"`
}

type SubStationLetter struct {
	Sub SubToken `json:"sub"`
	Id  string   `json:"id"`
}

type SubToken struct {
	Token       string `json:"token"`
	DeciveBrand string `json:"decive_brand"`
}

type SubNationInfo struct {
	ItemCode string `json:"itemCode"`
	Nation   string `json:"nation"`
}

type TradeOrderInfo struct {
	TradeOrderId        string `json:"tradeOrderId"` //交易表主键自增ID
	AppUserId           string `json:"appUserId"`
	AssetsCode          string `json:"assetsCode"`
	ItemCode            string `json:"itemCode"`
	TradeBeginTimestamp string `json:"tradeBeginTimestamp"`
	TradeEndTimestamp   string `json:"tradeEndTimestamp"`
	BeginAssetsPrice    string `json:"beginAssetsPrice"`
	EndAssetsPrice      string `json:"endAssetsPrice"`
	UpDownType          string `json:"upDownType"`
	Amounts             string `json:"amounts"`
	IsDemo              string `json:"isDemo"`
}

type DayTime struct {
	Num string `json:"num" form:"num" validate:"required"`
}

type SubParams struct {
	Sub interface{} `json:"sub"`
	Id  string      `json:"id"`
}
