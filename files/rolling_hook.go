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

package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// RollingHook is a hook that will be invoked in rolling process.
// This interface has two method: BeforeRolling and AfterRolling.
// BeforeRolling will be called before rolling to next file.
// AfterRolling will be called after rolling to next file.
// You can do your job in both of methods.
type RollingHook interface {
	BeforeRolling()
	AfterRolling()
}

// ============================= default rolling hook =============================

// DefaultRollingHook is default rolling hook, which will do nothing on rolling to next file.
type DefaultRollingHook struct{}

// NewDefaultRollingHook returns a default rolling hook.
func NewDefaultRollingHook() *DefaultRollingHook {
	return &DefaultRollingHook{}
}

// BeforeRolling does nothing when rolling to next file.
func (drh *DefaultRollingHook) BeforeRolling() {
	// Do nothing...
}

// AfterRolling does nothing when rolling to next file.
func (drh *DefaultRollingHook) AfterRolling() {
	// Do nothing...
}

// ============================= life rolling hook =============================

// LifeBasedRollingHook is a life based rolling hook.
// It will tag a life on every file and if life runs out, this file will be removed.
type LifeBasedRollingHook struct {

	// DefaultRollingHook has default implement and BeforeRolling is reserved.
	*DefaultRollingHook

	// life is the life of every file.
	life time.Duration

	// directory is the target storing all files need to monitor.
	directory string
}

// NewLifeBasedRollingHook returns a LifeBasedRollingHook holder.
// life is the life of every file and directory is the target storing all files need to monitor.
func NewLifeBasedRollingHook(life time.Duration, directory string) *LifeBasedRollingHook {
	return &LifeBasedRollingHook{
		DefaultRollingHook: &DefaultRollingHook{},
		life:               life,
		directory:          directory,
	}
}

// AfterRolling checks life of all files and removes the dead files.
// Remember, it will do nothing if this directory could not be read and modified.
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
