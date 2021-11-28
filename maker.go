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
// Created at 2021/10/31 23:31:19

package logit

import (
	"context"
	"errors"
	"sync"
)

var (
	loggerMakers     = make(map[string]LoggerMaker, 4) // loggerMakers stores all registered logger makers.
	loggerMakersLock sync.RWMutex                      // loggerMakersLock is for concurrency-safe when using loggerMakers.
)

// LoggerMaker is for making a new logger.
type LoggerMaker interface {
	// MakeLogger makes a new logger using params and returns an error if something wrong happens.
	MakeLogger(ctx context.Context, params ...interface{}) (*Logger, error)
}

// RegisterLoggerMaker registers logger maker with name.
// Returns an error if failed.
func RegisterLoggerMaker(makerName string, maker LoggerMaker) error {
	loggerMakersLock.Lock()
	defer loggerMakersLock.Unlock()

	if maker == nil {
		return errors.New("logit: logger maker " + makerName + " is nil")
	}

	if _, ok := loggerMakers[makerName]; ok {
		return errors.New("logit: logger maker " + makerName + " has been registered")
	}

	loggerMakers[makerName] = maker
	return nil
}

// NewLoggerFromMaker creates logger from logger maker with params.
// Returns an error if failed.
func NewLoggerFromMaker(ctx context.Context, makerName string, params ...interface{}) (*Logger, error) {
	loggerMakersLock.RLock()
	defer loggerMakersLock.RUnlock()

	maker, ok := loggerMakers[makerName]
	if !ok {
		return nil, errors.New("logit: logger maker " + makerName + " not found")
	}

	return maker.MakeLogger(ctx, params...)
}
