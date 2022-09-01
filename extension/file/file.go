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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-logit/logit/support/global"
	"github.com/go-logit/logit/support/size"
)

var (
	now = global.CurrentTime
)

type File struct {
	config

	path string
	size size.ByteSize

	file *os.File
	ch   chan struct{}
	lock sync.Mutex
}

func New(path string) (*File, error) {
	f := &File{
		config: newDefaultConfig(),
		path:   path,
		ch:     make(chan struct{}, 1),
	}

	if err := f.mkdir(); err != nil {
		return nil, err
	}

	if err := f.openNewFile(); err != nil {
		return nil, err
	}

	go f.runCleanTask()
	return f, nil
}

func (f *File) mkdir() error {
	return os.MkdirAll(filepath.Dir(f.path), f.dirMode)
}

func (f *File) open() (*os.File, error) {
	return os.OpenFile(f.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, f.mode)
}

func (f *File) listBackups() ([]backup, error) {
	dir := filepath.Dir(f.path)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	baseName := filepath.Base(f.path)
	prefix, ext := backupPrefixAndExt(baseName)

	var backups []backup
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if filename == baseName {
			continue
		}

		notBackup := !strings.HasPrefix(filename, prefix) || !strings.HasSuffix(filename, ext)
		if notBackup {
			continue
		}

		t, err := parseBackupTime(filename, prefix, ext, f.timeFormat)
		if err != nil {
			continue
		}

		backups = append(backups, backup{
			path: filepath.Join(dir, filename),
			t:    t,
		})
	}

	sortBackups(backups)
	return backups, nil
}

func (f *File) filterStaleBackups(backups []backup) []string {
	staleBackups := make(map[string]struct{})

	maxBackups := f.maxBackups
	if maxBackups > 0 && uint(len(backups)) > maxBackups {
		for i := uint(0); i < uint(len(backups))-maxBackups; i++ {
			staleBackups[backups[i].path] = struct{}{}
		}
	}

	maxAge := f.maxAge
	if maxAge > 0 {
		deadline := now().Add(-maxAge)

		for _, backup := range backups {
			if !backup.Before(deadline) {
				break
			}

			staleBackups[backup.path] = struct{}{}
		}
	}

	var paths []string
	for k := range staleBackups {
		paths = append(paths, k)
	}

	return paths
}

func (f *File) removeBackups(backups []string) {
	for _, backup := range backups {
		os.Remove(backup)
	}
}

func (f *File) clean() {
	backups, err := f.listBackups()
	if err != nil {
		return
	}

	paths := f.filterStaleBackups(backups)
	f.removeBackups(paths)
}

func (f *File) runCleanTask() {
	for range f.ch {
		f.clean()
	}
}

func (f *File) triggerCleanTask() {
	select {
	case f.ch <- struct{}{}:
	default:
	}
}

func (f *File) openNewFile() error {
	file, err := f.open()
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}

	f.file = file
	f.size = size.ByteSize(info.Size())
	return nil
}

func (f *File) closeOldFile() error {
	if err := f.file.Close(); err != nil {
		return err
	}

	backupPath := backupPath(f.path, f.timeFormat)
	if err := os.Rename(f.path, backupPath); err != nil {
		return err
	}

	return nil
}

func (f *File) rotate() error {
	if err := f.closeOldFile(); err != nil {
		return err
	}

	if err := f.openNewFile(); err != nil {
		return err
	}

	f.triggerCleanTask()
	return nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	writeSize := size.ByteSize(len(p))
	if f.size+writeSize > f.maxSize {
		f.rotate() // Ignore rotating error so this p won't be discarded.
	}

	n, err = f.file.Write(p)
	f.size += size.ByteSize(n)

	return n, err
}

func (f *File) Sync() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	return f.file.Sync()
}

func (f *File) Close() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	if err := f.file.Sync(); err != nil {
		return err
	}

	close(f.ch)
	return f.file.Close()
}
