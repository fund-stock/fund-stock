# FundStock

## 引入的库-

> [gin-gonic/gin](https://github.com/gin-gonic/gin)   【Gin框架】
>
> [go-playground/validator](https://github.com/go-playground/validator)   【validator表单验证器】
>
> [concurrent-map](https://github.com/orcaman/concurrent-map)   【concurrent-map替换原生map解决并发读写】
>
> [gorm.io/gorm](https://gorm.io/gorm)   【Gorm数据查询工具】
>
> [go-redis/redis](https://github.com/go-redis/redis)   【go-redis缓存】
>
> [spf13/viper](https://github.com/spf13/viper)   【配置读取工具】
>
> [uber-go/zap](https://github.com/uber-go/zap)   【zap日志库】
>
> 日志TrackId链路追踪，包含mysql语句
>
---
# 目录结构

```text
.
├── app
│     ├── controllers
│     │     ├── client
│     │     │     ├── v1
│     │     │     ├── v2
│     │     │     └── v3
│     │     ├── email
│     │     │     └── v1
│     │     └── web
│     │         └── v1
│     ├── middlewares
│     │     ├── common
│     │     ├── v1
│     │     ├── v2
│     │     └── v3
│     ├── migration
│     ├── models
│     ├── requests
│     └── response
├── bin
├── bootstrap
├── config
├── logs
│     ├── binary-client
│     ├── binary-fund
│     ├── binary-stock
├── pkg
│     ├── abnormal
│     ├── address
│     ├── config
│     ├── echo
│     │     └── code
│     │         ├── en
│     │         ├── insa
│     │         └── zh
│     ├── email
│     ├── helpers
│     ├── larkbot
│     ├── logger
│     │     └── zapgorm2
│     ├── mysql
│     ├── notice
│     │     ├── bark
│     │     ├── mail
│     │     └── telegram
│     ├── output
│     ├── pprof
│     ├── recaptcha
│     ├── redis
│     ├── request
│     ├── utils
│     ├── validate
│     ├── ws_server
│     ├── wss-heart-beating
│     └── xjson
├── routes
│     ├── client_v1
│     ├── client_v2
│     └── client_v3
├── serve
│     ├── binary-api
│     ├── binary-email
│     │     ├── common
│     │     ├── config
│     │     └── queue
│     ├── binary-fund
│     │     ├── client
│     │     ├── config
│     │     └── params
│     ├── binary-stock
│     │     ├── client
│     │     ├── config
│     │     ├── params
│     │     └── response
│     │         └── qtimg
│     └── binary-timer
│         ├── task_plan
│         └── task_route
└── test

81 directories
```
---
- [x] `2023-09-25` 添加 `bark` 通知功能
- [x] `2023-09-26` 添加 `实时数据获取`


# TODO
- [ ] 分时任务
- [x] 盯盘助手
- [ ] 研报股评
- [ ] 每日监控
- [ ] 微信提醒
- [ ] 玩转组织
- [ ] AI智能