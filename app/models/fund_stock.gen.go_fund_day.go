package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoFundDayMgr struct {
	*_BaseMgr
}

// GoFundDayMgr open func
func GoFundDayMgr(db *gorm.DB) *_GoFundDayMgr {
	if db == nil {
		panic(fmt.Errorf("GoFundDayMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoFundDayMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_fund_day"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoFundDayMgr) Debug() *_GoFundDayMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoFundDayMgr) GetTableName() string {
	return "go_fund_day"
}

// Reset 重置gorm会话
func (obj *_GoFundDayMgr) Reset() *_GoFundDayMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoFundDayMgr) Get() (result GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoFundDayMgr) Gets() (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoFundDayMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoFundDayMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 代码
func (obj *_GoFundDayMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 名称
func (obj *_GoFundDayMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAmount amount获取 金额
func (obj *_GoFundDayMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithNav nav获取 最新净值
func (obj *_GoFundDayMgr) WithNav(nav float64) Option {
	return optionFunc(func(o *options) { o.query["nav"] = nav })
}

// WithDayTs day_ts获取 当天的时间戳
func (obj *_GoFundDayMgr) WithDayTs(dayTs int64) Option {
	return optionFunc(func(o *options) { o.query["day_ts"] = dayTs })
}

// WithDayAt day_at获取 当天的时间
func (obj *_GoFundDayMgr) WithDayAt(dayAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["day_at"] = dayAt })
}

// WithCreateAt create_at获取 创建时间
func (obj *_GoFundDayMgr) WithCreateAt(createAt int64) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 更新时间
func (obj *_GoFundDayMgr) WithUpdateAt(updateAt int64) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoFundDayMgr) GetByOption(opts ...Option) (result GoFundDay, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoFundDayMgr) GetByOptions(opts ...Option) (results []*GoFundDay, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoFundDayMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoFundDay, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_GoFundDayMgr) GetFromID(id int64) (result GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoFundDayMgr) GetBatchFromID(ids []int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 代码
func (obj *_GoFundDayMgr) GetFromCode(code string) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 代码
func (obj *_GoFundDayMgr) GetBatchFromCode(codes []string) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_GoFundDayMgr) GetFromName(name string) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_GoFundDayMgr) GetBatchFromName(names []string) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 金额
func (obj *_GoFundDayMgr) GetFromAmount(amount float64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 金额
func (obj *_GoFundDayMgr) GetBatchFromAmount(amounts []float64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromNav 通过nav获取内容 最新净值
func (obj *_GoFundDayMgr) GetFromNav(nav float64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`nav` = ?", nav).Find(&results).Error

	return
}

// GetBatchFromNav 批量查找 最新净值
func (obj *_GoFundDayMgr) GetBatchFromNav(navs []float64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`nav` IN (?)", navs).Find(&results).Error

	return
}

// GetFromDayTs 通过day_ts获取内容 当天的时间戳
func (obj *_GoFundDayMgr) GetFromDayTs(dayTs int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`day_ts` = ?", dayTs).Find(&results).Error

	return
}

// GetBatchFromDayTs 批量查找 当天的时间戳
func (obj *_GoFundDayMgr) GetBatchFromDayTs(dayTss []int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`day_ts` IN (?)", dayTss).Find(&results).Error

	return
}

// GetFromDayAt 通过day_at获取内容 当天的时间
func (obj *_GoFundDayMgr) GetFromDayAt(dayAt time.Time) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`day_at` = ?", dayAt).Find(&results).Error

	return
}

// GetBatchFromDayAt 批量查找 当天的时间
func (obj *_GoFundDayMgr) GetBatchFromDayAt(dayAts []time.Time) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`day_at` IN (?)", dayAts).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_GoFundDayMgr) GetFromCreateAt(createAt int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_GoFundDayMgr) GetBatchFromCreateAt(createAts []int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 更新时间
func (obj *_GoFundDayMgr) GetFromUpdateAt(updateAt int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_GoFundDayMgr) GetBatchFromUpdateAt(updateAts []int64) (results []*GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoFundDayMgr) FetchByPrimaryKey(id int64) (result GoFundDay, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFundDay{}).Where("`id` = ?", id).First(&result).Error

	return
}
