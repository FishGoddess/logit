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

package config

import (
	"fmt"
	"strings"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/rotate"
)

type WriterConfig struct {
	// Target is where the writer writes logs.
	// Values: "stdout", "stderr", or a file path like "./logit.log".
	Target string `json:"target" yaml:"target" toml:"target" bson:"target"`

	// FileRotate is log file should split and backup when satisfy some conditions.
	// It's useful in production so we recommend you to set it to true.
	// Only available when target is a file path.
	FileRotate bool `json:"file_rotate" yaml:"file_rotate" toml:"file_rotate" bson:"file_rotate"`

	// FileMaxSize is the max size of a log file.
	// If size of data in one output operation is bigger than this value, then file will rotate before writing,
	// which means file and its backups may be bigger than this value in size.
	// You can use common words like "100MB" or "1GB".
	// Only available when rotate is true.
	FileMaxSize string `json:"file_max_size" yaml:"file_max_size" toml:"file_max_size" bson:"file_max_size"`

	// FileMaxAge is the time that backups will live.
	// All backups reach max age will be removed automatically.
	// You can use common words like "7d" or "24h".
	// See time.Duration and time.ParseDuration.
	// Only available when rotate is true.
	FileMaxAge string `json:"file_max_age" yaml:"file_max_age" toml:"file_max_age" bson:"file_max_age"`

	// FileMaxBackups is the max count of file backups.
	// Only available when rotate is true.
	FileMaxBackups uint32 `json:"file_max_backups" yaml:"file_max_backups" toml:"file_max_backups" bson:"file_max_backups"`

	// BufferSize is the size of a buffer.
	// You can use common words like "512B" or "4KB".
	// Only available when mode is "buffer".
	BufferSize string `json:"buffer_size" yaml:"buffer_size" toml:"buffer_size" bson:"buffer_size"`

	// BatchSize is the size of a batch.
	// Only available when mode is "batch".
	BatchSize uint64 `json:"batch_size" yaml:"batch_size" toml:"batch_size" bson:"batch_size"`
}

func (wc *WriterConfig) parseFileOptions() ([]rotate.Option, error) {
	opts := make([]rotate.Option, 0, 4)

	if wc.FileMaxSize != "" {
		maxSize, err := parseByteSize(wc.FileMaxSize)
		if err != nil {
			return nil, err
		}

		opts = append(opts, rotate.WithMaxSize(maxSize))
	}

	if wc.FileMaxAge != "" {
		maxAge, err := parseTimeDuration(wc.FileMaxAge)
		if err != nil {
			return nil, err
		}

		opts = append(opts, rotate.WithMaxAge(maxAge))
	}

	if wc.FileMaxBackups > 0 {
		opts = append(opts, rotate.WithMaxBackups(wc.FileMaxBackups))
	}

	return opts, nil
}

func (wc *WriterConfig) appendTargetOptions(opts []logit.Option) ([]logit.Option, error) {
	target := strings.ToLower(wc.Target)

	if target == "" {
		return opts, nil
	}

	if target == "stdout" {
		opts = append(opts, logit.WithStdout())
		return opts, nil
	}

	if target == "stderr" {
		opts = append(opts, logit.WithStderr())
		return opts, nil
	}

	if !wc.FileRotate {
		opts = append(opts, logit.WithFile(wc.Target))
		return opts, nil
	}

	fileOpts, err := wc.parseFileOptions()
	if err != nil {
		return nil, err
	}

	opts = append(opts, logit.WithRotateFile(wc.Target, fileOpts...))
	return opts, nil
}

func (wc *WriterConfig) appendModeOptions(opts []logit.Option) ([]logit.Option, error) {
	if wc.BufferSize != "" {
		bufferSize, err := parseByteSize(wc.BufferSize)
		if err != nil {
			return nil, err
		}

		opts = append(opts, logit.WithBuffer(bufferSize))
	}

	if wc.BatchSize > 0 {
		opts = append(opts, logit.WithBatch(wc.BatchSize))
	}

	return opts, nil
}

// Options parses a writer config and returns a list of options.
// Return an error if parse failed.
func (wc *WriterConfig) Options() (opts []logit.Option, err error) {
	opts = make([]logit.Option, 0, 4)

	appendFuncs := []func(opts []logit.Option) ([]logit.Option, error){
		wc.appendTargetOptions, wc.appendModeOptions,
	}

	for _, append := range appendFuncs {
		opts, err = append(opts)
		if err != nil {
			return nil, err
		}
	}

	return opts, nil
}

type Config struct {
	// Level is the level of logger.
	// Values: debug, info, warn, error.
	Level string `json:"level" yaml:"level" toml:"level" bson:"level"`

	// Handler is how the handler handles the logs.
	// Values: "text", "json", "slog.text", "slog.json".
	// These handlers with "slog" prefix are from slog package of Go.
	// We recommend you to use our faster handlers, and feel free if you want to use slog's handlers.
	// Also, you can register your handlers to logit, see RegisterHandler.
	Handler string `json:"handler" yaml:"handler" toml:"handler" bson:"handler"`

	// Writer is the config of writer.
	Writer WriterConfig `json:"writer" yaml:"writer" toml:"writer" bson:"writer"`

	// WithSource adds source to logs if true.
	WithSource bool `json:"with_source" yaml:"with_source" toml:"with_source" bson:"with_source"`

	// WithPID adds pid to logs if true.
	WithPID bool `json:"with_pid" yaml:"with_pid" toml:"with_pid" bson:"with_pid"`

	// SyncTimer is the timer duration of syncing.
	// An empty string means syncing is manual.
	// You can use common words like "5m" or "60s".
	// See time.Duration and time.ParseDuration.
	SyncTimer string `json:"sync_timer" yaml:"sync_timer" toml:"sync_timer" bson:"sync_timer"`
}

func (c *Config) appendLevelOptions(opts []logit.Option) ([]logit.Option, error) {
	if c.Level == "" {
		return opts, nil
	}

	level := strings.ToLower(c.Level)

	if level == "debug" {
		opts = append(opts, logit.WithDebugLevel())
		return opts, nil
	}

	if level == "info" {
		opts = append(opts, logit.WithInfoLevel())
		return opts, nil
	}

	if level == "warn" {
		opts = append(opts, logit.WithWarnLevel())
		return opts, nil
	}

	if level == "error" {
		opts = append(opts, logit.WithErrorLevel())
		return opts, nil
	}

	return nil, fmt.Errorf("logit: level %s unknown", level)
}

func (c *Config) appendHandlerOptions(opts []logit.Option) ([]logit.Option, error) {
	if c.Handler == "" {
		return opts, nil
	}

	handler := strings.ToLower(c.Handler)

	newHandler, err := pickHandler(handler)
	if err != nil {
		return nil, err
	}

	opts = append(opts, logit.WithHandler(newHandler))
	return opts, nil
}

func (c *Config) appendFlagOptions(opts []logit.Option) ([]logit.Option, error) {
	if c.WithSource {
		opts = append(opts, logit.WithSource())
	}

	if c.WithPID {
		opts = append(opts, logit.WithPID())
	}

	return opts, nil
}

func (c *Config) appendSyncOptions(opts []logit.Option) ([]logit.Option, error) {
	if c.SyncTimer == "" {
		return opts, nil
	}

	syncTimer, err := parseTimeDuration(c.SyncTimer)
	if err != nil {
		return nil, err
	}

	opts = append(opts, logit.WithSyncTimer(syncTimer))
	return opts, nil
}

// Options parses a config and returns a list of options.
// Return an error if parse failed.
func (c *Config) Options() (opts []logit.Option, err error) {
	opts = make([]logit.Option, 0, 4)

	appendFuncs := []func(opts []logit.Option) ([]logit.Option, error){
		c.appendLevelOptions, c.appendHandlerOptions, c.appendFlagOptions, c.appendSyncOptions,
	}

	for _, append := range appendFuncs {
		opts, err = append(opts)
		if err != nil {
			return nil, err
		}
	}

	return opts, nil
}
