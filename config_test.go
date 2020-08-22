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
// Created at 2020/03/29 22:59:14

package logit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// 创建 TestParseConfigFile 测试案例的配置文件
func createParseConfigFileTestConfigFile(t *testing.T) string {

	// 创建配置文件
	configFile, err := ioutil.TempFile("", "TestParseConfigFile_*.conf")
	if err != nil {
		t.Fatal(err)
	}
	defer configFile.Close()

	// 写入配置内容
	configFile.WriteString(`
    {
        # 这是注释
        // 这也是注释
		"level": "debug",
		
		"caller": false,
		
		"handlers": {
		   "console": {
			   "timeFormat": "unix",
			   "encoder": "json"
		   },
		   "file": {
			   "path": "` + escapeString(filepath.Join(os.TempDir(), "logit.log")) + `"
		   }
		}
    }
	`)
	return configFile.Name()
}

// 测试解析配置文件的方法
func TestParseConfigFile(t *testing.T) {

	// 打开配置文件
	file, err := os.Open(createParseConfigFileTestConfigFile(t))
	if err != nil {
		t.Fatal(err)
	}

	// 解析配置文件
	conf, err := parseConfigFrom(file)
	if err != nil {
		t.Fatal("parseConfigFile 测试出现问题！", err)
	}

	t.Logf("%v\n", conf.Level)
	t.Logf("%v\n", conf.Caller)
	for name, parmas := range conf.Handlers {
		t.Logf("%s ==> %v\n", name, parmas)
	}
}

// 测试从 config 中解析日志处理器的方法
func TestParseHandlersFromConfig(t *testing.T) {

	handlers := parseHandlersFrom(config{
		Handlers: map[string]map[string]interface{}{
			"console": {
				"k1": "v1",
			},
			"file": {
				"path": escapeString(filepath.Join(os.TempDir(), "TestParseHandlersFromConfig.log")),
				"k2":   "v2",
			},
		},
	})
	for i, handler := range handlers {
		t.Logf("No.%d ==> %T\n", i+1, handler)
	}
}
