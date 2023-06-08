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
> [gostudys/huobi_contract_golang](https://github.com/gostudys/huobi_contract_golang)   【火币网封装库】
>
> 日志TrackId链路追踪，包含mysql语句
> 
---


### 清除es数据

```bash
curl -XDELETE http://10.10.10.10:9200/kline_log
```