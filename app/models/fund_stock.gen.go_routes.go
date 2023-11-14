package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _GoRoutesMgr struct {
	*_BaseMgr
}

// GoRoutesMgr open func
func GoRoutesMgr(db *gorm.DB) *_GoRoutesMgr {
	if db == nil {
		panic(fmt.Errorf("GoRoutesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_GoRoutesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("go_routes"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_GoRoutesMgr) Debug() *_GoRoutesMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_GoRoutesMgr) GetTableName() string {
	return "go_routes"
}

// Reset 重置gorm会话
func (obj *_GoRoutesMgr) Reset() *_GoRoutesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_GoRoutesMgr) Get() (result GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_GoRoutesMgr) Gets() (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_GoRoutesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_GoRoutesMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSort sort获取 排序
func (obj *_GoRoutesMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithType type获取 page-页面   api-接口
func (obj *_GoRoutesMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithIsMenu is_menu获取 是否根菜单1-是 0-否
func (obj *_GoRoutesMgr) WithIsMenu(isMenu int) Option {
	return optionFunc(func(o *options) { o.query["is_menu"] = isMenu })
}

// WithRoute route获取 访问路由地址
func (obj *_GoRoutesMgr) WithRoute(route string) Option {
	return optionFunc(func(o *options) { o.query["route"] = route })
}

// WithComponent component获取 页面组件地址
func (obj *_GoRoutesMgr) WithComponent(component string) Option {
	return optionFunc(func(o *options) { o.query["component"] = component })
}

// WithName name获取 路由名称
func (obj *_GoRoutesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithIcon icon获取 icon图标
func (obj *_GoRoutesMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithParentID parent_id获取 上级id
func (obj *_GoRoutesMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCreateBy create_by获取 创建者
func (obj *_GoRoutesMgr) WithCreateBy(createBy int) Option {
	return optionFunc(func(o *options) { o.query["create_by"] = createBy })
}

// WithStatus status获取 1-已启用   0-未启用
func (obj *_GoRoutesMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreatedAt created_at获取 创建时间
func (obj *_GoRoutesMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_GoRoutesMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 删除时间
func (obj *_GoRoutesMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// GetByOption 功能选项模式获取
func (obj *_GoRoutesMgr) GetByOption(opts ...Option) (result GoRoutes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_GoRoutesMgr) GetByOptions(opts ...Option) (results []*GoRoutes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_GoRoutesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]GoRoutes, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where(options.query)
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
func (obj *_GoRoutesMgr) GetFromID(id uint64) (result GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_GoRoutesMgr) GetBatchFromID(ids []uint64) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_GoRoutesMgr) GetFromSort(sort int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_GoRoutesMgr) GetBatchFromSort(sorts []int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 page-页面   api-接口
func (obj *_GoRoutesMgr) GetFromType(_type string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 page-页面   api-接口
func (obj *_GoRoutesMgr) GetBatchFromType(_types []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromIsMenu 通过is_menu获取内容 是否根菜单1-是 0-否
func (obj *_GoRoutesMgr) GetFromIsMenu(isMenu int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`is_menu` = ?", isMenu).Find(&results).Error

	return
}

// GetBatchFromIsMenu 批量查找 是否根菜单1-是 0-否
func (obj *_GoRoutesMgr) GetBatchFromIsMenu(isMenus []int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`is_menu` IN (?)", isMenus).Find(&results).Error

	return
}

// GetFromRoute 通过route获取内容 访问路由地址
func (obj *_GoRoutesMgr) GetFromRoute(route string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`route` = ?", route).Find(&results).Error

	return
}

// GetBatchFromRoute 批量查找 访问路由地址
func (obj *_GoRoutesMgr) GetBatchFromRoute(routes []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`route` IN (?)", routes).Find(&results).Error

	return
}

// GetFromComponent 通过component获取内容 页面组件地址
func (obj *_GoRoutesMgr) GetFromComponent(component string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`component` = ?", component).Find(&results).Error

	return
}

// GetBatchFromComponent 批量查找 页面组件地址
func (obj *_GoRoutesMgr) GetBatchFromComponent(components []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`component` IN (?)", components).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 路由名称
func (obj *_GoRoutesMgr) GetFromName(name string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 路由名称
func (obj *_GoRoutesMgr) GetBatchFromName(names []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 icon图标
func (obj *_GoRoutesMgr) GetFromIcon(icon string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`icon` = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量查找 icon图标
func (obj *_GoRoutesMgr) GetBatchFromIcon(icons []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`icon` IN (?)", icons).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 上级id
func (obj *_GoRoutesMgr) GetFromParentID(parentID int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 上级id
func (obj *_GoRoutesMgr) GetBatchFromParentID(parentIDs []int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCreateBy 通过create_by获取内容 创建者
func (obj *_GoRoutesMgr) GetFromCreateBy(createBy int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`create_by` = ?", createBy).Find(&results).Error

	return
}

// GetBatchFromCreateBy 批量查找 创建者
func (obj *_GoRoutesMgr) GetBatchFromCreateBy(createBys []int) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`create_by` IN (?)", createBys).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 1-已启用   0-未启用
func (obj *_GoRoutesMgr) GetFromStatus(status string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 1-已启用   0-未启用
func (obj *_GoRoutesMgr) GetBatchFromStatus(statuss []string) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建时间
func (obj *_GoRoutesMgr) GetFromCreatedAt(createdAt time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找 创建时间
func (obj *_GoRoutesMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_GoRoutesMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 更新时间
func (obj *_GoRoutesMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 删除时间
func (obj *_GoRoutesMgr) GetFromDeletedAt(deletedAt time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找 删除时间
func (obj *_GoRoutesMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_GoRoutesMgr) FetchByPrimaryKey(id uint64) (result GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIDxGoRoutesCreatedAt  获取多个内容
func (obj *_GoRoutesMgr) FetchIndexByIDxGoRoutesCreatedAt(createdAt time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// FetchIndexByIDxGoRoutesDeletedAt  获取多个内容
func (obj *_GoRoutesMgr) FetchIndexByIDxGoRoutesDeletedAt(deletedAt time.Time) (results []*GoRoutes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(GoRoutes{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
