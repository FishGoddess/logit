// Copyright 2025 FishGoddess. All Rights Reserved.
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
	// A new logger outputs logs to stdout.
	logger := logit.NewLogger()
	logger.Debug("log to stdout")

	// What if I want to output logs to stderr? Try WithStderr.
	logger = logit.NewLogger(logit.WithStderr())
	logger.Debug("log to stderr")

	// Also, you can use WithWriter to specify your own writer.
	logger = logit.NewLogger(logit.WithWriter(os.Stdout))
	logger.Debug("log to writer")

	// How to output logs to a file? Try WithFile and WithRotateFile.
	// Rotate file is useful in production, see _examples/file.go.
	logger = logit.NewLogger(logit.WithFile("logit.log"))
	logger.Debug("log to file")

	logger = logit.NewLogger(logit.WithRotateFile("logit.log"))
	logger.Debug("log to rotate file")
}
