package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoLoginLogMgr struct {
	*_BaseMgr
}

// GoLoginLogMgr open func
func GoLoginLogMgr(db *gorm.DB) *_GoLoginLogMgr {
	if db == nil {
		panic(fmt.Errorf("GoLoginLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoLoginLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_login_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoLoginLogMgr) Debug() *_GoLoginLogMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoLoginLogMgr) GetTableName() string {
	return "go_login_log"
}

// Reset 重置gorm会话
func (obj *_GoLoginLogMgr) Reset() *_GoLoginLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoLoginLogMgr) Get() (result GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoLoginLogMgr) Gets() (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoLoginLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoLoginLogMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAccountID account_id获取 账号id
func (obj *_GoLoginLogMgr) WithAccountID(accountID int) Option {
	return optionFunc(func(o *options) { o.query["account_id"] = accountID })
}

// WithType type获取 类型1：为总后台用户，2：为合作商用户
func (obj *_GoLoginLogMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithUsername username获取 用户名
func (obj *_GoLoginLogMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithRoleID role_id获取 角色id
func (obj *_GoLoginLogMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithIP ip获取 登录ip
func (obj *_GoLoginLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithAddress address获取 登录地址
func (obj *_GoLoginLogMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoLoginLogMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoLoginLogMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoLoginLogMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoLoginLogMgr) GetByOption(opts ...Option) (result GoLoginLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoLoginLogMgr) GetByOptions(opts ...Option) (results []*GoLoginLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoLoginLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoLoginLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where(options.query)
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
func (obj *_GoLoginLogMgr) GetFromID(id uint64) (result GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoLoginLogMgr) GetBatchFromID(ids []uint64) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAccountID 通过account_id获取内容 账号id
func (obj *_GoLoginLogMgr) GetFromAccountID(accountID int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`account_id` = ?", accountID).Find(&results).Error

	return
}

// GetBatchFromAccountID 批量查找 账号id
func (obj *_GoLoginLogMgr) GetBatchFromAccountID(accountIDs []int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`account_id` IN (?)", accountIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型1：为总后台用户，2：为合作商用户
func (obj *_GoLoginLogMgr) GetFromType(_type int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型1：为总后台用户，2：为合作商用户
func (obj *_GoLoginLogMgr) GetBatchFromType(_types []int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_GoLoginLogMgr) GetFromUsername(username string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_GoLoginLogMgr) GetBatchFromUsername(usernames []string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_GoLoginLogMgr) GetFromRoleID(roleID int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找 角色id
func (obj *_GoLoginLogMgr) GetBatchFromRoleID(roleIDs []int) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 登录ip
func (obj *_GoLoginLogMgr) GetFromIP(ip string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 登录ip
func (obj *_GoLoginLogMgr) GetBatchFromIP(ips []string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 登录地址
func (obj *_GoLoginLogMgr) GetFromAddress(address string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 登录地址
func (obj *_GoLoginLogMgr) GetBatchFromAddress(addresss []string) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoLoginLogMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoLoginLogMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoLoginLogMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoLoginLogMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoLoginLogMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoLoginLogMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoLoginLogMgr) FetchByPrimaryKey(id uint64) (result GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxGoLoginLogCreatedAt  获取多个内容
func (obj *_GoLoginLogMgr) FetchIndexByIDxGoLoginLogCreatedAt(createdAt time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// FetchIndexByIDxGoLoginLogDeletedAt  获取多个内容
func (obj *_GoLoginLogMgr) FetchIndexByIDxGoLoginLogDeletedAt(deletedAt time.Time) (results []*GoLoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoLoginLog{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
