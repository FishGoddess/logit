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
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/defaults"
	//"github.com/rs/zerolog"
	//"github.com/sirupsen/logrus"
	//"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

/*
$ go test -v ./_examples/performance_test.go -bench=. -benchtime=1s

goos: linux
goarch: amd64
cpu: AMD EPYC 7K62 48-Core Processor

BenchmarkLogitLogger-2                   1486184               810 ns/op               0 B/op          0 allocs/op
BenchmarkLogitLoggerTextHandler-2        1000000              1080 ns/op               0 B/op          0 allocs/op
BenchmarkLogitLoggerJsonHandler-2         847864              1393 ns/op             120 B/op          3 allocs/op
BenchmarkLogitLoggerPrint-2              1222302               981 ns/op              48 B/op          1 allocs/op
BenchmarkSlogLoggerTextHandler-2          725522              1629 ns/op               0 B/op          0 allocs/op
BenchmarkSlogLoggerJsonHandler-2          583214              2030 ns/op             120 B/op          3 allocs/op
BenchmarkZeroLogLogger-2                 1929276               613 ns/op               0 B/op          0 allocs/op
BenchmarkZapLogger-2                      976855              1168 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusLogger-2                   231723              4927 ns/op            2080 B/op         32 allocs/op

BenchmarkLogitFile-2                      624774              1935 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBuffer-2           1378076               873 ns/op               0 B/op          0 allocs/op
BenchmarkLogitFileWithBatch-2            1367479               883 ns/op               0 B/op          0 allocs/op
BenchmarkSlogFile-2                       407590              2944 ns/op               0 B/op          0 allocs/op
BenchmarkZeroLogFile-2                    634375              1810 ns/op               0 B/op          0 allocs/op
BenchmarkZapFile-2                        382790              2641 ns/op             216 B/op          2 allocs/op
BenchmarkLogrusFile-2                     174944              6491 ns/op            2080 B/op         32 allocs/op
*/

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLogger$ -benchtime=1s
func BenchmarkLogitLogger(b *testing.B) {
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTapeHandler(),
		logit.WithWriter(io.Discard),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitLoggerTextHandler$ -benchtime=1s
func BenchmarkLogitLoggerTextHandler(b *testing.B) {
	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTextHandler(),
		logit.WithWriter(io.Discard),
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
		logit.WithWriter(io.Discard),
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
		logit.WithTapeHandler(),
		logit.WithWriter(io.Discard),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Printf("info... %s=%s %s=%d %s=%.3f", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkSlogLoggerTextHandler$ -benchtime=1s
func BenchmarkSlogLoggerTextHandler(b *testing.B) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewTextHandler(io.Discard, opts)
	logger := slog.New(handler)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkSlogLoggerJsonHandler$ -benchtime=1s
func BenchmarkSlogLoggerJsonHandler(b *testing.B) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(io.Discard, opts)
	logger := slog.New(handler)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// // go test -v ./_examples/performance_test.go -bench=^BenchmarkZeroLogLogger$ -benchtime=1s
// func BenchmarkZeroLogLogger(b *testing.B) {
// 	zerolog.TimeFieldFormat = timeFormat
// 	logger := zerolog.New(io.Discard).Level(zerolog.InfoLevel).With().Timestamp().Logger()
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
// 	}
// }
//
// // go test -v ./_examples/performance_test.go -bench=^BenchmarkZapLogger$ -benchtime=1s
// func BenchmarkZapLogger(b *testing.B) {
// 	config := zap.NewProductionEncoderConfig()
// 	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 		enc.AppendString(t.Format(timeFormat))
// 	}
//
// 	encoder := zapcore.NewJSONEncoder(config)
// 	nopWriteSyncer := zapcore.AddSync(io.Discard)
// 	core := zapcore.NewCore(encoder, nopWriteSyncer, zapcore.InfoLevel)
//
// 	logger := zap.New(core)
// 	defer logger.Sync()
//
// 	logTask := func() {
// 		logger.Info("info...", zap.String("trace", "abcxxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
// 	}
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logTask()
// 	}
// }
//
// // go test -v ./_examples/performance_test.go -bench=^BenchmarkLogrusLogger$ -benchtime=1s
// func BenchmarkLogrusLogger(b *testing.B) {
// 	logger := logrus.New()
// 	logger.SetOutput(io.Discard)
// 	logger.SetLevel(logrus.InfoLevel)
// 	logger.SetFormatter(&logrus.JSONFormatter{
// 		TimestampFormat: timeFormat,
// 	})
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Info("info...")
// 	}
// }

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
	path := filepath.Join(b.TempDir(), b.Name())

	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTapeHandler(),
		logit.WithFile(path),
	)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkLogitFileWithBuffer$ -benchtime=1s
func BenchmarkLogitFileWithBuffer(b *testing.B) {
	path := filepath.Join(b.TempDir(), b.Name())

	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTapeHandler(),
		logit.WithFile(path),
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
	path := filepath.Join(b.TempDir(), b.Name())

	logger := logit.NewLogger(
		logit.WithInfoLevel(),
		logit.WithTapeHandler(),
		logit.WithFile(path),
		logit.WithBatch(64),
	)

	defer logger.Close()

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// go test -v ./_examples/performance_test.go -bench=^BenchmarkSlogFile$ -benchtime=1s
func BenchmarkSlogFile(b *testing.B) {
	path := filepath.Join(b.TempDir(), b.Name())

	file, _ := openFile(path)
	defer file.Close()

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewTextHandler(file, opts)
	logger := slog.New(handler)

	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("info...", "trace", "xxx", "id", 123, "pi", 3.14)
	}
}

// // go test -v ./_examples/performance_test.go -bench=^BenchmarkZeroLogFile$ -benchtime=1s
// func BenchmarkZeroLogFile(b *testing.B) {
// 	path := filepath.Join(b.TempDir(), b.Name())
//
// 	file, _ := openFile(path)
// 	defer file.Close()
//
// 	zerolog.TimeFieldFormat = timeFormat
// 	logger := zerolog.New(file).Level(zerolog.InfoLevel).With().Timestamp().Logger()
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logger.Info().Str("trace", "xxx").Int("id", 123).Float64("pi", 3.14).Msg("info...")
// 	}
// }
//
// // go test -v ./_examples/performance_test.go -bench=^BenchmarkZapFile$ -benchtime=1s
// func BenchmarkZapFile(b *testing.B) {
// 	path := filepath.Join(b.TempDir(), b.Name())
//
// 	file, _ := openFile(path)
// 	defer file.Close()
//
// 	config := zap.NewProductionEncoderConfig()
// 	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 		enc.AppendString(t.Format(timeFormat))
// 	}
//
// 	encoder := zapcore.NewJSONEncoder(config)
// 	writeSyncer := zapcore.AddSync(file)
// 	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
//
// 	logger := zap.New(core)
// 	defer logger.Sync()
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logger.Info("info...", zap.String("trace", "xxx"), zap.Int("id", 123), zap.Float64("pi", 3.14))
// 	}
// }
//
// // go test -v ./_examples/performance_test.go -bench=^BenchmarkLogrusFile$ -benchtime=1s
// func BenchmarkLogrusFile(b *testing.B) {
// 	path := filepath.Join(b.TempDir(), b.Name())
//
// 	file, _ := openFile(path)
// 	defer file.Close()
//
// 	logger := logrus.New()
// 	logger.SetOutput(file)
// 	logger.SetLevel(logrus.InfoLevel)
// 	logger.SetFormatter(&logrus.JSONFormatter{
// 		TimestampFormat: timeFormat,
// 	})
//
// 	b.ReportAllocs()
// 	b.StartTimer()
//
// 	for i := 0; i < b.N; i++ {
// 		logger.WithFields(map[string]interface{}{"trace": "xxx", "id": 123, "pi": 3.14}).Info("info...")
// 	}
// }
