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
	"io"
	"os"
	"path/filepath"

	"github.com/FishGoddess/logit/defaults"
	"github.com/FishGoddess/logit/handler"
	"github.com/FishGoddess/logit/rotate"
	"github.com/FishGoddess/logit/writer"
)

type Option func(conf *config)

func (o Option) applyTo(conf *config) {
	o(conf)
}

func WithDebugLevel() Option {
	return func(conf *config) {
		conf.level = levelDebug
	}
}

func WithInfoLevel() Option {
	return func(conf *config) {
		conf.level = levelInfo
	}
}

func WithWarnLevel() Option {
	return func(conf *config) {
		conf.level = levelWarn
	}
}

func WithErrorLevel() Option {
	return func(conf *config) {
		conf.level = levelError
	}
}

func WithPrintLevel() Option {
	return func(conf *config) {
		conf.level = levelPrint
	}
}

func WithOffLevel() Option {
	return func(conf *config) {
		conf.level = levelOff
	}
}

func WithWriter(w io.Writer) Option {
	newWriter := func() (io.Writer, error) {
		return w, nil
	}

	return func(conf *config) {
		conf.newWriter = newWriter
	}
}

func WithStdout() Option {
	newWriter := func() (io.Writer, error) {
		return os.Stdout, nil
	}

	return func(conf *config) {
		conf.newWriter = newWriter
	}
}

func WithStderr() Option {
	newWriter := func() (io.Writer, error) {
		return os.Stderr, nil
	}

	return func(conf *config) {
		conf.newWriter = newWriter
	}
}

func WithFile(path string) Option {
	newWriter := func() (io.Writer, error) {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, defaults.FileDirMode); err != nil {
			return nil, err
		}

		return defaults.OpenFile(path, defaults.FileMode)
	}

	return func(conf *config) {
		conf.newWriter = newWriter
	}
}

func WithRotateFile(path string, opts ...rotate.Option) Option {
	newWriter := func() (io.Writer, error) {
		return rotate.New(path, opts...)
	}

	return func(conf *config) {
		conf.newWriter = newWriter
	}
}

func WithBuffer(bufferSize uint64) Option {
	wrapWriter := func(w io.Writer) writer.Writer {
		return writer.Buffer(w, bufferSize)
	}

	return func(conf *config) {
		conf.wrapWriter = wrapWriter
	}
}

func WithBatch(batchSize uint64) Option {
	wrapWriter := func(w io.Writer) writer.Writer {
		return writer.Batch(w, batchSize)
	}

	return func(conf *config) {
		conf.wrapWriter = wrapWriter
	}
}

func WithHandler(newHandler NewHandlerFunc) Option {
	return func(conf *config) {
		conf.newHandler = newHandler
	}
}

func WithTextHandler() Option {
	return func(conf *config) {
		conf.newHandler = handler.NewTextHandler
	}
}

func WithJsonHandler() Option {
	return func(conf *config) {
		conf.newHandler = handler.NewJsonHandler
	}
}

func WithSource() Option {
	return func(conf *config) {
		conf.withSource = true
	}
}

func WithPID() Option {
	return func(conf *config) {
		conf.withPID = true
	}
}
