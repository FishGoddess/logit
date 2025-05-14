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
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/FishGoddess/logit/defaults"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNew$
func TestNew(t *testing.T) {
	path := filepath.Join(t.TempDir(), t.Name())

	f, err := New(path)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	data := []byte("水不要鱼")
	n, err := f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatalf("n %d != len(data) %d", n, len(data))
	}

	read, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	if string(read) != string(data) {
		t.Fatalf("string(read) %s != string(data) %s", read, data)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewExisting$
func TestNewExisting(t *testing.T) {
	path := filepath.Join(t.TempDir(), t.Name())

	err := os.WriteFile(path, []byte("水不要鱼"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	f, err := New(path)
	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	data := []byte("FishGoddess")
	n, err := f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatalf("n %d != len(data) %d", n, len(data))
	}

	read, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	want := "水不要鱼FishGoddess"
	if string(read) != want {
		t.Fatalf("string(read) %s != want %s", read, want)
	}
}

func countFiles(dir string) int {
	files, _ := os.ReadDir(dir)
	return len(files)
}

// go test -v -cover -count=1 -run=^TestFileRotate$
func TestFileRotate(t *testing.T) {
	second := int64(0)
	defaults.CurrentTime = func() time.Time {
		second++
		return time.Unix(second, 0)
	}

	dir := filepath.Join(t.TempDir(), t.Name())
	if err := os.RemoveAll(dir); err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(dir, "test.log")

	f, err := New(path)
	if err != nil {
		t.Fatal(err)
	}

	f.conf.maxSize = 4
	defer f.Close()

	data := []byte("test")
	n, err := f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatalf("n %d != len(data) %d", n, len(data))
	}

	read, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if string(read) != string(data) {
		t.Fatalf("string(read) %s != string(data) %s", read, data)
	}

	count := countFiles(dir)
	if count != 1 {
		t.Fatalf("count %d != 1", count)
	}

	data = []byte("burst")
	n, err = f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatalf("n %d != len(data) %d", n, len(data))
	}

	data = []byte("!!!")
	n, err = f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(data) {
		t.Fatalf("n %d != len(data) %d", n, len(data))
	}

	read, err = os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if string(read) != "!!!" {
		t.Fatalf("string(read) %s != '!!!'", read)
	}

	count = countFiles(dir)
	if count != 3 {
		t.Fatalf("count %d != 3", count)
	}

	second = 3
	defaults.CurrentTime = func() time.Time {
		second--
		return time.Unix(second, 0)
	}

	var bs []byte
	for second > 1 {
		backup := backupPath(path, f.conf.timeFormat)
		if bs, err = os.ReadFile(backup); err != nil {
			t.Fatal(err)
		}

		read = append(read, bs...)
	}

	if string(read) != "!!!bursttest" {
		t.Fatalf("string(read) %s != '!!!bursttest'", read)
	}
}
