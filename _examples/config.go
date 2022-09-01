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
	"fmt"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/extension/config"
)

func main() {
	// We provide a config which can be converted to option in logit.
	// It has many tags in fields, such json, yaml, toml, which means you can use config file to create logger.
	// You just need to define your config file then unmarshal your config file to this config.
	// Of course, you can embed this struct to your application config struct!
	cfg := config.Config{
		Level:         config.LevelDebug,
		TimeKey:       "log.time",
		LevelKey:      "log.level",
		MsgKey:        "log.msg",
		PIDKey:        "log.pid",
		FileKey:       "log.file",
		LineKey:       "log.line",
		FuncKey:       "log.func",
		TimeFormat:    config.UnixTimeFormat,
		WithPID:       false,
		WithCaller:    false,
		CallerDepth:   0,
		AutoSync:      "",
		Appender:      config.AppenderText,
		DebugAppender: "",
		InfoAppender:  "",
		WarnAppender:  "",
		ErrorAppender: "",
		PrintAppender: "",
		Writer: config.WriterConfig{
			Target:     config.WriterTargetStdout,
			Mode:       config.WriterModeDirect,
			Filename:   "",
			BufferSize: "4MB",
			BatchCount: 1024,
		},
		DebugWriter: config.WriterConfig{},
		InfoWriter:  config.WriterConfig{},
		WarnWriter:  config.WriterConfig{},
		ErrorWriter: config.WriterConfig{},
		PrintWriter: config.WriterConfig{},
	}

	// Once you got a config, use Options() to convert to option in logger.
	opts, err := cfg.Options()
	if err != nil {
		panic(err)
	}
	fmt.Println(opts)

	// Then you can create your logger by options.
	// Amazing!
	logger := logit.NewLogger(opts...)
	defer logger.Close()
	logger.Info("My mother is a config").Any("config", cfg).Log()
}
