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

package main

import (
	"context"
	"errors"

	"github.com/go-logit/logit"
)

type testLoggerMaker struct{}

func (tlm *testLoggerMaker) MakeLogger(ctx context.Context, params ...interface{}) (*logit.Logger, error) {
	if len(params) < 1 {
		return nil, errors.New("testLoggerMaker: len(params) < 1")
	}

	if params[0].(string) == "error" {
		return nil, errors.New("testLoggerMaker: params[0] isn't a string")
	}

	// Customize your creation of logger here.
	return logit.NewLogger(), nil
}

func main() {
	makeName := "testLoggerMaker"

	// RegisterLoggerMaker registers maker to logit with given name.
	err := logit.RegisterLoggerMaker(makeName, new(testLoggerMaker))
	if err != nil {
		panic(err)
	}

	// NewLoggerFromMaker creates logger from maker with given params.
	// Panic will be invoked if params is "error" because MakeLogger in testLoggerMaker has this logic.
	logger, err := logit.NewLoggerFromMaker(context.Background(), makeName, "xxx")
	if err != nil {
		panic(err)
	}

	logger.Info("I am made from logger maker!").End()
}
