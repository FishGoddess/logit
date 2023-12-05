// Copyright 2023 FishGoddess. All Rights Reserved.
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

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/defaults"
	//"github.com/rs/zerolog"
	//"github.com/sirupsen/logrus"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

/*
$ go test -v ./_examples/performance_test.go -bench=. -benchtime=1s
*/

const (
	timeFormat = "2006-01-02 15:04:05"
)

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerTextHandler$ -benchtime=1s
func BenchmarkLogitLoggerTextHandler(b *testing.B) {
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithWriter(ioutil.Discard),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerJsonHandler$ -benchtime=1s
func BenchmarkLogitLoggerJsonHandler(b *testing.B) {
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithJsonHandler(),
		logit.WithWriter(ioutil.Discard),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerPrint$ -benchtime=1s
func BenchmarkLogitLoggerPrint(b *testing.B) {
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithWriter(ioutil.Discard),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Printf("info... %s=%s %s=%d %s=%.3f", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

//// go test -v ./_examples/performance_test.go -bench=^BenchmarkZeroLogLogger$ -benchtime=1s
//func BenchmarkZeroLogLogger(b *testing.B) {
//	zerolog.TimeFieldFormat = timeFormat
//	logger := zerolog.New(ioutil.Discard).With().Timestamp().Logger()
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
//	nopWriteSyncer := zapcore.AddSync(ioutil.Discard)
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
//	logger.SetOutput(ioutil.Discard)
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

func openFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := defaults.OpenFileDir(dir, defaults.FileDirMode); err != nil {
		return nil, err
	}

	return defaults.OpenFile(path, 644)
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFile$ -benchtime=1s
func BenchmarkLogitFile(b *testing.B) {
	file := filepath.Join(b.TempDir(), b.Name())
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithFile(file),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithBuffer$ -benchtime=1s
func BenchmarkLogitFileWithBuffer(b *testing.B) {
	file := filepath.Join(b.TempDir(), b.Name())
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithFile(file),
		logit.WithBuffer(65536),
	)

	defer logger.Close()

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithBatch$ -benchtime=1s
func BenchmarkLogitFileWithBatch(b *testing.B) {
	file := filepath.Join(b.TempDir(), b.Name())
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithFile(file),
		logit.WithBatch(64),
	)

	defer logger.Close()

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
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
