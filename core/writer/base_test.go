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
// Created at 2021/07/11 14:06:21

package writer

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestWrapped$
func TestWrapped(t *testing.T) {
	writer := Wrapped(os.Stdout)

	_, ok := writer.(*wrappedWriter)
	if !ok {
		t.Error("Wrapped returns a non-wrappedWriter instance")
	}
}

// go test -v -cover -run=^TestBuffered$
func TestBuffered(t *testing.T) {
	writer := Buffered(os.Stdout)

	_, ok := writer.(*bufferedWriter)
	if !ok {
		t.Error("Buffered returns a non-bufferedWriter instance")
	}
}

// go test -v -cover -run=^TestBufferedWithSize$
func TestBufferedWithSize(t *testing.T) {
	writer := BufferedWithSize(os.Stdout, 1024)

	bw, ok := writer.(*bufferedWriter)
	if !ok {
		t.Error("Buffered returns a non-bufferedWriter instance")
	}

	if bw.bufferSize != 1024 {
		t.Errorf("bw.bufferSize %d is wrong", bw.bufferSize)
	}
}
