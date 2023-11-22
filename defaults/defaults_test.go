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

package defaults

import (
	"io"
	"testing"
)

// go test -v -cover -run=^TestMarshalToJson$
func TestMarshalToJson(t *testing.T) {
	marshaled, err := MarshalToJson([]int{1, 2, 3})
	if err != nil {
		t.Error(err)
	}

	if string(marshaled) != "[1,2,3]" {
		t.Errorf("string(marshaled) %s != [1,2,3]", marshaled)
	}

	marshaled, err = MarshalToJson(map[string]interface{}{"key": 666, "str": "abc"})
	if err != nil {
		t.Error(err)
	}

	if string(marshaled) != "{\"key\":666,\"str\":\"abc\"}" {
		t.Errorf("string(marshaled) %v != {\"key\":666,\"str\":\"abc\"}", marshaled)
	}
}

// go test -v -cover -run=^TestHandleError$
func TestHandleError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	HandleError("", nil)
	HandleError("", io.EOF)
}
