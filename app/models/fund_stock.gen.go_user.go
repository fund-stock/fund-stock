package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoUserMgr struct {
	*_BaseMgr
}

// GoUserMgr open func
func GoUserMgr(db *gorm.DB) *_GoUserMgr {
	if db == nil {
		panic(fmt.Errorf("GoUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoUserMgr) Debug() *_GoUserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoUserMgr) GetTableName() string {
	return "go_user"
}

// Reset 重置gorm会话
func (obj *_GoUserMgr) Reset() *_GoUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoUserMgr) Get() (result GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoUserMgr) Gets() (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoUserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOpenid openid获取 openid
func (obj *_GoUserMgr) WithOpenid(openid string) Option {
	return optionFunc(func(o *options) { o.query["openid"] = openid })
}

// WithNickname nickname获取 昵称
func (obj *_GoUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithAvatar avatar获取 头像
func (obj *_GoUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithUsername username获取 用户名
func (obj *_GoUserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_GoUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithMobile mobile获取 手机号
func (obj *_GoUserMgr) WithMobile(mobile int64) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithStatus status获取 状态1：为正常 -1：为冻结
func (obj *_GoUserMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoUserMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoUserMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoUserMgr) GetByOption(opts ...Option) (result GoUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoUserMgr) GetByOptions(opts ...Option) (results []*GoUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where(options.query)
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
func (obj *_GoUserMgr) GetFromID(id uint64) (result GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoUserMgr) GetBatchFromID(ids []uint64) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOpenid 通过openid获取内容 openid
func (obj *_GoUserMgr) GetFromOpenid(openid string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`openid` = ?", openid).Find(&results).Error

	return
}

// GetBatchFromOpenid 批量查找 openid
func (obj *_GoUserMgr) GetBatchFromOpenid(openids []string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`openid` IN (?)", openids).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_GoUserMgr) GetFromNickname(nickname string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_GoUserMgr) GetBatchFromNickname(nicknames []string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_GoUserMgr) GetFromAvatar(avatar string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_GoUserMgr) GetBatchFromAvatar(avatars []string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_GoUserMgr) GetFromUsername(username string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_GoUserMgr) GetBatchFromUsername(usernames []string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_GoUserMgr) GetFromPassword(password string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_GoUserMgr) GetBatchFromPassword(passwords []string) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_GoUserMgr) GetFromMobile(mobile int64) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号
func (obj *_GoUserMgr) GetBatchFromMobile(mobiles []int64) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态1：为正常 -1：为冻结
func (obj *_GoUserMgr) GetFromStatus(status int) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态1：为正常 -1：为冻结
func (obj *_GoUserMgr) GetBatchFromStatus(statuss []int) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoUserMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoUserMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoUserMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoUserMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoUserMgr) FetchByPrimaryKey(id uint64) (result GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxGoAccountCreatedAt  获取多个内容
func (obj *_GoUserMgr) FetchIndexByIDxGoAccountCreatedAt(createdAt time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// FetchIndexByIDxGoAccountDeletedAt  获取多个内容
func (obj *_GoUserMgr) FetchIndexByIDxGoAccountDeletedAt(deletedAt time.Time) (results []*GoUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoUser{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
