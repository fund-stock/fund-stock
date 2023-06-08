package requests

type SetKline struct {
	AssetsCode  string `json:"assets_code" form:"assets_code" validate:"required"`
	KlineSource string `json:"kline_source" form:"kline_source" validate:"required"`
}

type KlineSourceList struct {
	AssetsCode  string `json:"assets_code" form:"assets_code"`
	KlineSource string `json:"kline_source" form:"kline_source"`
	PageNum     int64  `json:"pageNum" form:"pageNum"`   // 第几页
	PageSize    int64  `json:"pageSize" form:"pageSize"` // 每页多少条
}
