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

package size

import (
	"errors"
	"strconv"
	"strings"
)

const (
	B ByteSize = 1 << (10 * iota)
	KB
	MB
	GB
)

// ByteSize is size of byte.
type ByteSize = uint64

// parseByteSize parse size with given unit information.
func parseByteSize(size string, unit string, unitSize ByteSize, bitUnit bool) (ByteSize, error) {
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

// ParseByteSize parses byte size in string.
// You should add unit in your size string, like 4MB, 512K, 64.
// The unit will be byte if size string is just a number.
// General units is GB, G, MB, M, KB, K, B and you can see all of them is byte unit.
// If your size string is like 64kb, the result parsed will be 8KB (64kb = 8KB).
func ParseByteSize(size string) (ByteSize, error) {
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
		return parseByteSize(size, "G", GB, bitUnit)
	}

	if strings.HasSuffix(size, "M") {
		return parseByteSize(size, "M", MB, bitUnit)
	}

	if strings.HasSuffix(size, "K") {
		return parseByteSize(size, "K", KB, bitUnit)
	}

	return parseByteSize(size, "", B, bitUnit)
}
