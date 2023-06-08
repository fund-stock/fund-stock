package response

// 弹窗事件

type PopupMsg struct {
	Id           int64  `json:"id"`            //弹窗ID
	ItemCode     string `json:"item_code"`     //项目编码
	LanguageCode string `json:"language_code"` //语言，模板只区分语言，不区分国家，方便后台管理
	Name         string `json:"name"`          //模板名称
	Title        string `json:"title"`         //弹窗标题
	Body         string `json:"body"`          //弹窗内容
	Img          string `json:"img"`           //图片
	Confirm      string `json:"confirm"`       //确认按钮文字
	Type         int    `json:"type"`          //类型1直接关闭，2-h5跳转事件，3-app内部跳转事件 4-js事件
	Redirect     string `json:"redirect"`      //重定向h5页面，或者app内部地址
	ExpireTs     int64  `json:"expire_ts"`     //模板到期时间
	Remarks      string `json:"remarks"`       //备注
}
