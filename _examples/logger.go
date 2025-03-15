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
	"io"

	"github.com/FishGoddess/logit"
)

func main() {
	// Default() will return the default logger.
	// You can new a logger or just use the default logger.
	logger := logit.Default()
	logger.Info("nothing carried")

	// Use With() to carry some args in logger.
	// All logs output by this logger will carry these args.
	logger = logger.With("carry", 666, "who", "me")

	logger.Info("see what are carried")
	logger.Error("error carried", "err", io.EOF)

	// Use WithGroup() to group args in logger.
	// All logs output by this logger will group args.
	logger = logger.WithGroup("xxx")

	logger.Info("what group")
	logger.Error("error group", "err", io.EOF)

	// If you want to check if one level can be logged, try this:
	if logger.DebugEnabled() {
		logger.Debug("debug enabled")
	}

	// We provide some old-school logging methods.
	// They are using info level by default.
	// If you want to change the level, see defaults.LevelPrint.
	logger.Printf("printf %s log", "formatted")
	logger.Print("print log")
	logger.Println("println log")

	// Some useful method:
	logger.Sync()
	logger.Close()
}
