# 📝 logit

[![Go Doc](_icons/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit)
[![License](_icons/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Coverage](_icons/coverage.svg)](_icons/coverage.svg)
![Test](https://github.com/FishGoddess/logit/actions/workflows/test.yml/badge.svg)

**logit** 是一个基于级别控制的高性能纯结构化日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

### 🥇 功能特性

* 兼容标准库 Handler 的扩展设计，并且提供了更高的性能。
* 支持日志级别控制，一共有四个日志级别，分别是 debug，info，warn，error。
* 支持键值对形式的结构化日志记录，同时对格式化操作也有支持。
* 支持以 Text/Json 形式输出日志信息，方便对日志进行解析。
* 支持异步回写日志，提供高性能缓冲写出器模块，减少 IO 的访问次数。
* 提供调优使用的全局配置，对一些高级配置更贴合实际业务的需求。
* 加入 Context 机制，更优雅地使用日志，并支持业务分组划分。
* 支持拦截器模式，可以从 context 注入外部常量或变量，简化日志输出流程。
* 支持错误监控，可以很方便地进行错误统计和告警。
* 支持日志按大小自动分割，并支持按照时间和数量自动清理。
* 支持多种配置文件序列化成 option，比如 json/yaml/toml/bson，然后创建日志记录器。

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

> 由于 Go1.21 增加了 slog 日志包，基本确定了 Go 风格的日志 API，所以 logit v1.5.0 版本开始也调整为类似的风格，并且提供更多的功能和更高的性能。

### 🚀 安装方式

```bash
$ go get -u github.com/FishGoddess/logit
```

### 📖 参考案例

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

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ make bench
$ make benchfile
```

```bash
goos: linux
goarch: amd64
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLogitLoggerTextHandler-2        1000000              1089 ns/op               0 B/op          0 allocs/op
BenchmarkLogitLoggerJsonHandler-2         800017              1437 ns/op             120 B/op          3 allocs/op
BenchmarkLogitLoggerPrint-2               751623              1567 ns/op              48 B/op          1 allocs/op
BenchmarkSlogLoggerTextHandler-2          725522              1629 ns/op               0 B/op          0 allocs/op
BenchmarkSlogLoggerJsonHandler-2          583214              2030 ns/op             120 B/op          3 allocs/op
BenchmarkZeroLogLogger-2                 1929276               613 ns/op               0 B/op          0 allocs/op
BenchmarkZapLogger-2                      976855              1168 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusLogger-2                   231723              4927 ns/op            2080 B/op         32 allocs/op

BenchmarkLogitFile-2                      454489              2366 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBuffer-2           1038120              1154 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBatch-2            1026002              1179 ns/op               0 B/op          0 allocs/op
BenchmarkSlogFile-2                       407590              2944 ns/op               0 B/op          0 allocs/op
BenchmarkZeroLogFile-2                    634375              1810 ns/op               0 B/op          0 allocs/op
BenchmarkZapFile-2                        382790              2641 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusFile-2                     174944              6491 ns/op            2080 B/op         32 allocs/op
```

> 注：WithBuffer 和 WithBatch 分别是使用了缓冲器和批量写入的方式进行测试。

> 测试文件：[_examples/performance_test.go](./_examples/performance_test.go)

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

[![Star History Chart](https://api.star-history.com/svg?repos=fishgoddess/logit&type=Date)](https://star-history.com/#fishgoddess/logit&Date)
