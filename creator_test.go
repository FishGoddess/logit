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

package logit

import (
	"errors"
	"testing"
)

type testLoggerCreator struct{}

func (tlm *testLoggerCreator) CreateLogger(params ...interface{}) (*Logger, error) {
	if len(params) < 1 {
		return nil, errors.New("testLoggerCreator: len(params) < 1")
	}

	if params[0].(string) == "error" {
		return nil, errors.New("testLoggerCreator: params[0] isn't a string")
	}

	return nil, nil
}

// go test -v -cover -run=^TestRegisterLoggerCreator$
func TestRegisterLoggerCreator(t *testing.T) {
	name := "testLoggerCreator"

	if _, ok := loggerCreators[name]; ok {
		t.Errorf("logger creator %s has been registered", name)
	}

	creator := new(testLoggerCreator)
	err := RegisterLoggerCreator(name, creator)
	if err != nil {
		t.Error(err)
	}

	loggerCreator, ok := loggerCreators[name]
	if !ok {
		t.Errorf("logger creator %s not found", name)
	}

	if loggerCreator != creator {
		t.Errorf("logger creator %p != make %p", loggerCreator, creator)
	}
}

// go test -v -cover -run=^TestNewLoggerFromCreator$
func TestNewLoggerFromCreator(t *testing.T) {
	name := "testLoggerCreator"
	creator := new(testLoggerCreator)
	loggerCreators[name] = creator

	_, err := creator.CreateLogger("")
	if err != nil {
		t.Error(err)
	}

	_, err = NewLoggerFromCreator(name, "")
	if err != nil {
		t.Error(err)
	}

	_, err = creator.CreateLogger("error")
	if err == nil {
		t.Error("createLogger should return an error")
	}

	_, err = NewLoggerFromCreator(name, "error")
	if err == nil {
		t.Error("NewLoggerFromCreator should return an error")
	}
}
