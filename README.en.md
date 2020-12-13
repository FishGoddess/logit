# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)

**logit** is an easy-to-use, also level-based and config file first logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md)

[Introduction Video on BiliBili](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 Features

* Modularization design, easy to extend your logger with encoders and writers
* Level-based logging, and there are four levels to use
* Config file supports, you can use a config file to change you logger flexibility even it has been a binary
* Enable or disable Logger, you can disable or switch to a higher level in your production environment
* Log file supports, and you can customer the name of your log file
* Time rolling supports, such as one day one log file
* File size rolling supports, such as one 64 MB one log file
* Count rolling supports, such as 1000 logging operations one log file
* High-performance supports, by avoiding to call runtime.Caller
* Time format supports, you can format time in your way
* Log as Json string supports, by using provided JsonLoggerHandler

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

> Currently, stable version is v0.2.x, but a brand-new version v0.3.x is alpha with a more elegant design!

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
* [encoder](./_examples/encoder.go)
* [writer](./_examples/writer.go)

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=10s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **7513623** | **1612 ns/op** | **384 B/op** | **8 allocs/op** |
| golog | 4569554 | 2631 ns/op | 712 B/op | 24 allocs/op |
| zap | 3891336 | 3084 ns/op | 448 B/op | 16 allocs/op |
| logrus | 2089682 | 5769 ns/op | 1633 B/op | 52 allocs/op |

> Environment：R7-4700U CPU @ 2.0 GHZ，16 GB RAM

**Notice:**

**1. Fetching file info will call runtime.Caller, which is expensive.**
**However, we think file info is useful in check errors,**
**so we keep this feature, and provide a switch to turn off it for high-performance.**

**2. v0.0.7 and lower versions use some functions of fmt, and these functions is expensive**
**because of reflect (for judging the parameter v interface{}). Actually, these judgements**
**are redundant in a logger. The more effective output is used in v0.0.8 and higher versions.**

**3. After checking the benchmarks of v0.0.8 version, we found that time format takes a lots of time**
**because of time.Time.AppendFormat. In v0.0.11 and higher versions, we use time cache mechanism to**
**reduce the times of time format. However, is it worth to replace time format operation with concurrent competition?**
**The answer is no, so we cancel this mechanism in v0.1.1-alpha and higher versions.**

**4. You should know that some APIs like DebugF can't reach high performance as the same as others because of reflection,**
**however, their performances are not as bad as we think:**

| test case | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| logit | 7513623 | 1612 ns/op | 384 B/op | 8 allocs/op |
| **logit-reflection** | **6042254** | **1984 ns/op** | **424 B/op** | **12 allocs/op** |

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description | link |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | An easy-to-use and low-coupling email service | [Github](https://github.com/avino-plan/postar) / [Gitee](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | A high-performance and distributed cache middleware | [Github](https://github.com/FishGoddess/kafo) / [Gitee](https://gitee.com/FishGoddess/kafo) |
