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

package logit

import (
	"context"
	"testing"
)

// go test -v -cover -run=^TestNewContextWithKey$
func TestNewContextWithKey(t *testing.T) {
	key := "key"
	logger := NewLogger()
	ctx := NewContextWithKey(context.Background(), key, logger)

	value := ctx.Value(key)
	if value == nil {
		t.Fatal("value == nil")
	}

	contextLogger, ok := value.(*Logger)
	if !ok {
		t.Fatalf("value type %T is wrong", value)
	}

	if contextLogger != logger {
		t.Fatalf("contextLogger %+v != logger %+v", contextLogger, logger)
	}
}

// go test -v -cover -run=^TestFromContextWithKey$
func TestFromContextWithKey(t *testing.T) {
	ctx := context.Background()

	key := "key"
	logger := FromContextWithKey(ctx, key)

	if logger == nil {
		t.Fatal("logger == nil")
	}

	logger = NewLogger()
	contextLogger := FromContextWithKey(context.WithValue(ctx, key, logger), key)

	if contextLogger != logger {
		t.Fatalf("contextLogger %+v != logger %+v", contextLogger, logger)
	}
}

// go test -v -cover -run=^TestNewContext$
func TestNewContext(t *testing.T) {
	logger := NewLogger()
	ctx := NewContext(context.Background(), logger)

	value := ctx.Value(contextKey{})
	if value == nil {
		t.Fatal("value == nil")
	}

	contextLogger, ok := value.(*Logger)
	if !ok {
		t.Fatalf("value type %T is wrong", value)
	}

	if contextLogger != logger {
		t.Fatalf("contextLogger %+v != logger %+v", contextLogger, logger)
	}
}

// go test -v -cover -run=^TestFromContext$
func TestFromContext(t *testing.T) {
	ctx := context.Background()
	logger := FromContext(ctx)

	if logger == nil {
		t.Fatal("logger == nil")
	}

	logger = NewLogger()
	contextLogger := FromContext(context.WithValue(ctx, contextKey{}, logger))

	if contextLogger != logger {
		t.Fatalf("contextLogger %+v != logger %+v", contextLogger, logger)
	}
}
