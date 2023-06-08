package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _GoFundMgr struct {
	*_BaseMgr
}

// GoFundMgr open func
func GoFundMgr(db *gorm.DB) *_GoFundMgr {
	if db == nil {
		panic(fmt.Errorf("GoFundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoFundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_fund"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoFundMgr) Debug() *_GoFundMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoFundMgr) GetTableName() string {
	return "go_fund"
}

// Reset 重置gorm会话
func (obj *_GoFundMgr) Reset() *_GoFundMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoFundMgr) Get() (result GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoFundMgr) Gets() (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoFundMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoFund{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoFundMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProductID product_id获取 产品id
func (obj *_GoFundMgr) WithProductID(productID string) Option {
	return optionFunc(func(o *options) { o.query["product_id"] = productID })
}

// WithCode code获取 代码
func (obj *_GoFundMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 名称
func (obj *_GoFundMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAmount amount获取 金额
func (obj *_GoFundMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithNav nav获取 最新净值
func (obj *_GoFundMgr) WithNav(nav float64) Option {
	return optionFunc(func(o *options) { o.query["nav"] = nav })
}

// WithStatus status获取 状态：0-未启用，1-已启用
func (obj *_GoFundMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateAt create_at获取 创建时间
func (obj *_GoFundMgr) WithCreateAt(createAt int64) Option {
	return optionFunc(func(o *options) { o.query["create_at"] = createAt })
}

// WithUpdateAt update_at获取 更新时间
func (obj *_GoFundMgr) WithUpdateAt(updateAt int64) Option {
	return optionFunc(func(o *options) { o.query["update_at"] = updateAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoFundMgr) GetByOption(opts ...Option) (result GoFund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoFundMgr) GetByOptions(opts ...Option) (results []*GoFund, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoFundMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoFund, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where(options.query)
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
func (obj *_GoFundMgr) GetFromID(id int64) (result GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoFundMgr) GetBatchFromID(ids []int64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromProductID 通过product_id获取内容 产品id
func (obj *_GoFundMgr) GetFromProductID(productID string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`product_id` = ?", productID).Find(&results).Error

	return
}

// GetBatchFromProductID 批量查找 产品id
func (obj *_GoFundMgr) GetBatchFromProductID(productIDs []string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`product_id` IN (?)", productIDs).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 代码
func (obj *_GoFundMgr) GetFromCode(code string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 代码
func (obj *_GoFundMgr) GetBatchFromCode(codes []string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_GoFundMgr) GetFromName(name string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_GoFundMgr) GetBatchFromName(names []string) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 金额
func (obj *_GoFundMgr) GetFromAmount(amount float64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 金额
func (obj *_GoFundMgr) GetBatchFromAmount(amounts []float64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromNav 通过nav获取内容 最新净值
func (obj *_GoFundMgr) GetFromNav(nav float64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`nav` = ?", nav).Find(&results).Error

	return
}

// GetBatchFromNav 批量查找 最新净值
func (obj *_GoFundMgr) GetBatchFromNav(navs []float64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`nav` IN (?)", navs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：0-未启用，1-已启用
func (obj *_GoFundMgr) GetFromStatus(status int) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态：0-未启用，1-已启用
func (obj *_GoFundMgr) GetBatchFromStatus(statuss []int) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateAt 通过create_at获取内容 创建时间
func (obj *_GoFundMgr) GetFromCreateAt(createAt int64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`create_at` = ?", createAt).Find(&results).Error

	return
}

// GetBatchFromCreateAt 批量查找 创建时间
func (obj *_GoFundMgr) GetBatchFromCreateAt(createAts []int64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`create_at` IN (?)", createAts).Find(&results).Error

	return
}

// GetFromUpdateAt 通过update_at获取内容 更新时间
func (obj *_GoFundMgr) GetFromUpdateAt(updateAt int64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`update_at` = ?", updateAt).Find(&results).Error

	return
}

// GetBatchFromUpdateAt 批量查找 更新时间
func (obj *_GoFundMgr) GetBatchFromUpdateAt(updateAts []int64) (results []*GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`update_at` IN (?)", updateAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoFundMgr) FetchByPrimaryKey(id int64) (result GoFund, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoFund{}).Where("`id` = ?", id).First(&result).Error

	return
}
