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
	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/rotate"
)

func main() {
	// AS we know, you can use WithFile to output logs to a file.
	logger := logit.NewLogger(logit.WithFile("logit.log"))
	logger.Debug("debug to file")

	// However, a single file stored all logs isn't enough in production.
	// Sometimes we want a log file has a limit size and count of files not greater than a number.
	// So we provide a rotate file to do this thing.
	logger = logit.NewLogger(logit.WithRotateFile("logit.log"))
	defer logger.Close()

	logger.Debug("debug to rotate file")

	// Maybe you have noticed that WithRotateFile can pass some rotate.Option.
	// These options are used to setup the rotate file.
	opts := []rotate.Option{
		rotate.WithMaxSize(128 * rotate.MB),
		rotate.WithMaxAge(30 * rotate.Day),
		rotate.WithMaxBackups(60),
	}

	logger = logit.NewLogger(logit.WithRotateFile("logit.log", opts...))
	defer logger.Close()

	logger.Debug("debug to rotate file with rotate options")

	// See rotate.File if you want to use this magic in other scenes.
	file, err := rotate.New("logit.log")
	if err != nil {
		panic(err)
	}

	defer file.Close()
}
