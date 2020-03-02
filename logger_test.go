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
    "os"
    "testing"
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
func TestLoggerWarning(t *testing.T) {
    logger := NewLogger(os.Stdout, DebugLevel)
    logger.Warning("这是 warning 信息。。。")
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
    logger.Disable()
    logger.Debug("2. 这是 debug 信息。。。")
    logger.Enable()
    logger.Debug("3. 这是 debug 信息。。。")
}

// 测试日志记录器的级别控制是否可用
func TestLoggerLevel(t *testing.T) {
    logger := NewLogger(os.Stdout, WarningLevel)
    logger.Info("这条 info 级别的内容可以显示吗？")
    logger.Error("这条 error 级别的内容可以显示吗？")
    logger.Warning("这条 warning 级别的内容可以显示吗？")
}

// 测试更改日志级别是否可用
func TestLoggerChangeLevelTo(t *testing.T) {
    logger := NewStdoutLogger(WarningLevel)
    logger.Info("Log level is warning, so info message will not be logged!")

    logger.ChangeLevelTo(InfoLevel)
    logger.Info("Now info message will be logged!")

    logger.ChangeLevelTo(ErrorLevel)
    logger.Warning("Now only error messages will be logged!")
}
