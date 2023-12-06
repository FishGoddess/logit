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

type WriterConfig struct {
	// Target is where the writer writes logs.
	// Values: "stdout", "stderr", "file".
	Target string `json:"target" yaml:"target" toml:"target" bson:"target"`

	// Mode is how the writer writes logs.
	// Values: "direct", "buffer", "batch".
	// Direct means writer writes logs without any buffer or batch, which is one log one writing operation.
	// Buffer means writer will keep logs in a buffer and write these logs once until the buffer is full.
	// Batch means writer will keep logs in a batch and write these logs once until the count of logs in this batch >= batch size.
	// Both of buffer and batch have better performance in writing logs.
	// However, they will lose some logs if the program crashed before syncing.
	Mode string `json:"mode" yaml:"mode" toml:"mode" bson:"mode"`

	// BufferSize is the size of a buffer.
	// You can use common words like "512B" or "4KB".
	// Only available when mode is "buffer".
	BufferSize string `json:"buffer_size" yaml:"buffer_size" toml:"buffer_size" bson:"buffer_size"`

	// BatchSize is the size of a batch.
	// Only available when mode is "batch".
	BatchSize uint64 `json:"batch_size" yaml:"batch_size" toml:"batch_size" bson:"batch_size"`
}

type FileConfig struct {
	// Path is the path (or prefix) of log file.
	Path string `json:"path" yaml:"path" toml:"path" bson:"path"`

	// Rotate is log file should split and backup when satisfy some conditions.
	// It's useful in production so we recommend you to set it to true.
	Rotate bool `json:"rotate" yaml:"rotate" toml:"rotate" bson:"rotate"`

	// MaxSize is the max size of a log file.
	// If size of data in one output operation is bigger than this value, then file will rotate before writing,
	// which means file and its backups may be bigger than this value in size.
	// You can use common words like "100MB" or "1GB".
	// Only available when rotate is true.
	MaxSize string `json:"max_size" yaml:"max_size" toml:"max_size" bson:"max_size"`

	// MaxAge is the time that backups will live.
	// All backups reach max age will be removed automatically.
	// You can use common words like "7d" or "24h".
	// See time.Duration and time.ParseDuration.
	// Only available when rotate is true.
	MaxAge string `json:"max_age" yaml:"max_age" toml:"max_age" bson:"max_age"`

	// MaxBackups is the max count of file backups.
	// Only available when rotate is true.
	MaxBackups uint32 `json:"max_backups" yaml:"max_backups" toml:"max_backups" bson:"max_backups"`
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

	// File is the config of file.
	// Only available when the target of writer is "file".
	File FileConfig `json:"file" yaml:"file" toml:"file" bson:"file"`

	// WithSource adds source to logs if true.
	WithSource bool `json:"with_source" yaml:"with_source" toml:"with_source" bson:"with_source"`

	// WithPID adds pid to logs if true.
	WithPID bool `json:"with_pid" yaml:"with_pid" toml:"with_pid" bson:"with_pid"`

	// AutoSync is the frequency of syncing.
	// An empty string means syncing is manual.
	// You can use common words like "5m" or "60s".
	// See time.Duration and time.ParseDuration.
	AutoSync string `json:"auto_sync" yaml:"auto_sync" toml:"auto_sync" bson:"auto_sync"`
}
