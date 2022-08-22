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
	"context"
	"io"
	"os"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core"
)

func main() {
	// Create a new logger for use.
	// Default level is debug, so all logs will be logged.
	// Invoke Close() isn't necessary in all situations.
	// If logger's writer has buffer or something like that, it's better to invoke Close() for flushing buffer or something else.
	logger := logit.NewLogger()
	//defer logger.Close()

	// Then, you can log anything you want.
	// Remember, logs will be ignored if their level is smaller than logger's level.
	// Log() will do some finishing work, so this invocation is necessary.
	logger.Debug("This is a debug message").Log()
	logger.Info("This is a info message").Log()
	logger.Warn("This is a warn message").Log()
	logger.Error("This is a error message").Log()
	logger.Error("This is a %s message, with format", "error").Log() // Format with params.

	// As you know, we provide some levels: debug, info, warn, error, off.
	// The lowest is debug and the highest is off.
	// If you want to change the level of your logger, do it at creating.
	logger = logit.NewLogger(logit.Options().WithWarnLevel())
	logger.Debug("This is a debug message, but ignored").Log()
	logger.Info("This is a info message, but ignored").Log()
	logger.Warn("This is a warn message, not ignored").Log()
	logger.Error("This is a error message, not ignored").Log()

	// Also, we provide some "old school" log method :)
	// (Don't mistake~ I love old school~)
	logger.Printf("This is a log %s, and it's for compatibility", "printed")
	logger.Print("This is a log printed, and it's for compatibility", 123)
	logger.Println("This is a log printed, and it's for compatibility", 666)

	// If you want to log with some fields, try this:
	user := struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   666,
		Name: "FishGoddess",
		Age:  3,
	}
	logger.Warn("This is a structured message").Any("user", user).Json("userJson", user).Log()
	logger.Error("This is a structured message").Error("err", io.EOF).Int("trace", 123).Log()

	// You may notice logit.Options() which returns an options list.
	// Here is some of them:
	options := logit.Options()
	options.WithCaller()                          // Let logs carry caller information.
	options.WithLevelKey("lvl")                   // Change logger's level key to "lvl".
	options.WithWriter(os.Stderr)                 // Change logger's writer to os.Stderr without buffer or batch.
	options.WithBufferWriter(os.Stderr)           // Change logger's writer to os.Stderr with buffer.
	options.WithBatchWriter(os.Stderr)            // Change logger's writer to os.Stderr with batch.
	options.WithErrorWriter(os.Stderr)            // Change logger's error writer to os.Stderr without buffer or batch.
	options.WithTimeFormat("2006-01-02 15:04:05") // Change the format of time (Only the log's time will apply it).

	// You can bind context with logger and use it as long as you can get the context.
	ctx := logit.NewContext(context.Background(), logger)
	logger = logit.FromContext(ctx)
	logger.Info("Logger from context").Log()

	// You can initialize the global logger if you don't want to use an independent logger.
	// WithCallerDepth will set the depth of caller, and default is core.CallerDepth.
	// Functions in global logger are wrapped so depth of caller should be increased 1.
	// You can specify your depth if you wrap again or have something else reasons.
	logger = logit.NewLogger(options.WithCallerDepth(core.CallerDepth + 1))
	logit.SetGlobal(logger)
	logit.Info("Info from logit").Log()

	// We don't recommend you to call logit.SetGlobal unless you really need to call.
	// Instead, we recommend you to call logger.SetToGlobal to set one logger to global if you need.
	logger.SetToGlobal()
	logit.Println("Println from logit")
}
