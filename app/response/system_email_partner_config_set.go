package response

import "time"

type SystemEmailPartnerConfigSet struct {
	Id              int64     `json:"id"`               // 三方邮箱默认配置表ID
	Channel         string    `json:"channel"`          //三方渠道
	ChannelHost     string    `json:"channel_host"`     //服务器
	ChannelPort     string    `json:"channel_port"`     //渠道端口号
	ChannelUsername string    `json:"channel_username"` //用户名
	ChannelPassword string    `json:"channel_password"` //密码
	Email           string    `json:"email"`            //发信人邮箱
	MaxNumber       uint      `json:"max_number"`       //每日最多使用条数
	Valid           uint      `json:"valid"`            //有效状态 1有效 0失效
	CreatDate       time.Time `json:"creat_date"`       //创建时间
	UpdateDate      time.Time `json:"update_date"`      //修改时间
	Creator         string    `json:"creator"`          //创建人
	Updater         string    `json:"updater"`          //修改人
	Remarks         string    `json:"remarks"`          //备注
	Nation          string    `json:"nation"`           //绑定国家
}
