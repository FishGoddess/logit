// Copyright 2024 FishGoddess. All Rights Reserved.
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

package defaults

import (
	"log/slog"
	"os"
	"time"
)

var (
	// CurrentTime returns the current time with time.Time.
	CurrentTime = time.Now

	// HandleError handles an error passed to it.
	// You can collect all errors and count them for reporting.
	// Notice that this function is called synchronously, so don't do too many things in it.
	HandleError = func(label string, err error) {}
)

var (
	// CallerDepth is the depth of caller.
	// See runtime.Caller.
	CallerDepth = 4

	// LevelPrint is the level used for printing logs.
	LevelPrint = slog.LevelInfo
)

var (
	// MinBufferSize is the min buffer size used in bytes.
	MinBufferSize = 1 * 1024

	// MaxBufferSize is the max buffer size used in bytes.
	MaxBufferSize = 16 * 1024
)

var (
	// FileMode is the permission bits
	FileMode os.FileMode = 0644

	// FileDirMode is the permission bits of directory.
	FileDirMode os.FileMode = 0755
)

var (
	// OpenFile opens a file of path with given mode.
	OpenFile = func(path string, mode os.FileMode) (*os.File, error) {
		return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, mode)
	}

	// OpenFileDir opens a dir of path with given mode.
	OpenFileDir = func(path string, mode os.FileMode) error {
		return os.MkdirAll(path, mode)
	}
)
