package response

import (
	"goapi/pkg/helpers"
)

// 数据源表，一个数据源对应多个资产，不同的数据源有不同的休市时间

type TradingAssetsSource struct {
	ItemCode        string             `json:"item_code"`            // 应用编码
	AssetsCode      string             `json:"assets_code"`          // 资产编码
	KlineSource     string             `json:"current_kline_source"` // k线来源,数据源
	SubscribeSymbol string             `json:"subscribe_symbol"`     //数据源对应的订阅符号
	SleepTime       string             `json:"sleep_time"`           //休市时间 {"start":"6-04:00","ts_long":"172800"} start开始时间，从周几开始几点开始，休息多少秒
	IsValid         int                `json:"is_valid"`             //是否有效 1是 0否
	Remark          string             `json:"remark"`               //数据源备注
	Creator         string             `json:"creator"`              //创建人
	Updater         string             `json:"updater"`              //更新人
	CreateDate      helpers.TimeNormal `json:"create_date"`          //创建日期
	UpdateDate      helpers.TimeNormal `json:"update_date"`          //更新日期
	AssetsLogoUrl   string             `json:"assets_logo_url"`      //资产图标
	Sort            int64              `json:"sort"`                 //排序
	AssetsName      string             `json:"assets_name"`          //资产名称
}

type ListTradingAssetsSource struct {
	TradingAssetsSource
	AllowKlineSource    interface{} `json:"allow_kline_source"`    // 允许 k线来源,数据源
	CurrentSourceStatus int64       `json:"current_source_status"` // 当前数据源状态
}
