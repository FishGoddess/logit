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
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/02/29 16:41:41

package logit

import (
	"strconv"
	"sync"
	"testing"
)

// 测试日志记录器的 Debug 方法
func TestLoggerDebug(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(DebugLevel)
	logger.Debug("这是 debug 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerInfo(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(InfoLevel)
	logger.Info("这是 info 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerWarn(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(WarnLevel)
	logger.Warn("这是 warn 信息。。。")
}

// 测试日志记录器的 Debug 方法
func TestLoggerError(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(ErrorLevel)
	logger.Error("这是 error 信息。。。")
}

// 测试调整日志记录器为运行状态的方法
func TestLoggerEnable(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(DebugLevel)
	logger.Debug("1. 这是 debug 信息。。。")
	level := logger.ChangeLevelTo(OffLevel)
	logger.Debug("2. 这是 debug 信息。。。")
	logger.ChangeLevelTo(level)
	logger.Debug("3. 这是 debug 信息。。。")
}

// 测试日志记录器的级别控制是否可用
func TestLevel(t *testing.T) {
	logger := NewLogger()
	logger.Debug(logger.Level().String())
	logger.Info("这条 info 级别的内容可以显示吗？")
	logger.Warn("这条 warn 级别的内容可以显示吗？")
	logger.Error("这条 error 级别的内容可以显示吗？")
}

// 测试更改日志级别是否可用
func TestLoggerChangeLevelTo(t *testing.T) {
	logger := NewLogger()
	logger.Info("Logger's level is warn, so info message will not be logged!")

	logger.ChangeLevelTo(InfoLevel)
	logger.Info("Now info message will be logged!")

	logger.ChangeLevelTo(ErrorLevel)
	logger.Warn("Now only error messages will be logged!")
}

// 测试文件信息显示的开关是否可用
func TestLoggerEnableAndDisableFileInfo(t *testing.T) {
	logger := NewLogger()
	logger.Warn("没有文件信息！")
	logger.EnableFileInfo()
	logger.Warn("有文件信息？是否正确？")
	logger.DisableFileInfo()
	logger.Warn("现在应该没有文件信息了吧！")
}

// 测试并发情况下使用 Logger
func TestLoggerInConcurrency(t *testing.T) {

	logger := NewLogger()

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

// 测试带格式化的日志输出方法
func TestLoggerOutputWithFormat(t *testing.T) {
	logger := NewLogger()
	logger.ChangeLevelTo(DebugLevel)
	logger.EnableFileInfo()
	logger.DebugF("Debugf... %d %s %.3f", 123, "幸福呢", 123.123456)
	logger.InfoF("Infof... %d %s %.3f", 123, "幸福呢", 123.123456)
	logger.WarnF("Warnf... %d %s %.3f", 123, "幸福呢", 123.123456)
	logger.ErrorF("Errorf... %d %s %.3f", 123, "幸福呢", 123.123456)
}
