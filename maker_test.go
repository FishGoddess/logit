// Copyright 2021 Ye Zi Jie. All Rights Reserved.
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
// Created at 2021/11/01 00:17:08

package logit

import (
	"context"
	"errors"
	"testing"
)

type testLoggerMaker struct{}

func (tlm *testLoggerMaker) MakeLogger(ctx context.Context, params ...interface{}) (*Logger, error) {
	if len(params) < 1 {
		return nil, errors.New("testLoggerMaker: len(params) < 1")
	}

	if params[0].(string) == "error" {
		return nil, errors.New("testLoggerMaker: params[0] isn't a string")
	}

	return nil, nil
}

// go test -v -cover -run=^TestRegisterLoggerMaker$
func TestRegisterLoggerMaker(t *testing.T) {
	makerName := "testLoggerMaker"

	if _, ok := loggerMakers[makerName]; ok {
		t.Errorf("logger maker %s has been registered", makerName)
	}

	maker := new(testLoggerMaker)
	err := RegisterLoggerMaker(makerName, maker)
	if err != nil {
		t.Error(err)
	}

	loggerMaker, ok := loggerMakers[makerName]
	if !ok {
		t.Errorf("logger maker %s not found", makerName)
	}

	if loggerMaker != maker {
		t.Errorf("logger maker %p != make %p", loggerMaker, maker)
	}
}

// go test -v -cover -run=^TestNewLoggerFromMaker$
func TestNewLoggerFromMaker(t *testing.T) {
	makerName := "testLoggerMaker"
	maker := new(testLoggerMaker)
	loggerMakers[makerName] = maker

	_, err := maker.MakeLogger(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	_, err = NewLoggerFromMaker(context.Background(), makerName, "")
	if err != nil {
		t.Error(err)
	}

	_, err = maker.MakeLogger(context.Background(), "error")
	if err == nil {
		t.Error("makeLogger should return an error")
	}

	_, err = NewLoggerFromMaker(context.Background(), makerName, "error")
	if err == nil {
		t.Error("NewLoggerFromMaker should return an error")
	}
}
