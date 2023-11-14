package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoAccountMgr struct {
	*_BaseMgr
}

// GoAccountMgr open func
func GoAccountMgr(db *gorm.DB) *_GoAccountMgr {
	if db == nil {
		panic(fmt.Errorf("GoAccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoAccountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_account"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoAccountMgr) Debug() *_GoAccountMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoAccountMgr) GetTableName() string {
	return "go_account"
}

// Reset 重置gorm会话
func (obj *_GoAccountMgr) Reset() *_GoAccountMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoAccountMgr) Get() (result GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoAccountMgr) Gets() (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoAccountMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoAccountMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户名
func (obj *_GoAccountMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_GoAccountMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithLevel level获取 账号类型，1为管理员，2为商户
func (obj *_GoAccountMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithRoleID role_id获取 (角色id)
func (obj *_GoAccountMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithMobile mobile获取 手机号码
func (obj *_GoAccountMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithStatus status获取 状态1：为正常 -1：为冻结
func (obj *_GoAccountMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoAccountMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoAccountMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoAccountMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoAccountMgr) GetByOption(opts ...Option) (result GoAccount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoAccountMgr) GetByOptions(opts ...Option) (results []*GoAccount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoAccountMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoAccount, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where(options.query)
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
func (obj *_GoAccountMgr) GetFromID(id uint64) (result GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoAccountMgr) GetBatchFromID(ids []uint64) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_GoAccountMgr) GetFromUsername(username string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_GoAccountMgr) GetBatchFromUsername(usernames []string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_GoAccountMgr) GetFromPassword(password string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_GoAccountMgr) GetBatchFromPassword(passwords []string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 账号类型，1为管理员，2为商户
func (obj *_GoAccountMgr) GetFromLevel(level int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 账号类型，1为管理员，2为商户
func (obj *_GoAccountMgr) GetBatchFromLevel(levels []int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 (角色id)
func (obj *_GoAccountMgr) GetFromRoleID(roleID int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找 (角色id)
func (obj *_GoAccountMgr) GetBatchFromRoleID(roleIDs []int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号码
func (obj *_GoAccountMgr) GetFromMobile(mobile string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号码
func (obj *_GoAccountMgr) GetBatchFromMobile(mobiles []string) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态1：为正常 -1：为冻结
func (obj *_GoAccountMgr) GetFromStatus(status int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态1：为正常 -1：为冻结
func (obj *_GoAccountMgr) GetBatchFromStatus(statuss []int) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoAccountMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoAccountMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoAccountMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoAccountMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoAccountMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoAccountMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoAccountMgr) FetchByPrimaryKey(id uint64) (result GoAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoAccount{}).Where("`id` = ?", id).First(&result).Error

	return
}
