# go-admin-team 公共代码库

### 功能

- [x] log 组件
- [x] 缓存(支持 memory)
- [x] 队列(支持 memory)
- [x] 日志写入 writer
- [x] 日志插件 logrus
- [x] 日志插件 zap
- [x] 大文件分割写入

- 暂时移除对 redis 的支持

---

## 配置文件读取

使用 `Setup` 函数初始化配置：

```go
package main

// import "github.com/GoAdminTeam/go-admin-core/logger"

source := config.FileSource("config.json")
config.Setup(source, func() {
    // 回调函数逻辑
})
```

## 日志记录

使用 Log 和 Logf 方法记录日志：

```go
package main

// import "github.com/GoAdminTeam/go-admin-core/logger"

logger := logger.NewDefaultLogger()
logger.Log(logger.INFO, "This is an info message")
logger.Logf(logger.ERROR, "This is an error message: %s", "error details")
```
