# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)

**logit** 是一个基于级别控制的高性能日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

[B站上的介绍视频](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 功能特性

* 独特的日志输出模块设计，使用 encoder 和 writer 装载特定的模块，实现扩展功能
* 支持日志级别控制，一共有四个日志级别，分别是 debug，info，warn 和 error
* 支持开启或者关闭日志功能，线上环境可以关闭或调高日志级别
* 支持记录日志到文件中，并且可以自定义日志文件名
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
	logit.Debug("Hello, I am debug %d!", 2) // Ignore because default level is info
	logit.Info("Hello, I am info %d!", 2)
	logit.Warn("Hello, I am warn %d!", 2)
	logit.Error("Hello, I am error %d!", 2)

	// logit.Me() returns a completed logger for use

	// Set level to debug
	logit.Me().SetLevel(logit.DebugLevel)

	// Log won't carry caller information in default
	// So, try NeedCaller if you need
	logit.Me().NeedCaller(true)
	logit.Info("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().SetDebugEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetInfoEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().SetDebugWriter(os.Stdout)
	logit.Me().SetInfoWriter(os.Stdout)
	logit.Me().SetWarnWriter(os.Stdout)
	logit.Me().SetErrorWriter(os.Stdout)
}
```

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [encoder](./_examples/encoder.go)
* [writer](./_examples/writer.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=3s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3950809** | **917 ns/op** | **128 B/op** | **4 allocs/op** |
| golog | 4569554 | 2631 ns/op | 712 B/op | 24 allocs/op |
| zap | 3891336 | 3084 ns/op | 448 B/op | 16 allocs/op |
| logrus | 2089682 | 5769 ns/op | 1633 B/op | 52 allocs/op |

> 测试环境：R7-5800X CPU @ 3.8 GHZ，32 GB RAM

**注意：**

**1. 输出文件信息会有运行时操作（runtime.Caller 方法），非常影响性能，**
**如果你更在乎性能，那我们也提供了一个选项可以关闭文件信息的查询！**

**2. 值得注意的是，格式化的性能达不到这个水平，因为还是使用了反射技术，但是性能依旧是不差的：**

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| logit | 3950809 | 917 ns/op | 128 B/op | 4 allocs/op |
| **logit-使用反射技术** | **2984533** | **1197 ns/op** | **168 B/op** | **8 allocs/op** |

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 | 链接 |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | 一个极易上手的低耦合高性能邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | 一个高性能的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo) |
