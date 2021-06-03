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
// Created at 2020/03/02 20:51:29

package main

import (
	"os"
	"path/filepath"
	"testing"
	//"time"

	"github.com/FishGoddess/logit"
	//"github.com/kataras/golog"
	//"github.com/sirupsen/logrus"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

// 时间格式化字符串
const timeFormat = "2006-01-02 15:04:05"

type nopWriter struct{}

func (w *nopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

/*
BenchmarkLogitLogger-16          3950809               917 ns/op             128 B/op          4 allocs/op

BenchmarkGologLogger-16          2191954              1646 ns/op             713 B/op         24 allocs/op

BenchmarkZapLogger-16            1650854              2182 ns/op             449 B/op         16 allocs/op

BenchmarkLogrusLogger-16          922870              3898 ns/op            1634 B/op         52 allocs/op

**********************************************************************************************************

BenchmarkLogitFile-16             486379              7202 ns/op             128 B/op          4 allocs/op

BenchmarkGologFile-16             264644             13595 ns/op             713 B/op         24 allocs/op

BenchmarkZapFile-16               421682              8523 ns/op             449 B/op         16 allocs/op

BenchmarkLogrusFile-16            346075             10300 ns/op            1634 B/op         52 allocs/op
*/

// 测试 logit 日志记录器的速度
func BenchmarkLogitLogger(b *testing.B) {

	// 测试用的日志记录器
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.SetWriter(&nopWriter{})

	// 测试用的日志任务
	logTask := func() {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warn("warning...")
		logger.Error("error...")
	}
	//logTask := func() {
	//	logger.Debug("debug%s", "...")
	//	logger.Info("info%s", "...")
	//	logger.Warn("warning%s", "...")
	//	logger.Error("error%s", "...")
	//}

	// 开始性能测试
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// 测试 golog 日志记录器的速度
//func BenchmarkGologLogger(b *testing.B) {
//
//	logger := golog.New()
//	logger.SetOutput(&nopWriter{})
//	logger.SetLevel("debug")
//	logger.SetTimeFormat(timeFormat)
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	// 开始性能测试
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// 测试 zap 日志记录器的速度
//func BenchmarkZapLogger(b *testing.B) {
//
//	// 测试用的日志记录器
//	config := zap.NewProductionEncoderConfig()
//	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//		enc.AppendString(t.Format(timeFormat))
//	}
//	encoder := zapcore.NewConsoleEncoder(config)
//	nopWriteSyncer := zapcore.AddSync(&nopWriter{})
//	core := zapcore.NewCore(encoder, nopWriteSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	// 开始性能测试
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// 测试 logrus 日志记录器的速度
//func BenchmarkLogrusLogger(b *testing.B) {
//
//	logger := logrus.New()
//	logger.SetOutput(&nopWriter{})
//	logger.SetLevel(logrus.DebugLevel)
//	logger.SetFormatter(&logrus.TextFormatter{
//		TimestampFormat: timeFormat,
//	})
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}

// ******************************************************

// 创建文件
func createFileOf(filePath string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filePath), 0644)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// 测试 logit 文件日志记录器的速度
func BenchmarkLogitFile(b *testing.B) {

	file, _ := createFileOf("Z:/BenchmarkLogitFile.log")
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.SetWriter(file)

	// 测试用的日志任务
	logTask := func() {
		logger.Debug("debug...")
		logger.Info("info...")
		logger.Warn("warning...")
		logger.Error("error...")
	}

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// 测试 golog 文件日志记录器的速度
//func BenchmarkGologFile(b *testing.B) {
//
//	logger := golog.New()
//	file, _ := createFileOf("Z:/BenchmarkGologFile.log")
//	logger.SetOutput(file)
//	logger.SetLevel("debug")
//	logger.SetTimeFormat(timeFormat)
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	// 开始性能测试
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// 测试 zap 文件日志记录器的速度
//func BenchmarkZapFile(b *testing.B) {
//
//	// 测试用的日志记录器
//	config := zap.NewProductionEncoderConfig()
//	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//		enc.AppendString(t.Format(timeFormat))
//	}
//	encoder := zapcore.NewConsoleEncoder(config)
//	file, _ := createFileOf("Z:/BenchmarkZapFile.log")
//	writeSyncer := zapcore.AddSync(file)
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	// 开始性能测试
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// 测试 logrus 文件日志记录器的速度
//func BenchmarkLogrusFile(b *testing.B) {
//
//	logger := logrus.New()
//	file, _ := createFileOf("Z:/BenchmarkLogrusFile.log")
//	logger.SetOutput(file)
//	logger.SetLevel(logrus.DebugLevel)
//	logger.SetFormatter(&logrus.TextFormatter{
//		TimestampFormat: timeFormat,
//	})
//
//	// 测试用的日志任务
//	logTask := func() {
//		logger.Debug("debug...")
//		logger.Info("info...")
//		logger.Warn("warning...")
//		logger.Error("error...")
//	}
//
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
