package response

import (
	"goapi/pkg/helpers"
)

// 资产表

type TradingAssets struct {
	Id            int64              `json:"id"`              //资产表主键ID
	ItemCode      string             `json:"item_code"`       //应用编码
	AssetsName    string             `json:"assets_name"`     //资产名称
	AssetsCode    string             `json:"assets_code"`     //资产编码
	IsValid       int                `json:"is_valid"`        //是否有效 1是 0否
	IsShow        int                `json:"is_show"`         //是否展示 1是 0否 删除标识符
	AssetsLogoUrl string             `json:"assets_logo_url"` //资产logo链接
	Sort          int                `json:"sort"`            //排序
	CreateDate    helpers.TimeNormal `json:"create_date"`     //创建日期
	UpdateDate    helpers.TimeNormal `json:"update_date"`     //更新日期
	Creator       string             `json:"creator"`         //创建人
	Updater       string             `json:"updater"`         //更新人
	ReturnRate    float64            `json:"return_rate"`     //收益率
}

type TradingAssetsMonitoring struct {
	ID                  int64   `json:"id" gorm:"column:id"`                                     // 资产表主键ID
	AssetsCode          string  `json:"assets_code" gorm:"column:assets_code"`                   // 资产编码
	PointNumber         int     `json:"point_number" gorm:"column:point_number"`                 // 保留小数点位数
	IsMonitoring        int     `json:"is_monitoring" gorm:"column:is_monitoring"`               // 是否有效1-是 0-否
	ShockAmplitude      float64 `json:"shock_amplitude" gorm:"column:shock_amplitude"`           // 震荡幅度 TODO 暂时无用
	VolatilityAmplitude float64 `json:"volatility_amplitude" gorm:"column:volatility_amplitude"` // 波动幅度
	VolatilityFrequency int64   `json:"volatility_frequency" gorm:"column:volatility_frequency"` // 震荡频率
	AssetsInsertFrame   string  `json:"assets_insert_frame"`
}

type TradingAssetsMonitoringConfig struct {
	ID                  int64   `json:"id" gorm:"column:id"`                                     // 资产表主键ID
	AssetsCode          string  `json:"assets_code" gorm:"column:assets_code"`                   // 资产编码
	PointNumber         int     `json:"point_number" gorm:"column:point_number"`                 // 保留小数点位数
	IsMonitoring        int     `json:"is_monitoring" gorm:"column:is_monitoring"`               // 是否有效1-是 0-否
	ShockAmplitude      float64 `json:"shock_amplitude" gorm:"column:shock_amplitude"`           // 震荡幅度
	VolatilityAmplitude float64 `json:"volatility_amplitude" gorm:"column:volatility_amplitude"` // 波动幅度
	VolatilityFrequency int64   `json:"volatility_frequency" gorm:"column:volatility_frequency"` // 震荡频率
	AssetsInsertFrame   AssetsInsertFrame
}

type AssetsInsertFrame struct {
	AreaTimer               int     `json:"areaTimer"`               // 设置区域时长
	IntervalTimer           int     `json:"intervalTimer"`           // 设置间隔触发时间
	InAdvanceTimer          int64   `json:"inAdvanceTimer"`          // 提前判断时间 设置提前触发时间（拿这个时间间隔作为判断条件）
	LossThreshold           float64 `json:"lossThreshold"`           // 触发属性：区域平台亏损金额>=xx美金
	EarlyExecutionTimer     int64   `json:"earlyExecutionTimer"`     // 设置插针提前执行时间
	EarlyExecutionFrequency int     `json:"earlyExecutionFrequency"` // 设置插针提前执行频率
	RecoverTimer            int64   `json:"recoverTimer"`            // 设置插针恢复时间
	RecoverFrequency        int     `json:"recoverFrequency"`        // 设置插针恢复频率
}

type AssetsCodeArray struct {
	AssetsCode  string `json:"assets_code"`  //资产编码
	IsWhitelist string `json:"is_whitelist"` //是否白名单1是 0否，白名单的资产不跟随自动变化赔率
}

//  K线数据表

type KlineTs struct {
	Ts          int64   `json:"ts"`           //  kilne返回日期时间戳
	CClose      float64 `json:"c_close"`      //  结束峰值
	KlineSource string  `json:"kline_source"` //  k线来源 huobi
}
