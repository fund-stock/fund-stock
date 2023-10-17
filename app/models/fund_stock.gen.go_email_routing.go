package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoEmailRoutingMgr struct {
	*_BaseMgr
}

// GoEmailRoutingMgr open func
func GoEmailRoutingMgr(db *gorm.DB) *_GoEmailRoutingMgr {
	if db == nil {
		panic(fmt.Errorf("GoEmailRoutingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoEmailRoutingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_email_routing"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoEmailRoutingMgr) Debug() *_GoEmailRoutingMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoEmailRoutingMgr) GetTableName() string {
	return "go_email_routing"
}

// Reset 重置gorm会话
func (obj *_GoEmailRoutingMgr) Reset() *_GoEmailRoutingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoEmailRoutingMgr) Get() (result GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoEmailRoutingMgr) Gets() (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoEmailRoutingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 邮箱路由节点配置表ID
func (obj *_GoEmailRoutingMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEmail email获取 发信人邮箱
func (obj *_GoEmailRoutingMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithChannel channel获取 三方渠道
func (obj *_GoEmailRoutingMgr) WithChannel(channel string) Option {
	return optionFunc(func(o *options) { o.query["channel"] = channel })
}

// WithChannelUsername channel_username获取 渠道用户名
func (obj *_GoEmailRoutingMgr) WithChannelUsername(channelUsername string) Option {
	return optionFunc(func(o *options) { o.query["channel_username"] = channelUsername })
}

// WithChannelPassword channel_password获取 渠道密码
func (obj *_GoEmailRoutingMgr) WithChannelPassword(channelPassword string) Option {
	return optionFunc(func(o *options) { o.query["channel_password"] = channelPassword })
}

// WithChannelHost channel_host获取 渠道服务器
func (obj *_GoEmailRoutingMgr) WithChannelHost(channelHost string) Option {
	return optionFunc(func(o *options) { o.query["channel_host"] = channelHost })
}

// WithChannelPort channel_port获取 渠道端口号
func (obj *_GoEmailRoutingMgr) WithChannelPort(channelPort int) Option {
	return optionFunc(func(o *options) { o.query["channel_port"] = channelPort })
}

// WithAvailableNumber available_number获取 剩余可用条数
func (obj *_GoEmailRoutingMgr) WithAvailableNumber(availableNumber int) Option {
	return optionFunc(func(o *options) { o.query["available_number"] = availableNumber })
}

// WithMaxNumber max_number获取 每日最多使用条数
func (obj *_GoEmailRoutingMgr) WithMaxNumber(maxNumber int) Option {
	return optionFunc(func(o *options) { o.query["max_number"] = maxNumber })
}

// WithRateSuccess rate_success获取 成功率
func (obj *_GoEmailRoutingMgr) WithRateSuccess(rateSuccess float64) Option {
	return optionFunc(func(o *options) { o.query["rate_success"] = rateSuccess })
}

// WithRateFail rate_fail获取 失败率
func (obj *_GoEmailRoutingMgr) WithRateFail(rateFail float64) Option {
	return optionFunc(func(o *options) { o.query["rate_fail"] = rateFail })
}

// WithSort sort获取 优先排序
func (obj *_GoEmailRoutingMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithValid valid获取 有效状态 1有效 0失效
func (obj *_GoEmailRoutingMgr) WithValid(valid int) Option {
	return optionFunc(func(o *options) { o.query["valid"] = valid })
}

// WithCreatDate creat_date获取 创建时间
func (obj *_GoEmailRoutingMgr) WithCreatDate(creatDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["creat_date"] = creatDate })
}

// WithUpdateDate update_date获取 修改时间
func (obj *_GoEmailRoutingMgr) WithUpdateDate(updateDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_date"] = updateDate })
}

// WithCreator creator获取 创建人
func (obj *_GoEmailRoutingMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithUpdater updater获取 修改人
func (obj *_GoEmailRoutingMgr) WithUpdater(updater string) Option {
	return optionFunc(func(o *options) { o.query["updater"] = updater })
}

// WithRemarks remarks获取 备注
func (obj *_GoEmailRoutingMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithType type获取 类型：1验证码通知，2营销邮件
func (obj *_GoEmailRoutingMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_GoEmailRoutingMgr) GetByOption(opts ...Option) (result GoEmailRouting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoEmailRoutingMgr) GetByOptions(opts ...Option) (results []*GoEmailRouting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoEmailRoutingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoEmailRouting, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where(options.query)
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

// GetFromID 通过id获取内容 邮箱路由节点配置表ID
func (obj *_GoEmailRoutingMgr) GetFromID(id int64) (result GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 邮箱路由节点配置表ID
func (obj *_GoEmailRoutingMgr) GetBatchFromID(ids []int64) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 发信人邮箱
func (obj *_GoEmailRoutingMgr) GetFromEmail(email string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 发信人邮箱
func (obj *_GoEmailRoutingMgr) GetBatchFromEmail(emails []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromChannel 通过channel获取内容 三方渠道
func (obj *_GoEmailRoutingMgr) GetFromChannel(channel string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel` = ?", channel).Find(&results).Error

	return
}

// GetBatchFromChannel 批量查找 三方渠道
func (obj *_GoEmailRoutingMgr) GetBatchFromChannel(channels []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel` IN (?)", channels).Find(&results).Error

	return
}

// GetFromChannelUsername 通过channel_username获取内容 渠道用户名
func (obj *_GoEmailRoutingMgr) GetFromChannelUsername(channelUsername string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_username` = ?", channelUsername).Find(&results).Error

	return
}

// GetBatchFromChannelUsername 批量查找 渠道用户名
func (obj *_GoEmailRoutingMgr) GetBatchFromChannelUsername(channelUsernames []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_username` IN (?)", channelUsernames).Find(&results).Error

	return
}

// GetFromChannelPassword 通过channel_password获取内容 渠道密码
func (obj *_GoEmailRoutingMgr) GetFromChannelPassword(channelPassword string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_password` = ?", channelPassword).Find(&results).Error

	return
}

// GetBatchFromChannelPassword 批量查找 渠道密码
func (obj *_GoEmailRoutingMgr) GetBatchFromChannelPassword(channelPasswords []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_password` IN (?)", channelPasswords).Find(&results).Error

	return
}

// GetFromChannelHost 通过channel_host获取内容 渠道服务器
func (obj *_GoEmailRoutingMgr) GetFromChannelHost(channelHost string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_host` = ?", channelHost).Find(&results).Error

	return
}

// GetBatchFromChannelHost 批量查找 渠道服务器
func (obj *_GoEmailRoutingMgr) GetBatchFromChannelHost(channelHosts []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_host` IN (?)", channelHosts).Find(&results).Error

	return
}

// GetFromChannelPort 通过channel_port获取内容 渠道端口号
func (obj *_GoEmailRoutingMgr) GetFromChannelPort(channelPort int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_port` = ?", channelPort).Find(&results).Error

	return
}

// GetBatchFromChannelPort 批量查找 渠道端口号
func (obj *_GoEmailRoutingMgr) GetBatchFromChannelPort(channelPorts []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`channel_port` IN (?)", channelPorts).Find(&results).Error

	return
}

// GetFromAvailableNumber 通过available_number获取内容 剩余可用条数
func (obj *_GoEmailRoutingMgr) GetFromAvailableNumber(availableNumber int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`available_number` = ?", availableNumber).Find(&results).Error

	return
}

// GetBatchFromAvailableNumber 批量查找 剩余可用条数
func (obj *_GoEmailRoutingMgr) GetBatchFromAvailableNumber(availableNumbers []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`available_number` IN (?)", availableNumbers).Find(&results).Error

	return
}

// GetFromMaxNumber 通过max_number获取内容 每日最多使用条数
func (obj *_GoEmailRoutingMgr) GetFromMaxNumber(maxNumber int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`max_number` = ?", maxNumber).Find(&results).Error

	return
}

// GetBatchFromMaxNumber 批量查找 每日最多使用条数
func (obj *_GoEmailRoutingMgr) GetBatchFromMaxNumber(maxNumbers []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`max_number` IN (?)", maxNumbers).Find(&results).Error

	return
}

// GetFromRateSuccess 通过rate_success获取内容 成功率
func (obj *_GoEmailRoutingMgr) GetFromRateSuccess(rateSuccess float64) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`rate_success` = ?", rateSuccess).Find(&results).Error

	return
}

// GetBatchFromRateSuccess 批量查找 成功率
func (obj *_GoEmailRoutingMgr) GetBatchFromRateSuccess(rateSuccesss []float64) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`rate_success` IN (?)", rateSuccesss).Find(&results).Error

	return
}

// GetFromRateFail 通过rate_fail获取内容 失败率
func (obj *_GoEmailRoutingMgr) GetFromRateFail(rateFail float64) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`rate_fail` = ?", rateFail).Find(&results).Error

	return
}

// GetBatchFromRateFail 批量查找 失败率
func (obj *_GoEmailRoutingMgr) GetBatchFromRateFail(rateFails []float64) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`rate_fail` IN (?)", rateFails).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 优先排序
func (obj *_GoEmailRoutingMgr) GetFromSort(sort int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 优先排序
func (obj *_GoEmailRoutingMgr) GetBatchFromSort(sorts []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromValid 通过valid获取内容 有效状态 1有效 0失效
func (obj *_GoEmailRoutingMgr) GetFromValid(valid int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`valid` = ?", valid).Find(&results).Error

	return
}

// GetBatchFromValid 批量查找 有效状态 1有效 0失效
func (obj *_GoEmailRoutingMgr) GetBatchFromValid(valids []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`valid` IN (?)", valids).Find(&results).Error

	return
}

// GetFromCreatDate 通过creat_date获取内容 创建时间
func (obj *_GoEmailRoutingMgr) GetFromCreatDate(creatDate time.Time) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`creat_date` = ?", creatDate).Find(&results).Error

	return
}

// GetBatchFromCreatDate 批量查找 创建时间
func (obj *_GoEmailRoutingMgr) GetBatchFromCreatDate(creatDates []time.Time) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`creat_date` IN (?)", creatDates).Find(&results).Error

	return
}

// GetFromUpdateDate 通过update_date获取内容 修改时间
func (obj *_GoEmailRoutingMgr) GetFromUpdateDate(updateDate time.Time) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`update_date` = ?", updateDate).Find(&results).Error

	return
}

// GetBatchFromUpdateDate 批量查找 修改时间
func (obj *_GoEmailRoutingMgr) GetBatchFromUpdateDate(updateDates []time.Time) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`update_date` IN (?)", updateDates).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 创建人
func (obj *_GoEmailRoutingMgr) GetFromCreator(creator string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 创建人
func (obj *_GoEmailRoutingMgr) GetBatchFromCreator(creators []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromUpdater 通过updater获取内容 修改人
func (obj *_GoEmailRoutingMgr) GetFromUpdater(updater string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`updater` = ?", updater).Find(&results).Error

	return
}

// GetBatchFromUpdater 批量查找 修改人
func (obj *_GoEmailRoutingMgr) GetBatchFromUpdater(updaters []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`updater` IN (?)", updaters).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_GoEmailRoutingMgr) GetFromRemarks(remarks string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_GoEmailRoutingMgr) GetBatchFromRemarks(remarkss []string) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型：1验证码通知，2营销邮件
func (obj *_GoEmailRoutingMgr) GetFromType(_type int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型：1验证码通知，2营销邮件
func (obj *_GoEmailRoutingMgr) GetBatchFromType(_types []int) (results []*GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoEmailRoutingMgr) FetchByPrimaryKey(id int64) (result GoEmailRouting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoEmailRouting{}).Where("`id` = ?", id).First(&result).Error

	return
}
