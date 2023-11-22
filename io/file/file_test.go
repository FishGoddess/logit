package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/FishGoddess/logit/defaults"
	"github.com/FishGoddess/logit/io/size"
)

// go test -v -cover -run=^TestNew$
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
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("n %d != len(data) %d", n, len(data))
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	if string(read) != string(data) {
		t.Errorf("string(read) %s != string(data) %s", read, data)
	}
}

// go test -v -cover -run=^TestNewExisting$
func TestNewExisting(t *testing.T) {
	path := filepath.Join(t.TempDir(), t.Name())

	err := ioutil.WriteFile(path, []byte("水不要鱼"), 0644)
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
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("n %d != len(data) %d", n, len(data))
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	want := "水不要鱼FishGoddess"
	if string(read) != want {
		t.Errorf("string(read) %s != want %s", read, want)
	}
}

func countFiles(dir string) int {
	files, _ := os.ReadDir(dir)
	return len(files)
}

// go test -v -cover -run=^TestFileRotate$
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

	f.maxSize = 4 * size.B
	defer f.Close()

	data := []byte("test")
	n, err := f.Write(data)
	if err != nil {
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("n %d != len(data) %d", n, len(data))
	}

	read, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	if string(read) != string(data) {
		t.Errorf("string(read) %s != string(data) %s", read, data)
	}

	count := countFiles(dir)
	if count != 1 {
		t.Errorf("count %d != 1", count)
	}

	data = []byte("burst")
	n, err = f.Write(data)
	if err != nil {
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("n %d != len(data) %d", n, len(data))
	}

	data = []byte("!!!")
	n, err = f.Write(data)
	if err != nil {
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("n %d != len(data) %d", n, len(data))
	}

	read, err = ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	if string(read) != "!!!" {
		t.Errorf("string(read) %s != '!!!'", read)
	}

	count = countFiles(dir)
	if count != 3 {
		t.Errorf("count %d != 3", count)
	}

	second = 3
	defaults.CurrentTime = func() time.Time {
		second--
		return time.Unix(second, 0)
	}

	for second > 1 {
		var bs []byte
		bs, err = ioutil.ReadFile(backupPath(path, f.timeFormat))
		if err != nil {
			t.Error(err)
		}

		read = append(read, bs...)
	}

	if string(read) != "!!!bursttest" {
		t.Errorf("string(read) %s != '!!!bursttest'", read)
	}
}
