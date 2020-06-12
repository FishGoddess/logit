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
// Created at 2020/06/12 22:08:31

package writer

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type RollingHook interface {
	BeforeRolling()
	AfterRolling()
}

// ============================= default rolling hook =============================
type DefaultRollingHook struct{}

func (drh *DefaultRollingHook) BeforeRolling() {
	// Do nothing...
}

func (drh *DefaultRollingHook) AfterRolling() {
	// Do nothing...
}

// ============================= life rolling hook =============================

type LifeBasedRollingHook struct {
	*DefaultRollingHook

	life time.Duration

	directory string
}

func NewLifeBasedRollingHook(life time.Duration, directory string) *LifeBasedRollingHook {
	return &LifeBasedRollingHook{
		DefaultRollingHook: &DefaultRollingHook{},
		life:               life,
		directory:          directory,
	}
}

func (lrh *LifeBasedRollingHook) AfterRolling() {

	fileInfos, err := ioutil.ReadDir(lrh.directory)
	if err != nil {
		return
	}

	now := time.Now()
	for _, file := range fileInfos {
		if !file.IsDir() && now.Sub(file.ModTime()) >= lrh.life {
			os.Remove(filepath.Join(lrh.directory, file.Name()))
		}
	}
}
