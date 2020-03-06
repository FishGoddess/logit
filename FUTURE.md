## ✒ 未来版本的新特性 (Features in future version)

### v0.0.7
* 重构日志输出的模块，抛弃了标准库的 log 设计
* 增加日志处理器模块，支持用户自定义日志处理逻辑，大大地提高了扩展能力
* 支持不输出文件信息，避免 runtime.Caller 方法的调用，大大地提高了性能
* 支持调整时间格式化输出，让用户自定义时间输出的格式

### v0.0.6
* 支持按照文件大小自动划分日志文件
* 修复 nextFilename 中随机数生成重复的问题，设置了纳秒时钟作为种子
* ~~结合上面几点，以 “并发、缓冲” 为特点进行设计，技术使用 writer 接口进行包装~~
    > 取消这个特性是因为，经过实验，性能并没有改善多少，两个方案，一个是使用队列进行缓冲，
    > 开启一个 Go 协程写出数据，一个是不缓冲，直接开启多个 Go 协程进行写出数据。
    > 第一个方案中，性能反而下降了一倍，估计和内存分配有关，最重要的就是，如果队列还有缓冲数据，
    > 但是程序崩了，就有可能导致数据丢失，日志往往就需要最新最后的数据，而丢失的也正是最新最后的数据，
    > 所以这个方案直接否决。第二个方案中，性能几乎没有提升，而且导致日志输出的顺序不固定，也有可能丢失。
    > 综合上述，取消这个并发化日志输出的特性。

### v0.0.5
* 支持将日志输出到文件
* 支持按照时间间隔自动划分日志文件

### v0.0.4
* 修改 **LogLevel** 类型为 **LoggerLevel**，命名更符合意义
* 更改了部分源文件的命名，也是为了更符合实际的意义

### v0.0.3
* 让信息输出支持占位符，比如 %d 之类的
* 修复 Logit 日志调用方法的调用深度问题，之前直接通过 logit 调用时文件信息会显示错误

### v0.0.2
* 扩展 Logger 的使用方法，主要是创建日志记录器一类的方法
* 扩展 logit 的全局使用方法，增加一个默认的日志记录器
* 支持更改日志级别，参考 Logger#ChangeLevelTo 方法
* 修复 Logger#log 方法中漏加读锁导致并发安全的问题

### v0.0.1
* 实现最简单的日志输出功能
* 支持四种日志级别：_debug_, _info_, _warn_, _error_
* 对应四种日志级别分别有四个方法
* ~~给日志输出增加颜色显示~~
    > 取消颜色是因为考虑到线上生产环境主要使用文件，这个终端颜色显示的特性不是这么必须。
    > 如果要实现，还要针对不同的操作系统处理，代价大于价值，所以废弃这个新特性。