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
	"os"
	"strings"

	"github.com/FishGoddess/logit"
)

// WriterConfig stores all configs of writer.
type WriterConfig struct {
	// Target is where the writer writes to.
	// Values: stdout, stderr, file.
	Target string `json:"target" yaml:"target" toml:"target" bson:"target"`

	// Mode is how the writer writes to.
	// Values: direct, buffer, batch.
	Mode string `json:"mode" yaml:"mode" toml:"mode" bson:"mode"`

	// BufferSize is the buffer size of buffer writer.
	// You can use common words like 512B or 4KB.
	// Only available when mode is buffer.
	BufferSize string `json:"buffer_size" yaml:"buffer_size" toml:"buffer_size" bson:"buffer_size"`

	// BatchCount is the batch count of batch writer.
	// Only available when mode is batch.
	BatchCount uint `json:"batch_count" yaml:"batch_count" toml:"batch_count" bson:"batch_count"`

	// Filename is the name of file.
	// Only available when target is file and rotate_file.
	Filename string `json:"filename" yaml:"filename" toml:"filename" bson:"filename"`

	// DirMode is the permission mode of directory.
	// Only available when target is file and rotate_file.
	DirMode os.FileMode `json:"dir_mode" yaml:"dir_mode" toml:"dir_mode" bson:"dir_mode"`

	// FileMode is the permission mode of file.
	// Only available when target is file and rotate_file.
	FileMode os.FileMode `json:"file_mode" yaml:"file_mode" toml:"file_mode" bson:"file_mode"`

	// TimeFormat is the format of time.
	// Only available when target is rotate_file.
	// Values: unix, ...
	TimeFormat string `json:"time_format" yaml:"time_format" toml:"time_format" bson:"time_format"`

	// MaxSize is the max size of file.
	// If size of data in one write is bigger than MaxSize, then file will rotate and write it,
	// which means file and its backup may bigger than MaxSize in size.
	// You can use common words like 64MB or 1GB.
	// Only available when target is rotate_file.
	MaxSize string `json:"max_size" yaml:"max_size" toml:"max_size" bson:"max_size"`

	// MaxAge is the time that backup will live.
	// All backups reach MaxAge will be removed automatically.
	// You can use common words like 7d or 24h.
	// See time.Duration and time.ParseDuration.
	// Only available when target is rotate_file.
	MaxAge string `json:"max_age" yaml:"max_age" toml:"max_age" bson:"max_age"`

	// MaxBackups is the max count of backups.
	// Only available when target is rotate_file.
	MaxBackups int `json:"max_backups" yaml:"max_backups" toml:"max_backups" bson:"max_backups"`
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
	ErrorKey string `json:"error_key" yaml:"error_key" toml:"error_key" bson:"error_key"`

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

// Options returns a slice of logit.Option for creating logit.Logger.
// Returns an error if something wrong happens.
func (c *Config) Options() ([]logit.Option, error) {
	if c == nil {
		return nil, nil
	}

	opts := logit.Options()
	result := make([]logit.Option, 0, 16)

	if c.Level != "" {
		levelOption, err := parseLevel(c.Level)
		if err != nil {
			return nil, err
		}

		result = append(result, levelOption)
	}

	if c.TimeKey != "" {
		result = append(result, opts.WithTimeKey(c.TimeKey))
	}

	if c.LevelKey != "" {
		result = append(result, opts.WithLevelKey(c.LevelKey))
	}

	if c.MsgKey != "" {
		result = append(result, opts.WithMsgKey(c.MsgKey))
	}

	if c.PIDKey != "" {
		result = append(result, opts.WithPIDKey(c.PIDKey))
	}

	if c.FileKey != "" {
		result = append(result, opts.WithFileKey(c.FileKey))
	}

	if c.LineKey != "" {
		result = append(result, opts.WithLineKey(c.LineKey))
	}

	if c.FuncKey != "" {
		result = append(result, opts.WithFuncKey(c.FuncKey))
	}

	if c.ErrorKey != "" {
		result = append(result, opts.WithErrorKey(c.ErrorKey))
	}

	if strings.TrimSpace(c.TimeFormat) != "" {
		result = append(result, opts.WithTimeFormat(parseTimeFormat(c.TimeFormat)))
	}

	if c.WithPID {
		result = append(result, opts.WithPID())
	}

	if c.WithCaller {
		result = append(result, opts.WithCaller())
	}

	if c.CallerDepth > 0 {
		result = append(result, opts.WithCallerDepth(c.CallerDepth))
	}

	if strings.TrimSpace(c.AutoSync) != "" {
		frequency, err := parseTimeDuration(c.AutoSync)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithAutoSync(frequency))
	}

	if strings.TrimSpace(c.Appender) != "" {
		apdr, err := parseAppender(c.Appender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithAppender(apdr))
	}

	if strings.TrimSpace(c.DebugAppender) != "" {
		apdr, err := parseAppender(c.DebugAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithDebugAppender(apdr))
	}

	if strings.TrimSpace(c.InfoAppender) != "" {
		apdr, err := parseAppender(c.InfoAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithInfoAppender(apdr))
	}

	if strings.TrimSpace(c.WarnAppender) != "" {
		apdr, err := parseAppender(c.WarnAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithWarnAppender(apdr))
	}

	if strings.TrimSpace(c.ErrorAppender) != "" {
		apdr, err := parseAppender(c.ErrorAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithErrorAppender(apdr))
	}

	if strings.TrimSpace(c.PrintAppender) != "" {
		apdr, err := parseAppender(c.PrintAppender)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithPrintAppender(apdr))
	}

	if c.Writer.Target != "" {
		w, err := parseWriter(c.Writer)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithWriter(w))
	}

	if c.DebugWriter.Target != "" {
		w, err := parseWriter(c.DebugWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithDebugWriter(w))
	}

	if c.InfoWriter.Target != "" {
		w, err := parseWriter(c.InfoWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithInfoWriter(w))
	}

	if c.WarnWriter.Target != "" {
		w, err := parseWriter(c.WarnWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithWarnWriter(w))
	}

	if c.ErrorWriter.Target != "" {
		w, err := parseWriter(c.ErrorWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithErrorWriter(w))
	}

	if c.PrintWriter.Target != "" {
		w, err := parseWriter(c.PrintWriter)
		if err != nil {
			return nil, err
		}

		result = append(result, opts.WithPrintWriter(w))
	}

	return result, nil
}
