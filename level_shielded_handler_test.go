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
// Created at 2020/05/27 21:10:13

package logit

import "testing"

// 测试屏蔽日志级别的日志处理器
func TestLevelShieldedHandler(t *testing.T) {
	logger := NewLoggerFromPath("./_examples/level_shielded_handler.conf")
	logger.Debug("debug 有几条？")
	logger.Info("info 有几条？")
	logger.Warn("warn 有几条？")
	logger.Error("error 有几条？")
}
