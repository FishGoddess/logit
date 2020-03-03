# 📝 logit

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**logit** 是一个简单易用并且是基于级别控制的日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md).

### 🥇 功能特性

* 支持日志级别控制，目前一共有四个日志级别
* 支持开启或者关闭日志功能，线上环境可以关闭或调高日志级别

_历史版本的特性请参考 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请参考 [FUTURE.md](./FUTURE.md)。_

### 🚀 安装方式

唯一需要的依赖就是 [Golang 运行环境](https://golang.org).

> Go modules

```bash
$ go get -u github.com/FishGoddess/logit
```

您也可以直接编辑 go.mod 文件，然后执行 _**go build**_。

```bash
module your_project_name

go 1.14

require (
    github.com/FishGoddess/logit v0.0.4
)
```

> Go path

```bash
$ go get -u github.com/FishGoddess/logit
```

logit 没有任何其他额外的依赖，纯使用 [Golang 标准库](https://golang.org) 完成。

```go
package main

import (
    "github.com/FishGoddess/logit"
)

func main() {
    
    // Log as you want.
    logit.Debug("I am a debug message! But I will not be logged in default level!")
    logit.Info("I am an info message!")
    logit.Warn("I am a warn message!")
    logit.Error("I am an error message!")
    
    // Change logger level.
    logit.ChangeLevelTo(logit.DebugLevel)

    // If you want format your message, just add arguments!
    logit.Info("format info message! id = %d, content = %s", 1, "info!")
}
```

### 📖 参考案例

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [enable_disable](./_examples/enable_disable.go)
* [change_log_level](./_examples/change_log_level.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v -bench=. -benchtime=20s
```

| 测试 | 单位时间内运行次数 (large is better) |  ns/op (small is better) | B/op (small is better) | allocs/op (small is better) |
| -----------|--------|-------------|-------------|-------------|
| **[logit](./_examples/benchmarks_test.go)** | 4405342 | 5409 ns/op | 904 B/op | 12 allocs/op |
| [logrus](./_examples/benchmarks_test.go) | 2990408 | 7991 ns/op | 1633 B/op | 52 allocs/op |
| [Golang log](./_examples/benchmarks_test.go) | 5308578 | 4539 ns/op | 920 B/op | 12 allocs/op |
| [Golog](./_examples/benchmarks_test.go) | 15536137 | 1556 ns/op | 232 B/op | 16 allocs/op |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

_由于目前的 logit 是基于 Golang log 的，所以成绩相比更差，后续会重新设计内部日志输出模块，所以当前成绩仅供参考！_

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 |
| -----------|--------|-------------|
|  |  |  |

