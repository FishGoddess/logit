# 📝 logit

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**logit** is an easy-to-use and level-based logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* Modularization design, easy to extend your logger with wrapper and handler
* Level-based logging, and there are four levels to use
* Log Function supports, it is a better way to output a very long log
* Enable or disable Logger, you can disable or switch to a higher level in your production environment
* Log file supports, and you can customer the name of your log file
* Duration rolling supports, which means it will roll to a new log file by duration automatically, such as one day one log file
* File size rolling supports, which means it will roll to a new log file by file size automatically, such as one 64 MB one log file
* Log handler supports, you can extend logger with your own log handler easily
* High-performance supports, by avoiding to call runtime.Caller
* Time format supports, you can format time in your way

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to get more information._

### 🚀 Installation

The only requirement is the [Golang Programming Language](https://golang.org).

> Go modules

```bash
$ go get -u github.com/FishGoddess/logit
```

Or edit your project's go.mod file and execute _**go build**_.

```bash
module your_project_name

go 1.14

require (
    github.com/FishGoddess/logit v0.0.10
)
```

> Go path

```bash
$ go get -u github.com/FishGoddess/logit
```

logit has no more external dependencies.

```go
package main

import (
    "math/rand"
    "strconv"
    "time"
    
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

    // If you want to output log with file info, try this:
    logit.EnableFileInfo()
    logit.Info("Show file info!")

    // If you have a long log and it is made of many variables, try this:
    // The msg is the return value of msgGenerator.
    logit.DebugFunc(func() string {
        // Use time as the source of random number generator.
        r := rand.New(rand.NewSource(time.Now().Unix()))
        return "debug rand int: " + strconv.Itoa(r.Intn(100))
    })
}
```

### 📖 Examples

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [enable_disable](./_examples/enable_disable.go)
* [change_log_level](./_examples/change_log_level.go)
* [log_to_file](./_examples/log_to_file.go)
* [wrapper](./_examples/wrapper.go)
* [handler](./_examples/logger_handler.go)

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=1s
```

> Benchmark file：[_examples/benchmarks_test.go](./_examples/benchmarks_test.go)

| test case | times ran (large is better) |  ns/op (small is better) | features | extension |
| -----------|--------|-------------|-------------|-------------|
| **logit** | &nbsp; 572947 | 1939 ns/op | powerful | high |
| zap | &nbsp; 401043 | 2793 ns/op | normal | normal |
| logrus | &nbsp; 158262 | 7751 ns/op | normal | normal |
| golog | &nbsp; 751064 | 1614 ns/op | normal | normal |
| golang log | 1000000 | 1019 ns/op | not good | none |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

**Notice:**

**1. Fetching file info will call runtime.Caller, which is expensive.**
**However, we think file info is useful in check errors,**
**so we keep this feature, and provide a switch to turn off it for high-performance.**

**2. v0.0.7 and lower versions use some functions of fmt, and these functions is expensive**
**because of reflect (for judging the parameter v interface{}). Actually, these judgements**
**are redundant in a logger. The more effective output will be used in v0.0.8 and higher versions.**

**3. After checking the benchmarks of v0.0.8 version, we found that time format takes a lots of time**
**because of time.Time.AppendFormat. We are finding a more effective way to fix it, maybe fixed in higher version!**

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

