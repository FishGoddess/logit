# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/logit/actions/workflows/test.yml/badge.svg)

**logit** 是一个基于级别控制的高性能纯结构化日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 独特的日志模块设计，使用 appender 和 writer 装载特定的模块，实现扩展功能。
* 支持日志级别控制，一共有五个日志级别，分别是 debug，info，warn，error，print 和 off。
* 支持键值对形式的结构化日志记录，同时对格式化操作也有支持。
* 支持以 Text/Json 形式输出日志信息，方便对日志进行解析。
* 支持异步回写日志，提供高性能缓冲写出器模块，减少 IO 的访问次数。
* 提供调优使用的全局配置，对一些高级配置更贴合实际业务的需求。
* 分级别追加日志数据，分级别写出日志数据，推荐将 error 级别的日志单独处理和存储。
* 加入 Context 机制，更优雅地使用日志，并支持业务域划分。
* 支持拦截器模式，可以从 context 注入外部常量或变量，简化日志输出流程。
* 支持错误监控，可以很方便地进行错误统计和告警。
* 支持日志按大小自动分割，并支持按照时间和数量自动清理。
* 支持多种配置文件序列化成 option，比如 json/yaml/toml/bson，然后创建日志记录器。

_原谅我，先建了个组织，后面发现没啥用，又转移回个人仓库了。。。_

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

```bash
$ go get -u github.com/FishGoddess/logit
```

### 📖 参考案例

```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/support/global"
)

func main() {
	// Create a new logger for use.
	// Default level is debug, so all logs will be logged.
	// Invoke Close() isn't necessary in all situations.
	// If logger's writer has buffer or something like that, it's better to invoke Close() for syncing buffer or something else.
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want.
	// Remember, logs will be ignored if their level is smaller than logger's level.
	// Log() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").Log()
	logger.Info("This is an info message").Log()
	logger.Warn("This is a warn message").Log()
	logger.Error("This is an error message").Log()
	logger.Error("This is a %s message, with format", "error").Log() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").Log()
	logger.Info("This is an info message, but ignored").Log()
	logger.Warn("This is a warn message, not ignored").Log()
	logger.Error("This is an error message, not ignored").Log()

	// Also, we provide some "old school" log method :)
	// (Don't mistake~ I love old school~)
	logger.Printf("This is a log %s, and it's for compatibility", "printed")
	logger.Print("This is a log printed, and it's for compatibility", 123)
	logger.Println("This is a log printed, and it's for compatibility", 666)

	// If you want to log with some fields, try this:
	user := struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   666,
		Name: "FishGoddess",
		Age:  3,
	}

	logger.Warn("This is a structured message").Any("user", user).Json("userJson", user).Log()
	logger.Error("This is a structured message").Error("err", io.EOF).Int("trace", 123).Log()

	// Tired to write Error("err", io.EOF)?
	// Try Err method!
	logger.Error("This is a structured message, too").Err(io.EOF).Int("trace", 456).Log()

	// You may notice logit.Options() which returns an options list.
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information.
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl".
	options.WithWriter(os.Stderr)                 // Change logger's writer to os.Stderr without buffer or batch.
	options.WithBufferWriter(os.Stderr)           // Change logger's writer to os.Stderr with buffer.
	options.WithBatchWriter(os.Stderr)            // Change logger's writer to os.Stderr with batch.
	options.WithErrorWriter(os.Stderr)            // Change logger's error writer to os.Stderr without buffer or batch.
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it).

	// You can bind context with logger and use it as long as you can get the context.
	ctx := logit.NewContext(context.Background(), logger)
	logger = logit.FromContext(ctx)
	logger.Info("Logger from context").Log()

	// You can initialize the global logger if you don't want to use an independent logger.
	// WithCallerDepth will set the depth of caller, and default is core.CallerDepth.
	// Functions in global logger are wrapped so depth of caller should be increased 1.
	// You can specify your depth if you wrap again or have something else reasons.
	logger = logit.NewLogger(options.WithCallerDepth(global.CallerDepth + 1))
	logit.SetGlobal(logger)
	logit.Info("Info from logit").Log()

	// We don't recommend you to call logit.SetGlobal unless you really need to call.
	// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
	logger.SetToGlobal()
	logit.Println("Println from logit")
}
```

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ make bench
$ make benchfile
```

| 测试（输出到内存） | 单位时间内运行次数 (越大越好) | 每个操作消耗时间 (越小越好)       | B/op (越小越好)                     | allocs/op (越小越好)              |
|-----------|------------------|-----------------------|---------------------------------|-------------------------------|
| **logit** | **707851**       | **&nbsp; 1704 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog   | 706714           | &nbsp; 1585 ns/op     | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap       | 389608           | &nbsp; 4688 ns/op     | &nbsp; 865 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus    | &nbsp; 69789     | 17142 ns/op           | 8885 B/op                       | 136 allocs/op                 |

| 测试（输出到文件）     | 单位时间内运行次数 (越大越好) | 每个操作消耗时间 (越小越好)       | B/op (越小越好)                     | allocs/op (越小越好)              |
|---------------|------------------|-----------------------|---------------------------------|-------------------------------|
| **logit**     | **636033**       | **&nbsp; 1822 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| **logit-不缓冲** | **354542**       | **&nbsp; 3502 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog       | 354676           | &nbsp; 3440 ns/op     | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap           | 195354           | &nbsp; 6843 ns/op     | &nbsp; 865 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus        | &nbsp; 58030     | 21088 ns/op           | 8885 B/op                       | 136 allocs/op                 |

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)
> 
> 测试环境：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD，Linux/Manjaro

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目     | 作者          | 描述               | 链接                                                                                         |
|--------|-------------|------------------|--------------------------------------------------------------------------------------------|
| postar | avino-plan  | 一个极易上手的低耦合通用邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo   | FishGoddess | 一个简单的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo)   |

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 `logit` 是用该计划下的 Idea / GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=logit" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
