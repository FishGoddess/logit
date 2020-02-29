# 📝 logit

**logit** is a easy-to-use and level-based logger for [GoLang](https://golang.org) applications.

[阅读中文版的 Read me](./README.md).

### 🥇 Features

* level-based logging, and there are four levels to use

### 🚀 Installation

The only requirement is the [Golang Programming Language](https://golang.org).

> Go modules

```bash
$ go get github.com/FishGoddess/logit@v0.0.1
```

Or edit your project's go.mod file and execute _**go build**_.

```bash
module your_project_name

go 1.14

require (
    github.com/FishGoddess/logit v0.0.1
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
    "os"
    
    "github.com/FishGoddess/logit"
)

func main() {
    logger := logit.NewLogger(os.Stdout, logit.DebugLevel)

    // Then you will be easy to log!
    logger.Debug("this is a debug message!")
    logger.Info("this is a info message!")
    logger.Warning("this is a warning message!")
    logger.Error("this is a error message!")
}
```

### 📖 Examples

* [basic](./_examples/basic.go)
* [enable_disable](./_examples/enable_disable.go)

_Check more examples in [_examples](./_examples)._ 

### 🔥 Benchmarks

```bash
$ go test -v -bench=. -benchtime=20s
```

| test case | times ran (large is better) |  ns/op (small is better) | B/op (small is better) | allocs/op (small is better) |
| -----------|--------|-------------|-------------|-------------|
| **[logit](./logger_test.go)** | 4800000 | 5062 ns/op | 864 B/op | 8 allocs/op |
| [Golang log](./logger_test.go) | 5400000 | 4730 ns/op | 928 B/op | 12 allocs/op |

### 👥 Contributing

If you find that something is not working as expected please open an _**issue**_.

### 📦 Projects using logit

| Project | Author | Description |
| -----------|--------|-------------|
|  |  |  |

