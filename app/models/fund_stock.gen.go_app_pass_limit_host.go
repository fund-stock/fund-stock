package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoAppPassLimitHostMgr struct {
	*_BaseMgr
}

// GoAppPassLimitHostMgr open func
func GoAppPassLimitHostMgr(db *gorm.DB) *_GoAppPassLimitHostMgr {
	if db == nil {
		panic(fmt.Errorf("GoAppPassLimitHostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoAppPassLimitHostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_app_pass_limit_host"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoAppPassLimitHostMgr) Debug() *_GoAppPassLimitHostMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoAppPassLimitHostMgr) GetTableName() string {
	return "go_app_pass_limit_host"
}

// Reset 重置gorm会话
func (obj *_GoAppPassLimitHostMgr) Reset() *_GoAppPassLimitHostMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoAppPassLimitHostMgr) Get() (result GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoAppPassLimitHostMgr) Gets() (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoAppPassLimitHostMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 服务器访问白名单表主键自增ID
func (obj *_GoAppPassLimitHostMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPassHost pass_host获取 活跃的域名
func (obj *_GoAppPassLimitHostMgr) WithPassHost(passHost string) Option {
	return optionFunc(func(o *options) { o.query["pass_host"] = passHost })
}

// WithCreateDate create_date获取 创建日期
func (obj *_GoAppPassLimitHostMgr) WithCreateDate(createDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_date"] = createDate })
}

// WithUpdateDate update_date获取 修改日期
func (obj *_GoAppPassLimitHostMgr) WithUpdateDate(updateDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_date"] = updateDate })
}

// WithCreator creator获取 添加人ID
func (obj *_GoAppPassLimitHostMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithUpdater updater获取 更新人ID
func (obj *_GoAppPassLimitHostMgr) WithUpdater(updater string) Option {
	return optionFunc(func(o *options) { o.query["updater"] = updater })
}

// WithItemCode item_code获取 关联项目编码
func (obj *_GoAppPassLimitHostMgr) WithItemCode(itemCode string) Option {
	return optionFunc(func(o *options) { o.query["item_code"] = itemCode })
}

// WithIsShow is_show获取 是否展示 1是 0否
func (obj *_GoAppPassLimitHostMgr) WithIsShow(isShow int) Option {
	return optionFunc(func(o *options) { o.query["is_show"] = isShow })
}

// WithIsDelete is_delete获取 是否删除 1是 0否
func (obj *_GoAppPassLimitHostMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithPassLevel pass_level获取 活跃等级
func (obj *_GoAppPassLimitHostMgr) WithPassLevel(passLevel int) Option {
	return optionFunc(func(o *options) { o.query["pass_level"] = passLevel })
}

// WithReleaseSource release_source获取 放行源 (放行ip来源标注）
func (obj *_GoAppPassLimitHostMgr) WithReleaseSource(releaseSource string) Option {
	return optionFunc(func(o *options) { o.query["release_source"] = releaseSource })
}

// WithReleaseModule release_module获取 放行模块 (对应系统模块)
func (obj *_GoAppPassLimitHostMgr) WithReleaseModule(releaseModule string) Option {
	return optionFunc(func(o *options) { o.query["release_module"] = releaseModule })
}

// GetByOption 功能选项模式获取
func (obj *_GoAppPassLimitHostMgr) GetByOption(opts ...Option) (result GoAppPassLimitHost, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoAppPassLimitHostMgr) GetByOptions(opts ...Option) (results []*GoAppPassLimitHost, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoAppPassLimitHostMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoAppPassLimitHost, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where(options.query)
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

// GetFromID 通过id获取内容 服务器访问白名单表主键自增ID
func (obj *_GoAppPassLimitHostMgr) GetFromID(id int64) (result GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 服务器访问白名单表主键自增ID
func (obj *_GoAppPassLimitHostMgr) GetBatchFromID(ids []int64) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPassHost 通过pass_host获取内容 活跃的域名
func (obj *_GoAppPassLimitHostMgr) GetFromPassHost(passHost string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`pass_host` = ?", passHost).Find(&results).Error

	return
}

// GetBatchFromPassHost 批量查找 活跃的域名
func (obj *_GoAppPassLimitHostMgr) GetBatchFromPassHost(passHosts []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`pass_host` IN (?)", passHosts).Find(&results).Error

	return
}

// GetFromCreateDate 通过create_date获取内容 创建日期
func (obj *_GoAppPassLimitHostMgr) GetFromCreateDate(createDate time.Time) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`create_date` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找 创建日期
func (obj *_GoAppPassLimitHostMgr) GetBatchFromCreateDate(createDates []time.Time) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`create_date` IN (?)", createDates).Find(&results).Error

	return
}

// GetFromUpdateDate 通过update_date获取内容 修改日期
func (obj *_GoAppPassLimitHostMgr) GetFromUpdateDate(updateDate time.Time) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`update_date` = ?", updateDate).Find(&results).Error

	return
}

// GetBatchFromUpdateDate 批量查找 修改日期
func (obj *_GoAppPassLimitHostMgr) GetBatchFromUpdateDate(updateDates []time.Time) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`update_date` IN (?)", updateDates).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容 添加人ID
func (obj *_GoAppPassLimitHostMgr) GetFromCreator(creator string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找 添加人ID
func (obj *_GoAppPassLimitHostMgr) GetBatchFromCreator(creators []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromUpdater 通过updater获取内容 更新人ID
func (obj *_GoAppPassLimitHostMgr) GetFromUpdater(updater string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`updater` = ?", updater).Find(&results).Error

	return
}

// GetBatchFromUpdater 批量查找 更新人ID
func (obj *_GoAppPassLimitHostMgr) GetBatchFromUpdater(updaters []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`updater` IN (?)", updaters).Find(&results).Error

	return
}

// GetFromItemCode 通过item_code获取内容 关联项目编码
func (obj *_GoAppPassLimitHostMgr) GetFromItemCode(itemCode string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`item_code` = ?", itemCode).Find(&results).Error

	return
}

// GetBatchFromItemCode 批量查找 关联项目编码
func (obj *_GoAppPassLimitHostMgr) GetBatchFromItemCode(itemCodes []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`item_code` IN (?)", itemCodes).Find(&results).Error

	return
}

// GetFromIsShow 通过is_show获取内容 是否展示 1是 0否
func (obj *_GoAppPassLimitHostMgr) GetFromIsShow(isShow int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`is_show` = ?", isShow).Find(&results).Error

	return
}

// GetBatchFromIsShow 批量查找 是否展示 1是 0否
func (obj *_GoAppPassLimitHostMgr) GetBatchFromIsShow(isShows []int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`is_show` IN (?)", isShows).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 是否删除 1是 0否
func (obj *_GoAppPassLimitHostMgr) GetFromIsDelete(isDelete int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除 1是 0否
func (obj *_GoAppPassLimitHostMgr) GetBatchFromIsDelete(isDeletes []int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromPassLevel 通过pass_level获取内容 活跃等级
func (obj *_GoAppPassLimitHostMgr) GetFromPassLevel(passLevel int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`pass_level` = ?", passLevel).Find(&results).Error

	return
}

// GetBatchFromPassLevel 批量查找 活跃等级
func (obj *_GoAppPassLimitHostMgr) GetBatchFromPassLevel(passLevels []int) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`pass_level` IN (?)", passLevels).Find(&results).Error

	return
}

// GetFromReleaseSource 通过release_source获取内容 放行源 (放行ip来源标注）
func (obj *_GoAppPassLimitHostMgr) GetFromReleaseSource(releaseSource string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`release_source` = ?", releaseSource).Find(&results).Error

	return
}

// GetBatchFromReleaseSource 批量查找 放行源 (放行ip来源标注）
func (obj *_GoAppPassLimitHostMgr) GetBatchFromReleaseSource(releaseSources []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`release_source` IN (?)", releaseSources).Find(&results).Error

	return
}

// GetFromReleaseModule 通过release_module获取内容 放行模块 (对应系统模块)
func (obj *_GoAppPassLimitHostMgr) GetFromReleaseModule(releaseModule string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`release_module` = ?", releaseModule).Find(&results).Error

	return
}

// GetBatchFromReleaseModule 批量查找 放行模块 (对应系统模块)
func (obj *_GoAppPassLimitHostMgr) GetBatchFromReleaseModule(releaseModules []string) (results []*GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`release_module` IN (?)", releaseModules).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoAppPassLimitHostMgr) FetchByPrimaryKey(id int64) (result GoAppPassLimitHost, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAppPassLimitHost{}).Where("`id` = ?", id).First(&result).Error

	return
}
