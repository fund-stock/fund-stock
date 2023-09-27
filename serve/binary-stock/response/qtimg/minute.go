package qtimg

type Minute struct {
	List     []MinuteInfo
	PrePrice float64 // 昨日的收盘价
}

type MinuteInfo struct {
	Time  string  // 时间
	Price float64 // 当前价
	Vol   int64   // 总成交量
	Ave   float64 //  均价
}

type MinuteWWW struct {
	Code int        `json:"code"`
	Data MinuteData `json:"data"`
}

type MinuteData struct {
	Info struct {
		Data struct {
			Data []string `json:"data"`
		} `json:"data"`
		QT struct {
			Info []string `json:"info"`
		} `json:"qt"`
	} `json:"info"`
}
