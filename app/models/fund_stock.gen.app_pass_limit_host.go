package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AppPassLimitHostMgr struct {
	*_BaseMgr
}

// AppPassLimitHostMgr open func
func AppPassLimitHostMgr(db *gorm.DB) *_AppPassLimitHostMgr {
	if db == nil {
		panic(fmt.Errorf("AppPassLimitHostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AppPassLimitHostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("app_pass_limit_host"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AppPassLimitHostMgr) Debug() *_AppPassLimitHostMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AppPassLimitHostMgr) GetTableName() string {
	return "app_pass_limit_host"
}

// Reset 重置gorm会话
func (obj *_AppPassLimitHostMgr) Reset() *_AppPassLimitHostMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AppPassLimitHostMgr) Get() (result AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AppPassLimitHostMgr) Gets() (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AppPassLimitHostMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 app服务器访问白名单表主键自增ID
func (obj *_AppPassLimitHostMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPassHost pass_host获取 活跃的域名
func (obj *_AppPassLimitHostMgr) WithPassHost(passHost string) Option {
	return optionFunc(func(o *options) { o.query["pass_host"] = passHost })
}

// WithCreateDate create_date获取 创建日期
func (obj *_AppPassLimitHostMgr) WithCreateDate(createDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_date"] = createDate })
}

// WithUpdateDate update_date获取 修改日期
func (obj *_AppPassLimitHostMgr) WithUpdateDate(updateDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_date"] = updateDate })
}

// WithCreator creator获取 添加人ID
func (obj *_AppPassLimitHostMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithUpdater updater获取 更新人ID
func (obj *_AppPassLimitHostMgr) WithUpdater(updater string) Option {
	return optionFunc(func(o *options) { o.query["updater"] = updater })
}

// WithItemCode item_code获取 关联项目编码
func (obj *_AppPassLimitHostMgr) WithItemCode(itemCode string) Option {
	return optionFunc(func(o *options) { o.query["item_code"] = itemCode })
}

// WithIsShow is_show获取 是否展示 1是 0否
func (obj *_AppPassLimitHostMgr) WithIsShow(isShow int) Option {
	return optionFunc(func(o *options) { o.query["is_show"] = isShow })
}

// WithIsDelete is_delete获取 是否删除 1是 0否
func (obj *_AppPassLimitHostMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithPassLevel pass_level获取 活跃等级
func (obj *_AppPassLimitHostMgr) WithPassLevel(passLevel int) Option {
	return optionFunc(func(o *options) { o.query["pass_level"] = passLevel })
}

// GetByOption 功能选项模式获取
func (obj *_AppPassLimitHostMgr) GetByOption(opts ...Option) (result AppPassLimitHost, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AppPassLimitHostMgr) GetByOptions(opts ...Option) (results []*AppPassLimitHost, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AppPassLimitHostMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AppPassLimitHost, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where(options.query)
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

// GetFromID 通过id获取内容 app服务器访问白名单表主键自增ID
func (obj *_AppPassLimitHostMgr) GetFromID(id int64) (result AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 app服务器访问白名单表主键自增ID
func (obj *_AppPassLimitHostMgr) GetBatchFromID(ids []int64) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPassHost 通过pass_host获取内容 活跃的域名
func (obj *_AppPassLimitHostMgr) GetFromPassHost(passHost string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`pass_host` = ?", passHost).Find(&results).Error

	return
}

// GetBatchFromPassHost 批量查找 活跃的域名
func (obj *_AppPassLimitHostMgr) GetBatchFromPassHost(passHosts []string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`pass_host` IN (?)", passHosts).Find(&results).Error

	return
}

// GetFromCreateDate 通过create_date获取内容 创建日期
func (obj *_AppPassLimitHostMgr) GetFromCreateDate(createDate time.Time) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`create_date` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找 创建日期
func (obj *_AppPassLimitHostMgr) GetBatchFromCreateDate(createDates []time.Time) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`create_date` IN (?)", createDates).Find(&results).Error

	return
}

// GetFromUpdateDate 通过update_date获取内容 修改日期
func (obj *_AppPassLimitHostMgr) GetFromUpdateDate(updateDate time.Time) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`update_date` = ?", updateDate).Find(&results).Error

	return
}

// GetBatchFromUpdateDate 批量查找 修改日期
func (obj *_AppPassLimitHostMgr) GetBatchFromUpdateDate(updateDates []time.Time) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`update_date` IN (?)", updateDates).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 添加人ID
func (obj *_AppPassLimitHostMgr) GetFromCreator(creator string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 添加人ID
func (obj *_AppPassLimitHostMgr) GetBatchFromCreator(creators []string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromUpdater 通过updater获取内容 更新人ID
func (obj *_AppPassLimitHostMgr) GetFromUpdater(updater string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`updater` = ?", updater).Find(&results).Error

	return
}

// GetBatchFromUpdater 批量查找 更新人ID
func (obj *_AppPassLimitHostMgr) GetBatchFromUpdater(updaters []string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`updater` IN (?)", updaters).Find(&results).Error

	return
}

// GetFromItemCode 通过item_code获取内容 关联项目编码
func (obj *_AppPassLimitHostMgr) GetFromItemCode(itemCode string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`item_code` = ?", itemCode).Find(&results).Error

	return
}

// GetBatchFromItemCode 批量查找 关联项目编码
func (obj *_AppPassLimitHostMgr) GetBatchFromItemCode(itemCodes []string) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`item_code` IN (?)", itemCodes).Find(&results).Error

	return
}

// GetFromIsShow 通过is_show获取内容 是否展示 1是 0否
func (obj *_AppPassLimitHostMgr) GetFromIsShow(isShow int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`is_show` = ?", isShow).Find(&results).Error

	return
}

// GetBatchFromIsShow 批量查找 是否展示 1是 0否
func (obj *_AppPassLimitHostMgr) GetBatchFromIsShow(isShows []int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`is_show` IN (?)", isShows).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 是否删除 1是 0否
func (obj *_AppPassLimitHostMgr) GetFromIsDelete(isDelete int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除 1是 0否
func (obj *_AppPassLimitHostMgr) GetBatchFromIsDelete(isDeletes []int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromPassLevel 通过pass_level获取内容 活跃等级
func (obj *_AppPassLimitHostMgr) GetFromPassLevel(passLevel int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`pass_level` = ?", passLevel).Find(&results).Error

	return
}

// GetBatchFromPassLevel 批量查找 活跃等级
func (obj *_AppPassLimitHostMgr) GetBatchFromPassLevel(passLevels []int) (results []*AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`pass_level` IN (?)", passLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AppPassLimitHostMgr) FetchByPrimaryKey(id int64) (result AppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AppPassLimitHost{}).Where("`id` = ?", id).First(&result).Error

	return
}
