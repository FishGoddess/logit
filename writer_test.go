// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/12/01 00:19:00

package logit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// testChecker is for testing.
type testChecker struct{}

// Check returns if n is odd.
func (tc *testChecker) Check(fw *FileWriter, n int) bool {
	return n&1 == 0
}

// prepareTestDirAndName prepares directory and name of file for testing.
func prepareTestDirAndName(t *testing.T) (testDir string, name string) {

	testDir = filepath.Join(os.TempDir(), "TestFileWriter")
	err := os.Mkdir(testDir, 0644)
	if err != nil {
		t.Fatal(err)
	}
	return testDir, filepath.Join(testDir, time.Now().Format("20060102150405.log"))
}

// lengthOf returns the count of names exclude "." and "..".
func lengthOf(names []os.FileInfo) int {
	count := 0
	for _, info := range names {
		if info.Name() != "." && info.Name() != ".." {
			count++
		}
	}
	return count
}

// checkNamesLength compares the number of file in testDir with except and
// fatal if not match.
func checkNamesLength(t *testing.T, testDir string, except int) {

	names, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Fatal(err)
	}

	length := lengthOf(names)
	if length != except {
		t.Fatalf("length %d of names is wrong", length)
	}
}

// checkFileContent compares the content of file with except and
// fatal if not match.
func checkFileContent(t *testing.T, file string, except string) {

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}

	if string(fileBytes) != except {
		t.Fatalf("content %s of file is wrong", string(fileBytes))
	}
}

// go test -v -cover -run=^TestFileWriter$
func TestFileWriter(t *testing.T) {

	testDir, name := prepareTestDirAndName(t)
	t.Log("name of file is", name)

	writer, err := NewFileWriter(name, &testChecker{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := writer.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	writer.Write([]byte("abc"))
	writer.Write([]byte("d"))
	writer.Write([]byte("efg"))
	checkNamesLength(t, testDir, 1)
	checkFileContent(t, name, "abcdefg")

	writer.Write([]byte("1234"))
	checkNamesLength(t, testDir, 2)
	checkFileContent(t, name, "1234")
	checkFileContent(t, writer.nextName(), "abcdefg")

	if _, err = writer.Write([]byte("???")); err != nil {
		t.Fatal(err)
	}
}
