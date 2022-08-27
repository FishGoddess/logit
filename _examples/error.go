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

package main

import (
	"errors"
	"fmt"

	"github.com/go-logit/logit"
	"github.com/go-logit/logit/core"
)

// uselessWriter is a demo writer to demonstrate the error handling function.
type uselessWriter struct{}

func (uw uselessWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("always error in writing")
}

func main() {
	// You can specify a function to handle errors happens in logger.
	// For example, you can count these errors and report them to team members by email.
	core.HandleError = func(name string, err error) {
		fmt.Printf("%s received an error: %+v\n", name, err)
	}

	// Let's log something to see what happen.
	logger := logit.NewLogger(logit.Options().WithWriter(&uselessWriter{}))
	logger.Info("See what happen?").Log()
}
