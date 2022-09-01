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
	"strings"
	"time"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core/appender"
	"github.com/go-logit/logit/core/writer"
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
	WriterTargetStdout = "stdout"
	WriterTargetStderr = "stderr"
	WriterTargetFile   = "file"

	WriterModeDirect = "direct"
	WriterModeBuffer = "buffer"
	WriterModeBatch  = "batch"
)

const (
	UnixTimeFormat = "unix"
)

// WriterConfig stores all configs of writer.
type WriterConfig struct {
	// Target is where the writer writes to.
	// Values: stdout, stderr, file.
	Target string `json:"target" yaml:"target" toml:"target" bson:"target"`

	// Mode is how the writer writes to.
	// Values: direct, buffer, batch.
	Mode string `json:"mode" yaml:"mode" toml:"mode" bson:"mode"`

	// Filename is the name of file.
	// Only available when target is file.
	Filename string `json:"filename" yaml:"filename" toml:"filename" bson:"filename"`

	// BufferSize is the buffer size of buffer writer.
	// Only available when mode is buffer.
	BufferSize string `json:"buffer_size" yaml:"buffer_size" toml:"buffer_size" bson:"buffer_size"`

	// BatchCount is the batch count of batch writer.
	// Only available when mode is batch.
	BatchCount uint `json:"batch_count" yaml:"batch_count" toml:"batch_count" bson:"batch_count"`
}

// Config stores all configs of logger.
// You can embed it to your application config.
type Config struct {
	// Level is the level of logger.
	// Values: debug, info, warn, error, print, off.
	Level string `json:"level" yaml:"level" toml:"level" bson:"level"`

	// These keys are standard fields in log.
	TimeKey  string `json:"time_key" yaml:"time_key" toml:"time_key" bson:"time_key"`
	LevelKey string `json:"level_key" yaml:"level_key" toml:"level_key" bson:"level_key"`
	MsgKey   string `json:"msg_key" yaml:"msg_key" toml:"msg_key" bson:"msg_key"`
	PIDKey   string `json:"pid_key" yaml:"pid_key" toml:"pid_key" bson:"pid_key"`
	FileKey  string `json:"file_key" yaml:"file_key" toml:"file_key" bson:"file_key"`
	LineKey  string `json:"line_key" yaml:"line_key" toml:"line_key" bson:"line_key"`
	FuncKey  string `json:"func_key" yaml:"func_key" toml:"func_key" bson:"func_key"`

	// TimeFormat is the format of time.
	// Values: unix, ...
	TimeFormat string `json:"time_format" yaml:"time_format" toml:"time_format" bson:"time_format"`

	// These flags are standard fields in log.
	WithPID     bool `json:"with_pid" yaml:"with_pid" toml:"with_pid" bson:"with_pid"`
	WithCaller  bool `json:"with_caller" yaml:"with_caller" toml:"with_caller" bson:"with_caller"`
	CallerDepth int  `json:"caller_depth" yaml:"caller_depth" toml:"caller_depth" bson:"caller_depth"`

	// AutoSync is the frequency of syncing.
	// You can use common words like 5m or 60s.
	// See time.Duration and time.ParseDuration.
	AutoSync string `json:"auto_sync" yaml:"auto_sync" toml:"auto_sync" bson:"auto_sync"`

	// Appender is the appender of logger.
	// Values: text, json.
	Appender      string `json:"appender" yaml:"appender" toml:"appender" bson:"appender"`
	DebugAppender string `json:"debug_appender" yaml:"debug_appender" toml:"debug_appender" bson:"debug_appender"`
	InfoAppender  string `json:"info_appender" yaml:"info_appender" toml:"info_appender" bson:"info_appender"`
	WarnAppender  string `json:"warn_appender" yaml:"warn_appender" toml:"warn_appender" bson:"warn_appender"`
	ErrorAppender string `json:"error_appender" yaml:"error_appender" toml:"error_appender" bson:"error_appender"`
	PrintAppender string `json:"print_appender" yaml:"print_appender" toml:"print_appender" bson:"print_appender"`

	// Writer is the writer of logger.
	// See WriterConfig.
	Writer      WriterConfig `json:"writer" yaml:"writer" toml:"writer" bson:"writer"`
	DebugWriter WriterConfig `json:"debug_writer" yaml:"debug_writer" toml:"debug_writer" bson:"debug_writer"`
	InfoWriter  WriterConfig `json:"info_writer" yaml:"info_writer" toml:"info_writer" bson:"info_writer"`
	WarnWriter  WriterConfig `json:"warn_writer" yaml:"warn_writer" toml:"warn_writer" bson:"warn_writer"`
	ErrorWriter WriterConfig `json:"error_writer" yaml:"error_writer" toml:"error_writer" bson:"error_writer"`
	PrintWriter WriterConfig `json:"print_writer" yaml:"print_writer" toml:"print_writer" bson:"print_writer"`
}

// New returns a pointer to config.
func New() *Config {
	return new(Config)
}

func (c *Config) createFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}

// parseLevel returns the level option of c.
func (c *Config) parseLevel(level string) (logit.Option, error) {
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

// parseTimeFormat returns the format of time.
func (c *Config) parseTimeFormat(format string) string {
	if strings.ToLower(format) == UnixTimeFormat {
		return appender.UnixTimeFormat
	}

	return format
}

// parseAutoSync returns the frequency of syncing.
func (c *Config) parseAutoSync(frequency string) (time.Duration, error) {
	return time.ParseDuration(frequency)
}

// parseAppender returns the appender of appenderName.
func (c *Config) parseAppender(name string) (appender.Appender, error) {
	switch strings.ToLower(name) {
	case AppenderText:
		return appender.Text(), nil
	case AppenderJson:
		return appender.Json(), nil
	default:
		return nil, fmt.Errorf("appender %s unknown", name)
	}
}

// parseWriter returns the writer of writerName.
func (c *Config) parseWriter(wc WriterConfig) (io.Writer, error) {
	var w io.Writer

	switch strings.ToLower(wc.Target) {
	case WriterTargetStdout:
		w = os.Stdout
	case WriterTargetStderr:
		w = os.Stderr
	case WriterTargetFile:
		f, err := c.createFile(wc.Filename)
		if err != nil {
			return nil, err
		}

		w = f
	}

	switch strings.ToLower(wc.Mode) {
	case WriterModeDirect:
		w = writer.Wrap(w)
	case WriterModeBuffer:
		if wc.BufferSize != "" {
			s, err := size.ParseByteSize(wc.BufferSize)
			if err != nil {
				return nil, err
			}

			w = writer.BufferWithSize(w, s)
		} else {
			w = writer.Buffer(w)
		}
	case WriterModeBatch:
		if wc.BatchCount > 0 {
			w = writer.BatchWithCount(w, wc.BatchCount)
		} else {
			w = writer.Batch(w)
		}
	}

	return w, nil
}

// Options returns a slice of logit.Option for creating logit.Logger.
// Returns an error if something wrong happens.
func (c *Config) Options() ([]logit.Option, error) {
	if c == nil {
		return nil, nil
	}

	options := logit.Options()
	result := make([]logit.Option, 0, 16)

	if c.Level != "" {
		levelOption, err := c.parseLevel(c.Level)
		if err != nil {
			return nil, err
		}

		result = append(result, levelOption)
	}

	if c.TimeKey != "" {
		result = append(result, options.WithTimeKey(c.TimeKey))
	}

	if c.LevelKey != "" {
		result = append(result, options.WithLevelKey(c.LevelKey))
	}

	if c.MsgKey != "" {
		result = append(result, options.WithMsgKey(c.MsgKey))
	}

	if c.PIDKey != "" {
		result = append(result, options.WithPIDKey(c.PIDKey))
	}

	if c.FileKey != "" {
		result = append(result, options.WithFileKey(c.FileKey))
	}

	if c.LineKey != "" {
		result = append(result, options.WithLineKey(c.LineKey))
	}

	if c.FuncKey != "" {
		result = append(result, options.WithFuncKey(c.FuncKey))
	}

	if strings.TrimSpace(c.TimeFormat) != "" {
		result = append(result, options.WithTimeFormat(c.parseTimeFormat(c.TimeFormat)))
	}

	if c.WithPID {
		result = append(result, options.WithPID())
	}

	if c.WithCaller {
		result = append(result, options.WithCaller())
	}

	if c.CallerDepth > 0 {
		result = append(result, options.WithCallerDepth(c.CallerDepth))
	}

	if strings.TrimSpace(c.AutoSync) != "" {
		frequency, err := c.parseAutoSync(c.AutoSync)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithAutoSync(frequency))
	}

	if strings.TrimSpace(c.Appender) != "" {
		apdr, err := c.parseAppender(c.Appender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithAppender(apdr))
	}

	if strings.TrimSpace(c.DebugAppender) != "" {
		apdr, err := c.parseAppender(c.DebugAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithDebugAppender(apdr))
	}

	if strings.TrimSpace(c.InfoAppender) != "" {
		apdr, err := c.parseAppender(c.InfoAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithInfoAppender(apdr))
	}

	if strings.TrimSpace(c.WarnAppender) != "" {
		apdr, err := c.parseAppender(c.WarnAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWarnAppender(apdr))
	}

	if strings.TrimSpace(c.ErrorAppender) != "" {
		apdr, err := c.parseAppender(c.ErrorAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithErrorAppender(apdr))
	}

	if strings.TrimSpace(c.PrintAppender) != "" {
		apdr, err := c.parseAppender(c.PrintAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithPrintAppender(apdr))
	}

	if c.Writer.Target != "" {
		w, err := c.parseWriter(c.Writer)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWriter(w))
	}

	if c.DebugWriter.Target != "" {
		w, err := c.parseWriter(c.DebugWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithDebugWriter(w))
	}

	if c.InfoWriter.Target != "" {
		w, err := c.parseWriter(c.InfoWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithInfoWriter(w))
	}

	if c.WarnWriter.Target != "" {
		w, err := c.parseWriter(c.WarnWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithWarnWriter(w))
	}

	if c.ErrorWriter.Target != "" {
		w, err := c.parseWriter(c.ErrorWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithErrorWriter(w))
	}

	if c.PrintWriter.Target != "" {
		w, err := c.parseWriter(c.PrintWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, options.WithPrintWriter(w))
	}

	return result, nil
}
