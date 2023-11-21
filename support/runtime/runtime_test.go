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

package runtime

import (
	"os"
	"runtime"
	"testing"
)

// go test -v -cover -run=^TestPID$
func TestPID(t *testing.T) {
	pid := PID()
	osPID := os.Getpid()

	if pid != osPID {
		t.Errorf("pid %d is wrong with %d", pid, osPID)
	}
}

// go test -v -cover -run=^TestCaller$
func TestCaller(t *testing.T) {
	pc, correctFile, correctLine, ok := runtime.Caller(0)
	if !ok {
		t.Errorf("runtime.Caller failed")
	}

	_, file, line, function := Caller(1)
	if file != correctFile {
		t.Errorf("Caller returns wrong file %s", file)
	}

	if line != correctLine+5 {
		t.Errorf("Caller returns wrong line %d", line)
	}

	fc := runtime.FuncForPC(pc)
	if fc == nil {
		t.Error("runtime.FuncForPC(pc) == nil")
	}

	if function != fc.Name() {
		t.Errorf("Caller returns wrong function %s", function)
	}
}
