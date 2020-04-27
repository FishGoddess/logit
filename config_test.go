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
// Created at 2020/03/29 22:59:14

package logit

import "testing"

// 测试解析配置文件的方法
func TestParseConfigFile(t *testing.T) {

	conf, err := parseConfigFile("./_examples/logger.conf")
	if err != nil {
		t.Fatal("parseConfigFile 测试出现问题！")
	}

	t.Logf("%v\n", conf.Level)
	t.Logf("%v\n", conf.Caller)
	for name, parmas := range conf.Handlers {
		t.Logf("%s ==> %v\n", name, parmas)
	}
}

// 测试从 config 中解析日志处理器的方法
func TestParseHandlersFromConfig(t *testing.T) {

	handlers := parseHandlersFromConfig(config{
		Handlers: map[string]map[string]interface{}{
			"console": {
				"k1": "v1",
			},
			"file": {
				"path": "Z:/TestParseHandlersFromConfig.log",
				"k2":   "v2",
			},
		},
	})
	for i, handler := range handlers {
		t.Logf("No.%d ==> %T\n", i+1, handler)
	}
}
