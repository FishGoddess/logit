// Copyright 2022 FishGoddess. All Rights Reserved.
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

package core

import (
	"io"
	"testing"
)

// go test -v -cover -run=^TestParseByteSize$
func TestParseByteSize(t *testing.T) {
	cases := map[string]ByteSize{
		"128":    128 * B,
		"256B":   256 * B,
		"64K":    64 * KB,
		"16k":    16 * KB,
		"512KB":  512 * KB,
		"1024kB": 1024 * KB,
		"64M":    64 * MB,
		"16m":    16 * MB,
		"512MB":  512 * MB,
		"1024mB": 1024 * MB,
		"64G":    64 * GB,
		"16g":    16 * GB,
		"512GB":  512 * GB,
		"1024gB": 1024 * GB,
		"512Kb":  512 / 8 * KB,
		"1024kb": 1024 / 8 * KB,
		"512Mb":  512 / 8 * MB,
		"1024mb": 1024 / 8 * MB,
		"512Gb":  512 / 8 * GB,
		"1024gb": 1024 / 8 * GB,
	}

	for str, rightSize := range cases {
		parsedSize, err := ParseByteSize(str)
		if err != nil {
			t.Error(err)
		}

		if parsedSize != rightSize {
			t.Errorf("str %s parsedSize %d != rightSize %d", str, parsedSize, rightSize)
		}
	}

	_, err := ParseByteSize("   ")
	if err == nil {
		t.Error("ParseByteSize '   ' should be error")
	}

	_, err = ParseByteSize("xxx")
	if err == nil {
		t.Error("ParseByteSize 'xxx' should be error")
	}
}

// go test -v -cover -run=^TestHandleError$
func TestHandleError(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Error(r)
		}
	}()

	OnError = nil
	HandleError("", nil)
	HandleError("", io.EOF)

	OnError = func(name string, err error) {
		// ...
	}
	HandleError("", nil)
	HandleError("", io.EOF)
}
