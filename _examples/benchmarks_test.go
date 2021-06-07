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

	"github.com/FishGoddess/logit"
)

/*
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=3s

BenchmarkLogitLogger-16                  3775916               949 ns/op             128 B/op          4 allocs/op

BenchmarkLogitLoggerWithFormat-16        2931703              1233 ns/op             168 B/op          8 allocs/op

BenchmarkZapLogger-16                    1674750              2143 ns/op             449 B/op         16 allocs/op

BenchmarkGologLogger-16                  2223093              1619 ns/op             713 B/op         24 allocs/op

BenchmarkLogrusLogger-16                  899808              3968 ns/op            1634 B/op         52 allocs/op

******************************************************************************************************************

BenchmarkLogitFile-16                    3556720              1009 ns/op             129 B/op          4 allocs/op

BenchmarkLogitFileWithoutBuffer-16        499887              7176 ns/op             128 B/op          4 allocs/op

BenchmarkZapFile-16                       409000              8580 ns/op             449 B/op         16 allocs/op

BenchmarkGologFile-16                     257083             13884 ns/op             713 B/op         24 allocs/op

BenchmarkLogrusFile-16                    327198             10699 ns/op            1634 B/op         52 allocs/op
*/

const (
	timeFormat = "2006-01-02 15:04:05"
)

type nopWriter struct{}

func (w *nopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitLogger$ -benchtime=3s
func BenchmarkLogitLogger(b *testing.B) {

	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.Encoders().SetEncoder(logit.NewTextEncoder(timeFormat))
	logger.Writers().SetWriter(&nopWriter{})

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

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitLoggerWithReflection$ -benchtime=3s
func BenchmarkLogitLoggerWithReflection(b *testing.B) {

	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.Encoders().SetEncoder(logit.NewTextEncoder(timeFormat))
	logger.Writers().SetWriter(&nopWriter{})

	logTask := func() {
		logger.Debug("debug%s", "...")
		logger.Info("info%s", "...")
		logger.Warn("warning%s", "...")
		logger.Error("error%s", "...")
	}

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkZapLogger$ -benchtime=3s
//func BenchmarkZapLogger(b *testing.B) {
//
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
//
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkGologLogger$ -benchtime=3s
//func BenchmarkGologLogger(b *testing.B) {
//
//	logger := golog.New()
//	logger.SetOutput(&nopWriter{})
//	logger.SetLevel("debug")
//	logger.SetTimeFormat(timeFormat)
//
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
//
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogrusLogger$ -benchtime=3s
//func BenchmarkLogrusLogger(b *testing.B) {
//
//	logger := logrus.New()
//	logger.SetOutput(&nopWriter{})
//	logger.SetLevel(logrus.DebugLevel)
//	logger.SetFormatter(&logrus.TextFormatter{
//		TimestampFormat: timeFormat,
//	})
//
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

func createFileOf(filePath string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filePath), 0644)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitFile$ -benchtime=3s
func BenchmarkLogitFile(b *testing.B) {

	file, _ := createFileOf("Z:/" + b.Name() + ".log")
	writer := logit.NewBufferedWriter(file)
	defer writer.Close()
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.Encoders().SetEncoder(logit.NewTextEncoder(timeFormat))
	logger.Writers().SetWriter(writer)

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

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitFileWithoutBuffer$ -benchtime=3s
func BenchmarkLogitFileWithoutBuffer(b *testing.B) {

	file, _ := createFileOf("Z:/" + b.Name() + ".log")
	logger := logit.NewLogger()
	logger.SetLevel(logit.DebugLevel)
	logger.Encoders().SetEncoder(logit.NewTextEncoder(timeFormat))
	logger.Writers().SetWriter(file)

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

//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkZapFile$ -benchtime=3s
//func BenchmarkZapFile(b *testing.B) {
//
//	file, _ := createFileOf("Z:/" + b.Name() + ".log")
//	config := zap.NewProductionEncoderConfig()
//	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//		enc.AppendString(t.Format(timeFormat))
//	}
//	encoder := zapcore.NewConsoleEncoder(config)
//	writeSyncer := zapcore.AddSync(file)
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//
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
//
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkGologFile$ -benchtime=3s
//func BenchmarkGologFile(b *testing.B) {
//
//	file, _ := createFileOf("Z:/" + b.Name() + ".log")
//	logger := golog.New()
//	logger.SetOutput(file)
//	logger.SetLevel("debug")
//	logger.SetTimeFormat(timeFormat)
//
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
//
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogrusFile$ -benchtime=3s
//func BenchmarkLogrusFile(b *testing.B) {
//
//	file, _ := createFileOf("Z:/" + b.Name() + ".log")
//	logger := logrus.New()
//	logger.SetOutput(file)
//	logger.SetLevel(logrus.DebugLevel)
//	logger.SetFormatter(&logrus.TextFormatter{
//		TimestampFormat: timeFormat,
//	})
//
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
