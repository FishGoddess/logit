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

package main

import (
	"encoding/json"
	"os"

	"github.com/FishGoddess/logit"
	"github.com/FishGoddess/logit/extension/config"
)

// newConfig reads config from a json file.
func newConfig() (*config.Config, error) {
	conf := new(config.Config)

	bs, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}

func main() {
	// We know you may use a configuration file in your program to setup some resources including logging.
	// After thinking about serval kinds of configurations like "yaml", "toml" and "json", we decide to support all of them.
	// As you can see, there are many tags on Config's fields like "yaml" and "toml", so you can unmarshal to a config from one of them.
	conf, err := newConfig()
	if err != nil {
		panic(err)
	}

	opts, err := conf.Options()
	if err != nil {
		panic(err)
	}

	// Use options to create a logger.
	logger := logit.NewLogger(opts...)
	defer logger.Close()

	logger.Info("logging from config", "conf", conf)
}
