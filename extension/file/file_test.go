package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {
	dir, _ := ioutil.TempDir("", "TestNewFile")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "New")

	f, err := New(path)
	if err != nil {
		t.Fatalf("NewWriter error: %v", err)
	}

	data := []byte("水不要鱼")
	n, err := f.Write(data)
	if err != nil {
		t.Error(err)
	}

	if n != len(data) {
		t.Errorf("Write() = %v; want %v", n, len(data))
	}

	out, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	if string(out) != string(data) {
		t.Errorf("ReadFile = %q; want %q", out, data)
	}
}

func TestExistingFile(t *testing.T) {
	dir, _ := ioutil.TempDir("", "TestNewFile")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "New")
	ioutil.WriteFile(path, []byte("水不要鱼，"), 0644)

	f, err := New(path)
	if err != nil {
		t.Fatalf("NewWriter error: %v", err)
	}

	data := []byte("FishGoddess")
	n, err := f.Write(data)
	if err != nil {
		t.Errorf("Write error: %v", err)
	}

	if n != len(data) {
		t.Errorf("Write() = %v; want %v", n, len(data))
	}

	out, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	want := "水不要鱼FishGoddess"
	if string(out) != want {
		t.Errorf("ReadFile = %q; want %q", out, want)
	}
}

func countFiles(dir string) int {
	files, _ := ioutil.ReadDir(dir)
	return len(files)
}

func TestRotate(t *testing.T) {
	now = func() time.Time {
		return time.Unix(0, 0)
	}

	dir, _ := ioutil.TempDir("", "TestRotate")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "rock.log")

	f, err := New(path)
	if err != nil {
		t.Error(err)
	}

	defer f.Close()

	b := []byte("boo!")
	n, err := f.Write(b)
	if err != nil || n != len(b) {
		t.Errorf("Write(%q) = (%v, %v); want (%v, %v)", b, n, err, len(b), nil)
	}

	out, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	if string(out) != string(b) {
		t.Errorf("ReadFile = %q; want %q", out, b)
	}

	if got := countFiles(dir); got != 1 {
		t.Errorf("File count = %v; want 1", got)
	}

	b2 := []byte("foooooo!")
	n, err = f.Write(b2)
	if err != nil || n != len(b2) {
		t.Errorf("Write(%q) = (%v, %v); want (%v, %v)", b2, n, err, len(b2), nil)
	}

	out, err = ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	if string(out) != string(b2) {
		t.Errorf("ReadFile = %q; want %q", out, b2)
	}

	out, err = ioutil.ReadFile(f.backupPath(path))
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}

	if string(out) != string(b) {
		t.Errorf("ReadFile = %q; want %q", out, b)
	}

	if got := countFiles(dir); got != 2 {
		t.Errorf("File count = %v; want 2", got)
	}
}
