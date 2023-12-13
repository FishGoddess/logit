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
	"os"

	"github.com/FishGoddess/logit"
)

func main() {
	// As you can see, NewLogger can use some options to create a logger.
	logger := logit.NewLogger(logit.WithDebugLevel())
	logger.Debug("debug log")

	// We provide some options for different scenes and all options have prefix "With".
	// Change logger level:
	logit.WithDebugLevel()
	logit.WithInfoLevel()
	logit.WithWarnLevel()
	logit.WithDebugLevel()

	// Change logger handler:
	logit.WithHandler("xxx")
	logit.WithStandardHandler()
	logit.WithTextHandler()
	logit.WithJsonHandler()

	// Change handler writer:
	logit.WithWriter(os.Stdout)
	logit.WithStdout()
	logit.WithStderr()
	logit.WithFile("")
	logit.WithRotateFile("")

	// Some useful flags:
	logit.WithSource()
	logit.WithPID()

	// More options can be found in logit package which have prefix "With".
	// What's more? We provide a options pack that we think it's useful in production.
	// It outputs logs to a rotate file using batch write, so you should call Sync() or Close() when shutdown.
	opts := logit.ProductionOptions()

	logger = logit.NewLogger(opts...)
	defer logger.Close()

	logger.Info("log from production options")
	logger.Error("error log from production options")
}
