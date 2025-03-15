// Copyright 2025 FishGoddess. All Rights Reserved.
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

package handler

import (
	"fmt"
	"io"
	"log/slog"
	"sync"
)

const (
	Tape = "tape"
	Text = "text"
	Json = "json"
)

var (
	newHandlers = map[string]NewHandlerFunc{
		Tape: func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return NewTapeHandler(w, opts)
		},
		Text: func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return slog.NewTextHandler(w, opts)
		},
		Json: func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return slog.NewJSONHandler(w, opts)
		},
	}
)

var (
	newHandlersLock sync.RWMutex
)

// NewHandlerFunc is a function for creating slog.Handler with w and opts.
type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

// Get gets new handler func with name and returns an error if failed.
func Get(name string) (NewHandlerFunc, error) {
	newHandlersLock.RLock()
	defer newHandlersLock.RUnlock()

	if newHandler, ok := newHandlers[name]; ok {
		return newHandler, nil
	}

	return nil, fmt.Errorf("logit: handler %s not found", name)
}

// Register registers newHandler with name.
func Register(name string, newHandler NewHandlerFunc) error {
	newHandlersLock.Lock()
	defer newHandlersLock.Unlock()

	if _, registered := newHandlers[name]; registered {
		return fmt.Errorf("logit: handler %s has been registered", name)
	}

	newHandlers[name] = newHandler
	return nil
}
