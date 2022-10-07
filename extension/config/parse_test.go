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

package config

import (
	"testing"
	"time"

	"github.com/FishGoddess/logit/support/global"
)

// go test -v -cover -run=^TestParseTimeFormat$
func TestParseTimeFormat(t *testing.T) {
	cases := map[string]string{
		"unix":           global.UnixTimeFormat,
		"":               "",
		"123":            "123",
		"20060102150405": "20060102150405",
	}

	for input, expect := range cases {
		output := parseTimeFormat(input)
		if output != expect {
			t.Errorf("output %s != expect %s", output, expect)
		}
	}
}

// go test -v -cover -run=^parseTimeDuration$
func TestParseTimeDuration(t *testing.T) {
	cases := map[string]time.Duration{
		"1s":     time.Second,
		"1330s":  1330 * time.Second,
		"45m20s": 45*time.Minute + 20*time.Second,
		"8h20m":  8*time.Hour + 20*time.Minute,
		"14h50s": 14*time.Hour + 50*time.Second,
		"7d":     7 * 24 * time.Hour,
		"14D":    14 * 24 * time.Hour,
	}

	for input, expect := range cases {
		output, err := parseTimeDuration(input)
		if err != nil {
			t.Error(err)
		}

		if output != expect {
			t.Errorf("output %s != expect %s", output, expect)
		}
	}
}
