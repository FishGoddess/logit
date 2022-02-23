# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/go-logit/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit** 是一个基于级别控制的高性能纯结构化日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 在看了一些优秀日志库的设计之后，我越发觉得 logit 非常烂，尤其是和 zerolog 对比之后，简直不堪入目。这让我夜不能寐，最后在小黑屋中完成了新的设计方案。

[Read me in English](./README.en.md)

~~[B站上的介绍视频](https://www.bilibili.com/video/BV14t4y1y7rF)~~

### 🥇 功能特性

* 独特的日志模块设计，使用 appender 和 writer 装载特定的模块，实现扩展功能。
* 支持日志级别控制，一共有五个日志级别，分别是 debug，info，warn，error，print 和 off。
* 支持键值对形式的结构化日志记录，同时对格式化操作也有支持。
* 支持以 Text/Json 形式输出日志信息，方便对日志进行解析。
* 支持异步回写日志，提供高性能缓冲写出器模块，减少 IO 的访问次数。
* 提供调优使用的全局配置，对一些高级配置更贴合实际业务的需求。
* 分级别追加日志数据，分级别写出日志数据，推荐将 error 级别的日志单独处理和存储。
* 加入 Context 机制，更优雅地使用日志，并支持业务域划分。
* 以插件的形式支持配置方式初始化，比如 yaml 配置插件可以通过 yaml 配置文件初始化等。
* 支持拦截器模式，可以从 context 注入外部常量或变量，简化日志输出流程。

> 配置插件以接口注册的形式支持，具体有哪些支持的配置插件请查看 [Github](https://github.com/go-logit) / [码云](https://gitee.com/go-logit)

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

```bash
$ go get -u github.com/go-logit/logit
```

### 📖 参考案例

```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/go-logit/logit"
)

func main() {
	// Create a new logger for use.
	// Default level is debug, so all logs will be logged.
	// Invoke Close() isn't necessary in all situations.
	// If logger's writer has buffer or something like that, it's better to invoke Close() for flushing buffer or something else.
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want.
	// Remember, logs will be ignored if their level is smaller than logger's level.
	// End() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").End()
	logger.Info("This is a info message").End()
	logger.Warn("This is a warn message").End()
	logger.Error("This is a error message").End()
	logger.Error("This is a %s message, with format", "error").End() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").End()
	logger.Info("This is a info message, but ignored").End()
	logger.Warn("This is a warn message, not ignored").End()
	logger.Error("This is a error message, not ignored").End()

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
	logger.Warn("This is a structured message").Any("user", user).Json("userJson", user).End()
	logger.Error("This is a structured message").Error("err", io.EOF).Int("trace", 123).End()

	// You may notice logit.Options() which returns an options list.
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information.
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl".
	options.WithWriter(os.Stderr, true)           // Change logger's writer to os.Stderr with buffer.
	options.WithErrorWriter(os.Stderr, false)     // Change logger's error writer to os.Stderr without buffer.
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it).

	// You can bind context with logger and use it as long as you can get the context.
	ctx := logit.NewContext(context.Background(), logger)
	logger = logit.FromContext(ctx)
}
```

* [basic](./_examples/basic.go)
* [options](./_examples/options.go)
* [appender](./_examples/appender.go)
* [writer](./_examples/writer.go)
* [global](./_examples/global.go)
* [context](./_examples/context.go)
* [creator](./_examples/creator.go)
* [caller](./_examples/caller.go)
* [interceptor](./_examples/interceptor.go)

_所有的使用案例都在 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/performance_test.go -bench=. -benchtime=1s
```

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

| 测试（输出到内存） | 单位时间内运行次数 (越大越好) | 每个操作消耗时间 (越小越好)       | B/op (越小越好)                     | allocs/op (越小越好)              |
|-----------|------------------|-----------------------|---------------------------------|-------------------------------|
| **logit** | **799759**       | **&nbsp; 1373 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog   | 922863           | &nbsp; 1244 ns/op     | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap       | 413701           | &nbsp; 2824 ns/op     | &nbsp; 897 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus    | 105238           | 11474 ns/op           | 7411 B/op                       | 128 allocs/op                 |

| 测试（输出到文件）     | 单位时间内运行次数 (越大越好) | 每个操作消耗时间 (越小越好)       | B/op (越小越好)                     | allocs/op (越小越好)              |
|---------------|------------------|-----------------------|---------------------------------|-------------------------------|
| **logit**     | **599862**       | **&nbsp; 1768 ns/op** | **&nbsp; 901 B/op**             | **&nbsp; &nbsp; 0 allocs/op** |
| **logit-不缓冲** | **148113**       | **&nbsp; 7773 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog       | 159962           | &nbsp; 7472 ns/op     | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap           | 130405           | &nbsp; 9137 ns/op     | &nbsp; 897 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus        | &nbsp; 65202     | 18439 ns/op           | 7410 B/op                       | 128 allocs/op                 |

> 测试环境：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目     | 作者          | 描述               | 链接                                                                                         |
|--------|-------------|------------------|--------------------------------------------------------------------------------------------|
| postar | avino-plan  | 一个极易上手的低耦合通用邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo   | FishGoddess | 一个简单的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo)   |

最后，我想感谢 JetBrains 公司的 **free JetBrains Open Source license(s)**，因为 `logit` 是用该计划下的 Idea / GoLand 完成开发的。

<a href="https://www.jetbrains.com/?from=logit" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
