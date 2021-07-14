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
// Created at 2021/07/14 23:31:47

package core

var (
	// LogMallocSize is the pre-malloc size of a new Log data.
	// If your logs are extremely long, such as 4000 bytes, you can set it to 4096 to avoid re-malloc.
	LogMallocSize = 512 // 512 Bytes

	// WriterBufferedSize is the default size of buffered writer.
	// If your logs are extremely long, such as 16KB, you can set it to 2048 to avoid re-malloc.
	WriterBufferedSize = 16 * 1024 // 16 KB
)
