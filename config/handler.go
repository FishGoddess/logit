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

package config

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"

	"github.com/FishGoddess/logit"
)

var (
	newHandlers = map[string]NewHandlerFunc{
		"text": func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return logit.NewTextHandler(w, opts)
		},
		"json": func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return logit.NewJsonHandler(w, opts)
		},
		"slog.text": func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return slog.NewTextHandler(w, opts)
		},
		"slog.json": func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return slog.NewJSONHandler(w, opts)
		},
	}

	newHandlersLock sync.RWMutex
)

type NewHandlerFunc func(w io.Writer, opts *slog.HandlerOptions) slog.Handler

func newHandler(name string, w io.Writer, opts *slog.HandlerOptions) (slog.Handler, error) {
	newHandlersLock.RLock()
	defer newHandlersLock.RUnlock()

	if newHandler, ok := newHandlers[name]; ok {
		return newHandler(w, opts), nil
	}

	var handlerNames strings.Builder
	for name := range newHandlers {
		handlerNames.WriteString(name)
		handlerNames.WriteString(",")
	}

	return nil, fmt.Errorf("logit: handler %s not found, available handlers are %s", name, handlerNames.String())
}

func RegisterHandler(name string, newHandler NewHandlerFunc) error {
	newHandlersLock.Lock()
	defer newHandlersLock.Unlock()

	if _, registered := newHandlers[name]; registered {
		return fmt.Errorf("logit: handler %s has been registered", name)
	}

	newHandlers[name] = newHandler
	return nil
}
