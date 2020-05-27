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

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// 创建 TestLevelShieldedHandler 测试案例的配置文件
func createLevelShieldedHandlerTestConfigFile(t *testing.T) string {

	// 创建配置文件
	configFile, err := ioutil.TempFile("", "createLevelShieldedHandlerTestConfigFile_*.conf")
	if err != nil {
		t.Fatal(err)
	}
	defer configFile.Close()

	// 写入配置内容
	configFile.WriteString(`
		"handlers": {
			"!debug": {
				"file": {
					"path": "` + escapeString(filepath.Join(os.TempDir(), "non-debug.log")) + `"
				}
			},
			"!info": {
				"file": {
					"path": "` + escapeString(filepath.Join(os.TempDir(), "non-info.log")) + `"
				}
			},
			"!warn": {
				"file": {
					"path": "` + escapeString(filepath.Join(os.TempDir(), "non-warn.log")) + `"
				}
			},
			"!error": {
				"file": {
					"path": "` + escapeString(filepath.Join(os.TempDir(), "non-error.log")) + `"
				}
			},
			"console": {
				"encoder": "json"
			}
		}
	`)
	return configFile.Name()
}

// 测试屏蔽日志级别的日志处理器
func TestLevelShieldedHandler(t *testing.T) {
	logger := NewLoggerFromPath(createLevelShieldedHandlerTestConfigFile(t))
	logger.Debug("debug 有几条？")
	logger.Info("info 有几条？")
	logger.Warn("warn 有几条？")
	logger.Error("error 有几条？")
}
