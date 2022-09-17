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

package file

import (
	"os"
	"time"

	"github.com/go-logit/logit/support/size"
)

// config stores some fields of file.
type config struct {
	// mode is the permission mode of file.
	mode os.FileMode

	// dirMode is the permission mode of file directory.
	dirMode os.FileMode

	// timeFormat is the time format of backup path.
	timeFormat string

	// maxSize is the max size of file.
	// If size of data in one write is bigger than maxSize, then file will rotate and write it,
	// which means file and its backup may bigger than maxSize in size.
	maxSize size.ByteSize

	// maxAge is the time that backup will live.
	// All backups reach maxAge will be removed automatically.
	maxAge time.Duration

	// maxBackups is the max count of backups.
	maxBackups int
}

// newDefaultConfig returns a default config.
func newDefaultConfig() config {
	return config{
		mode:       0644,
		dirMode:    0755,
		timeFormat: "20060102150405",
		maxSize:    256 * size.MB,
		maxAge:     0,
		maxBackups: 0,
	}
}
