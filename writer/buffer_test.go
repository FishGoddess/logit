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
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBuffer$
func TestBuffer(t *testing.T) {
	writer := Buffer(os.Stdout, 1024)

	if writer == nil {
		t.Fatal("writer == nil")
	}

	if writer.maxBufferSize != 1024 {
		t.Fatalf("writer.maxBufferSize %d is wrong", writer.maxBufferSize)
	}

	newWriter := Buffer(writer, 4096)
	if newWriter != writer {
		t.Fatal("newWriter is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBufferWriter$
func TestBufferWriter(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))

	writer := Buffer(buffer, 65536)
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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBufferWriterClose$
func TestBufferWriterClose(t *testing.T) {
	writer := Buffer(os.Stdout, 4096)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}

	writer = Buffer(os.Stderr, 4096)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
