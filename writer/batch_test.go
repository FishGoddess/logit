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

// go test -v -cover -run=^TestBatchWriter$
func TestBatchWriter(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))

	writer := newBatchWriter(buffer, 10)
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

// go test -v -cover -run=^TestBatchWriterCount$
func TestBatchWriterCount(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))

	writer := newBatchWriter(buffer, 10)
	defer writer.Close()

	for i := 0; i < 10; i++ {
		writer.Write([]byte("1"))
	}

	if buffer.String() != "" {
		t.Fatalf("buffer.String() %s != ''", buffer.String())
	}

	writer.Sync()
	if buffer.String() != "1111111111" {
		t.Fatalf("buffer.String() %s != '1111111111'", buffer.String())
	}

	buffer.Reset()
	for i := 0; i < 12; i++ {
		writer.Write([]byte("1"))
	}

	if buffer.String() != "1111111111" {
		t.Fatalf("buffer.String() %s != '1111111111'", buffer.String())
	}

	writer.Sync()
	if buffer.String() != "111111111111" {
		t.Fatalf("buffer.String() %s != '111111111111'", buffer.String())
	}
}

// go test -v -cover -run=^TestBatchWriterClose$
func TestBatchWriterClose(t *testing.T) {
	writer := newBatchWriter(os.Stdout, 10)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}

	writer = newBatchWriter(os.Stderr, 10)
	for i := 0; i < 10; i++ {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
