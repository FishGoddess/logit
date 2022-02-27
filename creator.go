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
	"sync"
)

var (
	loggerCreators     = make(map[string]LoggerCreator, 4) // loggerCreators stores all registered logger creators.
	loggerCreatorsLock sync.RWMutex                        // loggerCreatorsLock is for concurrency-safe when using loggerCreators.
)

// LoggerCreator is for creating a new logger.
type LoggerCreator interface {
	// CreateLogger creates a new logger using params and returns an error if failed.
	CreateLogger(params ...interface{}) (*Logger, error)
}

// RegisterLoggerCreator registers logger creator with name and returns an error if failed.
func RegisterLoggerCreator(name string, creator LoggerCreator) error {
	loggerCreatorsLock.Lock()
	defer loggerCreatorsLock.Unlock()

	if creator == nil {
		return errors.New("logit: logger creator " + name + " is nil")
	}

	if _, ok := loggerCreators[name]; ok {
		return errors.New("logit: logger creator " + name + " has been registered")
	}

	loggerCreators[name] = creator
	return nil
}

// NewLoggerFromCreator creates logger from logger creator with params and returns an error if failed.
func NewLoggerFromCreator(name string, params ...interface{}) (*Logger, error) {
	loggerCreatorsLock.RLock()
	defer loggerCreatorsLock.RUnlock()

	creator, ok := loggerCreators[name]
	if !ok {
		return nil, errors.New("logit: logger creator " + name + " not found")
	}

	return creator.CreateLogger(params...)
}
