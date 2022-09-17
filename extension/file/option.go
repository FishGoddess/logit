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

package file

import (
	"os"
	"time"

	"github.com/go-logit/logit/support/size"
)

// Option will set something to config.
type Option func(c *config)

// Apply applies option to config.
func (o Option) Apply(c *config) {
	o(c)
}

// WithMode sets mode to config.
func WithMode(mode os.FileMode) Option {
	return func(c *config) {
		c.mode = mode
	}
}

// WithDirMode sets dir mode to config.
func WithDirMode(mode os.FileMode) Option {
	return func(c *config) {
		c.dirMode = mode
	}
}

// WithTimeFormat sets time format to config.
func WithTimeFormat(format string) Option {
	return func(c *config) {
		c.timeFormat = format
	}
}

// WithMaxSize sets max size to config.
func WithMaxSize(size size.ByteSize) Option {
	return func(c *config) {
		c.maxSize = size
	}
}

// WithMaxAge sets max age to config.
func WithMaxAge(age time.Duration) Option {
	return func(c *config) {
		c.maxAge = age
	}
}

// WithMaxBackups sets max backups to config.
func WithMaxBackups(count int) Option {
	return func(c *config) {
		if count < 0 {
			count = 0
		}

		c.maxBackups = count
	}
}
