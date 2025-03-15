// Copyright 2025 FishGoddess. All Rights Reserved.
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

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBackupBefore$
func TestBackupBefore(t *testing.T) {
	b := backup{t: time.Unix(2, 0)}

	if b.before(time.Unix(1, 0)) {
		t.Fatalf("b.before(time.Unix(1, 0)) returns false")
	}

	if b.before(time.Unix(2, 0)) {
		t.Fatalf("b.before(time.Unix(2, 0)) returns false")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSortBackups$
func TestSortBackups(t *testing.T) {
	backups := []backup{
		{t: time.Unix(2, 0)},
		{t: time.Unix(1, 0)},
		{t: time.Unix(4, 0)},
		{t: time.Unix(0, 0)},
		{t: time.Unix(3, 0)},
	}

	sortBackups(backups)

	for i, backup := range backups {
		if backup.t.Unix() != int64(i) {
			t.Fatalf("backup.t.Unix() %d != int64(i) %d", backup.t.Unix(), int64(i))
		}
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBackupPrefixAndExt$
func TestBackupPrefixAndExt(t *testing.T) {
	prefix, ext := backupPrefixAndExt("test.log")

	want := "test" + backupSeparator
	if prefix != want {
		t.Fatalf("prefix %s != want %s", prefix, want)
	}

	want = ".log"
	if ext != want {
		t.Fatalf("ext %s != want %s", ext, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestBackupPath$
func TestBackupPath(t *testing.T) {
	defaults.CurrentTime = func() time.Time {
		return time.Unix(1, 0).In(time.UTC)
	}

	path := backupPath("test.log", "20060102150405")
	want := "test.19700101000001.log"
	if path != want {
		t.Fatalf("path %s != want %s", path, want)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestParseBackupTime$
func TestParseBackupTime(t *testing.T) {
	defaults.CurrentTime = func() time.Time {
		return time.Now().In(time.UTC)
	}

	filename := "test.19700101000001.log"
	prefix := "test."
	ext := ".log"
	timeFormat := "20060102150405"

	backupTime, err := parseBackupTime(filename, prefix, ext, timeFormat)
	if err != nil {
		t.Fatal(err)
	}

	if backupTime.Unix() != 1 {
		t.Fatalf("backupTime.Unix() %d != 1", backupTime.Unix())
	}
}
