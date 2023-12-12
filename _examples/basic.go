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

package main

import (
	"context"
	"fmt"

	"github.com/FishGoddess/logit"
)

func main() {
	// Use default logger to log.
	// By default, logs will be output to stdout.
	logit.Default().Info("hello from logit", "key", 123)

	// Use a new logger to log.
	// By default, logs will be output to stdout.
	logger := logit.NewLogger()

	logger.Debug("new version of logit", "version", "1.5.0-alpha", "date", 20231122)
	logger.Error("new version of logit", "version", "1.5.0-alpha", "date", 20231122)

	type user struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	u := user{123456, "fishgoddess"}
	logger.Info("user information", "user", u, "pi", 3.14)

	// Yep, I know you want to output logs to a file, try WithFile option.
	// The path in WithFile is where the log file will be stored.
	// Also, it's a good choice to call logger.Close() when program shutdown.
	logger = logit.NewLogger(logit.WithFile("./logit.log"))
	defer logger.Close()

	logger.Info("check where I'm logged", "file", "logit.log")

	// What if I want to use default logger and output logs to a file? Try SetDefault.
	// It sets a logger to default and you can use it by package function or Default().
	logit.SetDefault(logger)
	logit.Default().Warn("this is from default logger", "pi", 3.14, "default", true)

	// If you want to change level of logger to info, try WithInfoLevel.
	// Other levels is similar to info level.
	logger = logit.NewLogger(logit.WithInfoLevel())

	logger.Debug("debug logs will be ignored")
	logger.Info("info logs can be logged")

	// If you want to pass logger by context, use NewContext and FromContext.
	ctx := logit.NewContext(context.Background(), logger)

	logger = logit.FromContext(ctx)
	logger.Info("logger from context", "from", "context")

	// Don't want to panic when new a logger? Try NewLoggerGracefully.
	logger, err := logit.NewLoggerGracefully(logit.WithFile(""))
	if err != nil {
		fmt.Println("new logger gracefully failed:", err)
	}
}
