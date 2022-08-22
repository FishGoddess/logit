// Copyright 2022 FishGoddess. All Rights Reserved.
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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	//"time"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core/appender"
	//"github.com/rs/zerolog"
	//"github.com/sirupsen/logrus"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

/*
$ go test -v ./_examples/performance_test.go -bench=. -benchtime=1s

BenchmarkLogitLoggerWithTextAppender-16    856958              1362 ns/op               0 B/op          0 allocs/op

BenchmarkLogitLoggerWithJsonAppender-16    799759              1373 ns/op               0 B/op          0 allocs/op

BenchmarkLogitLoggerWithFormat-16          666484              1740 ns/op              40 B/op          4 allocs/op

BenchmarkZeroLogLogger-16                  922863              1244 ns/op               0 B/op          0 allocs/op

BenchmarkZapLogger-16                      413701              2824 ns/op             897 B/op          8 allocs/op

BenchmarkLogrusLogger-16                   105238             11474 ns/op            7411 B/op        128 allocs/op

******************************************************************************************************************

BenchmarkLogitFileWithTextAppender-16     631435              1751 ns/op             851 B/op          0 allocs/op

BenchmarkLogitFileWithJsonAppender-16     599862              1768 ns/op             896 B/op          0 allocs/op

BenchmarkLogitFileWithoutBuffer-16        148113              7773 ns/op               0 B/op          0 allocs/op

BenchmarkZeroLogFile-16                   159962              7472 ns/op               0 B/op          0 allocs/op

BenchmarkZapFile-16                       130405              9137 ns/op             897 B/op          8 allocs/op

BenchmarkLogrusFile-16                     65202             18439 ns/op            7410 B/op        128 allocs/op
*/

const (
	timeFormat = "2006-01-02 15:04:05"
)

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerWithTextAppender$ -benchtime=1s
func BenchmarkLogitLoggerWithTextAppender(b *testing.B) {
	options := logit.Options()

	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Text()),
		options.WithWriter(ioutil.Discard),
		options.WithTimeFormat(timeFormat),
	)

	logTask := func() {
		logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerWithJsonAppender$ -benchtime=1s
func BenchmarkLogitLoggerWithJsonAppender(b *testing.B) {
	options := logit.Options()

	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Json()),
		options.WithWriter(ioutil.Discard),
		options.WithTimeFormat(timeFormat),
	)

	logTask := func() {
		logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerWithFormat$ -benchtime=1s
func BenchmarkLogitLoggerWithFormat(b *testing.B) {
	options := logit.Options()

	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Text()),
		options.WithWriter(ioutil.Discard),
		options.WithTimeFormat(timeFormat),
	)

	logTask := func() {
		logger.Debug("debug%s", "...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info%s", "...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning%s", "...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error%s", "...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerPrint$ -benchtime=1s
func BenchmarkLogitLoggerPrint(b *testing.B) {
	options := logit.Options()

	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Text()),
		options.WithWriter(ioutil.Discard),
		options.WithTimeFormat(timeFormat),
	)

	logTask := func() {
		logger.Println("debug", "trace", "xxx", "id", 123, "pi", 3.14)
		logger.Println("info", "trace", "xxx", "id", 123, "pi", 3.14)
		logger.Println("warn", "trace", "xxx", "id", 123, "pi", 3.14)
		logger.Println("error", "trace", "xxx", "id", 123, "pi", 3.14)
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

//// go test -v ./_examples/performance_test.go -bench=^BenchmarkZeroLogLogger$ -benchtime=1s
//func BenchmarkZeroLogLogger(b *testing.B) {
//	zerolog.TimeFieldFormat = timeFormat
//	logger := zerolog.New(&nopWriter{}).With().Timestamp().Logger()
//
//	logTask := func() {
//		logger.Debug().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("debug...")
//		logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
//		logger.Warn().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("warning...")
//		logger.Error().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error...")
//	}
//
//	b.ReportAllocs()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// go test -v ./_examples/performance_test.go -bench=^BenchmarkZapLogger$ -benchtime=1s
//func BenchmarkZapLogger(b *testing.B) {
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
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogrusLogger$ -benchtime=1s
//func BenchmarkLogrusLogger(b *testing.B) {
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
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}

// *******************************************************************************

func createFileOf(filePath string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(filePath), 0644)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithTextAppender$ -benchtime=1s
func BenchmarkLogitFileWithTextAppender(b *testing.B) {
	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
	defer file.Close()

	options := logit.Options()
	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Text()),
		options.WithBufferWriter(file),
		options.WithTimeFormat(timeFormat),
	)
	defer logger.Close()

	logTask := func() {
		logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithJsonAppender$ -benchtime=1s
func BenchmarkLogitFileWithJsonAppender(b *testing.B) {
	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
	defer file.Close()

	options := logit.Options()
	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Json()),
		options.WithBufferWriter(file),
		options.WithTimeFormat(timeFormat),
	)
	defer logger.Close()

	logTask := func() {
		logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithoutBuffer$ -benchtime=1s
func BenchmarkLogitFileWithoutBuffer(b *testing.B) {
	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
	defer file.Close()

	options := logit.Options()
	logger := logit.NewLogger(
		options.WithDebugLevel(),
		options.WithAppender(appender.Text()),
		options.WithWriter(file),
		options.WithTimeFormat(timeFormat),
	)

	logTask := func() {
		logger.Debug("debug...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Info("info...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Warn("warning...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
		logger.Error("error...").String("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Log()
	}

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logTask()
	}
}

//// go test -v ./_examples/performance_test.go -bench=^BenchmarkZeroLogFile$ -benchtime=1s
//func BenchmarkZeroLogFile(b *testing.B) {
//	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
//	zerolog.TimeFieldFormat = timeFormat
//	logger := zerolog.New(file).With().Timestamp().Logger()
//
//	logTask := func() {
//		logger.Debug().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("debug...")
//		logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
//		logger.Warn().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("warning...")
//		logger.Error().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("error...")
//	}
//
//	b.ReportAllocs()
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// go test -v ./_examples/performance_test.go -bench=^BenchmarkZapFile$ -benchtime=1s
//func BenchmarkZapFile(b *testing.B) {
//	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
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
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
//
//// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogrusFile$ -benchtime=1s
//func BenchmarkLogrusFile(b *testing.B) {
//	file, _ := createFileOf(filepath.Join(b.TempDir(), b.Name()))
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
//	for i := 0; i < b.N; i++ {
//		logTask()
//	}
//}
