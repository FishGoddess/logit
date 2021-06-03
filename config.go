// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/25 23:06:59

package logit

import (
	"fmt"
)

var (
	// encoders store all encoders provided.
	// Actually, this field is for me, not for you, ha:)
	encoders = map[string]Encoder{
		"text": NewTextEncoder(TimeFormat),
		"json": NewJsonEncoder(TimeFormat),
	}

	// defaultConfig is default config.
	defaultConfig = Config{
		LogFileName:      "logit.log",
		MaxLogFileNumber: 100,
		MaxLogFileSize:   1024,
	}
)

// Config is the configuration of logit.
type Config struct {

	// LogFileName is the name of log file.
	LogFileName string

	// MaxLogFileNumber is the max number of log files.
	MaxLogFileNumber int

	// MaxLogFileSize is the max size of one log file.
	// The unit is MB.
	MaxLogFileSize int64
}

// DefaultConfig returns a default config.
func DefaultConfig() Config {
	return defaultConfig
}

// checkConfig checks if this config is legal.
func checkConfig(config Config) {

	if config.MaxLogFileNumber < 1 {
		panic(fmt.Errorf("config.MaxFileNumber %d is less than 1", config.MaxLogFileNumber))
	}

	if config.MaxLogFileSize < 1 {
		panic(fmt.Errorf("config.MaxLogFileSize %dMB is less than 1MB", config.MaxLogFileSize))
	}
}

// ======================================== for convenience ========================================

// levelOf returns the real level of passed level and an error if not found.
func levelOf(level string) (Level, error) {
	for k, v := range levels {
		if v == level {
			return k, nil
		}
	}
	return OffLevel, fmt.Errorf("level \"%s\" doesn't exist! Try: debug, info, warn, error, off", level)
}

// encoderOf returns the encoder called name.
// If the encoder doesn't exist, an error will be returned.
func encoderOf(name string) (Encoder, error) {
	encoder, ok := encoders[name]
	if !ok {
		return nil, fmt.Errorf("encoder \"%s\" doesn't exist! Try: text or json", name)
	}
	return encoder, nil
}
