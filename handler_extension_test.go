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
	"io/ioutil"
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

func TestNewDurationRollingHandlerWithOptions(t *testing.T) {

	// 准备测试环境
	testDirectory, err := ioutil.TempDir("", "TestNewDurationRollingHandlerWithOptions_*")
	if err != nil {
		t.Fatal(err)
	}

	// 创建日志处理器
	handler := NewDurationRollingHandlerWithOptions(testDirectory, time.Second, TextEncoder(), "", RollingHandlerOptions{
		nameGenerator: func(directory string, now time.Time) string {
			return filepath.Join(directory, now.Format("2006年01月02日-15点04分05秒.log"))
		},
		rollingHook: files.NewLifeBasedRollingHook(testDirectory, 500*time.Millisecond),
	})

	// TODO 这个测试通不过，发现 rollingHook 中的 now.Sub(file.ModTime()) 是负数。。。
	// TODO 这个 RollingHook 机制需要取消掉，因为涉及到太多原本的代码了，考虑使用定时任务解决
	logger := NewLogger(DebugLevel, handler)
	logger.Info("1")
	time.Sleep(900 * time.Millisecond)
	logger.Info("2")
	time.Sleep(1 * time.Second)
	logger.Info("3")
	time.Sleep(1200 * time.Millisecond)
	logger.Info("4")

	infos, err := ioutil.ReadDir(testDirectory)
	if err != nil || len(infos) != 2 {
		t.Log("len(infos) == 2 ? ", len(infos) == 2)
		t.Fatal(err)
	}
}
