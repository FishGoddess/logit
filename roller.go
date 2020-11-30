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
// Created at 2020/12/01 00:18:24

package logit

import "os"

// Roller is an interface of rolling a file writer.
// File writer will call TimeToRoll() to know if time to roll before writing,
// and if true, the Roll() will be called.
type Roller interface {

	// TimeToRoll returns true if need rolling or false.
	// Although file is a pointer, you shouldn't change it in this method.
	// Remember, file in this method should be read only.
	TimeToRoll(fw *FileWriter) bool

	// Roll will roll this file and returns an error if failed.
	// Although file is a pointer, you shouldn't change it in this method.
	// Return an os.File instance will be fine.
	Roll(fw *FileWriter) (*os.File, error)
}
