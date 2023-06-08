package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SystemEmailRoutingNodeMgr struct {
	*_BaseMgr
}

// SystemEmailRoutingNodeMgr open func
func SystemEmailRoutingNodeMgr(db *gorm.DB) *_SystemEmailRoutingNodeMgr {
	if db == nil {
		panic(fmt.Errorf("SystemEmailRoutingNodeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SystemEmailRoutingNodeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("system_email_routing_node"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_SystemEmailRoutingNodeMgr) Debug() *_SystemEmailRoutingNodeMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SystemEmailRoutingNodeMgr) GetTableName() string {
	return "system_email_routing_node"
}

// Reset 重置gorm会话
func (obj *_SystemEmailRoutingNodeMgr) Reset() *_SystemEmailRoutingNodeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SystemEmailRoutingNodeMgr) Get() (result SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SystemEmailRoutingNodeMgr) Gets() (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SystemEmailRoutingNodeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 邮箱路由节点配置表ID
func (obj *_SystemEmailRoutingNodeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEmail email获取 发信人邮箱
func (obj *_SystemEmailRoutingNodeMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithChannel channel获取 三方渠道
func (obj *_SystemEmailRoutingNodeMgr) WithChannel(channel string) Option {
	return optionFunc(func(o *options) { o.query["channel"] = channel })
}

// WithChannelUsername channel_username获取 渠道用户名
func (obj *_SystemEmailRoutingNodeMgr) WithChannelUsername(channelUsername string) Option {
	return optionFunc(func(o *options) { o.query["channel_username"] = channelUsername })
}

// WithChannelPassword channel_password获取 渠道密码
func (obj *_SystemEmailRoutingNodeMgr) WithChannelPassword(channelPassword string) Option {
	return optionFunc(func(o *options) { o.query["channel_password"] = channelPassword })
}

// WithChannelHost channel_host获取 渠道服务器
func (obj *_SystemEmailRoutingNodeMgr) WithChannelHost(channelHost string) Option {
	return optionFunc(func(o *options) { o.query["channel_host"] = channelHost })
}

// WithChannelPort channel_port获取 渠道端口号
func (obj *_SystemEmailRoutingNodeMgr) WithChannelPort(channelPort int) Option {
	return optionFunc(func(o *options) { o.query["channel_port"] = channelPort })
}

// WithAvailableNumber available_number获取 剩余可用条数
func (obj *_SystemEmailRoutingNodeMgr) WithAvailableNumber(availableNumber int) Option {
	return optionFunc(func(o *options) { o.query["available_number"] = availableNumber })
}

// WithMaxNumber max_number获取 每日最多使用条数
func (obj *_SystemEmailRoutingNodeMgr) WithMaxNumber(maxNumber int) Option {
	return optionFunc(func(o *options) { o.query["max_number"] = maxNumber })
}

// WithRateSuccess rate_success获取 成功率
func (obj *_SystemEmailRoutingNodeMgr) WithRateSuccess(rateSuccess float64) Option {
	return optionFunc(func(o *options) { o.query["rate_success"] = rateSuccess })
}

// WithRateFail rate_fail获取 失败率
func (obj *_SystemEmailRoutingNodeMgr) WithRateFail(rateFail float64) Option {
	return optionFunc(func(o *options) { o.query["rate_fail"] = rateFail })
}

// WithSort sort获取 优先排序
func (obj *_SystemEmailRoutingNodeMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithValid valid获取 有效状态 1有效 0失效
func (obj *_SystemEmailRoutingNodeMgr) WithValid(valid int) Option {
	return optionFunc(func(o *options) { o.query["valid"] = valid })
}

// WithCreatDate creat_date获取 创建时间
func (obj *_SystemEmailRoutingNodeMgr) WithCreatDate(creatDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["creat_date"] = creatDate })
}

// WithUpdateDate update_date获取 修改时间
func (obj *_SystemEmailRoutingNodeMgr) WithUpdateDate(updateDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_date"] = updateDate })
}

// WithCreator creator获取 创建人
func (obj *_SystemEmailRoutingNodeMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithUpdater updater获取 修改人
func (obj *_SystemEmailRoutingNodeMgr) WithUpdater(updater string) Option {
	return optionFunc(func(o *options) { o.query["updater"] = updater })
}

// WithRemarks remarks获取 备注
func (obj *_SystemEmailRoutingNodeMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithItemCode item_code获取 应用编码
func (obj *_SystemEmailRoutingNodeMgr) WithItemCode(itemCode string) Option {
	return optionFunc(func(o *options) { o.query["item_code"] = itemCode })
}

// WithType type获取 类型：1验证码通知，2营销邮件
func (obj *_SystemEmailRoutingNodeMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_SystemEmailRoutingNodeMgr) GetByOption(opts ...Option) (result SystemEmailRoutingNode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SystemEmailRoutingNodeMgr) GetByOptions(opts ...Option) (results []*SystemEmailRoutingNode, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SystemEmailRoutingNodeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SystemEmailRoutingNode, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where(options.query)
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
func (obj *_SystemEmailRoutingNodeMgr) GetFromID(id int64) (result SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 邮箱路由节点配置表ID
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromID(ids []int64) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 发信人邮箱
func (obj *_SystemEmailRoutingNodeMgr) GetFromEmail(email string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 发信人邮箱
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromEmail(emails []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromChannel 通过channel获取内容 三方渠道
func (obj *_SystemEmailRoutingNodeMgr) GetFromChannel(channel string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel` = ?", channel).Find(&results).Error

	return
}

// GetBatchFromChannel 批量查找 三方渠道
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromChannel(channels []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel` IN (?)", channels).Find(&results).Error

	return
}

// GetFromChannelUsername 通过channel_username获取内容 渠道用户名
func (obj *_SystemEmailRoutingNodeMgr) GetFromChannelUsername(channelUsername string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_username` = ?", channelUsername).Find(&results).Error

	return
}

// GetBatchFromChannelUsername 批量查找 渠道用户名
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromChannelUsername(channelUsernames []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_username` IN (?)", channelUsernames).Find(&results).Error

	return
}

// GetFromChannelPassword 通过channel_password获取内容 渠道密码
func (obj *_SystemEmailRoutingNodeMgr) GetFromChannelPassword(channelPassword string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_password` = ?", channelPassword).Find(&results).Error

	return
}

// GetBatchFromChannelPassword 批量查找 渠道密码
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromChannelPassword(channelPasswords []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_password` IN (?)", channelPasswords).Find(&results).Error

	return
}

// GetFromChannelHost 通过channel_host获取内容 渠道服务器
func (obj *_SystemEmailRoutingNodeMgr) GetFromChannelHost(channelHost string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_host` = ?", channelHost).Find(&results).Error

	return
}

// GetBatchFromChannelHost 批量查找 渠道服务器
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromChannelHost(channelHosts []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_host` IN (?)", channelHosts).Find(&results).Error

	return
}

// GetFromChannelPort 通过channel_port获取内容 渠道端口号
func (obj *_SystemEmailRoutingNodeMgr) GetFromChannelPort(channelPort int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_port` = ?", channelPort).Find(&results).Error

	return
}

// GetBatchFromChannelPort 批量查找 渠道端口号
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromChannelPort(channelPorts []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`channel_port` IN (?)", channelPorts).Find(&results).Error

	return
}

// GetFromAvailableNumber 通过available_number获取内容 剩余可用条数
func (obj *_SystemEmailRoutingNodeMgr) GetFromAvailableNumber(availableNumber int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`available_number` = ?", availableNumber).Find(&results).Error

	return
}

// GetBatchFromAvailableNumber 批量查找 剩余可用条数
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromAvailableNumber(availableNumbers []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`available_number` IN (?)", availableNumbers).Find(&results).Error

	return
}

// GetFromMaxNumber 通过max_number获取内容 每日最多使用条数
func (obj *_SystemEmailRoutingNodeMgr) GetFromMaxNumber(maxNumber int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`max_number` = ?", maxNumber).Find(&results).Error

	return
}

// GetBatchFromMaxNumber 批量查找 每日最多使用条数
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromMaxNumber(maxNumbers []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`max_number` IN (?)", maxNumbers).Find(&results).Error

	return
}

// GetFromRateSuccess 通过rate_success获取内容 成功率
func (obj *_SystemEmailRoutingNodeMgr) GetFromRateSuccess(rateSuccess float64) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`rate_success` = ?", rateSuccess).Find(&results).Error

	return
}

// GetBatchFromRateSuccess 批量查找 成功率
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromRateSuccess(rateSuccesss []float64) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`rate_success` IN (?)", rateSuccesss).Find(&results).Error

	return
}

// GetFromRateFail 通过rate_fail获取内容 失败率
func (obj *_SystemEmailRoutingNodeMgr) GetFromRateFail(rateFail float64) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`rate_fail` = ?", rateFail).Find(&results).Error

	return
}

// GetBatchFromRateFail 批量查找 失败率
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromRateFail(rateFails []float64) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`rate_fail` IN (?)", rateFails).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 优先排序
func (obj *_SystemEmailRoutingNodeMgr) GetFromSort(sort int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 优先排序
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromSort(sorts []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromValid 通过valid获取内容 有效状态 1有效 0失效
func (obj *_SystemEmailRoutingNodeMgr) GetFromValid(valid int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`valid` = ?", valid).Find(&results).Error

	return
}

// GetBatchFromValid 批量查找 有效状态 1有效 0失效
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromValid(valids []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`valid` IN (?)", valids).Find(&results).Error

	return
}

// GetFromCreatDate 通过creat_date获取内容 创建时间
func (obj *_SystemEmailRoutingNodeMgr) GetFromCreatDate(creatDate time.Time) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`creat_date` = ?", creatDate).Find(&results).Error

	return
}

// GetBatchFromCreatDate 批量查找 创建时间
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromCreatDate(creatDates []time.Time) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`creat_date` IN (?)", creatDates).Find(&results).Error

	return
}

// GetFromUpdateDate 通过update_date获取内容 修改时间
func (obj *_SystemEmailRoutingNodeMgr) GetFromUpdateDate(updateDate time.Time) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`update_date` = ?", updateDate).Find(&results).Error

	return
}

// GetBatchFromUpdateDate 批量查找 修改时间
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromUpdateDate(updateDates []time.Time) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`update_date` IN (?)", updateDates).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 创建人
func (obj *_SystemEmailRoutingNodeMgr) GetFromCreator(creator string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 创建人
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromCreator(creators []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromUpdater 通过updater获取内容 修改人
func (obj *_SystemEmailRoutingNodeMgr) GetFromUpdater(updater string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`updater` = ?", updater).Find(&results).Error

	return
}

// GetBatchFromUpdater 批量查找 修改人
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromUpdater(updaters []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`updater` IN (?)", updaters).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_SystemEmailRoutingNodeMgr) GetFromRemarks(remarks string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromRemarks(remarkss []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromItemCode 通过item_code获取内容 应用编码
func (obj *_SystemEmailRoutingNodeMgr) GetFromItemCode(itemCode string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`item_code` = ?", itemCode).Find(&results).Error

	return
}

// GetBatchFromItemCode 批量查找 应用编码
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromItemCode(itemCodes []string) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`item_code` IN (?)", itemCodes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型：1验证码通知，2营销邮件
func (obj *_SystemEmailRoutingNodeMgr) GetFromType(_type int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型：1验证码通知，2营销邮件
func (obj *_SystemEmailRoutingNodeMgr) GetBatchFromType(_types []int) (results []*SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SystemEmailRoutingNodeMgr) FetchByPrimaryKey(id int64) (result SystemEmailRoutingNode, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SystemEmailRoutingNode{}).Where("`id` = ?", id).First(&result).Error

	return
}
