package migration

type Fund struct {
	ID       int64   `json:"id" gorm:"column:id"`
	Code     string  `json:"code" gorm:"column:code"`           // 代码
	Name     string  `json:"name" gorm:"column:name"`           // 名称
	Amount   float64 `json:"amount" gorm:"column:amount"`       // 金额
	Nav      float64 `json:"nav" gorm:"column:nav"`             // 最新净值
	Status   int     `json:"status" gorm:"column:status"`       // 状态：0-未启用，1-已启用
	CreateAt int64   `json:"create_at" gorm:"column:create_at"` // 创建时间
	UpdateAt int64   `json:"update_at" gorm:"column:update_at"` // 更新时间
}

func (m *Fund) TableName() string {
	return "go_fund"
}
