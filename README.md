# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)

**logit** 是一个简单易用并且是基于级别控制和配置文件的日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

[B站上的介绍视频](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 功能特性

* 独特的日志输出模块设计，使用 wrapper 和 handler 装载特定的模块，实现扩展功能
* 支持日志级别控制，一共有四个日志级别，分别是 debug，info，warn 和 error
* 支持配置文件，可以让用户在项目编译成二进制之后还能灵活地控制日志的设置
* 支持日志记录函数，使用回调的形式获取日志内容，对长日志内容的组织逻辑会更清晰
* 支持开启或者关闭日志功能，线上环境可以关闭或调高日志级别
* 支持记录日志到文件中，并且可以自定义日志文件名
* 支持按照时间间隔进行自动划分日志文件，比如每一天划分一个日志文件
* 支持按照文件大小进行自动划分日志文件，比如每 64 MB 划分一个日志文件
* 增加日志处理器模块，支持用户自定义日志处理逻辑，具有很高的扩展能力
* 支持不输出文件信息，避免 runtime.Caller 方法的调用，具有很高的性能
* 支持调整时间格式化输出，让用户自定义时间输出的格式
* 支持以 Json 形式输出日志信息，更方便后续对日志进行解析

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

> 目前稳定的是 v0.2.x 版本，但是 v0.3.x 版本已经出了第一个体验版，这是一个全新设计的版本，废除了很多冗余设计！

### 🚀 安装方式

```bash
$ go get github.com/FishGoddess/logit
```

### 📖 参考案例

```go
package main

import (
	"os"

	"github.com/FishGoddess/logit"
)

func main() {

	// There are four levels can be logged
	logit.Debug("Hello, I am debug!") // Ignore because default level is info
	logit.Info("Hello, I am info!")
	logit.Warn("Hello, I am warn!")
	logit.Error("Hello, I am error!")

	// You can format log with some parameters if you want
	logit.DebugF("Hello, I am debugF %d!", 2) // Ignore because default level is info
	logit.InfoF("Hello, I am infoF %d!", 2)
	logit.WarnF("Hello, I am warnF %d!", 2)
	logit.ErrorF("Hello, I am errorF %d!", 2)

	// logit.Me() returns a completed logger for use

	// Set level to debug
	logit.Me().SetLevel(logit.DebugLevel)

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logit.Me().NeedCaller(true)

	// Set format of time in log
	logit.Me().TimeFormat("2006/01/02 15:04:05")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().SetEncoder(logit.JsonEncoder())
	logit.Me().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().SetDebugEncoder(logit.JsonEncoder())
	logit.Me().SetInfoEncoder(logit.JsonEncoder())
	logit.Me().SetWarnEncoder(logit.JsonEncoder())
	logit.Me().SetErrorEncoder(logit.JsonEncoder())
	logit.Me().SetDebugWriter(os.Stdout)
	logit.Me().SetInfoWriter(os.Stdout)
	logit.Me().SetWarnWriter(os.Stdout)
	logit.Me().SetErrorWriter(os.Stdout)
}
```

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=10s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **7513623** | **1612 ns/op** | **384 B/op** | **8 allocs/op** |
| golog | 4569554 | 2631 ns/op | 712 B/op | 24 allocs/op |
| zap | 3891336 | 3084 ns/op | 448 B/op | 16 allocs/op |
| logrus | 2089682 | 5769 ns/op | 1633 B/op | 52 allocs/op |

> 测试环境：R7-4700U CPU @ 2.0 GHZ，16 GB RAM

**注意：**

**1. 输出文件信息会有运行时操作（runtime.Caller 方法），非常影响性能，**
**但是这个功能感觉还是比较实用的，尤其是在查找错误的时候，所以我们还是加了这个功能！**
**如果你更在乎性能，那我们也提供了一个选项可以关闭文件信息的查询！**

**2. v0.0.7 及以前版本的日志输出使用了 fmt 包的一些方法，经过性能检测发现这些方法存在大量使用反射的**
**行为，主要体现在对参数 v interface{} 进行类型检测的逻辑上，而日志输出都是字符串，这一个**
**判断是可以省略的，可以减少很多运行时操作时间！v0.0.8 版本开始使用了更有效率的输出方式！**

**3. 经过对 v0.0.8 版本的性能检测，发现时间格式化操作消耗了接近一半的处理时间，**
**主要体现在 time.Time.AppendFormat 的调用上。在 v0.0.11 版本中使用了时间缓存机制进行优化，**
**目前存在一个疑惑就是使用并发竞争去换取时间格式化的性能消耗究竟值不值得？**
**答案是不值得，我们在 v0.1.1-alpha 及更高版本中取消了这个时间缓存机制。**

**4. 值得注意的是，DebugF 一类带格式化的 API 性能达不到这个水平，因为还是使用了反射技术，但是性能依旧是不差的：**

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| logit | 7513623 | 1612 ns/op | 384 B/op | 8 allocs/op |
| **logit-使用反射技术** | **6042254** | **1984 ns/op** | **424 B/op** | **12 allocs/op** |

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 | 链接 |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | 一个极易上手的低耦合高性能邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | 一个高性能的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo) |
