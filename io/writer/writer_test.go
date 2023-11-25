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

package writer

import (
	"os"
	"testing"
)

// go test -v -cover -run=^TestNotStdoutAndStderr$
func TestNotStdoutAndStderr(t *testing.T) {
	if notStdoutAndStderr(os.Stdout) {
		t.Error("notStdoutAndStderr(os.Stdout) returns true")
	}

	if notStdoutAndStderr(os.Stderr) {
		t.Error("notStdoutAndStderr(os.Stderr) returns true")
	}
}

// go test -v -cover -run=^TestWrapped$
func TestWrapped(t *testing.T) {
	writer := Wrap(os.Stdout)

	if _, ok := writer.(*wrapWriter); !ok {
		t.Error("Wrap returns a non-wrapWriter instance")
	}
}

// go test -v -cover -run=^TestBuffer$
func TestBuffer(t *testing.T) {
	writer := Buffer(os.Stdout, 1024)

	if _, ok := writer.(*bufferWriter); !ok {
		t.Error("Buffer returns a non-bufferWriter instance")
	}
}

// go test -v -cover -run=^TestBufferWithSize$
func TestBufferWithSize(t *testing.T) {
	writer := Buffer(os.Stdout, 1024)

	bw, ok := writer.(*bufferWriter)
	if !ok {
		t.Error("Buffer returns a non-bufferWriter instance")
	}

	if bw.maxBufferSize != 1024 {
		t.Errorf("bw.maxBufferSize %d is wrong", bw.maxBufferSize)
	}
}
