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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/core/writer"
	"github.com/go-logit/logit/extension/file"
	"github.com/go-logit/logit/support/global"
	"github.com/go-logit/logit/support/size"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
	LevelPrint = "print"
	LevelOff   = "off"
)

const (
	AppenderText = "text"
	AppenderJson = "json"
)

const (
	WriterTargetStdout     = "stdout"
	WriterTargetStderr     = "stderr"
	WriterTargetFile       = "file"
	WriterTargetRotateFile = "rotate_file"

	WriterModeDirect = "direct"
	WriterModeBuffer = "buffer"
	WriterModeBatch  = "batch"
)

const (
	Day            = global.Day
	UnixTimeFormat = global.UnixTimeFormat
)

func mkdir(path string, mode os.FileMode) error {
	return os.MkdirAll(filepath.Dir(path), mode)
}

func open(path string, mode os.FileMode) (*os.File, error) {
	return global.OpenFile(path, mode)
}

// parseLevel parses level option from level.
func parseLevel(level string) (logit.Option, error) {
	switch strings.ToLower(level) {
	case LevelDebug:
		return logit.Options().WithDebugLevel(), nil
	case LevelInfo:
		return logit.Options().WithInfoLevel(), nil
	case LevelWarn:
		return logit.Options().WithWarnLevel(), nil
	case LevelError:
		return logit.Options().WithErrorLevel(), nil
	case LevelPrint:
		return logit.Options().WithPrintLevel(), nil
	case LevelOff:
		return logit.Options().WithOffLevel(), nil
	default:
		return nil, fmt.Errorf("level %s unknown", level)
	}
}

// parseTimeFormat parses the format of time.
func parseTimeFormat(format string) string {
	if strings.ToLower(format) == UnixTimeFormat {
		return global.UnixTimeFormat
	}

	return format
}

// parseTimeDuration parses the time duration from s.
func parseTimeDuration(s string) (time.Duration, error) {
	if strings.HasSuffix(s, "d") || strings.HasSuffix(s, "D") {
		s = strings.TrimSuffix(s, "d")
		s = strings.TrimSuffix(s, "D")

		days, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, err
		}

		return time.Duration(days) * Day, nil
	}

	return time.ParseDuration(s)
}

// parseAppender parses appender from config.
func parseAppender(name string) (appender.Appender, error) {
	switch strings.ToLower(name) {
	case AppenderText:
		return appender.Text(), nil
	case AppenderJson:
		return appender.Json(), nil
	default:
		return nil, fmt.Errorf("appender %s unknown", name)
	}
}

// parseWriterTarget parses write target from config.
func parseWriterTarget(wc WriterConfig) (io.Writer, error) {
	switch strings.ToLower(wc.Target) {
	case WriterTargetStdout:
		return os.Stdout, nil
	case WriterTargetStderr:
		return os.Stderr, nil
	case WriterTargetFile:
		var dirMode, fileMode os.FileMode = 0755, 0644
		if wc.DirMode > 0 {
			dirMode = wc.DirMode
		}

		if wc.FileMode > 0 {
			fileMode = wc.FileMode
		}

		if err := mkdir(wc.Filename, dirMode); err != nil {
			return nil, err
		}

		return open(wc.Filename, fileMode)
	case WriterTargetRotateFile:
		var opts []file.Option
		if wc.DirMode > 0 {
			opts = append(opts, file.WithDirMode(wc.DirMode))
		}

		if wc.FileMode > 0 {
			opts = append(opts, file.WithMode(wc.FileMode))
		}

		if wc.TimeFormat != "" {
			opts = append(opts, file.WithTimeFormat(wc.TimeFormat))
		}

		if wc.MaxSize != "" {
			maxSize, err := size.ParseByteSize(wc.MaxSize)
			if err != nil {
				return nil, err
			}

			opts = append(opts, file.WithMaxSize(maxSize))
		}

		if wc.MaxAge != "" {
			maxAge, err := parseTimeDuration(wc.MaxAge)
			if err != nil {
				return nil, err
			}

			opts = append(opts, file.WithMaxAge(maxAge))
		}

		if wc.MaxBackups > 0 {
			opts = append(opts, file.WithMaxBackups(wc.MaxBackups))
		}

		return file.New(wc.Filename, opts...)
	default:
		return nil, fmt.Errorf("writer target %s unknown", wc.Target)
	}
}

// parseWriterMode parses write mode from config.
func parseWriterMode(wc WriterConfig, w io.Writer) (io.Writer, error) {
	switch strings.ToLower(wc.Mode) {
	case WriterModeDirect:
		return writer.Wrap(w), nil
	case WriterModeBuffer:
		if wc.BufferSize == "" {
			return writer.Buffer(w), nil
		}

		s, err := size.ParseByteSize(wc.BufferSize)
		if err != nil {
			return nil, err
		}

		return writer.BufferWithSize(w, s), nil
	case WriterModeBatch:
		if wc.BatchCount <= 0 {
			return writer.Batch(w), nil
		}

		return writer.BatchWithCount(w, wc.BatchCount), nil
	default:
		return nil, fmt.Errorf("writer mode %s unknown", wc.Mode)
	}
}

// parseWriter parses writer from config.
func parseWriter(wc WriterConfig) (io.Writer, error) {
	w, err := parseWriterTarget(wc)
	if err != nil {
		return nil, err
	}

	return parseWriterMode(wc, w)
}
