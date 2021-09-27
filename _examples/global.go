// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2021/07/14 23:46:31

package main

import (
	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/core"
	"github.com/FishGoddess/logit/core/writer"
)

func main() {
	// There are some global settings for optimizations, and you can set all of them in need.
	//
	//     import "github.com/FishGoddess/logit/core"
	//
	// All global settings are stored in package core.

	// 1. LogMallocSize (The pre-malloc size of a new Log data)
	// If your logs are extremely long, such as 4000 bytes, you can set it to 4096 to avoid re-malloc.
	core.LogMallocSize = 4096 // 4096 Bytes

	// 2. WriterBufferedSize (The default size of buffered writer)
	// If your logs are extremely long, such as 16KB, you can set it to 2048 to avoid re-malloc.
	core.WriterBufferedSize = 32 * writer.KB

	// After setting global settings, just use Logger as normal.
	logger := logit.NewLogger()
	defer logger.Close()

	logger.Info("set global settings").Int("LogMallocSize", core.LogMallocSize).Int("WriterBufferedSize", core.WriterBufferedSize).End()
}
