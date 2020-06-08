# 📝 logit

[![License](_icon/license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![Go Doc](_icon/godoc.svg)](https://pkg.go.dev/github.com/FishGoddess/logit?tab=doc)


**logit** 是一个简单易用并且是基于级别控制和配置文件的日志库，可以应用于所有的 [GoLang](https://golang.org) 应用程序中。

[Read me in English](./README.en.md)

[B站上的介绍视频](https://www.bilibili.com/video/BV14t4y1y7rF)

### 🥇 功能特性

* 独特的日志输出模块设计，使用 wrapper 和 handler 装载特定的模块，实现扩展功能
* 支持日志级别控制，一共有四个日志级别，分别是 debug，info，warn 和 error
* 支持配置文件，可以让用户在项目编译成二进制之后还能灵活地控制日志的设置
* 支持日志记录函数，使用回调的形式获取日志内容，对长日志内容的组织逻辑会更清晰
* 支持开启或者关闭日志功能，线上环境可以关闭或调高日志级别
* 支持记录日志到文件中，并且可以自定义日志文件名
* 支持按照时间间隔进行自动划分日志文件，比如每一天划分一个日志文件
* 支持按照文件大小进行自动划分日志文件，比如每 64 MB 划分一个日志文件
* 增加日志处理器模块，支持用户自定义日志处理逻辑，具有很高的扩展能力
* 支持不输出文件信息，避免 runtime.Caller 方法的调用，具有很高的性能
* 支持调整时间格式化输出，让用户自定义时间输出的格式
* 支持以 Json 形式输出日志信息，更方便后续对日志进行解析

_历史版本的特性请查看 [HISTORY.md](./HISTORY.md)。未来版本的新特性和计划请查看 [FUTURE.md](./FUTURE.md)。_

> v0.1.x 及以下版本已经停止维护，请尽快升级到 v0.2.x 版本！您将感受到全新的使用体验，并可以享受长期的更新和维护！

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
    github.com/FishGoddess/logit v0.2.5-alpha
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
	"math/rand"
	"strconv"
	"time"

	"github.com/FishGoddess/logit"
)

func main() {

	// Log messages with four levels.
	logit.Debug("I am a debug message!")
	logit.Info("I am an info message!")
	logit.Warn("I am a warn message!")
	logit.Error("I am an error message!")

	// Notice that logit has blocked some methods for more refreshing method list.
	// If you want to use some higher level methods, you should call logit.Me() to
	// get the fully functional logger, then call what you want to call.
	// For example, if you want to output log with file info, try this:
	logit.Me().EnableFileInfo()
	logit.Info("Show file info!")

	// If you have a long log and it is made of many variables, try this:
	// The msg is the return value of msgGenerator.
	logit.DebugFunc(func() string {
		// Use time as the source of random number generator.
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return "debug rand int: " + strconv.Itoa(r.Intn(100))
	})

	// Or you can use formatting method like this:
	logit.Debugf("This is a debug msg with %d params: %s, %s", 2, "msgFormat", "msgParams")
	logit.Infof("This is a debug msg with %d params: %s, %s", 2, "msgFormat", "msgParams")
	logit.Warnf("This is a debug msg with %d params: %s, %s", 2, "msgFormat", "msgParams")
	logit.Errorf("This is a debug msg with %d params: %s, %s", 2, "msgFormat", "msgParams")

	// If a config file "logit.conf" in "./", then logit will load it automatically.
	// This is more convenience to use config file and logger.
}
```

### 📖 参考案例

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [level_and_disable](./_examples/level_and_disable.go)
* [config_file](./_examples/config_file.go)
* [handler](./_examples/handler.go)
* [writer](./_examples/writer.go)
* [log_to_file](./_examples/log_to_file.go)

_更多使用案例请查看 [_examples](./_examples) 目录。_

_配置文件模板请查看 [_examples/config](./_examples/config) 目录。_

### 🔥 性能测试

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=10s
```

> 测试文件：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| **logit** | **6429907** | **1855 ns/op** | 384 B/op | 8 allocs/op |
| golog | 3361483 | 3589 ns/op | 712 B/op | 24 allocs/op |
| zap | 2971119 | 4066 ns/op | 448 B/op | 16 allocs/op |
| logrus | 1553419 | 7869 ns/op | 1633 B/op | 52 allocs/op |

> 测试环境：I7-6700HQ CPU @ 2.6 GHZ，16 GB RAM

**注意：**

**1. 输出文件信息会有运行时操作（runtime.Caller 方法），非常影响性能，**
**但是这个功能感觉还是比较实用的，尤其是在查找错误的时候，所以我们还是加了这个功能！**
**如果你更在乎性能，那我们也提供了一个选项可以关闭文件信息的查询！**

**2. v0.0.7 及以前版本的日志输出使用了 fmt 包的一些方法，经过性能检测发现这些方法存在大量使用反射的**
**行为，主要体现在对参数 v interface{} 进行类型检测的逻辑上，而日志输出都是字符串，这一个**
**判断是可以省略的，可以减少很多运行时操作时间！v0.0.8 版本开始使用了更有效率的输出方式！**

**3. 经过对 v0.0.8 版本的性能检测，发现时间格式化操作消耗了接近一半的处理时间，**
**主要体现在 time.Time.AppendFormat 的调用上。在 v0.0.11 版本中使用了时间缓存机制进行优化，**
**目前存在一个疑惑就是使用并发竞争去换取时间格式化的性能消耗究竟值不值得？**
**答案是不值得，我们在 v0.1.1-alpha 及更高版本中取消了这个时间缓存机制。**

**4. 值得注意的是，Debugf 一类带格式化的 API 性能达不到这个水平，因为还是使用了反射技术，但是性能依旧是不差的：**

| 测试 | 单位时间内运行次数 (越大越好) |  每个操作消耗时间 (越小越好) | B/op (越小越好) | allocs/op (越小越好) |
| -----------|--------|-------------|-------------|-------------|
| logit | **6429907** | **1855 ns/op** | 384 B/op | 8 allocs/op |
| **logit-使用反射技术 ** | **5288931** | **2334 ns/op** | 424 B/op | 12 allocs/op |

### 👥 贡献者

如果您觉得 logit 缺少您需要的功能，请不要犹豫，马上参与进来，发起一个 _**issue**_。

### 📦 使用 logit 的项目

| 项目 | 作者 | 描述 |
| -----------|--------|-------------|
|  |  |  |

