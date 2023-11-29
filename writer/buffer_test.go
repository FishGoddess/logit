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
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

// go test -v -cover -run=^TestBufferWriter$
func TestBufferWriter(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))

	writer := newBufferWriter(buffer, defaults.BufferSize)
	defer writer.Close()

	writer.Write([]byte("abc"))
	writer.Sync()

	if buffer.String() != "abc" {
		t.Fatalf("writing abc but found %s in buffer", buffer.String())
	}

	writer.Write([]byte("123"))
	writer.Write([]byte(".!?"))
	writer.Write([]byte("+-*/"))
	writer.Close()
	time.Sleep(time.Second)

	if buffer.String() != "abc123.!?+-*/" {
		t.Fatalf("writing abc123.!?+-*/ but found %s in buffer", buffer.String())
	}
}

// go test -v -cover -run=^TestBufferWriterClose$
func TestBufferWriterClose(t *testing.T) {
	writer := newBufferWriter(os.Stdout, 4096)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}

	writer = newBufferWriter(os.Stderr, 4096)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
