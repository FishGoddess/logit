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

package pkg

import (
	"os"
	"runtime"
)

var (
	pid = os.Getpid() // The pid of current process.
)

// Pid returns the pid of current process.
func Pid() int {
	return pid
}

// Caller returns the caller information of depth.
func Caller(depth int) (file string, line int) {
	if _, file, line, ok := runtime.Caller(depth); ok {
		return file, line
	}

	return "unknown file", -1
}
