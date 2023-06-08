package response

import "time"

// app服务器访问白名单表

type AppPassLimitHost struct {
	Id         int64     `json:"id"`          //app服务器访问白名单表主键自增ID
	PassHost   string    `json:"pass_host"`   //活跃的域名
	CreateDate time.Time `json:"create_date"` //创建日期
	UpdateDate time.Time `json:"update_date"` //修改日期
	Creator    string    `json:"creator"`     //添加人ID
	Updater    string    `json:"updater"`     //更新人ID
	ItemCode   string    `json:"item_code"`   //关联项目编码
	IsShow     int       `json:"is_show"`     //是否展示 1是 0否
	IsDelete   int       `json:"is_delete"`   //是否删除 1是 0否
	PassLevel  int       `json:"pass_level"`  //活跃等级
}
