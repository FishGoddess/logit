// Copyright 2024 FishGoddess. All Rights Reserved.
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
	"context"
)

type contextKey struct{}

// NewContext wraps context with logger and returns a new context.
func NewContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext gets logger from context and returns the default logger if missed.
func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return logger
	}

	return Default()
}
