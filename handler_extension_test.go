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
// Created at 2020/03/03 22:48:34

package logit

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/FishGoddess/logit/files"
)

// 测试创建文件日志处理器
func TestNewFileHandler(t *testing.T) {

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("创建文件日志处理器测试出现问题！")
		}
	}()

	logger := NewLogger(DebugLevel, NewFileHandler(filepath.Join(os.TempDir(), "test.log"), TextEncoder(), ""))
	for i := 0; i < 100; i++ {
		logger.Info("我是第 " + strconv.Itoa(i) + " 条日志！")
	}

	logger = NewLogger(DebugLevel, NewFileHandler("https://test.io", TextEncoder(), ""))
}

// 测试创建随时间间隔滚动的文件日志处理器
func TestNewDurationRollingHandler(t *testing.T) {
	logger := NewLogger(DebugLevel, NewDurationRollingHandler(os.TempDir(), time.Second, TextEncoder(), ""))
	for i := 0; i < 5; i++ {
		logger.Info("1. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
		time.Sleep(time.Second)
		logger.Info("2. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
	}
}

// 测试按照文件大小自动划分日志文件的日志处理器
func TestNewSizeRollingHandler(t *testing.T) {
	logger := NewLogger(DebugLevel, NewSizeRollingHandler(os.TempDir(), 64*files.KB, TextEncoder(), ""))
	for i := 0; i < 2000; i++ {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warn("warn...")
		logger.Error("error...")
	}
}

// 测试创建空的 RollingHandlerOptions
func TestRollingHandlerOptions(t *testing.T) {
	options := RollingHandlerOptions{}
	t.Log(options)
	if options.nameGenerator != nil {
		t.Fatal("nameGenerator 应该是 nil 的，结果不是。。。")
	}
	if options.rollingHook != nil {
		t.Fatal("rollingHook 应该是 nil 的，结果不是。。。")
	}
}
