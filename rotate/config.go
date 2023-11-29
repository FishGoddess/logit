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

package rotate

import (
	"time"
)

const (
	MB  = 1024 * 1024
	Day = 24 * time.Hour
)

// config stores some fields of file.
type config struct {
	// timeFormat is the time format of backup path.
	timeFormat string

	// maxSize is the max size of file.
	// If size of data in one write is bigger than maxSize, then file will rotate and write it,
	// which means file and its backup may bigger than maxSize in size.
	maxSize uint64

	// maxAge is the time that backup will live.
	// All backups reach maxAge will be removed automatically.
	maxAge time.Duration

	// maxBackups is the max count of backups.
	maxBackups int
}

// newDefaultConfig returns a default config.
func newDefaultConfig() config {
	return config{
		timeFormat: "20060102150405",
		maxSize:    256 * MB,
		maxAge:     90 * Day,
		maxBackups: 100,
	}
}
