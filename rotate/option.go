// Copyright 2023 FishGoddess. All Rights Reserved.
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

package rotate

import (
	"time"
)

// Option sets some fields to config.
type Option func(c *config)

func (o Option) apply(c *config) {
	o(c)
}

// WithMaxSize sets max size to config.
func WithMaxSize(size uint64) Option {
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
func WithMaxBackups(backups uint32) Option {
	return func(c *config) {
		c.maxBackups = backups
	}
}
