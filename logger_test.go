// Copyright 2020 Ye Zi Jie. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/02/29 16:41:41

package logit

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
    "testing"
    "time"
)

// 测试日志记录器的 Debug 方法
func TestLoggerDebug(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Debug("这是 debug 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerInfo(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Info("这是 info 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerWarn(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Warn("这是 warn 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerError(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Error("这是 error 信息。。。")
}

// 测试调整日志记录器为运行状态的方法
func TestLoggerEnable(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Debug("1. 这是 debug 信息。。。")
    logger.ChangeLevelTo(OffLevel)
    logger.Debug("2. 这是 debug 信息。。。")
    logger.ChangeLevelTo(DebugLevel)
    logger.Debug("3. 这是 debug 信息。。。")
}

// 测试日志记录器的级别控制是否可用
func TestLevel(t *testing.T) {
    logger := NewLogger(os.Stdout, WarnLevel)
    logger.Warn(logger.Level().String())
    logger.Info("这条 info 级别的内容可以显示吗？")
    logger.Error("这条 error 级别的内容可以显示吗？")
    logger.Warn("这条 warn 级别的内容可以显示吗？")
}

// 测试更改日志级别是否可用
func TestLoggerChangeLevelTo(t *testing.T) {
    logger := NewLogger(os.Stdout, WarnLevel)
    logger.Info("Log level is warn, so info message will not be logged!")

    logger.ChangeLevelTo(InfoLevel)
    logger.Info("Now info message will be logged!")

    logger.ChangeLevelTo(ErrorLevel)
    logger.Warn("Now only error messages will be logged!")
}

// 测试使用日志处理器创建日志记录器的方法
func TestNewLoggerWithHandlers(t *testing.T) {

    defer func() {
        err := recover()
        if err == nil {
            t.Fatal("没有传入日志处理器，应该报错的，但是没有报！")
        }
    }()

    NewLoggerWithHandlers(os.Stdout, DebugLevel)
}

// 测试文件信息显示的开关是否可用
func TestLoggerEnableAndDisableFileInfo(t *testing.T) {
    logger := NewLogger(os.Stdout, WarnLevel)
    logger.Warn("没有文件信息！")
    logger.EnableFileInfo()
    logger.Warn("有文件信息？是否正确？")
    logger.DisableFileInfo()
    logger.Warn("现在应该没有文件信息了吧！")
}

// 测试增加处理器是否可用
func TestLoggerAddHandlersAndSetHandlers(t *testing.T) {
    logger := NewLogger(os.Stdout, InfoLevel)
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))

    handlers1 := func(logger *Logger, level Level, now time.Time, msg string) bool {
        logger.writer.Write([]byte("第一个日志处理器！\n"))
        return true
    }

    handlers2 := func(logger *Logger, level Level, now time.Time, msg string) bool {
        logger.writer.Write([]byte("第二个日志处理器！\n"))
        return true
    }

    logger.AddHandlers(handlers1, handlers2)
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))

    ok := logger.SetHandlers()
    if ok {
        t.Fatal("SetHandlers 应该返回 false！")
    }
    logger.SetHandlers(handlers1, handlers2)
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))
}

// 测试更改时间格式化标准的方法
func TestLoggerSetFormatOfTime(t *testing.T) {
    logger := NewLogger(os.Stdout, InfoLevel)
    logger.Info("当前时间格式化信息！" + logger.FormatOfTime())
    logger.SetFormatOfTime("2006年01月02日 15点04分05秒")
    logger.Info("更改之后的时间格式化信息！" + logger.FormatOfTime())
}

// 测试获取日志处理器的方法
func TestLoggerHandlers(t *testing.T) {
    logger := NewStdoutLogger(InfoLevel)
    handlers := logger.Handlers()

    // 需要判断地址是否一样，一样说明有问题
    if &handlers[0] == &logger.handlers[0] {
        t.Fatal("handlers 获取的数据和底层数据地址一样！这是有问题的！")
    }

    // 显示每个日志处理器的地址
    logger.Info("handlers 个数：" + strconv.Itoa(len(handlers)))
    logger.InfoFunc(func() string {
        builder := strings.Builder{}
        for _, handler := range handlers {
            builder.WriteString(fmt.Sprintf("%p ", &handler))
        }
        return builder.String()
    })

    // 尝试非法篡改日志处理器
    handlers[0] = func(logger *Logger, level Level, now time.Time, msg string) bool {
        fmt.Println("哈哈哈被非法篡改了！！！")
        return false
    }
    logger.Warn("这条信息需要显示出来！")
}

// 测试更改写出器的方法
func TestLoggerChangeWriterTo(t *testing.T) {
    logger := NewStdoutLogger(DebugLevel)
    fmt.Println(logger.writer)
    logger.Debug("哈哈哈")

    file, _ := os.Create("Z:/TestLoggerChangeWriterTo.log")
    logger.ChangeWriterTo(file)
    fmt.Println(logger.writer)
    logger.Debug("哈哈哈文件？")
}

// 测试并发情况下使用 Logger
func TestLoggerInConcurrency(t *testing.T) {

    logger := NewFileLogger("Z:/TestLoggerInConcurrency.log", DebugLevel)
    //logger := NewStdoutLogger(DebugLevel)

    group := sync.WaitGroup{}
    for i := 0; i < 100; i++ {
        group.Add(1)
        go func(num int) {
            if num == 30 || num == 60 {
                logger.ChangeLevelTo(InfoLevel)
            }
            logger.Info(strconv.Itoa(num))
            if num == 60 || num == 90 {
                logger.ChangeWriterTo(os.Stdout)
            }
            group.Done()
        }(i)
    }

    group.Wait()
}
