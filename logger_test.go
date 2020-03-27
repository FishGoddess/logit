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
)

// 测试日志记录器的 Debug 方法
func TestLoggerDebug(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Debug("这是 debug 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerInfo(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Info("这是 info 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerWarn(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Warn("这是 warn 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerError(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Error("这是 error 信息。。。")
}

// 测试调整日志记录器为运行状态的方法
func TestLoggerEnable(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Debug("1. 这是 debug 信息。。。")
    logger.ChangeLevelTo(OffLevel)
    logger.Debug("2. 这是 debug 信息。。。")
    logger.ChangeLevelTo(DebugLevel)
    logger.Debug("3. 这是 debug 信息。。。")
}

// 测试日志记录器的级别控制是否可用
func TestLevel(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Debug(logger.Level().String())
    logger.Info("这条 info 级别的内容可以显示吗？")
    logger.Warn("这条 warn 级别的内容可以显示吗？")
    logger.Error("这条 error 级别的内容可以显示吗？")
}

// 测试更改日志级别是否可用
func TestLoggerChangeLevelTo(t *testing.T) {
    logger := NewLogger(WarnLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Info("Logger's level is warn, so info message will not be logged!")

    logger.ChangeLevelTo(InfoLevel)
    logger.Info("Now info message will be logged!")

    logger.ChangeLevelTo(ErrorLevel)
    logger.Warn("Now only error messages will be logged!")
}

// 测试文件信息显示的开关是否可用
func TestLoggerEnableAndDisableFileInfo(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Warn("没有文件信息！")
    logger.EnableFileInfo()
    logger.Warn("有文件信息？是否正确？")
    logger.DisableFileInfo()
    logger.Warn("现在应该没有文件信息了吧！")
}

type myHandler struct{}

// Customize your own handler.
func (mh *myHandler) Handle(log *Log) bool {
    os.Stdout.WriteString("handler1: " + log.Msg() + "\n")
    return true
}

// 测试增加处理器是否可用
func TestLoggerAddHandlersAndSetHandlers(t *testing.T) {
    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))

    if len(logger.handlers) != 1 {
        t.Fatal("处理器个数不正确！")
    }

    logger.AddHandlers(&myHandler{})
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))

    if len(logger.handlers) != 2 {
        t.Fatal("处理器个数不正确！")
    }

    ok := logger.SetHandlers()
    if ok {
        t.Fatal("SetHandlers 应该返回 false！")
    }

    logger.SetHandlers(&myHandler{})
    logger.Info("当前的日志处理器：" + fmt.Sprintf("%v", logger.handlers))

    if len(logger.handlers) != 1 {
        t.Fatal("处理器个数不正确！")
    }
}

// 测试获取日志处理器的方法
func TestLoggerHandlers(t *testing.T) {
    logger := NewDevelopLogger()
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
    handlers[0] = &myHandler{}
    logger.Warn("这条信息需要显示出来！")
}

// 测试并发情况下使用 Logger
func TestLoggerInConcurrency(t *testing.T) {

    logger := NewLogger(DebugLevel, NewDefaultHandler(os.Stdout, NewDefaultEncoder(DefaultTimeFormat)))

    group := sync.WaitGroup{}
    for i := 0; i < 100; i++ {
        group.Add(1)
        go func(num int) {
            if num == 30 || num == 60 {
                logger.ChangeLevelTo(InfoLevel)
            }
            logger.Info(strconv.Itoa(num))
            group.Done()
        }(i)
    }

    group.Wait()
}
