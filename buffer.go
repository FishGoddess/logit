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

package logit

import (
	"sync"

	"github.com/FishGoddess/logit/defaults"
)

var bufferPool = sync.Pool{
	New: func() any {
		bs := make([]byte, 0, defaults.MinBufferSize)
		return &bs
	},
}

func newBuffer() *[]byte {
	return bufferPool.Get().(*[]byte)
}

func freeBuffer(b *[]byte) {
	// Return only smaller buffers for reducing peak allocation.
	if cap(*b) <= defaults.MaxBufferSize {
		*b = (*b)[:0]
		bufferPool.Put(b)
	}
}
