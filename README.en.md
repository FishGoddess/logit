# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)

**logit** is a level-based and high-performance logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md)

[Introduction Video on BiliBili](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 Features

* Modularization design, easy to extend your logger with encoders and writers
* Level-based logging, and there are four levels to use
* Enable or disable Logger, you can disable or switch to a higher level in your production environment
* Log file supports, and you can customer the name of your log file
* Time rolling supports, such as one day one log file
* File size rolling supports, such as one 64 MB one log file
* Count rolling supports, such as 1000 logging operations one log file
* High-performance supports, by avoiding to call runtime.Caller
* Time format supports, you can format time in your way
* Log as Json string supports, by using provided JsonLoggerHandler

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

> The brand-new version v0.4.x is developing with a more elegant design!

### 🚀 Installation

```bash
$ go get github.com/FishGoddess/logit
```

### 📖 Examples

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
	// So, try SetNeedCaller if you need
	logit.Me().SetNeedCaller(true)
	logit.Info("I need caller!")

	// Set encoder and writer
	// Actually, every level has own encoder and writer
	// This way will set encoder and writer of all levels to the same one
	logit.Me().Encoders().SetEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Writers().SetWriter(os.Stdout)

	// We also provide some functions to set encoder and writer of each level
	logit.Me().Encoders().SetDebugEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetInfoEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetWarnEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Encoders().SetErrorEncoder(logit.NewJsonEncoder("2006-01-02 15:04:05"))
	logit.Me().Writers().SetDebugWriter(os.Stdout)
	logit.Me().Writers().SetInfoWriter(os.Stdout)
	logit.Me().Writers().SetWarnWriter(os.Stdout)
	logit.Me().Writers().SetErrorWriter(os.Stdout)
}
```

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [encoder](./_examples/encoder.go)
* [writer](./_examples/writer.go)

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=3s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3950809** | **917 ns/op** | **128 B/op** | **4 allocs/op** |
| golog | 4569554 | 2631 ns/op | 712 B/op | 24 allocs/op |
| zap | 3891336 | 3084 ns/op | 448 B/op | 16 allocs/op |
| logrus | 2089682 | 5769 ns/op | 1633 B/op | 52 allocs/op |

> Environment：R7-5800X CPU @ 3.8 GHZ，32 GB RAM

**Notice:**

**1. Fetching file info will call runtime.Caller, which is expensive,**
**and we keep this feature, and provide a switch to turn off it for high-performance.**

**2. You should know that format can't reach high performance as the same as others because of reflection,**
**however, their performances are not as bad as we think:**

| test case | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| logit | 3950809 | 917 ns/op | 128 B/op | 4 allocs/op |
| **logit-reflection** | **2984533** | **1197 ns/op** | **168 B/op** | **8 allocs/op** |

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description | link |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | An easy-to-use and low-coupling email service | [Github](https://github.com/avino-plan/postar) / [Gitee](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | A high-performance and distributed cache middleware | [Github](https://github.com/FishGoddess/kafo) / [Gitee](https://gitee.com/FishGoddess/kafo) |
