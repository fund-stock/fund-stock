package qtimg

// Resp 结构体，包含顶级的响应信息
type Resp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data StockData `json:"data"`
}

// StockData 结构体，包含与股票相关的数据
type StockData struct {
	StockInfo StockInfo `json:"info"`
}

// StockInfo 结构体，包含"info"字段的数据
type StockInfo struct {
	Day     [][]string  `json:"day"`
	Qt      QtData      `json:"qt"`
	MxPrice MxPriceData `json:"mx_price"`
	Prec    string      `json:"prec"`
	Version string      `json:"version"`
}

// QtData 结构体，包含"qt"字段的数据
type QtData struct {
	Info   []string `json:"info"`
	Market []string `json:"market"`
}

// MxPriceData 结构体，包含"mx_price"字段的数据
type MxPriceData struct {
	Mx    MxData    `json:"mx"`
	Price PriceData `json:"price"`
}

// MxData 结构体，包含"mx"字段的数据
type MxData struct {
	Data     []interface{} `json:"data"`
	Timeline []interface{} `json:"timeline"`
}

// PriceData 结构体，包含"price"字段的数据
type PriceData struct {
	Data []interface{} `json:"data"`
}
