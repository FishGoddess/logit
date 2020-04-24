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
// Created at 2020/03/03 22:48:34

package logit

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/FishGoddess/logit/writer"
)

// 测试创建文件日志记录器
func TestNewFileLogger(t *testing.T) {

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("创建文件日志记录器测试出现问题！")
		}
	}()

	logger := NewFileLogger("Z:/test.log")
	for i := 0; i < 100; i++ {
		logger.Info("我是第 " + strconv.Itoa(i) + " 条日志！")
	}

	logger = NewFileLogger("https://test.io")
}

// 测试创建随时间间隔滚动的文件日志记录器
func TestNewDurationRollingLogger(t *testing.T) {

	logger := NewDurationRollingLogger("Z:/", time.Second)
	for i := 0; i < 10; i++ {
		logger.Info("1. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
		time.Sleep(time.Second)
		logger.Info("2. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
	}
}

// 测试按天自动划分日志文件的日志记录器
func TestNewDayRollingLogger(t *testing.T) {

	logger := NewDayRollingLogger("Z:/")
	logger.Info("1. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
	time.Sleep(time.Second)
	logger.Info("2. info!!!!!!!! " + strconv.FormatInt(time.Now().Unix(), 10))
}

// 测试按照文件大小自动划分日志文件的日志记录器
func TestNewSizeRollingLogger(t *testing.T) {

	logger := NewSizeRollingLogger("Z:/", 64*writer.KB)
	for i := 0; i < 1000; i++ {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warn("warn...")
		logger.Error("error...")
	}
}

// 测试按照默认文件大小自动划分日志文件的日志记录器
func TestNewDefaultSizeRollingLogger(t *testing.T) {

	logger := NewDefaultSizeRollingLogger("Z:/")
	for i := 0; i < 1000; i++ {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warn("warn...")
		logger.Error("error...")
	}
}

// 测试输出日志是从函数中生成的几个方法
func TestLoggerLogFunction(t *testing.T) {
	logger := NewProductionLogger(os.Stdout)
	logger.ChangeLevelTo(DebugLevel)
	logger.DebugFunc(func() string {
		return "debug rand int: " + strconv.Itoa(rand.Intn(100))
	})
	logger.InfoFunc(func() string {
		return "info rand int: " + strconv.Itoa(rand.Intn(100))
	})
	logger.WarnFunc(func() string {
		return "warn rand int: " + strconv.Itoa(rand.Intn(100))
	})
	logger.ErrorFunc(func() string {
		return "error rand int: " + strconv.Itoa(rand.Intn(100))
	})

	// test escaping
	logger.Info(`test "double quotes"\t\b \u0003 \u0019 !!!!`)
}

// 测试从配置文件中创建一个 logger
func TestNewLoggerFromConfigFile(t *testing.T) {
	logger := NewLoggerFromConfigFile("./_examples/logger.cfg")
	logger.Info("Does it work?")
}
