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
// Created at 2020/12/01 00:19:00

package writer

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/go-logit/logit/core"
)

// go test -v -cover -run=^TestBufferedWriter$
func TestBufferedWriter(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0, 4096))
	writer := newBufferedWriter(buffer, core.WriterBufferedSize)
	writer.AutoFlush(time.Millisecond)
	defer writer.Close()

	writer.Write([]byte("abc"))
	writer.Flush()
	if buffer.String() != "abc" {
		t.Errorf("writing abc but found %s in buffer", buffer.String())
	}

	writer.Write([]byte("123"))
	writer.Write([]byte(".!?"))
	writer.Write([]byte("+-*/"))
	time.Sleep(time.Second)

	if buffer.String() != "abc123.!?+-*/" {
		t.Errorf("writing abc123.!?+-*/ but found %s in buffer", buffer.String())
	}
}

// go test -v -cover -run=^TestBufferedWriterClose$
func TestBufferedWriterClose(t *testing.T) {
	writer := newBufferedWriter(os.Stdout, 4096)
	for i := 0; i < 10; i++ {
		err := writer.Close()
		if err != nil {
			t.Error(err)
		}
	}

	writer = newBufferedWriter(os.Stderr, 4096)
	for i := 0; i < 10; i++ {
		err := writer.Close()
		if err != nil {
			t.Error(err)
		}
	}
}
