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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/FishGoddess/logit/support/global"
	"github.com/FishGoddess/logit/support/size"
)

var (
	// CurrentTime returns current time in time.Time.
	CurrentTime = global.CurrentTime
)

// File is a file which supports rotating automatically.
// It has max size and file will rotate if size exceeds max size.
// It has max age and max backups, so rotated files will be controlled in quantity which is beneficial to space.
type File struct {
	// config stores all configurations of file.
	config

	// path is the path of file.
	path string

	// size is the current size of writing in file.
	size size.ByteSize

	file *os.File
	ch   chan struct{}
	lock sync.Mutex
}

// New returns a new file.
func New(path string, opts ...Option) (*File, error) {
	f := newFile(path, opts)

	if err := f.mkdir(); err != nil {
		return nil, err
	}

	if err := f.openNewFile(); err != nil {
		return nil, err
	}

	go f.runCleanTask()
	return f, nil
}

// newFile creates a new file.
func newFile(path string, opts []Option) *File {
	c := newDefaultConfig()

	for _, opt := range opts {
		opt.Apply(&c)
	}

	return &File{config: c, path: path, ch: make(chan struct{}, 1)}
}

func (f *File) mkdir() error {
	return os.MkdirAll(filepath.Dir(f.path), f.dirMode)
}

func (f *File) open() (*os.File, error) {
	return global.OpenFile(f.path, f.mode)
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

func (f *File) removeStaleBackups(backups []backup) {
	staleBackups := make(map[string]struct{}, 16)

	if f.maxBackups > 0 {
		for i := 0; i < len(backups)-f.maxBackups; i++ {
			staleBackups[backups[i].path] = struct{}{}
		}
	}

	if f.maxAge > 0 {
		deadline := CurrentTime().Add(-f.maxAge)

		for _, backup := range backups {
			if !backup.before(deadline) {
				break
			}

			staleBackups[backup.path] = struct{}{}
		}
	}

	for backup := range staleBackups {
		os.Remove(backup)
	}
}

func (f *File) clean() {
	backups, err := f.listBackups()
	if err != nil {
		return
	}

	f.removeStaleBackups(backups)
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

func (f *File) nextBackupPath() (string, error) {
	backupPath := backupPath(f.path, f.timeFormat)

	_, err := os.Stat(backupPath)
	if os.IsNotExist(err) {
		return backupPath, nil
	}

	if err != nil {
		return "", err
	}

	// Backup path conflict...
	return "", fmt.Errorf("logit: extension.file wants a backup path %s but conflict", backupPath)
}

func (f *File) closeOldFile() (err error) {
	backupPath, err := f.nextBackupPath()
	if err != nil {
		return err
	}

	fileClosed := false

	defer func() {
		if err != nil && fileClosed {
			f.openNewFile() // Reopen closed file.
		}
	}()

	if err = f.file.Close(); err != nil {
		return err
	}

	fileClosed = true
	if err = os.Rename(f.path, backupPath); err != nil {
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

// Write writes len(p) bytes from p to the underlying data stream.
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

// Sync syncs data to the underlying io device.
func (f *File) Sync() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	return f.file.Sync()
}

// Close closes file and returns an error if failed.
func (f *File) Close() error {
	f.lock.Lock()
	defer f.lock.Unlock()

	if err := f.file.Sync(); err != nil {
		return err
	}

	close(f.ch)
	return f.file.Close()
}
