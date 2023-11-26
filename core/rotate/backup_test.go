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
	"testing"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

// go test -v -cover -run=^TestBackupBefore$
func TestBackupBefore(t *testing.T) {
	b := backup{t: time.Unix(2, 0)}
	if b.before(time.Unix(1, 0)) {
		t.Errorf("b.before(time.Unix(1, 0)) returns false")
	}

	if b.before(time.Unix(2, 0)) {
		t.Errorf("b.before(time.Unix(2, 0)) returns false")
	}
}

// go test -v -cover -run=^TestSortBackups$
func TestSortBackups(t *testing.T) {
	backups := []backup{
		{t: time.Unix(2, 0)}, {t: time.Unix(1, 0)}, {t: time.Unix(4, 0)}, {t: time.Unix(0, 0)}, {t: time.Unix(3, 0)},
	}

	sortBackups(backups)
	for i, backup := range backups {
		if backup.t.Unix() != int64(i) {
			t.Errorf("backup.t.Unix() %d != int64(i) %d", backup.t.Unix(), int64(i))
		}
	}
}

// go test -v -cover -run=^TestBackupPrefixAndExt$
func TestBackupPrefixAndExt(t *testing.T) {
	prefix, ext := backupPrefixAndExt("test.log")
	if prefix != "test"+backupSeparator {
		t.Errorf("prefix %s != 'test'+backupSeparator %s", prefix, "test"+backupSeparator)
	}

	if ext != ".log" {
		t.Errorf("ext %s != '.log'", ext)
	}
}

// go test -v -cover -run=^TestBackupPath$
func TestBackupPath(t *testing.T) {
	defaults.TimeLocation = time.UTC

	defaults.CurrentTime = func() time.Time {
		return time.Unix(1, 0)
	}

	path := backupPath("test.log", "20060102150405")

	if path != "test.19700101000001.log" {
		t.Errorf("path %s != 'test.19700101080001.log'", path)
	}
}

// go test -v -cover -run=^TestParseBackupTime$
func TestParseBackupTime(t *testing.T) {
	defaults.TimeLocation = time.UTC

	filename := "test.19700101000001.log"
	prefix := "test."
	ext := ".log"
	timeFormat := "20060102150405"

	backupTime, err := parseBackupTime(filename, prefix, ext, timeFormat)
	if err != nil {
		t.Error(err)
	}

	if backupTime.Unix() != 1 {
		t.Errorf("backupTime.Unix() %d != 1", backupTime.Unix())
	}
}
