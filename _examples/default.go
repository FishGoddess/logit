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
	"fmt"
	"log/slog"
	"time"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/defaults"
)

func main() {
	// We set a defaults package that setups all shared fields.
	// For example, if you want to customize the time getter:
	defaults.CurrentTime = func() time.Time {
		// Return a fixed time for example.
		return time.Unix(666, 0).In(time.Local)
	}

	logit.Print("println log is info level")

	// If you want change the level of old-school logging methods:
	defaults.LevelPrint = slog.LevelDebug

	logit.Print("println log is debug level now")

	// More fields see defaults package.
	defaults.HandleError = func(label string, err error) {
		fmt.Printf("%s: %+n\n", label, err)
	}
}
