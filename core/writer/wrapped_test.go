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
// Created at 2021/07/11 14:04:11

package writer

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestWrappedWriterClose$
func TestWrappedWriterClose(t *testing.T) {
	writer := newWrappedWriter(os.Stdout)
	for i := 0; i < 10; i++ {
		err := writer.Close()
		if err != nil {
			t.Error(err)
		}
	}

	writer = newWrappedWriter(os.Stderr)
	for i := 0; i < 10; i++ {
		err := writer.Close()
		if err != nil {
			t.Error(err)
		}
	}
}
