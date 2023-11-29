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
	"reflect"
	"testing"
)

// go test -v -cover -run=^TestLevelPeel$
func TestLevelPeel(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		want  slog.Level
	}{
		{name: "debug", level: LevelDebug, want: slog.LevelDebug},
		{name: "info", level: LevelInfo, want: slog.LevelInfo},
		{name: "warn", level: LevelWarn, want: slog.LevelWarn},
		{name: "error", level: LevelError, want: slog.LevelError},
		{name: "print", level: LevelPrint, want: slog.Level(LevelPrint)},
		{name: "off", level: LevelOff, want: slog.Level(LevelOff)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.Peel(); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Level.Peel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -cover -run=^TestLevelString$
func TestLevelString(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		want  string
	}{
		{name: "debug", level: LevelDebug, want: "debug"},
		{name: "info", level: LevelInfo, want: "info"},
		{name: "warn", level: LevelWarn, want: "warn"},
		{name: "error", level: LevelError, want: "error"},
		{name: "print", level: LevelPrint, want: "print"},
		{name: "off", level: LevelOff, want: "off"},
		{name: "unknown", level: 1997, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.want {
				t.Fatalf("Level.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -cover -run=^TestParseLevel$
func TestParseLevel(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    Level
		wantErr bool
	}{
		{name: "debug", str: "debug", want: LevelDebug, wantErr: false},
		{name: "info", str: "info", want: LevelInfo, wantErr: false},
		{name: "warn", str: "warn", want: LevelWarn, wantErr: false},
		{name: "error", str: "error", want: LevelError, wantErr: false},
		{name: "print", str: "print", want: LevelPrint, wantErr: false},
		{name: "off", str: "off", want: LevelOff, wantErr: false},
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
