package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemEmailSendLogMgr struct {
	*_BaseMgr
}

// SystemEmailSendLogMgr open func
func SystemEmailSendLogMgr(db *gorm.DB) *_SystemEmailSendLogMgr {
	if db == nil {
		panic(fmt.Errorf("SystemEmailSendLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemEmailSendLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_email_send_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_SystemEmailSendLogMgr) Debug() *_SystemEmailSendLogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemEmailSendLogMgr) GetTableName() string {
	return "system_email_send_log"
}

// Reset 重置gorm会话
func (obj *_SystemEmailSendLogMgr) Reset() *_SystemEmailSendLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SystemEmailSendLogMgr) Get() (result SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemEmailSendLogMgr) Gets() (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemEmailSendLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 邮箱发送记录表ID
func (obj *_SystemEmailSendLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSendingMailbox sending_mailbox获取 发件邮箱账号
func (obj *_SystemEmailSendLogMgr) WithSendingMailbox(sendingMailbox string) Option {
	return optionFunc(func(o *options) { o.query["sending_mailbox"] = sendingMailbox })
}

// WithReceiveEmail receive_email获取 收件箱账号
func (obj *_SystemEmailSendLogMgr) WithReceiveEmail(receiveEmail string) Option {
	return optionFunc(func(o *options) { o.query["receive_email"] = receiveEmail })
}

// WithSendTotal send_total获取 发送数量
func (obj *_SystemEmailSendLogMgr) WithSendTotal(sendTotal int) Option {
	return optionFunc(func(o *options) { o.query["send_total"] = sendTotal })
}

// WithAvailableNumber available_number获取 剩余可用条数
func (obj *_SystemEmailSendLogMgr) WithAvailableNumber(availableNumber int) Option {
	return optionFunc(func(o *options) { o.query["available_number"] = availableNumber })
}

// WithStatus status获取 发送状态：1成功 0失败
func (obj *_SystemEmailSendLogMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatDate creat_date获取 创建时间
func (obj *_SystemEmailSendLogMgr) WithCreatDate(creatDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["creat_date"] = creatDate })
}

// WithUpdateDate update_date获取 修改时间
func (obj *_SystemEmailSendLogMgr) WithUpdateDate(updateDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_date"] = updateDate })
}

// WithCreator creator获取 创建人
func (obj *_SystemEmailSendLogMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithUpdater updater获取 修改人
func (obj *_SystemEmailSendLogMgr) WithUpdater(updater string) Option {
	return optionFunc(func(o *options) { o.query["updater"] = updater })
}

// WithRemarks remarks获取 备注
func (obj *_SystemEmailSendLogMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithBody body获取 发送内容
func (obj *_SystemEmailSendLogMgr) WithBody(body string) Option {
	return optionFunc(func(o *options) { o.query["body"] = body })
}

// GetByOption 功能选项模式获取
func (obj *_SystemEmailSendLogMgr) GetByOption(opts ...Option) (result SystemEmailSendLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemEmailSendLogMgr) GetByOptions(opts ...Option) (results []*SystemEmailSendLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SystemEmailSendLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SystemEmailSendLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 邮箱发送记录表ID
func (obj *_SystemEmailSendLogMgr) GetFromID(id int64) (result SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 邮箱发送记录表ID
func (obj *_SystemEmailSendLogMgr) GetBatchFromID(ids []int64) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSendingMailbox 通过sending_mailbox获取内容 发件邮箱账号
func (obj *_SystemEmailSendLogMgr) GetFromSendingMailbox(sendingMailbox string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`sending_mailbox` = ?", sendingMailbox).Find(&results).Error

	return
}

// GetBatchFromSendingMailbox 批量查找 发件邮箱账号
func (obj *_SystemEmailSendLogMgr) GetBatchFromSendingMailbox(sendingMailboxs []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`sending_mailbox` IN (?)", sendingMailboxs).Find(&results).Error

	return
}

// GetFromReceiveEmail 通过receive_email获取内容 收件箱账号
func (obj *_SystemEmailSendLogMgr) GetFromReceiveEmail(receiveEmail string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`receive_email` = ?", receiveEmail).Find(&results).Error

	return
}

// GetBatchFromReceiveEmail 批量查找 收件箱账号
func (obj *_SystemEmailSendLogMgr) GetBatchFromReceiveEmail(receiveEmails []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`receive_email` IN (?)", receiveEmails).Find(&results).Error

	return
}

// GetFromSendTotal 通过send_total获取内容 发送数量
func (obj *_SystemEmailSendLogMgr) GetFromSendTotal(sendTotal int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`send_total` = ?", sendTotal).Find(&results).Error

	return
}

// GetBatchFromSendTotal 批量查找 发送数量
func (obj *_SystemEmailSendLogMgr) GetBatchFromSendTotal(sendTotals []int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`send_total` IN (?)", sendTotals).Find(&results).Error

	return
}

// GetFromAvailableNumber 通过available_number获取内容 剩余可用条数
func (obj *_SystemEmailSendLogMgr) GetFromAvailableNumber(availableNumber int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`available_number` = ?", availableNumber).Find(&results).Error

	return
}

// GetBatchFromAvailableNumber 批量查找 剩余可用条数
func (obj *_SystemEmailSendLogMgr) GetBatchFromAvailableNumber(availableNumbers []int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`available_number` IN (?)", availableNumbers).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 发送状态：1成功 0失败
func (obj *_SystemEmailSendLogMgr) GetFromStatus(status int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 发送状态：1成功 0失败
func (obj *_SystemEmailSendLogMgr) GetBatchFromStatus(statuss []int) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatDate 通过creat_date获取内容 创建时间
func (obj *_SystemEmailSendLogMgr) GetFromCreatDate(creatDate time.Time) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`creat_date` = ?", creatDate).Find(&results).Error

	return
}

// GetBatchFromCreatDate 批量查找 创建时间
func (obj *_SystemEmailSendLogMgr) GetBatchFromCreatDate(creatDates []time.Time) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`creat_date` IN (?)", creatDates).Find(&results).Error

	return
}

// GetFromUpdateDate 通过update_date获取内容 修改时间
func (obj *_SystemEmailSendLogMgr) GetFromUpdateDate(updateDate time.Time) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`update_date` = ?", updateDate).Find(&results).Error

	return
}

// GetBatchFromUpdateDate 批量查找 修改时间
func (obj *_SystemEmailSendLogMgr) GetBatchFromUpdateDate(updateDates []time.Time) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`update_date` IN (?)", updateDates).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 创建人
func (obj *_SystemEmailSendLogMgr) GetFromCreator(creator string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 创建人
func (obj *_SystemEmailSendLogMgr) GetBatchFromCreator(creators []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromUpdater 通过updater获取内容 修改人
func (obj *_SystemEmailSendLogMgr) GetFromUpdater(updater string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`updater` = ?", updater).Find(&results).Error

	return
}

// GetBatchFromUpdater 批量查找 修改人
func (obj *_SystemEmailSendLogMgr) GetBatchFromUpdater(updaters []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`updater` IN (?)", updaters).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_SystemEmailSendLogMgr) GetFromRemarks(remarks string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_SystemEmailSendLogMgr) GetBatchFromRemarks(remarkss []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromBody 通过body获取内容 发送内容
func (obj *_SystemEmailSendLogMgr) GetFromBody(body string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`body` = ?", body).Find(&results).Error

	return
}

// GetBatchFromBody 批量查找 发送内容
func (obj *_SystemEmailSendLogMgr) GetBatchFromBody(bodys []string) (results []*SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`body` IN (?)", bodys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemEmailSendLogMgr) FetchByPrimaryKey(id int64) (result SystemEmailSendLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailSendLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
