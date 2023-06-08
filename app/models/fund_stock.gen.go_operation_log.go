package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoOperationLogMgr struct {
	*_BaseMgr
}

// GoOperationLogMgr open func
func GoOperationLogMgr(db *gorm.DB) *_GoOperationLogMgr {
	if db == nil {
		panic(fmt.Errorf("GoOperationLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoOperationLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_operation_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoOperationLogMgr) Debug() *_GoOperationLogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoOperationLogMgr) GetTableName() string {
	return "go_operation_log"
}

// Reset 重置gorm会话
func (obj *_GoOperationLogMgr) Reset() *_GoOperationLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoOperationLogMgr) Get() (result GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoOperationLogMgr) Gets() (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoOperationLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoOperationLogMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 类型1：为总后台用户，2：为合作商用户
func (obj *_GoOperationLogMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithAccountID account_id获取 账号id
func (obj *_GoOperationLogMgr) WithAccountID(accountID int) Option {
	return optionFunc(func(o *options) { o.query["account_id"] = accountID })
}

// WithContent content获取 操作内容
func (obj *_GoOperationLogMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithIP ip获取 ip地址
func (obj *_GoOperationLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithAddress address获取 操作地址
func (obj *_GoOperationLogMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoOperationLogMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoOperationLogMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoOperationLogMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoOperationLogMgr) GetByOption(opts ...Option) (result GoOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoOperationLogMgr) GetByOptions(opts ...Option) (results []*GoOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoOperationLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoOperationLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where(options.query)
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
func (obj *_GoOperationLogMgr) GetFromID(id uint64) (result GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoOperationLogMgr) GetBatchFromID(ids []uint64) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型1：为总后台用户，2：为合作商用户
func (obj *_GoOperationLogMgr) GetFromType(_type int) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型1：为总后台用户，2：为合作商用户
func (obj *_GoOperationLogMgr) GetBatchFromType(_types []int) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromAccountID 通过account_id获取内容 账号id
func (obj *_GoOperationLogMgr) GetFromAccountID(accountID int) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`account_id` = ?", accountID).Find(&results).Error

	return
}

// GetBatchFromAccountID 批量查找 账号id
func (obj *_GoOperationLogMgr) GetBatchFromAccountID(accountIDs []int) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`account_id` IN (?)", accountIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 操作内容
func (obj *_GoOperationLogMgr) GetFromContent(content string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 操作内容
func (obj *_GoOperationLogMgr) GetBatchFromContent(contents []string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip地址
func (obj *_GoOperationLogMgr) GetFromIP(ip string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip地址
func (obj *_GoOperationLogMgr) GetBatchFromIP(ips []string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 操作地址
func (obj *_GoOperationLogMgr) GetFromAddress(address string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 操作地址
func (obj *_GoOperationLogMgr) GetBatchFromAddress(addresss []string) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoOperationLogMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoOperationLogMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoOperationLogMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoOperationLogMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoOperationLogMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoOperationLogMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoOperationLogMgr) FetchByPrimaryKey(id uint64) (result GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxGoOperationLogCreatedAt  获取多个内容
func (obj *_GoOperationLogMgr) FetchIndexByIDxGoOperationLogCreatedAt(createdAt time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// FetchIndexByIDxGoOperationLogDeletedAt  获取多个内容
func (obj *_GoOperationLogMgr) FetchIndexByIDxGoOperationLogDeletedAt(deletedAt time.Time) (results []*GoOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoOperationLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
