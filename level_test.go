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
	"log/slog"
	"testing"
)

// go test -v -cover -run=^TestStringifyLevel$
func TestStringifyLevel(t *testing.T) {
	tests := []struct {
		name  string
		level slog.Level
		want  string
	}{
		{name: "debug", level: levelDebug, want: "debug"},
		{name: "info", level: levelInfo, want: "info"},
		{name: "warn", level: levelWarn, want: "warn"},
		{name: "error", level: levelError, want: "error"},
		{name: "print", level: levelPrint, want: "print"},
		{name: "off", level: levelOff, want: "off"},
		{name: "unknown", level: 1997, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringifyLevel(tt.level); got != tt.want {
				t.Fatalf("stringifyLevel(tt.level) = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -cover -run=^TestParseLevel$
func TestParseLevel(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    slog.Level
		wantErr bool
	}{
		{name: "debug", str: "debug", want: levelDebug, wantErr: false},
		{name: "info", str: "info", want: levelInfo, wantErr: false},
		{name: "warn", str: "warn", want: levelWarn, wantErr: false},
		{name: "error", str: "error", want: levelError, wantErr: false},
		{name: "print", str: "print", want: levelPrint, wantErr: false},
		{name: "off", str: "off", want: levelOff, wantErr: false},
		{name: "unknown", str: "unknown", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.str)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Fatalf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
