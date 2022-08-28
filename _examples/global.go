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
	"encoding/json"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/support/global"
	"github.com/go-logit/logit/support/size"
)

func main() {
	// There are some global settings for optimizations, and you can set all of them in need.

	// 1. LogMallocSize (The pre-malloc size of a new Log data)
	// If your logs are extremely long, such as 4000 bytes, you can set it to 4096 to avoid re-malloc.
	global.LogMallocSize = 4 * size.MB

	// 2. WriterBufferSize (The default size of buffer writer)
	// If your logs are extremely long, such as 16 KB, you can set it to 2048 to avoid re-malloc.
	global.WriterBufferSize = 32 * size.KB

	// 3. MarshalToJson (The marshal function which marshal interface{} to json data)
	// Use std by default, and you can customize your marshal function.
	global.MarshalToJson = json.Marshal

	// After setting global settings, just use Logger as normal.
	logger := logit.NewLogger()
	defer logger.Close()

	logger.Info("set global settings").Uint64("LogMallocSize", global.LogMallocSize).Uint64("WriterBufferSize", global.WriterBufferSize).Log()
}
