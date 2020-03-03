# 📝 logit

[![License](./license.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

**logit** is a easy-to-use and level-based logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* Level-based logging, and there are four levels to use
* Enable or disable Logger, you can disable or switch to a higher level in your production environment

_Check [HISTORY.md](./HISTORY.md) and [FUTURE.md](./FUTURE.md) to more information._

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
    github.com/FishGoddess/logit v0.0.3
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
    "github.com/FishGoddess/logit"
)

func main() {
    
    // Log as you want.
    logit.Debug("I am a debug message! But I will not be logged in default level!")
    logit.Info("I am an info message!")
    logit.Warning("I am a warning message!")
    logit.Error("I am an error message!")
    
    // Change log level.
    logit.ChangeLevelTo(logit.DebugLevel)

    // If you want format your message, just add arguments!
    logit.Info("format info message! id = %d, content = %s", 1, "info!")
}
```

### 📖 Examples

* [basic](./_examples/basic.go)
* [logger](./_examples/logger.go)
* [enable_disable](./_examples/enable_disable.go)
* [change_log_level](./_examples/change_log_level.go)

_Check more examples in [_examples](./_examples)._

### 🔥 Benchmarks

```bash
$ go test -v -bench=. -benchtime=20s
```

| test case | times ran (large is better) |  ns/op (small is better) | B/op (small is better) | allocs/op (small is better) |
| -----------|--------|-------------|-------------|-------------|
| **[logit](./_examples/benchmarks_test.go)** | 4405342 | 5409 ns/op | 904 B/op | 12 allocs/op |
| [logrus](./_examples/benchmarks_test.go) | 2990408 | 7991 ns/op | 1633 B/op | 52 allocs/op |
| [Golang log](./_examples/benchmarks_test.go) | 5308578 | 4539 ns/op | 920 B/op | 12 allocs/op |
| [Golog](./_examples/benchmarks_test.go) | 15536137 | 1556 ns/op | 232 B/op | 16 allocs/op |

> Environment：I7-6700HQ CPU @ 2.6 GHZ, 16 GB RAM

_Logit is based on Golang log, so it looks like not better than Golang log. Don't worry, we will redesign the implement of log output!_

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

