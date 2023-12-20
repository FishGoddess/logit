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

package handler

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestAppendEscapedByte$
func TestAppendEscapedByte(t *testing.T) {
	testcases := []byte{'a', '0', '\n', '\t', '\\', '\b', '\f', '\r', '"'}
	want := "a0\\n\\t\\\\\\b\\f\\r\\\""

	buffer := make([]byte, 0, 16)
	for _, b := range testcases {
		buffer = appendEscapedByte(buffer, b)
	}

	if string(buffer) != want {
		t.Errorf("result %s is wrong", string(buffer))
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestAppendEscapedString$
func TestAppendEscapedString(t *testing.T) {
	testcases := []string{"a0国\n\t\\\b\f\r\""}
	want := "a0国\\n\\t\\\\\\b\\f\\r\\\""

	buffer := make([]byte, 0, 16)
	for _, str := range testcases {
		buffer = appendEscapedString(buffer, str)
	}

	if string(buffer) != want {
		t.Errorf("result %s is wrong", string(buffer))
	}
}
