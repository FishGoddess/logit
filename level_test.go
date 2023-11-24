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

package logit

import (
	"reflect"
	"testing"
)

// go test -v -cover -run=^TestParseLevel$
func TestParseLevel(t *testing.T) {
	type testCase struct {
		name string
		str  string
		want level
	}

	tests := []testCase{
		{
			name: "debug",
			str:  "debug",
			want: debugLevel,
		},
		{
			name: "info",
			str:  "info",
			want: infoLevel,
		},
		{
			name: "warn",
			str:  "warn",
			want: warnLevel,
		},
		{
			name: "error",
			str:  "error",
			want: errorLevel,
		},
		{
			name: "print",
			str:  "print",
			want: printLevel,
		},
		{
			name: "off",
			str:  "off",
			want: offLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLevel(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
