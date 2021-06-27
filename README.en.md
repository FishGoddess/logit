# 📝 logit

[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)
[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![License](_icon/build.svg)](_icon/build.svg)
[![License](_icon/coverage.svg)](_icon/coverage.svg)

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

	// Create a new logger first
	logger := logit.NewLogger()

	// There are four levels can be logged
	logger.Debug("Hello, I am debug!") // Ignore because default level is info
	logger.Info("Hello, I am info!")
	logger.Warn("Hello, I am warn!")
	logger.Error("Hello, I am error!", logit.KV{"err": "xxx", "id": 666}) // carry some values to log

	// You can format log with some parameters if you want
	logger.DebugF("Hello, I am debug %d!", 2) // Ignore because default level is info
	logger.InfoF("Hello, I am info %d!", 2)
	logger.WarnF("Hello, I am warn %d!", 2)
	logger.ErrorF("Hello, I am error %d!", 2)

	// Set level to debug
	logger.SetLevel(logit.DebugLevel)
	logger.Debug("Now debug logs will come up!")

	// Log won't carry caller information in default
	// So, try SetNeedCaller if you need
	logger.SetNeedCaller(true)
	logger.Info("I need caller!")

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

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=3s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case(output to memory) | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3185127** | **1125 ns/op** | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; 0 allocs/op** |
| zap | 1674750 | 2143 ns/op | &nbsp; 449 B/op | 16 allocs/op |
| zerolog | 5026804 | &nbsp; 1201 ns/op | &nbsp; &nbsp; &nbsp; 0 B/op | &nbsp; 0 allocs/op |
| logrus | &nbsp; 899808 | 3968 ns/op | 1634 B/op | 52 allocs/op |

| test case(output to file) | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **3556720** | **&nbsp; 1009 ns/op** | **&nbsp; 129 B/op** | **&nbsp; 4 allocs/op** |
| **logit-withoutBuffer** | **&nbsp; 499887** | **&nbsp; 7176 ns/op** | **&nbsp; 128 B/op** | **&nbsp; 4 allocs/op** |
| zap | &nbsp; 409000 | &nbsp; 8580 ns/op | &nbsp; 449 B/op | 16 allocs/op |
| zerolog | 506928 | &nbsp; 7633 ns/op | &nbsp; &nbsp; &nbsp; 0 B/op | &nbsp; 0 allocs/op |
| logrus | &nbsp; 327198 | 10699 ns/op | 1634 B/op | 52 allocs/op |

> Environment：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD

**Notice: You should know that format can't reach high performance as the same as others because of reflection,**
**however, their performances are not as bad as we think:**

| test case | times ran (large is better) |  ns/op (small is better) | B/op | allocs/op |
| -----------|--------|-------------|-------------|-------------|
| logit | 3775916 | &nbsp; 949 ns/op | 128 B/op | 4 allocs/op |
| **logit-useFormatLog** | **2931703** | **1233 ns/op** | **168 B/op** | **8 allocs/op** |

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description | link |
| -----------|--------|-------------| ---------------- |
| postar | avino-plan | An easy-to-use and low-coupling email service | [Github](https://github.com/avino-plan/postar) / [Gitee](https://gitee.com/avino-plan/postar) |
| kafo | FishGoddess | A high-performance and distributed cache middleware | [Github](https://github.com/FishGoddess/kafo) / [Gitee](https://gitee.com/FishGoddess/kafo) |
