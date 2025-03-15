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
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

const (
	backupSeparator = "."
)

type backup struct {
	path string
	t    time.Time
}

func (b backup) before(t time.Time) bool {
	return b.t.Before(t)
}

func sortBackups(backups []backup) {
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].before(backups[j].t)
	})
}

func backupPrefixAndExt(path string) (prefix string, ext string) {
	ext = filepath.Ext(path)
	prefix = path[:len(path)-len(ext)] + backupSeparator

	return prefix, ext
}

func backupPath(path string, timeFormat string) string {
	now := defaults.CurrentTime()
	name, ext := backupPrefixAndExt(path)

	if timeFormat != "" {
		return name + now.Format(timeFormat) + ext
	}

	return name + strconv.FormatInt(now.Unix(), 10) + ext
}

func parseBackupTime(filename string, prefix string, ext string, timeFormat string) (time.Time, error) {
	ts := filename[len(prefix) : len(filename)-len(ext)]

	if timeFormat != "" {
		return time.Parse(timeFormat, ts)
	}

	seconds, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(seconds, 0), nil
}
