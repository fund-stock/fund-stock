package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _GoStockMgr struct {
	*_BaseMgr
}

// GoStockMgr open func
func GoStockMgr(db *gorm.DB) *_GoStockMgr {
	if db == nil {
		panic(fmt.Errorf("GoStockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoStockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_stock"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoStockMgr) Debug() *_GoStockMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoStockMgr) GetTableName() string {
	return "go_stock"
}

// Reset 重置gorm会话
func (obj *_GoStockMgr) Reset() *_GoStockMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoStockMgr) Get() (result GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoStockMgr) Gets() (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoStockMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoStock{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoStockMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 代码
func (obj *_GoStockMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 名称
func (obj *_GoStockMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAmount amount获取 金额
func (obj *_GoStockMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithNav nav获取 最新净值
func (obj *_GoStockMgr) WithNav(nav float64) Option {
	return optionFunc(func(o *options) { o.query["nav"] = nav })
}

// WithStatus status获取 状态：0-未启用，1-已启用
func (obj *_GoStockMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_GoStockMgr) WithCreateAt(createAt int64) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 更新时间
func (obj *_GoStockMgr) WithUpdateAt(updateAt int64) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoStockMgr) GetByOption(opts ...Option) (result GoStock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoStockMgr) GetByOptions(opts ...Option) (results []*GoStock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoStockMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoStock, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where(options.query)
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
func (obj *_GoStockMgr) GetFromID(id int64) (result GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoStockMgr) GetBatchFromID(ids []int64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 代码
func (obj *_GoStockMgr) GetFromCode(code string) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 代码
func (obj *_GoStockMgr) GetBatchFromCode(codes []string) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_GoStockMgr) GetFromName(name string) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_GoStockMgr) GetBatchFromName(names []string) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 金额
func (obj *_GoStockMgr) GetFromAmount(amount float64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 金额
func (obj *_GoStockMgr) GetBatchFromAmount(amounts []float64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromNav 通过nav获取内容 最新净值
func (obj *_GoStockMgr) GetFromNav(nav float64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`nav` = ?", nav).Find(&results).Error

	return
}

// GetBatchFromNav 批量查找 最新净值
func (obj *_GoStockMgr) GetBatchFromNav(navs []float64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`nav` IN (?)", navs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：0-未启用，1-已启用
func (obj *_GoStockMgr) GetFromStatus(status int) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态：0-未启用，1-已启用
func (obj *_GoStockMgr) GetBatchFromStatus(statuss []int) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_GoStockMgr) GetFromCreateAt(createAt int64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_GoStockMgr) GetBatchFromCreateAt(createAts []int64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 更新时间
func (obj *_GoStockMgr) GetFromUpdateAt(updateAt int64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_GoStockMgr) GetBatchFromUpdateAt(updateAts []int64) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoStockMgr) FetchByPrimaryKey(id int64) (result GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByCode  获取多个内容
func (obj *_GoStockMgr) FetchIndexByCode(code string) (results []*GoStock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoStock{}).Where("`code` = ?", code).Find(&results).Error

	return
}
