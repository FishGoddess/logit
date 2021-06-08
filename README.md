# 📝 logit

[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)
[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icon/build.svg)](_icon/build.svg)
[![License](_icon/coverage.svg)](_icon/coverage.svg)

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

	// Create a new logger first
	logger := logit.NewLogger()

	// There are four levels can be logged, and you can format log with some parameters
	logger.DebugF("Hello, I am debug %d!", 2) // Ignore because default level is info
	logger.InfoF("Hello, I am info %d!", 2)
	logger.WarnF("Hello, I am warn %d!", 2)
	logger.ErrorF("Hello, I am error %d!", 2)

	// Set level to debug
	logger.SetLevel(logit.DebugLevel)
	logger.DebugF("Now debug logs will come up!")

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logger.SetNeedCaller(true)
	logger.InfoF("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logger.Encoders().SetEncoder(logit.NewTextEncoder("2006-01-02 15:04:05"))
	logger.Encoders().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logger.Writers().SetWriter(os.Stdout)
	logger.Writers().SetErrorWriter(os.Stderr)
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

| 测试（输出到内存） | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3775916** | **&nbsp; 949 ns/op** | **&nbsp; 128 B/op** | **&nbsp; 4 allocs/op** |
| zap | 1674750 | 2143 ns/op | &nbsp; 449 B/op | 16 allocs/op |
| golog | 2223093 | 1619 ns/op | &nbsp; 713 B/op | 24 allocs/op |
| logrus | &nbsp; 899808 | 3968 ns/op | 1634 B/op | 52 allocs/op |

| 测试（输出到文件） | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3556720** | **&nbsp; 1009 ns/op** | **&nbsp; 129 B/op** | **&nbsp; 4 allocs/op** |
| **logit-不使用缓冲写出器** | **&nbsp; 499887** | **&nbsp; 7176 ns/op** | **&nbsp; 128 B/op** | **&nbsp; 4 allocs/op** |
| zap | &nbsp; 409000 | &nbsp; 8580 ns/op | &nbsp; 449 B/op | 16 allocs/op |
| golog | &nbsp; 257083 | 13884 ns/op | &nbsp; 713 B/op | 24 allocs/op |
| logrus | &nbsp; 327198 | 10699 ns/op | 1634 B/op | 52 allocs/op |

> 测试环境：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD

**注意：格式化的性能达不到这个水平，因为还是使用了反射技术，但是性能依旧是不差的：**

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| logit | 3775916 | &nbsp; 949 ns/op | 128 B/op | 4 allocs/op |
| **logit-使用格式化日志** | **2931703** | **1233 ns/op** | **168 B/op** | **8 allocs/op** |

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 | 链接 |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | 一个极易上手的低耦合高性能邮件服务 | [Github](https://github.com/avino-plan/postar) / [码云](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | 一个高性能的轻量级分布式缓存中间件 | [Github](https://github.com/FishGoddess/kafo) / [码云](https://gitee.com/FishGoddess/kafo) |
