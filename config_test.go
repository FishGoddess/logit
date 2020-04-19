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

// 测试解析日志级别的方法
func TestParseLevel(t *testing.T) {
    level, err := ParseLevel("debug")
    if err != nil || level != DebugLevel {
        t.Fatal("ParseLevel 出现错误！")
    }

    level, err = ParseLevel("debug?")
    if err == nil {
        t.Fatal("ParseLevel 出现错误！")
    }
}

// 测试将 fileConfig 转换成 Config 的方法
func TestFileConfigToConfig(t *testing.T) {

    config, err := parseConfig(fileConfig{
        Level: "info",
        Handlers: map[string]map[string]interface{}{
            "json": {
                "timeFormat": "2006年01月02日 15点04分05秒",
            },
        },
    })
    if err != nil {
        t.Fatal(err.Error())
    }

    t.Logf("%v\n", config.Level)
    for _, handler := range config.Handlers {
        t.Logf("%T\n", handler)
    }

    logger := NewLoggerFrom(config)
    logger.Info("info 测试 TestFileConfigToConfig...")
    logger.Warn("warn 测试 TestFileConfigToConfig...")
}

// 测试解析配置文件的方法
func TestParseConfigFile(t *testing.T) {

    config, err := ParseConfigFile("./_examples/config/logit-config-template.cfg")
    if err != nil {
        t.Fatal(err.Error())
    }

    t.Logf("%v\n", config.Level)
    for _, handler := range config.Handlers {
        t.Logf("%T\n", handler)
    }

    logger := NewLoggerFrom(config)
    logger.Info("info 测试 TestFileConfigToConfig...")
    logger.Warn("warn 测试 TestFileConfigToConfig...")
}
