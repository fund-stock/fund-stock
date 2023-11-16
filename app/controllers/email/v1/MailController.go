package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/app/response"
	"goapi/pkg/echo"
	"goapi/pkg/email"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/validate"
	"goapi/serve/binary-email/common"
	"goapi/serve/binary-email/config"
	"goapi/serve/binary-email/queue"
)

// k线图服务

type MailController struct {
	BaseController
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) QueryMargin(c *gin.Context) {
	var params requests.QueryMargin
	var result response.SystemEmailRoutingNode
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	// 构建查询条件
	conditions := map[string]interface{}{
		"valid":  1,
		"nation": params.Nation,
		"email":  params.SenderName,
	}
	mysql.DB.Model(models.GoEmailRouting{}).Where("valid = :valid AND nation = :nation AND email = :email", conditions).Find(&result)
	if result.Id == 0 {
		echo.Error(c, "Failed", fmt.Sprintf("，国家：%v 不存在发信人 %v ", params.Nation, params.SenderName))
		return
	}
	echo.Success(c, gin.H{"available_number": result.AvailableNumber}, "")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) SendOneEmail(c *gin.Context) {
	var params requests.SendOneEmail
	var sendData email.Params
	var SendTotal response.SendTotal
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.MailAccount) == 0 {
		echo.Error(c, "ParamsError", "发件人邮箱不能为空")
		return
	}
	// 获取邮箱路由节点
	err, SystemEmailRoutingNode := config.GetOne(params.MailAccount)
	if err != nil {
		echo.Error(c, "ParamsError", err.Error())
		return
	}
	params.MailAccount = SystemEmailRoutingNode.Email
	if SystemEmailRoutingNode.AvailableNumber <= 0 {
		echo.Error(c, "Failed", fmt.Sprintf("%v 剩余可用条数为 %v", params.MailAccount, SystemEmailRoutingNode.AvailableNumber))
		return
	}
	sendTotalResult := cmap.New().Items()
	err = mysql.DB.Model(models.GoEmailLog{}).Where(map[string]interface{}{
		"sending_mailbox": params.MailAccount,
	}).Select("sum(send_total) AS sendTotal").
		Find(&sendTotalResult).Error
	if err != nil {
		echo.Error(c, "Failed", err.Error())
		return
	}
	// map转struct
	err = mapstructure.Decode(sendTotalResult, &SendTotal)
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	if (SendTotal.Total + len(params.ReceiveEmail)) >= SystemEmailRoutingNode.MaxNumber {
		echo.Error(c, "Failed", fmt.Sprintf("%v 每日最多使用 %v 条，您当前已使用：%v", params.MailAccount, SystemEmailRoutingNode.MaxNumber, SendTotal.Total))
		return
	}
	sendData.From = params.MailAccount
	sendData.To = params.ReceiveEmail
	sendData.ContentType = params.ContentType
	sendData.Subject = params.MainTitle
	sendData.Body = params.MessageContent
	sendData.Host = SystemEmailRoutingNode.ChannelHost
	sendData.Port = SystemEmailRoutingNode.ChannelPort
	sendData.Username = SystemEmailRoutingNode.ChannelUsername
	sendData.Password = SystemEmailRoutingNode.ChannelPassword
	err = email.SendOneEmail(sendData)
	if err != nil {
		// 统计失败率
		common.CreateSystemEmailSendLog(params, SystemEmailRoutingNode.AvailableNumber, 0, err.Error())
		go queue.StatisticsEmail(params.MailAccount, "fail")
		echo.Error(c, "Failed", err.Error())
		return
	}
	// 统计成功率
	go queue.StatisticsEmail(params.MailAccount, "success")
	// 剩余可用条数
	AvailableNumber := SystemEmailRoutingNode.AvailableNumber - 1
	err = mysql.DB.Model(models.GoEmailRouting{}).Where("id", SystemEmailRoutingNode.Id).Update("available_number", AvailableNumber).Error
	if err != nil {
		logger.Error(err)
	}
	common.CreateSystemEmailSendLog(params, AvailableNumber, 1, "")
	echo.Success(c, nil, "SUCCESS")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) SendBatchEmail(c *gin.Context) {
	var params requests.SendBatchEmail
	var sendData []string
	var SendTotal response.SendTotal
	_ = c.Bind(&params)
	// 数据验证
	if validate.ParamsError(c, params) {
		return
	}
	if len(params.MailAccount) == 0 {
		echo.Error(c, "ParamsError", "发件人邮箱不能为空")
		return
	}
	// 获取邮箱路由节点
	err, SystemEmailRoutingNode := config.GetOne(params.MailAccount)
	if err != nil {
		echo.Error(c, "ParamsError", err.Error())
		return
	}
	if SystemEmailRoutingNode.AvailableNumber <= 0 {
		echo.Error(c, "Failed", fmt.Sprintf("%v 剩余可用条数为 %v", params.MailAccount, SystemEmailRoutingNode.AvailableNumber))
		return
	}
	sendTotalResult := cmap.New().Items()
	err = mysql.DB.Model(models.GoEmailLog{}).Where(map[string]interface{}{
		"sending_mailbox": params.MailAccount,
	}).Select("sum(send_total) AS sendTotal").
		Find(&sendTotalResult).Error
	if err != nil {
		echo.Error(c, "Failed", err.Error())
		return
	}
	// map转struct
	err = mapstructure.Decode(sendTotalResult, &SendTotal)
	if err != nil {
		logger.Error(err)
		echo.Error(c, "Failed", err.Error())
		return
	}
	if (SendTotal.Total + len(params.ReceiveEmail)) >= SystemEmailRoutingNode.MaxNumber {
		echo.Error(c, "Failed", fmt.Sprintf("%v 每日最多使用 %v 条，您当前已使用：%v", params.MailAccount, SystemEmailRoutingNode.MaxNumber, SendTotal.Total))
		return
	}
	// 创建消息队列，将要发送的消息进行入队
	for _, item := range params.ReceiveEmail {
		marshal, err := json.Marshal(email.Params{
			From:     params.MailAccount,
			To:       item,
			Subject:  params.MainTitle,
			Body:     params.MessageContent,
			Host:     SystemEmailRoutingNode.ChannelHost,
			Port:     SystemEmailRoutingNode.ChannelPort,
			Username: SystemEmailRoutingNode.ChannelUsername,
			Password: SystemEmailRoutingNode.ChannelPassword,
		})
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
		sendData = append(sendData, string(marshal))
	}
	if len(sendData) > 0 {
		err = queue.Produce(sendData)
		if err != nil {
			echo.Error(c, "Failed", err.Error())
			return
		}
	}
	echo.Success(c, nil, "SUCCESS")
}

/**接口编号:【U01】
 * .发送邮件【POST】
 */
// MailController

func (h *MailController) ResetMargin(c *gin.Context) {
	var list []response.SystemEmailRoutingNode
	mysql.DB.Model(models.GoEmailRouting{}).Where("valid=1").Find(&list)
	for _, item := range list {
		go queue.InitStatisticsEmail(item.Email)
	}
	echo.Success(c, nil, "SUCCESS")
}
