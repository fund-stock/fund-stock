package common

import (
	"goapi/app/models"
	"goapi/app/requests"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"time"
)

// 添加操作记录

func CreateSystemEmailSendLog(params requests.SendOneEmail, AvailableNumber int, Status int, ErrMsg string) {
	var SystemEmailSendLog models.GoEmailLog
	SystemEmailSendLog.SendingMailbox = params.MailAccount
	SystemEmailSendLog.ReceiveEmail = params.ReceiveEmail
	SystemEmailSendLog.SendTotal = 1
	SystemEmailSendLog.AvailableNumber = AvailableNumber // 剩余可用条数
	SystemEmailSendLog.CreatDate = time.Now()
	SystemEmailSendLog.UpdateDate = time.Now()
	SystemEmailSendLog.Creator = "GoMail"
	SystemEmailSendLog.Updater = "GoMail"
	SystemEmailSendLog.Status = Status
	SystemEmailSendLog.Remarks = ErrMsg
	SystemEmailSendLog.Body = params.MessageContent
	err := mysql.DB.Model(models.GoEmailLog{}).Create(&SystemEmailSendLog).Error
	if err != nil {
		logger.Error(err)
	}
}
