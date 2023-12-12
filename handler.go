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
	"fmt"
	"io"
	"log/slog"
	"sync"
)

var (
	newHandlers = map[string]NewHandlerFunc{
		"text": newTextHandler,
		"json": newJsonHandler,
	}
)

var (
	newHandlersLock sync.RWMutex
)

// NewHandlerFunc is a function for creating slog.Handler with w and opts.
type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

func newTextHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(w, opts)
}

func newJsonHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(w, opts)
}

// PickHandler picks the handler with name and returns an error if failed.
func PickHandler(name string) (NewHandlerFunc, error) {
	newHandlersLock.RLock()
	defer newHandlersLock.RUnlock()

	if newHandler, ok := newHandlers[name]; ok {
		return newHandler, nil
	}

	return nil, fmt.Errorf("logit: handler %s unknown", name)
}

// RegisterHandler registers newHandler with name to logit.
func RegisterHandler(name string, newHandler NewHandlerFunc) error {
	newHandlersLock.Lock()
	defer newHandlersLock.Unlock()

	if _, registered := newHandlers[name]; registered {
		return fmt.Errorf("logit: handler %s has been registered", name)
	}

	newHandlers[name] = newHandler
	return nil
}

// Syncer is an interface that syncs data to somewhere.
type Syncer interface {
	Sync() error
}

// Writer is an interface which have sync, write and close functions.
type Writer interface {
	io.Writer
	Syncer
	io.Closer
}

type nilSyncer struct{}

func (nilSyncer) Sync() error {
	return nil
}

type nilCloser struct{}

func (nilCloser) Close() error {
	return nil
}
