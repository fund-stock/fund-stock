package qq

import (
	"bytes"
	JwEmail "github.com/jordan-wright/email"
	"goapi/pkg/config"
	"html/template"
	"net/smtp"
)

// SendEmail 使用第三方库发送邮件
func SendEmail(Info map[string]interface{}, toUser ...string) error {
	e := JwEmail.NewEmail()

	e.From = config.GetString("email.qq.user")
	e.To = toUser
	e.Subject = Info["title"].(string)

	t, err := template.ParseFiles("resource/template/email/template.html")
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	//作为变量传递给html模板
	err1 := t.Execute(body, struct {
		Email   []string
		Title   string
		Content template.HTML
	}{
		Email:   toUser,
		Title:   Info["title"].(string),
		Content: template.HTML(Info["content"].(string)),
	})

	if err1 != nil {
		return err1
	}
	// html形式的消息
	e.HTML = body.Bytes()
	return e.Send(config.GetString("email.qq.addr"), smtp.PlainAuth(
		"",
		config.GetString("email.qq.user"),
		config.GetString("email.qq.password"),
		config.GetString("email.qq.host"),
	))
}
