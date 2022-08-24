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

package core

import (
	"encoding/json"
	"time"
)

const (
	B ByteSize = 1 << (10 * iota)
	KB
	MB
	GB
)

var (
	// LogMallocSize is the pre-malloc size of a new Log data.
	// If your logs are extremely long, such as 4000 bytes/log, you can set it to 4KB to avoid re-malloc.
	LogMallocSize = 512 * B

	// WriterBufferSize is the default size of buffer writer.
	// If your logs are extremely long, such as 4 KB/log, you can set it to 512 KB to avoid re-malloc.
	WriterBufferSize = 64 * KB

	// WriterBatchCount is the default count of batch writer.
	WriterBatchCount = uint(128)

	// MarshalToJson marshals v to json bytes.
	// If you want to use your own way to marshal, change it to your own marshal function.
	MarshalToJson = json.Marshal

	// CallerDepth is the depth of caller.
	// See runtime.Caller.
	CallerDepth = 4

	// CurrentTime returns the current time with time.Time.
	CurrentTime = time.Now
)

// ByteSize is size of byte.
type ByteSize = uint64
