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
// Created at 2020/12/01 00:19:24

package logit

import (
	"testing"
	"time"
)

// go test -v -cover -run=^TestTimeChecker$
func TestTimeChecker(t *testing.T) {

	testDir, name := prepareTestDirAndName(t)
	t.Log("name of file is", name)

	checker := NewTimeChecker(time.Second)
	writer, err := NewFileWriter(name, checker)
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()

	for i := 1; i <= 3; i++ {
		time.Sleep(900 * time.Millisecond)
		writer.Write(make([]byte, 1))
		t.Logf("duration is %s, lastTime is %s", checker.duration, checker.lastTime)
		if i == 3 {
			checkNamesLength(t, testDir, 2)
			break
		}
		checkNamesLength(t, testDir, i)
	}
}

// go test -v -cover -run=^TestSizeChecker$
func TestSizeChecker(t *testing.T) {

	testDir, name := prepareTestDirAndName(t)
	t.Log("name of file is", name)

	checker := NewSizeChecker(1 * MB)
	writer, err := NewFileWriter(name, checker)
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()

	for i := 1; i <= 3; i++ {
		writer.Write(make([]byte, 1000*KB))
		t.Logf("limitedSize is %d, currentSize is %d", checker.limitedSize, checker.currentSize)
		checkNamesLength(t, testDir, i)
	}
}

// go test -v -cover -run=^TestCountChecker$
func TestCountChecker(t *testing.T) {

	testDir, name := prepareTestDirAndName(t)
	t.Log("name of file is", name)

	checker := NewCountChecker(3)
	writer, err := NewFileWriter(name, checker)
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()

	for i := 1; i <= 3; i++ {
		for j := 0; j < 2; j++ {
			writer.Write(make([]byte, 1))
		}
		t.Logf("limitedCount is %d, currentCount is %d", checker.limitedCount, checker.currentCount)
		checkNamesLength(t, testDir, i)
	}
}
