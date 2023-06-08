package requests

type QueryMargin struct {
	Nation     string `json:"nation" validate:"required"`     // 国家编码
	SenderName string `json:"senderName" validate:"required"` // 发件人名称
}

type SendOneEmail struct {
	ItemCode       string `json:"itemCode" validate:"required"`       // 项目编码
	MailAccount    string `json:"mailAccount" validate:"required"`    // 发邮件邮箱
	ContentType    string `json:"ContentType"`                        // 类型 text/html ，text/plain
	ReceiveEmail   string `json:"receiveEmail" validate:"required"`   // 收件邮箱
	MainTitle      string `json:"mainTitle" validate:"required"`      // 邮件标题
	MessageContent string `json:"messageContent" validate:"required"` // 邮件内容
}

type SendBatchEmail struct {
	ItemCode       string   `json:"itemCode" validate:"required"`
	MailAccount    string   `json:"mailAccount" validate:"required"`    // 发邮件邮箱
	ReceiveEmail   []string `json:"receiveEmail" validate:"required"`   // 收件邮箱
	MainTitle      string   `json:"mainTitle" validate:"required"`      // 邮件标题
	MessageContent string   `json:"messageContent" validate:"required"` // 邮件内容
}
