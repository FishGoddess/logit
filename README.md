# 📝 logit

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**logit** 是一个简单易用并且是基于级别控制的日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md).

### 🥇 功能特性

* 独特的日志输出模块设计，使用 wrapper 和 handler 装载特定的模块，实现扩展功能
* 支持日志级别控制，一共有四个日志级别，分别是 debug，info，warn 和 error。
* 支持开启或者关闭日志功能，线上环境可以关闭或调高日志级别
* 支持记录日志到文件中，自定义日志文件名
* 支持按照时间间隔进行自动划分日志文件，比如每一天划分一个日志文件
* 支持按照文件大小进行自动划分日志文件，比如每 64 MB 划分一个日志文件
* 增加日志处理器模块，支持用户自定义日志处理逻辑，具有很高的扩展能力
* 支持不输出文件信息，避免 runtime.Caller 方法的调用，具有很高的性能
* 支持调整时间格式化输出，让用户自定义时间输出的格式

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

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
    github.com/FishGoddess/logit v0.0.7
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
* [log_to_file](./_examples/log_to_file.go)
* [wrapper](./_examples/wrapper.go)
* [handler](./_examples/logger_handler.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=20s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| 测试 | 单位时间内运行次数 (large is better) |  ns/op (small is better) | B/op (small is better) | allocs/op (small is better) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | &nbsp; 4405342 | 5409 ns/op | &nbsp; 904 B/op | 12 allocs/op |
| **logit (关闭文件信息)** | 20341443 | 1130 ns/op | &nbsp; &nbsp; 32 B/op | &nbsp; 4 allocs/op |
| logrus | &nbsp; 2990408 | 7991 ns/op | 1633 B/op | 52 allocs/op |
| Golang log | &nbsp; 5308578 | 4539 ns/op | &nbsp; 920 B/op | 12 allocs/op |
| Golog | 15536137 | 1556 ns/op | &nbsp; 232 B/op | 16 allocs/op |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

**注意：golog 库是不会输出文件信息的，也就是少了运行时操作（runtime.Caller 方法），性能自然会高很多，**
**但是这个功能感觉还是比较实用的，尤其是在查找错误的时候，所以我们还是加了这个功能！**
**如果你更在乎性能，那我们也提供了一个选项可以关闭文件信息的查询（开发中）！**

_由于目前的 logit 是基于 Golang log 的，所以成绩相比更差，后续会重新设计内部日志输出模块，所以当前成绩仅供参考！_

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 |
| -----------|--------|-------------|
|  |  |  |

