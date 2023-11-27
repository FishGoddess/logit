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
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

const (
	backupSeparator = "."
)

// backup is the backup of file.
type backup struct {
	path string
	t    time.Time
}

// before returns if b.t is earlier than t.
func (b backup) before(t time.Time) bool {
	return b.t.Before(t)
}

// sortBackups sorts backups.
func sortBackups(backups []backup) {
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].before(backups[j].t)
	})
}

// backupPrefixAndExt returns the prefix and ext of path in backup form.
func backupPrefixAndExt(path string) (string, string) {
	ext := filepath.Ext(path)
	prefix := path[:len(path)-len(ext)] + backupSeparator
	return prefix, ext
}

// backupPath returns the backup path of path with time format.
func backupPath(path string, timeFormat string) string {
	now := defaults.CurrentTime().In(defaults.TimeLocation)
	name, ext := backupPrefixAndExt(path)

	if strings.ToLower(timeFormat) != "" {
		return name + now.Format(timeFormat) + ext
	}

	return name + strconv.FormatInt(now.Unix(), 10) + ext
}

// parseBackupTime parses backup time from filename and given time format.
func parseBackupTime(filename string, prefix string, ext string, timeFormat string) (time.Time, error) {
	ts := filename[len(prefix) : len(filename)-len(ext)]

	if strings.ToLower(timeFormat) != "" {
		return time.ParseInLocation(timeFormat, ts, defaults.TimeLocation)
	}

	seconds, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(seconds, 0).In(defaults.TimeLocation), nil
}
