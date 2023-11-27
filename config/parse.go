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
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

// parseByteSize parse size with given unit information.
func parseByteSizeWithUnit(size string, unit string, unitSize uint64, bitUnit bool) (uint64, error) {
	size = strings.TrimSuffix(size, unit)

	n, err := strconv.ParseUint(size, 10, 64)
	if err != nil {
		return 0, err
	}

	if bitUnit {
		return n / 8 * unitSize, nil
	}

	return n * unitSize, nil
}

// parseByteSize parses byte size in string.
// You should add unit in your size string, like "4MB", "512K", "64".
// The unit will be byte if size string is just a number.
// General units is GB, G, MB, M, KB, K, B and you can see all of them is byte unit.
// If your size string is like "64kb", the result parsed will be 8KB (64kb = 8KB).
func parseByteSize(size string) (uint64, error) {
	size = strings.TrimSpace(size)
	if size == "" {
		return 0, errors.New("logit: parse byte size from an empty string")
	}

	bitUnit := false
	if strings.HasSuffix(size, "b") {
		bitUnit = true
		size = strings.TrimSuffix(size, "b")
	} else {
		size = strings.TrimSuffix(size, "B")
	}

	size = strings.ToUpper(size)
	if strings.HasSuffix(size, "G") {
		return parseByteSizeWithUnit(size, "G", GB, bitUnit)
	}

	if strings.HasSuffix(size, "M") {
		return parseByteSizeWithUnit(size, "M", MB, bitUnit)
	}

	if strings.HasSuffix(size, "K") {
		return parseByteSizeWithUnit(size, "K", KB, bitUnit)
	}

	return parseByteSizeWithUnit(size, "", B, bitUnit)
}

func parseTimeDuration(s string) (time.Duration, error) {
	if strings.HasSuffix(s, "d") || strings.HasSuffix(s, "D") {
		s = strings.TrimSuffix(s, "d")
		s = strings.TrimSuffix(s, "D")

		days, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, err
		}

		return time.Duration(days) * 24 * time.Hour, nil
	}

	return time.ParseDuration(s)
}
