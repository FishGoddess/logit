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
// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/03/25 23:06:59

package logit

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

// config is the config mapping the config file.
type config struct {

	// Level is the logger level in string form.
	Level string `json:"level"`

	// Caller will determine weather you need caller info in log or not.
	// Notice that adding caller will call a runtime method which costs lots of time,
	// so set it to true only when you really need it.
	Caller bool `json:"caller"`

	// Handlers is the handlers you want to use mapping to config file.
	Handlers map[string]map[string]interface{} `json:"handlers"`
}

// jsonifyContent normalize content to a standard json string in bytes.
// We make a new config format which is transformed from Json.
// As we know, a standard Json string is strict in format. However,
// the worst thing when using a Json as config file is we can't add any
// comments so we have no idea that why we use this configuration in our config.
// So we must transform it. The first thing we do is let it supports comments.
// Also, the "{}" in Json string is redundant in a config file. After doing these,
// a reading-friendly config file is born, and it is not a standard Json string any more.
// This method jsonifyContent will fix this so that we can use Json parser to
// parse our config file.
func jsonifyContent(content []byte) []byte {

	// 我们的配置文件是支持注释的，而 Json 规范中并没有对注释的支持，所以我们需要对注释进行擦除
	// 注意：注释只能是单独起一行，并且以 # 开头
	var buffer []byte
	lines := bytes.Split(content, []byte("\n"))
	for _, line := range lines {
		if !bytes.HasPrefix(bytes.TrimSpace(line), []byte("#")) {
			buffer = append(buffer, line...)
		}
	}

	// 由于配置文件使用 Json 格式，而 Json 规范要求使用 {} 包裹内容，但这个 {} 不方便配置文件的阅读，
	// 所以我们设定在配置文件中不使用 {} 包裹，而是交给我们读取出来之后进行包裹
	return bytes.Join([][]byte{[]byte("{"), buffer, []byte("}")}, []byte(""))
}

// parseConfigFrom parses a config from reader which returns a mapping config.
// Return an error if something wrong happened.
func parseConfigFrom(reader io.Reader) (config, error) {

	// 配置文件一般不会太大，直接全部读取进内存
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return config{}, err
	}

	// 反序列化出配置
	conf := config{
		Level: "debug",
	}
	err = json.Unmarshal(jsonifyContent(content), &conf)
	if err != nil {
		return config{}, err
	}

	return conf, nil
}

// parseHandlersFrom parses all handlers in conf.
// Return a slice of all parsed handlers.
func parseHandlersFrom(conf config) []Handler {
	handlers := make([]Handler, 0, len(conf.Handlers)+2)
	for name, params := range conf.Handlers {
		handlers = append(handlers, handlerOf(name, params))
	}
	return handlers
}
