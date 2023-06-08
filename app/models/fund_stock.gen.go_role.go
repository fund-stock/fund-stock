package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoRoleMgr struct {
	*_BaseMgr
}

// GoRoleMgr open func
func GoRoleMgr(db *gorm.DB) *_GoRoleMgr {
	if db == nil {
		panic(fmt.Errorf("GoRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoRoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoRoleMgr) Debug() *_GoRoleMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoRoleMgr) GetTableName() string {
	return "go_role"
}

// Reset 重置gorm会话
func (obj *_GoRoleMgr) Reset() *_GoRoleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoRoleMgr) Get() (result GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoRoleMgr) Gets() (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoRoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoRoleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 角色名称
func (obj *_GoRoleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithRoutes routes获取 路由id,该角色所具有的路由
func (obj *_GoRoleMgr) WithRoutes(routes string) Option {
	return optionFunc(func(o *options) { o.query["routes"] = routes })
}

// WithDesc desc获取 角色描述
func (obj *_GoRoleMgr) WithDesc(desc string) Option {
	return optionFunc(func(o *options) { o.query["desc"] = desc })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoRoleMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoRoleMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoRoleMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoRoleMgr) GetByOption(opts ...Option) (result GoRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoRoleMgr) GetByOptions(opts ...Option) (results []*GoRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoRoleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoRole, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where(options.query)
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
func (obj *_GoRoleMgr) GetFromID(id uint64) (result GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoRoleMgr) GetBatchFromID(ids []uint64) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 角色名称
func (obj *_GoRoleMgr) GetFromName(name string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 角色名称
func (obj *_GoRoleMgr) GetBatchFromName(names []string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromRoutes 通过routes获取内容 路由id,该角色所具有的路由
func (obj *_GoRoleMgr) GetFromRoutes(routes string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`routes` = ?", routes).Find(&results).Error

	return
}

// GetBatchFromRoutes 批量查找 路由id,该角色所具有的路由
func (obj *_GoRoleMgr) GetBatchFromRoutes(routess []string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`routes` IN (?)", routess).Find(&results).Error

	return
}

// GetFromDesc 通过desc获取内容 角色描述
func (obj *_GoRoleMgr) GetFromDesc(desc string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`desc` = ?", desc).Find(&results).Error

	return
}

// GetBatchFromDesc 批量查找 角色描述
func (obj *_GoRoleMgr) GetBatchFromDesc(descs []string) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`desc` IN (?)", descs).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoRoleMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoRoleMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoRoleMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoRoleMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoRoleMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoRoleMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoRoleMgr) FetchByPrimaryKey(id uint64) (result GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxGoRoleCreatedAt  获取多个内容
func (obj *_GoRoleMgr) FetchIndexByIDxGoRoleCreatedAt(createdAt time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// FetchIndexByIDxGoRoleDeletedAt  获取多个内容
func (obj *_GoRoleMgr) FetchIndexByIDxGoRoleDeletedAt(deletedAt time.Time) (results []*GoRole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRole{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
