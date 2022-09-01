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

package config

import (
	"testing"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/support/global"
)

// go test -v -cover -run=^TestConfigOptions$
func TestConfigOptions(t *testing.T) {
	cfg := Config{
		Level:         LevelDebug,
		TimeKey:       "log.time",
		LevelKey:      "log.level",
		MsgKey:        "log.msg",
		PIDKey:        "log.pid",
		FileKey:       "log.file",
		LineKey:       "log.line",
		FuncKey:       "log.func",
		TimeFormat:    UnixTimeFormat,
		WithPID:       true,
		WithCaller:    true,
		CallerDepth:   global.CallerDepth,
		AutoSync:      "30s",
		Appender:      AppenderText,
		DebugAppender: AppenderText,
		InfoAppender:  AppenderText,
		WarnAppender:  AppenderText,
		ErrorAppender: AppenderText,
		PrintAppender: AppenderText,
		Writer: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		DebugWriter: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		InfoWriter: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		WarnWriter: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		ErrorWriter: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		PrintWriter: WriterConfig{
			Target:     WriterTargetStdout,
			Mode:       WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
	}

	opts, err := cfg.Options()
	if err != nil {
		t.Error(err)
	}

	logger := logit.NewLogger(opts...)
	defer logger.Close()

	logger.Info("My mother is a config").Any("config", cfg).Log()
}
