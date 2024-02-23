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

package config

import (
	"testing"
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestParseByteSize$
func TestParseByteSize(t *testing.T) {
	tests := []struct {
		name    string
		size    string
		want    uint64
		wantErr bool
	}{
		{name: "64", size: "64", want: 64, wantErr: false},
		{name: "128b", size: "128b", want: 16, wantErr: false},
		{name: "256B", size: "256B", want: 256, wantErr: false},
		{name: "1K", size: "1K", want: 1024, wantErr: false},
		{name: "2k", size: "2k", want: 2048, wantErr: false},
		{name: "4Kb", size: "4Kb", want: 512, wantErr: false},
		{name: "8kb", size: "8kb", want: 1024, wantErr: false},
		{name: "4KB", size: "4KB", want: 4096, wantErr: false},
		{name: "16kB", size: "16kB", want: 16384, wantErr: false},
		{name: "1M", size: "1M", want: 1024 * 1024, wantErr: false},
		{name: "2Mb", size: "2Mb", want: 2 * 1024 * 1024 / 8, wantErr: false},
		{name: "3MB", size: "3MB", want: 3 * 1024 * 1024, wantErr: false},
		{name: "20mB", size: "20mB", want: 20 * 1024 * 1024, wantErr: false},
		{name: "24m", size: "24m", want: 24 * 1024 * 1024, wantErr: false},
		{name: "48mb", size: "48mb", want: 48 * 1024 * 1024 / 8, wantErr: false},
		{name: "1G", size: "1G", want: 1024 * 1024 * 1024, wantErr: false},
		{name: "21Gb", size: "21Gb", want: 21 * 1024 * 1024 * 1024 / 8, wantErr: false},
		{name: "3GB", size: "3GB", want: 3 * 1024 * 1024 * 1024, wantErr: false},
		{name: "20gB", size: "20gB", want: 20 * 1024 * 1024 * 1024, wantErr: false},
		{name: "24g", size: "24g", want: 24 * 1024 * 1024 * 1024, wantErr: false},
		{name: "48gb", size: "48gb", want: 48 * 1024 * 1024 * 1024 / 8, wantErr: false},
		{name: "64x", size: "64x", want: 0, wantErr: true},
		{name: "''", size: "", want: 0, wantErr: true},
		{name: "M", size: "M", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseByteSize(tt.size)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseByteSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("parseByteSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestParseTimeDuration$
func TestParseTimeDuration(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    time.Duration
		wantErr bool
	}{
		{name: "12s", s: "12s", want: 12 * time.Second, wantErr: false},
		{name: "3m", s: "3m", want: 3 * time.Minute, wantErr: false},
		{name: "24h", s: "24h", want: 24 * time.Hour, wantErr: false},
		{name: "24h50m12s", s: "24h50m12s", want: 24*time.Hour + 50*time.Minute + 12*time.Second, wantErr: false},
		{name: "7d", s: "7d", want: 7 * 24 * time.Hour, wantErr: false},
		{name: "90D", s: "90D", want: 90 * 24 * time.Hour, wantErr: false},
		{name: "''", s: "", want: 0, wantErr: true},
		{name: "14", s: "14", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTimeDuration(tt.s)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("parseTimeDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
