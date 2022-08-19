# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/go-logit/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Build](_icons/build.svg)](_icons/build.svg)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)

**logit** is a level-based, high-performance and pure-structured logger for all [GoLang](https://golang.org)
applications.

> After reading some amazing logging libs, I found that logit is just a joke, especially comparing with zerolog, so I decided to redesign logit.

[阅读中文版的 Read me](./README.md)

~~[Introduction Video on BiliBili](https://www.bilibili.com/video/BV14t4y1y7rF)~~

### 🥇 Features

* Modularization design, easy to extend your logger with appender and writer.
* Level-based logging, and there are five levels to use: debug, info, warn, error, print, off.
* Key-Value structured log supports, also supporting format.
* Support logging as Text/Json string, by using provided appender.
* Asynchronous write back supports, providing high-performance Buffer writer to avoid IO accessing.
* Provide global optimized settings, let some settings feet your business.
* Every level has its own appender and writer, separating process error logs is recommended.
* Context binding supports, using logger is more elegant.
* Configuration plugins supports, ex: yaml plugin can create logger from yaml configuration file.
* Interceptor supports which can inject values from context.

> Configuration plugins is based on an interface, more information can be found in [Github](https://github.com/go-logit) / [Gitee](https://gitee.com/go-logit)

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

> The brand-new version v0.4.x is developing with a more elegant design!

### 🚀 Installation

```bash
$ go get -u github.com/go-logit/logit
```

### 📖 Examples

```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core"
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
	// Log() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").Log()
	logger.Info("This is a info message").Log()
	logger.Warn("This is a warn message").Log()
	logger.Error("This is a error message").Log()
	logger.Error("This is a %s message, with format", "error").Log() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").Log()
	logger.Info("This is a info message, but ignored").Log()
	logger.Warn("This is a warn message, not ignored").Log()
	logger.Error("This is a error message, not ignored").Log()

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
	logger.Info("Logger from context").Log()

	// You can initialize the global logger if you don't want to use an independent logger.
	// WithCallerDepth will set the depth of caller, and default is core.CallerDepth.
	// Functions in global logger are wrapped so depth of caller should be increased 1.
	// You can specify your depth if you wrap again or have something else reasons.
	logger = logit.NewLogger(options.WithCallerDepth(core.CallerDepth + 1))
	logit.SetGlobal(logger)
	logit.Info("Info from logit").Log()

	// We don't recommend you to call logit.SetGlobal unless you really need to call.
	// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
	logger.SetToGlobal()
	logit.Println("Println from logit")
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

_All examples can be found in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/performance_test.go -bench=. -benchtime=1s
```

> Benchmark file：[_examples/performance_test.go](./_examples/performance_test.go)

| test case(output to memory) | times ran (large is better) | ns/op (small is better) | B/op                            | allocs/op                     |
|-----------------------------|-----------------------------|-------------------------|---------------------------------|-------------------------------|
| **logit**                   | **799759**                  | **&nbsp; 1373 ns/op**   | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog                     | 922863                      | &nbsp; 1244 ns/op       | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap                         | 413701                      | &nbsp; 2824 ns/op       | &nbsp; 897 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus                      | 105238                      | 11474 ns/op             | 7411 B/op                       | 128 allocs/op                 |

| test case(output to file) | times ran (large is better) | ns/op (small is better) | B/op                            | allocs/op                     |
|---------------------------|-----------------------------|-------------------------|---------------------------------|-------------------------------|
| **logit**                 | **599862**                  | **&nbsp; 1768 ns/op**   | **&nbsp; 901 B/op**             | **&nbsp; &nbsp; 0 allocs/op** |
| **logit-notBuffer**       | **148113**                  | **&nbsp; 7773 ns/op**   | **&nbsp; &nbsp; &nbsp; 0 B/op** | **&nbsp; &nbsp; 0 allocs/op** |
| zerolog                   | 159962                      | &nbsp; 7472 ns/op       | &nbsp; &nbsp; &nbsp; 0 B/op     | &nbsp; &nbsp; 0 allocs/op     |
| zap                       | 130405                      | &nbsp; 9137 ns/op       | &nbsp; 897 B/op                 | &nbsp; &nbsp; 8 allocs/op     |
| logrus                    | &nbsp; 65202                | 18439 ns/op             | 7410 B/op                       | 128 allocs/op                 |

> Environment：R7-5800X CPU@3.8GHZ，32GB RAM，512GB SSD

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author      | Description                                     | link                                                                                          |
|---------|-------------|-------------------------------------------------|-----------------------------------------------------------------------------------------------|
| postar  | avino-plan  | An easy-to-use and low-coupling email service   | [Github](https://github.com/avino-plan/postar) / [Gitee](https://gitee.com/avino-plan/postar) |
| kafo    | FishGoddess | An easy-to-use and distributed cache middleware | [Github](https://github.com/FishGoddess/kafo) / [Gitee](https://gitee.com/FishGoddess/kafo)   |

At last, I want to thank JetBrains for **free JetBrains Open Source license(s)**, because `logit` is developed with Idea / GoLand under it.

<a href="https://www.jetbrains.com/?from=logit" target="_blank"><img src="./_icons/jetbrains.png" width="250"/></a>
