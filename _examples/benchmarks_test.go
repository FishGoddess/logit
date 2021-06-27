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
	//"github.com/rs/zerolog"
	//"github.com/sirupsen/logrus"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

/*
$ go test -v ./_examples/benchmarks_test.go -bench=. -benchtime=3s

BenchmarkLogitLoggerWithTextEncoder-16   3993361               905 ns/op             128 B/op          4 allocs/op

BenchmarkLogitLoggerWithJsonEncoder-16   1728166              2086 ns/op             424 B/op         16 allocs/op

BenchmarkLogitLoggerWithReflection-16    1487979              2418 ns/op             464 B/op         20 allocs/op

BenchmarkZapLogger-16                    1282438              2793 ns/op             897 B/op          8 allocs/op

BenchmarkZeroLogLogger-16                2972082              1201 ns/op               0 B/op          0 allocs/op

BenchmarkLogrusLogger-16                  312974             11451 ns/op            7411 B/op        128 allocs/op

******************************************************************************************************************

BenchmarkLogitFileWithTextEncoder-16     3159062              1100 ns/op             468 B/op          4 allocs/op

BenchmarkLogitFileWithJsonEncoder-16     1542493              2314 ns/op            1123 B/op         16 allocs/op

BenchmarkLogitFileWithoutBuffer-16        428478              8463 ns/op             424 B/op         16 allocs/op

BenchmarkZapFile-16                       367264              9348 ns/op             897 B/op          8 allocs/op

BenchmarkZeroLogFile-16                   461437              7633 ns/op               0 B/op          0 allocs/op

BenchmarkLogrusFile-16                    194550             18516 ns/op            7410 B/op        128 allocs/op
*/

const (
	timeFormat = "2006-01-02 15:04:05"
)

type nopWriter struct{}

func (w *nopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitLoggerWithTextEncoder$ -benchtime=3s
func BenchmarkLogitLoggerWithTextEncoder(b *testing.B) {



	logTask := func() {

	}

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitLoggerWithJsonEncoder$ -benchtime=3s
func BenchmarkLogitLoggerWithJsonEncoder(b *testing.B) {



	logTask := func() {

	}

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitLoggerWithReflection$ -benchtime=3s
func BenchmarkLogitLoggerWithReflection(b *testing.B) {



	logTask := func() {

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
//	encoder := zapcore.NewJSONEncoder(config)
//	nopWriteSyncer := zapcore.AddSync(&nopWriter{})
//	core := zapcore.NewCore(encoder, nopWriteSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//
//	logTask := func() {
//		logger.Debug("debug...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Info("info...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Warn("warning...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Error("error...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
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
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkZeroLogLogger$ -benchtime=3s
//func BenchmarkZeroLogLogger(b *testing.B) {
//
//	zerolog.TimeFieldFormat = timeFormat
//	logger := zerolog.New(&nopWriter{}).With().Timestamp().Logger()
//
//	logTask := func() {
//		logger.Debug().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("debug...")
//		logger.Info().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
//		logger.Warn().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("warning...")
//		logger.Error().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error...")
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
//	logger.SetFormatter(&logrus.JSONFormatter{
//		TimestampFormat: timeFormat,
//	})
//
//	logTask := func() {
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Debug("debug...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Info("info...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Warn("warning...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Error("error...")
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

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitFileWithTextEncoder$ -benchtime=3s
func BenchmarkLogitFileWithTextEncoder(b *testing.B) {

	file, _ := createFileOf("Z:/" + b.Name() + ".log")
	writer := logit.NewBufferedWriter(file)
	defer writer.Close()


	logTask := func() {

	}

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkLogitFileWithJsonEncoder$ -benchtime=3s
func BenchmarkLogitFileWithJsonEncoder(b *testing.B) {

	file, _ := createFileOf("Z:/" + b.Name() + ".log")
	writer := logit.NewBufferedWriter(file)
	defer writer.Close()


	logTask := func() {

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
	defer file.Close()

	logTask := func() {

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
//	encoder := zapcore.NewJSONEncoder(config)
//	writeSyncer := zapcore.AddSync(file)
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	defer logger.Sync()
//
//	logTask := func() {
//		logger.Debug("debug...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Info("info...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Warn("warning...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
//		logger.Error("error...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
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
//// go test -v ./_examples/benchmarks_test.go -bench=^BenchmarkZeroLogFile$ -benchtime=3s
//func BenchmarkZeroLogFile(b *testing.B) {
//
//	file, _ := createFileOf("Z:/" + b.Name() + ".log")
//	zerolog.TimeFieldFormat = timeFormat
//	logger := zerolog.New(file).With().Timestamp().Logger()
//
//	logTask := func() {
//		logger.Debug().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("debug...")
//		logger.Info().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
//		logger.Warn().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("warning...")
//		logger.Error().String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error...")
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
//	logger.SetFormatter(&logrus.JSONFormatter{
//		TimestampFormat: timeFormat,
//	})
//
//	logTask := func() {
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Debug("debug...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Info("info...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Warn("warning...")
//		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Error("error...")
//	}
//
//	b.ReportAllocs()
//	b.StartTimer()
//
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
