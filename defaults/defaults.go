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

package defaults

import (
	"encoding/json"
	"os"
	"time"

	"github.com/FishGoddess/logit/io/size"
)

const (
	// UnixTimeFormat is time format of unix.
	UnixTimeFormat = ""
)

var (
	// CallerDepth is the depth of caller.
	// See runtime.Caller.
	CallerDepth = 4

	// LogSize is the pre-malloc size of a new log.
	// If your logs are extremely long, such as 4000 bytes/log, you can set it to 4KB to avoid re-malloc.
	LogSize = 512 * size.B

	// WriterBufferSize is the default size of buffer writer.
	WriterBufferSize = 64 * size.KB

	// WriterBatchCount is the default count of batch writer.
	WriterBatchCount = uint(128)

	// TimeLocation is the location of time.
	TimeLocation = time.Local

	// CurrentTime returns the current time with time.Time.
	CurrentTime = time.Now

	// MarshalToJson marshals v to json bytes.
	// If you want to use your own way to marshal, change it to your own marshal function.
	MarshalToJson = json.Marshal

	// OpenFile opens a file of path with given mode.
	OpenFile = func(path string, mode os.FileMode) (*os.File, error) {
		return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, mode)
	}

	// HandleError handles an error passed to it.
	// You can collect all errors and count them for reporting.
	// Notice that this function is called synchronously, so don't do too many things in it.
	HandleError = func(label string, err error) {}
)
