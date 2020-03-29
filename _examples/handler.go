// Copyright 2020 Ye Zi Jie. All Rights Reserved.
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
//
// Author: fish
// Email: fishinlove@163.com
// Created at 2020/03/06 16:01:00

package main

import (
    "fmt"
    "os"

    "github.com/FishGoddess/logit"
)

type myHandler struct{}

// Customize your own handler.
func (mh *myHandler) Handle(log *logit.Log) bool {
    os.Stdout.Write([]byte("myHandler: "))
    os.Stdout.Write(logit.EncodeToJson(log, "")) // Try `os.Stdout.WriteString(log.Msg())` ?
    return true
}

// We recommend you to register your handler to logit, so that
// you can use your handler by logit.HandlerOf.
// See logit.HandlerOf.
func init() {
    logit.RegisterHandler("myHandler", func() logit.Handler {
        return &myHandler{}
    })
}

func main() {

    // Create a logger holder.
    // Default handler is logit.DefaultHandler.
    logger := logit.NewDevelopLogger()
    logger.Info("before adding handlers...")

    // Add handlers to logger.
    // There are two handlers in logger because logger has a default handler inside after creating.
    // See logit.DefaultHandler.
    logger.AddHandlers(&myHandler{}, logit.HandlerOf("json"))
    fmt.Println("fmt =========================================")
    logger.Info("after adding handlers...")

    // Set handlers to logger.
    // There are two handlers in logger because the default handler inside was removed.
    // If you register your handler to logit by logit.RegisterHandler, then you can
    // use your handler everywhere like this:
    logger.SetHandlers(logit.HandlerOf("myHandler"))
    fmt.Println("fmt =========================================")
    logger.Info("after setting handlers...")
}
