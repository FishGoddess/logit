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
	"os"
	"path/filepath"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/pkg/file"
)

func main() {
	// Logger will log everything to console by default.
	logger := logit.NewLogger()
	logger.Info("I log everything to console.").Log()

	// You can use WithWriter to change writer in logger.
	logger = logit.NewLogger(logit.Options().WithWriter(os.Stdout, false))
	logger.Info("I also log everything to console.").Log()

	// As we know, we always log everything to file in production.
	// So we provide a convenient way to create a file.
	logFile := filepath.Join(os.TempDir(), "test.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(file.MustNewFile(logFile), false))
	logger.Info("I log everything to file.").String("logFile", logFile).Log()
	logger.Close()

	// Also, as you can see, there is a parameter called withBuffer in WithWriter option.
	// It will use a buffer writer to write logs if withBuffer is true which will bring a huge performance improvement.
	logFile = filepath.Join(os.TempDir(), "test_buffer.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(file.MustNewFile(logFile), false))
	logger.Info("I log everything to file with buffer.").String("logFile", logFile).Log()
	logger.Close()
}
