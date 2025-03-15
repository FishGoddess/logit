// Copyright 2025 FishGoddess. All Rights Reserved.
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

package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/defaults"
)

func removeTimeAndSource(str string) string {
	str = strings.ReplaceAll(str, "\n", " ")
	strs := strings.Split(str, " ")

	var removed strings.Builder
	for _, s := range strs {
		if strings.HasPrefix(s, slog.TimeKey) {
			continue
		}

		if strings.HasPrefix(s, slog.SourceKey) {
			continue
		}

		removed.WriteString(s)
		removed.WriteString(" ")
	}

	return removed.String()
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestConfig$
func TestConfig(t *testing.T) {
	logitFile := filepath.Join(t.TempDir(), t.Name()+"_logit.log")
	slogFile := filepath.Join(t.TempDir(), t.Name()+"_slog.log")

	conf := Config{
		Level:   "debug",
		Handler: "text",
		Writer: WriterConfig{
			Target:         logitFile,
			FileRotate:     true,
			FileMaxSize:    "1GB",
			FileMaxAge:     "7d",
			FileMaxBackups: 30,
			BufferSize:     "64KB",
			BatchSize:      16,
		},
		WithSource: true,
		WithPID:    true,
		SyncTimer:  "1m",
	}

	opts, err := conf.Options()
	if err != nil {
		t.Fatal(err)
	}

	logger := logit.NewLogger(opts...)
	defer logger.Close()

	logger.Debug("debug msg", "key1", 1)
	logger.Info("info msg", "key2", 2)
	logger.Warn("warn msg", "key3", 3)
	logger.Error("error msg", "key4", 4)
	logger.Close()

	file, err := defaults.OpenFile(slogFile, 0644)
	if err != nil {
		t.Fatal(err)
	}

	handlerOpts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	slogLogger := slog.New(slog.NewTextHandler(file, handlerOpts)).With("pid", os.Getpid())

	slogLogger.Debug("debug msg", "key1", 1)
	slogLogger.Info("info msg", "key2", 2)
	slogLogger.Warn("warn msg", "key3", 3)
	slogLogger.Error("error msg", "key4", 4)

	gotBytes, err := os.ReadFile(logitFile)
	if err != nil {
		t.Fatal(err)
	}

	wantBytes, err := os.ReadFile(slogFile)
	if err != nil {
		t.Fatal(err)
	}

	got := removeTimeAndSource(string(gotBytes))
	want := removeTimeAndSource(string(wantBytes))

	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
