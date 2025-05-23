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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/FishGoddess/logit/defaults"
)

// File is a file which supports rotating automatically.
// It has max size and file will rotate if size exceeds max size.
// It has max age and max backups, so rotated files will be cleaned which is beneficial to space.
type File struct {
	conf *config

	file *os.File
	size uint64
	ch   chan struct{}

	lock sync.Mutex
}

// New returns a new rotate file.
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
	conf := newDefaultConfig(path)

	for _, opt := range opts {
		opt.applyTo(conf)
	}

	f := &File{
		conf: conf,
		ch:   make(chan struct{}, 1),
	}

	return f
}

func (f *File) mkdir() error {
	dir := filepath.Dir(f.conf.path)

	return defaults.OpenFileDir(dir, defaults.FileDirMode)
}

func (f *File) open() (*os.File, error) {
	return defaults.OpenFile(f.conf.path, defaults.FileMode)
}

func (f *File) listBackups() ([]backup, error) {
	dir := filepath.Dir(f.conf.path)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	baseName := filepath.Base(f.conf.path)
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

		t, err := parseBackupTime(filename, prefix, ext, f.conf.timeFormat)
		if err != nil {
			defaults.HandleError("rotate.parseBackupTime", err)
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

	if f.conf.maxBackups > 0 {
		exceeds := len(backups) - int(f.conf.maxBackups)
		for i := 0; i < exceeds; i++ {
			staleBackups[backups[i].path] = struct{}{}
		}
	}

	if f.conf.maxAge > 0 {
		deadline := defaults.CurrentTime().Add(-f.conf.maxAge)

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
	f.size = uint64(info.Size())
	return nil
}

func (f *File) nextBackupPath() (string, error) {
	backupPath := backupPath(f.conf.path, f.conf.timeFormat)

	_, err := os.Stat(backupPath)
	if os.IsNotExist(err) {
		return backupPath, nil
	}

	if err != nil {
		return "", err
	}

	// Backup path conflicted...
	err = fmt.Errorf("logit: rotate.file wants a backup path %s but conflicted", backupPath)
	return "", err
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
	err = os.Rename(f.conf.path, backupPath)
	return err
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

	writeSize := uint64(len(p))
	if f.size+writeSize > f.conf.maxSize {
		// Ignore rotating error so this p won't be discarded.
		if rotateErr := f.rotate(); rotateErr != nil {
			defaults.HandleError("rotate.File.rotate", rotateErr)
		}
	}

	n, err = f.file.Write(p)
	f.size += uint64(n)
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
