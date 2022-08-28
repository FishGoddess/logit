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
	"github.com/go-logit/logit/core/writer"
	"github.com/go-logit/logit/support/size"
)

func mustCreateFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	// Logger will log everything to console by default.
	logger := logit.NewLogger()
	logger.Info("I log everything to console.").Log()

	// You can use WithWriter to change writer in logger.
	logger = logit.NewLogger(logit.Options().WithWriter(os.Stdout))
	logger.Info("I also log everything to console.").Log()

	// As we know, we always log everything to file in production.
	logFile := filepath.Join(os.TempDir(), "test.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(mustCreateFile(logFile)))
	logger.Info("I log everything to file.").String("logFile", logFile).Log()
	logger.Close()

	// Also, as you can see, there is a parameter called withBuffer in WithWriter option.
	// It will use a buffer writer to write logs if withBuffer is true which will bring a huge performance improvement.
	logFile = filepath.Join(os.TempDir(), "test_buffer.log")
	fmt.Println(logFile)
	logger = logit.NewLogger(logit.Options().WithWriter(mustCreateFile(logFile)))
	logger.Info("I log everything to file with buffer.").String("logFile", logFile).Log()
	logger.Close()

	// We provide some high-performance file for you. Try these:
	writer.BufferWithSize(os.Stdout, 128*size.KB)
	writer.BatchWithCount(os.Stdout, 256)
	logit.Options().WithBufferWriter(os.Stdout)
	logit.Options().WithBatchWriter(os.Stdout)
}
