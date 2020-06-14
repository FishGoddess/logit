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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	// rollingHooks stores all rollingHooks registered.
	// mutexOfRollingHooks is for concurrency.
	rollingHooks = map[string]func(params map[string]interface{}) RollingHook{
		"default": newDefaultRollingHookFunc,
		"life":    newLifeBasedRollingHookFunc,
	}
	mutexOfRollingHooks = &sync.RWMutex{}

	// RollingHookIsExistedError is an error happening on repeating rollingHook name.
	RollingHookIsExistedError = errors.New("the name of rollingHook you want to register already exists! May be you should give it an another name")
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

// RegisterRollingHook registers your rollingHook to logit so that you can use them in config file.
// Return an error if the name is existed, and you should change another name for your rollingHook.
// Notice that newRollingHook has a parameter called params, which will be injected into newRollingHook
// by logit automatically. Different rollingHook may have different params, so what params should
// be injected into newRollingHook is dependent to specific rollingHook. Actually, this params is a
// mapping of config file. All params you write in config file will be injected here.
// For example, your config file is like this:
//
//     "rollingHook": {
//         "myRollingHook": {
//             "db": "127.0.0.1:3306",
//             "user": "me",
//             "password": "you guess?",
//             "maxConnections": 1024
//         }
//     }
//
// Then a map[string]interface{} {
//            "db": "127.0.0.1:3306",
//            "user": "me",
//            "password": "you guess?",
//            "maxConnections": 1024
//        } will be injected to params.
//
// So you can use these params written in config file.
func RegisterRollingHook(name string, newRollingHook func(params map[string]interface{}) RollingHook) error {
	mutexOfRollingHooks.Lock()
	defer mutexOfRollingHooks.Unlock()
	if _, ok := rollingHooks[name]; ok {
		return RollingHookIsExistedError
	}
	rollingHooks[name] = newRollingHook
	return nil
}

// rollingHookOf returns rollingHook whose name is given name.
// Notice that we use tips+exit mechanism to check the name.
// This is a more convenient way to use rollingHook (we think).
// so if the rollingHook doesn't exist, a tip will be printed and
// the program will exit with status code 11.
func rollingHookOf(name string, params map[string]interface{}) RollingHook {
	mutexOfRollingHooks.RLock()
	defer mutexOfRollingHooks.RUnlock()
	newRollingHook, ok := rollingHooks[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: The rollingHook \"%s\" doesn't exist! Please change it to another rollingHook.\n", name)
		os.Exit(11)
	}
	return newRollingHook(params)
}

// ================================ for code-readable ================================

// newDefaultRollingHookFunc creates a default rolling hook with given params.
// After registering to logit, you can use it in config file, try this:
//
//     "rollingHook": {
//         "default": {}
//     }
//
func newDefaultRollingHookFunc(params map[string]interface{}) RollingHook {
	return NewDefaultRollingHook()
}

// newLifeBasedRollingHookFunc creates a life based rolling hook with params.
// After registering to logit, you can use it in config file, try this:
//
//     "rollingHook": {
//         "life": {
//             "limit": 60,
//             "directory": "D:/logs"
//         }
//     }
//
// Notice that the unit of limit is second, and default value of limit is one day.
// Directory is the target storing all files, and it is a necessary parameter.
func newLifeBasedRollingHookFunc(params map[string]interface{}) RollingHook {

	// limit 是寿命，单位是秒，默认值是一天
	limit := 24 * 60 * 60
	if param, ok := params["limit"]; ok {
		limit = int(param.(float64))
	}

	// directory 是要监控的文件目录，一般配置的就是几个 rollingFile 中的 directory，否则没有意义
	// 另外，这个值是必须的，不然就没法正常工作
	directory, ok := params["directory"]
	if !ok {
		fmt.Fprintln(os.Stderr, "Error: The rollingHook life misses a necessary parameter \"directory\"!")
		os.Exit(12)
	}

	return NewLifeBasedRollingHook(directory.(string), time.Duration(limit)*time.Second)
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

// ============================= life based rolling hook =============================

// LifeBasedRollingHook is a life based rolling hook.
// It will tag a life on every file and if life runs out, this file will be removed.
type LifeBasedRollingHook struct {

	// DefaultRollingHook has default implement and BeforeRolling is reserved.
	*DefaultRollingHook

	// directory is the target storing all files need to monitor.
	directory string

	// life is the life of every file.
	life time.Duration
}

// NewLifeBasedRollingHook returns a LifeBasedRollingHook holder.
// life is the life of every file and directory is the target storing all files need to monitor.
func NewLifeBasedRollingHook(directory string, life time.Duration) *LifeBasedRollingHook {
	return &LifeBasedRollingHook{
		DefaultRollingHook: &DefaultRollingHook{},
		directory:          directory,
		life:               life,
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
