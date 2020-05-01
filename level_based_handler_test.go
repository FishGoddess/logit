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
// Email: fishinlove@163.com
// Created at 2020/04/27 23:40:12

package logit

import "testing"

// 测试基于 debug 日志级别的日志处理器
func TestDebugLevelHandler(t *testing.T) {
	logger := NewLoggerFromPath("./_examples/level_based_handler.conf")
	logger.Debug("debug 去哪了？")
	logger.Info("info 有一条？")
	logger.Warn("warn 有一条？")
	logger.Error("error 有一条？")
}

// 测试基于 info 日志级别的日志处理器
func TestInfoLevelHandler(t *testing.T) {
	logger := NewLoggerFromPath("./_examples/level_based_handler.conf")
	logger.Debug("debug 有一条？")
	logger.Info("info 去哪了？")
	logger.Warn("warn 有一条？")
	logger.Error("error 有一条？")
}

// 测试基于 warn 日志级别的日志处理器
func TestWarnLevelHandler(t *testing.T) {
	logger := NewLoggerFromPath("./_examples/level_based_handler.conf")
	logger.Debug("debug 有一条？")
	logger.Info("info 有一条？")
	logger.Warn("warn 去哪了？")
	logger.Error("error 有一条？")
}

// 测试基于 debug 日志级别的日志处理器
func TestErrorLevelHandler(t *testing.T) {
	logger := NewLoggerFromPath("./_examples/level_based_handler.conf")
	logger.Debug("debug 有一条？")
	logger.Info("info 有一条？")
	logger.Warn("warn 有一条？")
	logger.Error("error 去哪了？")
}
