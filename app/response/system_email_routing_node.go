package response

import "time"

type SystemEmailRoutingNode struct {
	Id              int64     `json:"id"`               //邮箱路由节点配置表ID
	Email           string    `json:"email"`            //发信人邮箱
	Channel         string    `json:"channel"`          //三方渠道
	ChannelUsername string    `json:"channel_username"` //渠道用户名
	ChannelPassword string    `json:"channel_password"` //渠道密码
	ChannelHost     string    `json:"channel_host"`     //渠道服务器
	ChannelPort     int       `json:"channel_port"`     //渠道端口号
	AvailableNumber int       `json:"available_number"` //剩余可用条数
	MaxNumber       int       `json:"max_number"`       //每日最多使用条数
	RateSuccess     float64   `json:"rate_success"`     //成功率
	RateFail        float64   `json:"rate_fail"`        //失败率
	Sort            uint      `json:"sort"`             //优先排序
	Valid           uint      `json:"valid"`            //有效状态 1有效 0失效
	CreatDate       time.Time `json:"creat_date"`       //创建时间
	UpdateDate      time.Time `json:"update_date"`      //修改时间
	Creator         string    `json:"creator"`          //创建人
	Updater         string    `json:"updater"`          //修改人
	Remarks         string    `json:"remarks"`          //备注
}
