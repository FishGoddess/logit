# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icons/build.svg)](_icons/build.svg)
[![License](_icons/coverage.svg)](_icons/coverage.svg)

**logit** 是一个基于级别控制的高性能日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

> 在看了一些优秀日志库的设计之后，我越发觉得 logit 非常烂，尤其是和 zerolog 对比之后，简直不堪入目。这让我夜不能寐，最后在小黑屋中完成了新的设计方案。

[Read me in English](./README.en.md)

[B站上的介绍视频](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 功能特性

* 独特的日志输出模块设计，使用 appender 和 writer 装载特定的模块，实现扩展功能
* 支持日志级别控制，一共有五个日志级别，分别是 debug，info，warn，error 和 off
* 自带缓冲写出器，并且可以自定义日志文件名
* 支持按照时间间隔进行自动分割日志文件，比如每一天分割一个日志文件
* 支持按照文件大小进行自动分割日志文件，比如每 64 MB 分割一个日志文件
* 支持按照日志记录次数进行自动分割日志文件，比如每记录 1000 条日志分割一个日志文件
* 支持不输出文件信息，避免 runtime.Caller 方法的调用，具有很高的性能
* 支持调整时间格式化输出，让用户自定义时间输出的格式
* 支持以 Json 形式输出日志信息，更方便后续对日志进行解析

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

> v0.4.x 版本已经在规划开发中，这是一个全新设计的版本，去掉了很多垃圾设计和功能！

### 🚀 安装方式

```bash
$ go get -u github.com/FishGoddess/logit
```

### 📖 参考案例

```go
package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// Create a new logger for use
	// Default level is debug, so all logs will be logged
	// Invoke Close() isn't necessary in all situations
	// If logger's writer has buffer or something like that, it's better to invoke Close() for flushing buffer or something else
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want
	// Remember, logs will be ignored if their level is smaller than logger's level
	// End() will do some finishing work, so this invocation is necessary
	logger.Debug("This is a debug message").End()
	logger.Info("This is a info message").End()
	logger.Warn("This is a warn message").End()
	logger.Error("This is a error message").End()
	logger.Error("This is a %s message, with format", "error").End() // Format with params

	// As you know, we provide some levels: debug, info, warn, error, off
	// The lowest is debug and the highest is off
	// If you want to change the level of your logger, do it at creating
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").End()
	logger.Info("This is a info message, but ignored").End()
	logger.Warn("This is a warn message, not ignored").End()
	logger.Error("This is a error message, not ignored").End()

	// You may notice logit.Options() which returns an options list
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl"
	options.WithWriter(os.Stderr)                 // Change logger's writer to os.Stderr
	options.WithBuffered(os.Stderr)               // Change logger's writer to os.Stderr with buffer
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it)
}
```

* [basic](./_examples/basic.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=1s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| 测试（输出到内存） | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **856915** | **&nbsp; 1385 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog | 922863 | &nbsp; 1244 ns/op | &nbsp; &nbsp; &nbsp; 0 B/op | &nbsp; &nbsp; 0 allocs/op |
| zap | 413701 | &nbsp; 2824 ns/op | &nbsp; 897 B/op | &nbsp; &nbsp; 8 allocs/op |
| logrus | 105238 | 11474 ns/op | 7411 B/op | 128 allocs/op |

| 测试（输出到文件） | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **521606** | **&nbsp; 1927 ns/op** | **1036 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| **logit-不缓冲** | **149965** | **&nbsp; 7704 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog | 159962 | &nbsp; 7472 ns/op | &nbsp; &nbsp; &nbsp; 0 B/op | &nbsp; &nbsp; 0 allocs/op |
| zap | 130405 | &nbsp; 9137 ns/op | &nbsp; 897 B/op | &nbsp; &nbsp; 8 allocs/op |
| logrus | &nbsp; 65202 | 18439 ns/op | 7410 B/op | 128 allocs/op |

> 测试环境：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 | 链接 |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | 一个极易上手的低耦合高性能邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | 一个高性能的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo) |
