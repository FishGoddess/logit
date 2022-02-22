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
	"context"
)

// contextKey is the type of context key.
type contextKey struct{}

var (
	// contextLogger is a non-nil but useless logger for FromContextWithKey() if missing.
	contextLogger = NewLogger(Options().WithOffLevel())
)

// NewContextWithKey wraps context with logger of key and returns a new context.
func NewContextWithKey(ctx context.Context, key interface{}, logger *Logger) context.Context {
	return context.WithValue(ctx, key, logger)
}

// FromContextWithKey gets logger from context with key and returns a discard logger if missing.
func FromContextWithKey(ctx context.Context, key interface{}) *Logger {
	if logger, ok := ctx.Value(key).(*Logger); ok {
		return logger
	}

	if globalLogger != nil {
		return globalLogger
	}
	return contextLogger
}

// NewContext wraps context with logger and returns a new context.
func NewContext(ctx context.Context, logger *Logger) context.Context {
	return NewContextWithKey(ctx, contextKey{}, logger)
}

// FromContext gets logger from context and returns a discard logger if missing.
func FromContext(ctx context.Context) *Logger {
	return FromContextWithKey(ctx, contextKey{})
}
