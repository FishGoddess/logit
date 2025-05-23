## ✒ 未来版本的新特性 (Features in future versions)

### v1.8.x

* [x] 提高单元测试覆盖率到 80%
* [x] 增加快速时钟，可以非常快速地查询时间
* [ ] 提高单元测试覆盖率到 90%
* [ ] 增加按天日期进行分裂的文件写出器

### v1.5.x

* [x] 考虑结合 Go1.21 新加的 slog 包，做一些生产化的适配和功能
* [x] 增加 context 适配功能
* [x] 增加 Production 现成配置，搭配 Option 方便使用
* [x] 完善 Logger 相关的单元测试
* [x] 调整 writer 包代码，完善单元测试
* [x] 调整 rotate 包代码，完善单元测试
* [x] 完善 config 包功能和单元测试
* [x] 完善性能测试
* [x] 完善示例代码
* [x] 增加属性解析器适配功能
* [x] 优化 handler 和 writer 设计
* [x] 增加一个高可读性且高性能的 handler 实现
* [x] TapeHandler 转义处理
* [x] 提高单元测试覆盖率到 70%

### v1.2.x

* [ ] ~~考虑增加类似于 scope 的机制，可以在 logit 中装载超过一个 global 的 logger~~
* [ ] ~~考虑增加 Production 之类不同级别的现成配置 options，方便使用~~

### v1.1.x

* [x] 调整 global logger 的设计，去除 caller + 1 的代码

### v1.0.0

* [x] 稳定的 API

### v0.5.x

* [x] ~~去掉讨厌的 End() 方法，考虑增加一个 Log 方法，然后直接以 logger.Log(log.Debug("test").Int64("userID", 123)) 这样的形式改造~~

> 尝试过这种 API 形式，但是实现起来有些别扭，使用起来也很奇怪，虽然没有出现 End 方法，但为了避免 End 方法引入了别的方法，写起来反而繁琐。。

* [x] 代码优化和重构，包括性能和逻辑上的优化，代码质量上的重构
* [x] 支持增加调用堆栈或者函数名的信息到日志
* [x] 增加批量写入的 Writer，区别于缓冲写入的 Writer，这个批量是按照数量进行缓冲的
* [x] 考虑到标准库的时间格式化耗费的时间和内存都很多，可能会有这方面优化的需求和场景，抽离出获取时间的函数
* [x] ~~缓冲写出器增加分段锁机制增加并发性能，考虑使用 chan 重构设计，可以参考下 zap 的异步处理~~

> 经过 zap 源码的研究，发现和现有 buffer writer 的逻辑区别不大，而之前使用 chan 尝试过，性能并没有很好，IO 次数还是一样多，容易出现堆积情况。

* [x] 增加错误监控回调函数
* [x] 增加配置项解析功能，提供配置文件支持
* [x] 加入日志存活天数的特性
* [x] 加入日志存活个数的特性
* [x] 按时间、文件大小自动切割日志文件，并支持过期清理机制

### v0.4.x

* [x] 尝试和 Context 机制结合，传递信息
* [x] 重新支持配置文件，考虑抽离配置文件模块为独立仓库，甚至是设计一种新的配置文件
* [x] depth 放开，可以支持用户包裹
* [x] Log 增加 Caller 方法，方便某些场景下获取调用信息
* [x] 兼容 log 包部分方法，比如 Printf
* [x] 考虑下 interceptor 的实现，在 Log 中加入 context，然后 End 时调用 interceptor 即可

### v0.3.x

经过一段时间的实际使用，发现了一些不太方便的地方需要进行调整。

* [x] Handler 的设计太过于抽象，导致很多日志库本身的功能实现过于剥离、插件化
* [x] 类 Json 的配置文件容易嵌套过多，不方便看（这也是 Handler 抽象程度太高导致的）
* [x] 部分 API 使用不太方便，特别是和 Handler 相关的一些功能
* [x] 有些 API 的使用频率确实很低，原本已经屏蔽了一些，但目前的 API 列表还不够清爽
* [x] duration 和 size 的设计导致没办法同时使用，而且加新特性会越来越臃肿
* [ ] 加入日志存活天数的特性
* [ ] 加入日志存活个数的特性
* [ ] ~~使用多个变量替代 map，避免哈希消耗性能~~

> 这么做还是无法避免一层映射，如果真的要避免映射，就得对 log 方法进行比较大幅度的改造。

> 总结：原本让我引以为傲的 Handler 在长期的使用下来发现很蛋疼，
> 优点是有，但麻烦也不少，所以需要改造！

### v0.0.x - v0.2.x

* [x] 实现基础的日志功能
* [ ] ~~引入 RollinHook 组件~~

> 取消这个特性是因为，它的代码入侵太严重了，并且会使代码设计变得很复杂，使用体验也会变差。
> 为了这样一个扩展特性要去改动核心特性的代码，有些喧宾夺主了，所以最后取消了这个组件。

* [ ] ~~修复配置文件中出现转义字符导致解析出错的问题~~

> 取消这个特性是因为，配置文件是用户写的，如果存在转义字符的问题，用户自行做转义会更适合一点。

* [ ] ~~增加 timeout_handler.go，里面是带超时功能的日志处理器包装器~~

> 取消这个特性是因为，一般在需要获取某个执行时间很长甚至可能一直阻塞的操作的结果时才需要超时，
> 对于日志输出而言，我们并不需要获取日志输出操作的结果，所以这个特性意义不大。
> 还有一个原因就是，实现超时需要使用并发，在超时的任务里终止某个任务，
> 而 Go 语言并没有提供可以停止并销毁一个 goroutine 的方法，所以即使超时了，也没有办法终止这个任务
> 甚至可能造成 goroutine 的阻塞。综合上述，取消这个超时功能的日志处理器包装器。

* [ ] ~~结合上面几点，以 “并发、缓冲” 为特点进行设计，技术使用 writer 接口进行包装~~

> 取消这个特性是因为，经过实验，性能并没有改善多少，两个方案，一个是使用队列进行缓冲，
> 开启一个 Go 协程写出数据，一个是不缓冲，直接开启多个 Go 协程进行写出数据。
> 第一个方案中，性能反而下降了一倍，估计和内存分配有关，最重要的就是，如果队列还有缓冲数据，
> 但是程序崩了，就有可能导致数据丢失，日志往往就需要最新最后的数据，而丢失的也正是最新最后的数据，
> 所以这个方案直接否决。第二个方案中，性能几乎没有提升，而且导致日志输出的顺序不固定，也有可能丢失。
> 综合上述，取消这个并发化日志输出的特性。

* [ ] ~~给日志输出增加颜色显示~~

> 取消颜色是因为考虑到线上生产环境主要使用文件，这个终端颜色显示的特性不是这么必须。
> 如果要实现，还要针对不同的操作系统处理，代价大于价值，所以废弃这个新特性。