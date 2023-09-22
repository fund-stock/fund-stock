package models

import (
	"time"
)

// AppPassLimitHost app服务器访问白名单表
type AppPassLimitHost struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"-"`        // app服务器访问白名单表主键自增ID
	PassHost   string    `gorm:"column:pass_host" json:"passHost"`     // 活跃的域名
	CreateDate time.Time `gorm:"column:create_date" json:"createDate"` // 创建日期
	UpdateDate time.Time `gorm:"column:update_date" json:"updateDate"` // 修改日期
	Creator    string    `gorm:"column:creator" json:"creator"`        // 添加人ID
	Updater    string    `gorm:"column:updater" json:"updater"`        // 更新人ID
	ItemCode   string    `gorm:"column:item_code" json:"itemCode"`     // 关联项目编码
	IsShow     int       `gorm:"column:is_show" json:"isShow"`         // 是否展示 1是 0否
	IsDelete   int       `gorm:"column:is_delete" json:"isDelete"`     // 是否删除 1是 0否
	PassLevel  int       `gorm:"column:pass_level" json:"passLevel"`   // 活跃等级
}

// TableName get sql table name.获取数据库表名
func (m *AppPassLimitHost) TableName() string {
	return "app_pass_limit_host"
}

// AppPassLimitHostColumns get sql column name.获取数据库列名
var AppPassLimitHostColumns = struct {
	ID         string
	PassHost   string
	CreateDate string
	UpdateDate string
	Creator    string
	Updater    string
	ItemCode   string
	IsShow     string
	IsDelete   string
	PassLevel  string
}{
	ID:         "id",
	PassHost:   "pass_host",
	CreateDate: "create_date",
	UpdateDate: "update_date",
	Creator:    "creator",
	Updater:    "updater",
	ItemCode:   "item_code",
	IsShow:     "is_show",
	IsDelete:   "is_delete",
	PassLevel:  "pass_level",
}

// GoAccount [...]
type GoAccount struct {
	ID             uint64         `gorm:"primaryKey;column:id" json:"-"`
	GoOperationLog GoOperationLog `gorm:"joinForeignKey:id;foreignKey:id;references:ID" json:"goOperationLogList"` // 操作记录表
	Username       string         `gorm:"column:username" json:"username"`                                         // 用户名
	Password       string         `gorm:"column:password" json:"password"`                                         // 密码
	Level          int            `gorm:"column:level" json:"level"`                                               // 账号类型，1为管理员，2为商户
	RoleID         int            `gorm:"column:role_id" json:"roleId"`                                            // (角色id)
	Mobile         string         `gorm:"column:mobile" json:"mobile"`                                             // 手机号码
	Status         int            `gorm:"column:status" json:"status"`                                             // 状态1：为正常 -1：为冻结
	CreatedAt      time.Time      `gorm:"column:created_at" json:"createdAt"`                                      // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updatedAt"`                                      // 更新时间
	DeletedAt      time.Time      `gorm:"column:deleted_at" json:"deletedAt"`                                      // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *GoAccount) TableName() string {
	return "go_account"
}

// GoAccountColumns get sql column name.获取数据库列名
var GoAccountColumns = struct {
	ID        string
	Username  string
	Password  string
	Level     string
	RoleID    string
	Mobile    string
	Status    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Username:  "username",
	Password:  "password",
	Level:     "level",
	RoleID:    "role_id",
	Mobile:    "mobile",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// GoFund [...]
type GoFund struct {
	ID        int64   `gorm:"primaryKey;column:id" json:"-"`
	ProductID string  `gorm:"column:product_id" json:"productId"` // 产品id
	Code      string  `gorm:"column:code" json:"code"`            // 代码
	Name      string  `gorm:"column:name" json:"name"`            // 名称
	Amount    float64 `gorm:"column:amount" json:"amount"`        // 金额
	Nav       float64 `gorm:"column:nav" json:"nav"`              // 最新净值
	Status    int     `gorm:"column:status" json:"status"`        // 状态：0-未启用，1-已启用
	CreateAt  int64   `gorm:"column:create_at" json:"createAt"`   // 创建时间
	UpdateAt  int64   `gorm:"column:update_at" json:"updateAt"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *GoFund) TableName() string {
	return "go_fund"
}

// GoFundColumns get sql column name.获取数据库列名
var GoFundColumns = struct {
	ID        string
	ProductID string
	Code      string
	Name      string
	Amount    string
	Nav       string
	Status    string
	CreateAt  string
	UpdateAt  string
}{
	ID:        "id",
	ProductID: "product_id",
	Code:      "code",
	Name:      "name",
	Amount:    "amount",
	Nav:       "nav",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// GoFundDay [...]
type GoFundDay struct {
	ID       int64     `gorm:"primaryKey;column:id" json:"-"`
	Code     string    `gorm:"column:code" json:"code"`          // 代码
	Name     string    `gorm:"column:name" json:"name"`          // 名称
	Amount   float64   `gorm:"column:amount" json:"amount"`      // 金额
	Nav      float64   `gorm:"column:nav" json:"nav"`            // 最新净值
	DayTs    int64     `gorm:"column:day_ts" json:"dayTs"`       // 当天的时间戳
	DayAt    time.Time `gorm:"column:day_at" json:"dayAt"`       // 当天的时间
	CreateAt int64     `gorm:"column:create_at" json:"createAt"` // 创建时间
	UpdateAt int64     `gorm:"column:update_at" json:"updateAt"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *GoFundDay) TableName() string {
	return "go_fund_day"
}

// GoFundDayColumns get sql column name.获取数据库列名
var GoFundDayColumns = struct {
	ID       string
	Code     string
	Name     string
	Amount   string
	Nav      string
	DayTs    string
	DayAt    string
	CreateAt string
	UpdateAt string
}{
	ID:       "id",
	Code:     "code",
	Name:     "name",
	Amount:   "amount",
	Nav:      "nav",
	DayTs:    "day_ts",
	DayAt:    "day_at",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// GoLoginLog [...]
type GoLoginLog struct {
	ID        uint64    `gorm:"primaryKey;column:id" json:"-"`
	AccountID int       `gorm:"column:account_id" json:"accountId"` // 账号id
	Type      int       `gorm:"column:type" json:"type"`            // 类型1：为总后台用户，2：为合作商用户
	Username  string    `gorm:"column:username" json:"username"`    // 用户名
	RoleID    int       `gorm:"column:role_id" json:"roleId"`       // 角色id
	IP        string    `gorm:"column:ip" json:"ip"`                // 登录ip
	Address   string    `gorm:"column:address" json:"address"`      // 登录地址
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *GoLoginLog) TableName() string {
	return "go_login_log"
}

// GoLoginLogColumns get sql column name.获取数据库列名
var GoLoginLogColumns = struct {
	ID        string
	AccountID string
	Type      string
	Username  string
	RoleID    string
	IP        string
	Address   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	AccountID: "account_id",
	Type:      "type",
	Username:  "username",
	RoleID:    "role_id",
	IP:        "ip",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// GoOperationLog 操作记录表
type GoOperationLog struct {
	ID        uint64    `gorm:"primaryKey;column:id" json:"-"`
	Type      int       `gorm:"column:type" json:"type"`            // 类型1：为总后台用户，2：为合作商用户
	AccountID int       `gorm:"column:account_id" json:"accountId"` // 账号id
	Content   string    `gorm:"column:content" json:"content"`      // 操作内容
	IP        string    `gorm:"column:ip" json:"ip"`                // ip地址
	Address   string    `gorm:"column:address" json:"address"`      // 操作地址
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *GoOperationLog) TableName() string {
	return "go_operation_log"
}

// GoOperationLogColumns get sql column name.获取数据库列名
var GoOperationLogColumns = struct {
	ID        string
	Type      string
	AccountID string
	Content   string
	IP        string
	Address   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Type:      "type",
	AccountID: "account_id",
	Content:   "content",
	IP:        "ip",
	Address:   "address",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// GoRole [...]
type GoRole struct {
	ID        uint64    `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`            // 角色名称
	Routes    string    `gorm:"column:routes" json:"routes"`        // 路由id,该角色所具有的路由
	Desc      string    `gorm:"column:desc" json:"desc"`            // 角色描述
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *GoRole) TableName() string {
	return "go_role"
}

// GoRoleColumns get sql column name.获取数据库列名
var GoRoleColumns = struct {
	ID        string
	Name      string
	Routes    string
	Desc      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Routes:    "routes",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// GoRoutes [...]
type GoRoutes struct {
	ID        uint64    `gorm:"primaryKey;column:id" json:"-"`
	Sort      int       `gorm:"column:sort" json:"sort"`            // 排序
	Type      string    `gorm:"column:type" json:"type"`            // page-页面   api-接口
	IsMenu    int       `gorm:"column:is_menu" json:"isMenu"`       // 是否根菜单1-是 0-否
	Route     string    `gorm:"column:route" json:"route"`          // 访问路由地址
	Component string    `gorm:"column:component" json:"component"`  // 页面组件地址
	Name      string    `gorm:"column:name" json:"name"`            // 路由名称
	Icon      string    `gorm:"column:icon" json:"icon"`            // icon图标
	ParentID  int       `gorm:"column:parent_id" json:"parentId"`   // 上级id
	CreateBy  int       `gorm:"column:create_by" json:"createBy"`   // 创建者
	Status    string    `gorm:"column:status" json:"status"`        // 1-已启用   0-未启用
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deletedAt"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *GoRoutes) TableName() string {
	return "go_routes"
}

// GoRoutesColumns get sql column name.获取数据库列名
var GoRoutesColumns = struct {
	ID        string
	Sort      string
	Type      string
	IsMenu    string
	Route     string
	Component string
	Name      string
	Icon      string
	ParentID  string
	CreateBy  string
	Status    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Sort:      "sort",
	Type:      "type",
	IsMenu:    "is_menu",
	Route:     "route",
	Component: "component",
	Name:      "name",
	Icon:      "icon",
	ParentID:  "parent_id",
	CreateBy:  "create_by",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// GoStock [...]
type GoStock struct {
	ID        int64   `gorm:"primaryKey;column:id" json:"-"`
	ProductID string  `gorm:"column:product_id" json:"productId"` // 产品id
	Code      string  `gorm:"column:code" json:"code"`            // 代码
	Name      string  `gorm:"column:name" json:"name"`            // 名称
	Amount    float64 `gorm:"column:amount" json:"amount"`        // 金额
	Nav       float64 `gorm:"column:nav" json:"nav"`              // 最新净值
	Status    int     `gorm:"column:status" json:"status"`        // 状态：0-未启用，1-已启用
	CreateAt  int64   `gorm:"column:create_at" json:"createAt"`   // 创建时间
	UpdateAt  int64   `gorm:"column:update_at" json:"updateAt"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *GoStock) TableName() string {
	return "go_stock"
}

// GoStockColumns get sql column name.获取数据库列名
var GoStockColumns = struct {
	ID        string
	ProductID string
	Code      string
	Name      string
	Amount    string
	Nav       string
	Status    string
	CreateAt  string
	UpdateAt  string
}{
	ID:        "id",
	ProductID: "product_id",
	Code:      "code",
	Name:      "name",
	Amount:    "amount",
	Nav:       "nav",
	Status:    "status",
	CreateAt:  "create_at",
	UpdateAt:  "update_at",
}

// GoStockDay [...]
type GoStockDay struct {
	ID       int64     `gorm:"primaryKey;column:id" json:"-"`
	Code     string    `gorm:"column:code" json:"code"`          // 代码
	Name     string    `gorm:"column:name" json:"name"`          // 名称
	Amount   float64   `gorm:"column:amount" json:"amount"`      // 金额
	Nav      float64   `gorm:"column:nav" json:"nav"`            // 最新净值
	DayTs    int64     `gorm:"column:day_ts" json:"dayTs"`       // 当天的时间戳
	DayAt    time.Time `gorm:"column:day_at" json:"dayAt"`       // 当天的时间
	CreateAt int64     `gorm:"column:create_at" json:"createAt"` // 创建时间
	UpdateAt int64     `gorm:"column:update_at" json:"updateAt"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *GoStockDay) TableName() string {
	return "go_stock_day"
}

// GoStockDayColumns get sql column name.获取数据库列名
var GoStockDayColumns = struct {
	ID       string
	Code     string
	Name     string
	Amount   string
	Nav      string
	DayTs    string
	DayAt    string
	CreateAt string
	UpdateAt string
}{
	ID:       "id",
	Code:     "code",
	Name:     "name",
	Amount:   "amount",
	Nav:      "nav",
	DayTs:    "day_ts",
	DayAt:    "day_at",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// GoStockLog 股票实时数据表
type GoStockLog struct {
	ID       int64     `gorm:"primaryKey;column:id" json:"-"`
	Code     string    `gorm:"column:code" json:"code"`          // 代码
	Name     string    `gorm:"column:name" json:"name"`          // 名称
	Amount   float64   `gorm:"column:amount" json:"amount"`      // 金额
	Nav      float64   `gorm:"column:nav" json:"nav"`            // 最新净值
	DayTs    int64     `gorm:"column:day_ts" json:"dayTs"`       // 当天的时间戳
	DayAt    time.Time `gorm:"column:day_at" json:"dayAt"`       // 当天的时间
	CreateAt int64     `gorm:"column:create_at" json:"createAt"` // 创建时间
	UpdateAt int64     `gorm:"column:update_at" json:"updateAt"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *GoStockLog) TableName() string {
	return "go_stock_log"
}

// GoStockLogColumns get sql column name.获取数据库列名
var GoStockLogColumns = struct {
	ID       string
	Code     string
	Name     string
	Amount   string
	Nav      string
	DayTs    string
	DayAt    string
	CreateAt string
	UpdateAt string
}{
	ID:       "id",
	Code:     "code",
	Name:     "name",
	Amount:   "amount",
	Nav:      "nav",
	DayTs:    "day_ts",
	DayAt:    "day_at",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// SystemEmailRoutingNode [...]
type SystemEmailRoutingNode struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"-"`                  // 邮箱路由节点配置表ID
	Email           string    `gorm:"column:email" json:"email"`                      // 发信人邮箱
	Channel         string    `gorm:"column:channel" json:"channel"`                  // 三方渠道
	ChannelUsername string    `gorm:"column:channel_username" json:"channelUsername"` // 渠道用户名
	ChannelPassword string    `gorm:"column:channel_password" json:"channelPassword"` // 渠道密码
	ChannelHost     string    `gorm:"column:channel_host" json:"channelHost"`         // 渠道服务器
	ChannelPort     int       `gorm:"column:channel_port" json:"channelPort"`         // 渠道端口号
	AvailableNumber int       `gorm:"column:available_number" json:"availableNumber"` // 剩余可用条数
	MaxNumber       int       `gorm:"column:max_number" json:"maxNumber"`             // 每日最多使用条数
	RateSuccess     float64   `gorm:"column:rate_success" json:"rateSuccess"`         // 成功率
	RateFail        float64   `gorm:"column:rate_fail" json:"rateFail"`               // 失败率
	Sort            int       `gorm:"column:sort" json:"sort"`                        // 优先排序
	Valid           int       `gorm:"column:valid" json:"valid"`                      // 有效状态 1有效 0失效
	CreatDate       time.Time `gorm:"column:creat_date" json:"creatDate"`             // 创建时间
	UpdateDate      time.Time `gorm:"column:update_date" json:"updateDate"`           // 修改时间
	Creator         string    `gorm:"column:creator" json:"creator"`                  // 创建人
	Updater         string    `gorm:"column:updater" json:"updater"`                  // 修改人
	Remarks         string    `gorm:"column:remarks" json:"remarks"`                  // 备注
	ItemCode        string    `gorm:"column:item_code" json:"itemCode"`               // 应用编码
	Type            int       `gorm:"column:type" json:"type"`                        // 类型：1验证码通知，2营销邮件
}

// TableName get sql table name.获取数据库表名
func (m *SystemEmailRoutingNode) TableName() string {
	return "system_email_routing_node"
}

// SystemEmailRoutingNodeColumns get sql column name.获取数据库列名
var SystemEmailRoutingNodeColumns = struct {
	ID              string
	Email           string
	Channel         string
	ChannelUsername string
	ChannelPassword string
	ChannelHost     string
	ChannelPort     string
	AvailableNumber string
	MaxNumber       string
	RateSuccess     string
	RateFail        string
	Sort            string
	Valid           string
	CreatDate       string
	UpdateDate      string
	Creator         string
	Updater         string
	Remarks         string
	ItemCode        string
	Type            string
}{
	ID:              "id",
	Email:           "email",
	Channel:         "channel",
	ChannelUsername: "channel_username",
	ChannelPassword: "channel_password",
	ChannelHost:     "channel_host",
	ChannelPort:     "channel_port",
	AvailableNumber: "available_number",
	MaxNumber:       "max_number",
	RateSuccess:     "rate_success",
	RateFail:        "rate_fail",
	Sort:            "sort",
	Valid:           "valid",
	CreatDate:       "creat_date",
	UpdateDate:      "update_date",
	Creator:         "creator",
	Updater:         "updater",
	Remarks:         "remarks",
	ItemCode:        "item_code",
	Type:            "type",
}

// SystemEmailSendLog [...]
type SystemEmailSendLog struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"-"`                  // 邮箱发送记录表ID
	SendingMailbox  string    `gorm:"column:sending_mailbox" json:"sendingMailbox"`   // 发件邮箱账号
	ReceiveEmail    string    `gorm:"column:receive_email" json:"receiveEmail"`       // 收件箱账号
	SendTotal       int       `gorm:"column:send_total" json:"sendTotal"`             // 发送数量
	AvailableNumber int       `gorm:"column:available_number" json:"availableNumber"` // 剩余可用条数
	Status          int       `gorm:"column:status" json:"status"`                    // 发送状态：1成功 0失败
	CreatDate       time.Time `gorm:"column:creat_date" json:"creatDate"`             // 创建时间
	UpdateDate      time.Time `gorm:"column:update_date" json:"updateDate"`           // 修改时间
	Creator         string    `gorm:"column:creator" json:"creator"`                  // 创建人
	Updater         string    `gorm:"column:updater" json:"updater"`                  // 修改人
	Remarks         string    `gorm:"column:remarks" json:"remarks"`                  // 备注
	Body            string    `gorm:"column:body" json:"body"`                        // 发送内容
}

// TableName get sql table name.获取数据库表名
func (m *SystemEmailSendLog) TableName() string {
	return "system_email_send_log"
}

// SystemEmailSendLogColumns get sql column name.获取数据库列名
var SystemEmailSendLogColumns = struct {
	ID              string
	SendingMailbox  string
	ReceiveEmail    string
	SendTotal       string
	AvailableNumber string
	Status          string
	CreatDate       string
	UpdateDate      string
	Creator         string
	Updater         string
	Remarks         string
	Body            string
}{
	ID:              "id",
	SendingMailbox:  "sending_mailbox",
	ReceiveEmail:    "receive_email",
	SendTotal:       "send_total",
	AvailableNumber: "available_number",
	Status:          "status",
	CreatDate:       "creat_date",
	UpdateDate:      "update_date",
	Creator:         "creator",
	Updater:         "updater",
	Remarks:         "remarks",
	Body:            "body",
}
