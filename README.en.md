# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/logit/actions/workflows/test.yml/badge.svg)

**logit** is a level-based, high-performance and pure-structured logger for all [GoLang](https://golang.org)
applications.

[阅读中文版的 Read me](./README.md)

### 🥇 Features

* Based on Handler in Go, and provide a better performance.
* Level-based logging, and there are four levels to use: debug, info, warn, error.
* Key-Value structured log supports, also supporting format.
* Support logging as Text/Json string, by using provided appender.
* Asynchronous write back supports, providing high-performance Buffer writer to avoid IO accessing.
* Provide global optimized settings, let some settings feet your business.
* Every level has its own appender and writer, separating process error logs is recommended.
* Context binding supports, using logger is more elegant.
* Configuration plugins supports, ex: yaml plugin can create logger from yaml configuration file.
* Interceptor supports which can inject values from context.
* Error handling supports which can let you count errors and report them.
* Rotate file supports, clean automatically if files are aged or too many.
* Config file supports, such as json/yaml/toml/bson.

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to know about more information._

> We redesigned api of logit after v1.5.0 because Go1.21 introduced slog package.

### 🚀 Installation

```bash
$ go get -u github.com/FishGoddess/logit
```

### 📖 Examples

```go
package main

import (
	"context"
	"fmt"

	"github.com/FishGoddess/logit"
)

func main() {
	// Use default logger to log.
	// By default, logs will be output to stdout.
	logit.Default().Info("hello from logit", "key", 123)

	// Use a new logger to log.
	// By default, logs will be output to stdout.
	logger := logit.NewLogger()

	logger.Debug("new version of logit", "version", "1.5.0-alpha", "date", 20231122)
	logger.Error("new version of logit", "version", "1.5.0-alpha", "date", 20231122)

	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	u := user{123456, "fishgoddess"}
	logger.Info("user information", "user", u, "pi", 3.14)

	// Yep, I know you want to output logs to a file, try WithFile option.
	// The path in WithFile is where the log file will be stored.
	// Also, it's a good choice to call logger.Close() when program shutdown.
	logger = logit.NewLogger(logit.WithFile("./logit.log"))
	defer logger.Close()

	logger.Info("check where I'm logged", "file", "logit.log")

	// What if I want to use default logger and output logs to a file? Try SetDefault.
	// It sets a logger to default and you can use it by package function or Default().
	logit.SetDefault(logger)
	logit.Default().Warn("this is from default logger", "pi", 3.14, "default", true)

	// If you want to change level of logger to info, try WithInfoLevel.
	// Other levels is similar to info level.
	logger = logit.NewLogger(logit.WithInfoLevel())

	logger.Debug("debug logs will be ignored")
	logger.Info("info logs can be logged")

	// If you want to pass logger by context, use NewContext and FromContext.
	ctx := logit.NewContext(context.Background(), logger)

	logger = logit.FromContext(ctx)
	logger.Info("logger from context", "from", "context")

	// Don't want to panic when new a logger? Try NewLoggerGracefully.
	logger, err := logit.NewLoggerGracefully(logit.WithFile(""))
	if err != nil {
		fmt.Println("new logger gracefully failed:", err)
	}
}
```

_More examples can be found in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ make bench
$ make benchfile
```

```bash
goos: linux
goarch: amd64
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLogitLogger-2                   1486184               810 ns/op               0 B/op          0 allocs/op
BenchmarkLogitLoggerTextHandler-2        1000000              1080 ns/op               0 B/op          0 allocs/op
BenchmarkLogitLoggerJsonHandler-2         847864              1393 ns/op             120 B/op          3 allocs/op
BenchmarkLogitLoggerPrint-2              1222302               981 ns/op              48 B/op          1 allocs/op
BenchmarkSlogLoggerTextHandler-2          725522              1629 ns/op               0 B/op          0 allocs/op
BenchmarkSlogLoggerJsonHandler-2          583214              2030 ns/op             120 B/op          3 allocs/op
BenchmarkZeroLogLogger-2                 1929276               613 ns/op               0 B/op          0 allocs/op
BenchmarkZapLogger-2                      976855              1168 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusLogger-2                   231723              4927 ns/op            2080 B/op         32 allocs/op

BenchmarkLogitFile-2                      624774              1935 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBuffer-2           1378076               873 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBatch-2            1367479               883 ns/op               0 B/op          0 allocs/op
BenchmarkSlogFile-2                       407590              2944 ns/op               0 B/op          0 allocs/op
BenchmarkZeroLogFile-2                    634375              1810 ns/op               0 B/op          0 allocs/op
BenchmarkZapFile-2                        382790              2641 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusFile-2                     174944              6491 ns/op            2080 B/op         32 allocs/op
```

> Notice: WithBuffer and WithBatch are using buffer writer and batch writer.

> Benchmarks: [_examples/performance_test.go](./_examples/performance_test.go)

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

[![Star History Chart](https://api.star-history.com/svg?repos=fishgoddess/logit&type=Date)](https://star-history.com/#fishgoddess/logit&Date)
