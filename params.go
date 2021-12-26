// Copyright (c) 2021 Jing-Ying Chen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bdd

import (
	"strconv"
	"strings"
)

type params struct {
	vals map[string]string
}

func Params(args []string) params {
	vals := map[string]string{}
	for _, arg := range args {
		pair := strings.SplitN(arg, ":", 2)
		key := strings.TrimSpace(pair[0])
		var val string
		if len(pair) == 2 {
			val = strings.TrimSpace(pair[1])
		}
		vals[key] = val
	}
	return params{vals}
}

func (args params) Get(name string, otherwise ...string) string {
	if val, ok := args.vals[name]; ok {
		return val
	}

	if len(otherwise) > 0 {
		return otherwise[0]
	}
	return ""
}

func (args params) GetInt(name string, otherwise ...int) int {
	if val, ok := args.vals[name]; ok {
		if ret, err := strconv.Atoi(val); err == nil {
			return ret
		}
	}

	if len(otherwise) > 0 {
		return otherwise[0]
	}
	return 0
}

func ParamString(args []string, key string, otherwise ...string) string {
	return Params(args).Get(key, otherwise...)
}

func ParamStrings(args []string, keys ...string) []string {
	ps := Params(args)
	ret := make([]string, len(keys))
	for i, key := range keys {
		ret[i] = ps.Get(key)
	}
	return ret
}

func ParamMap(args []string, keys ...string) map[string]string {
	ret := Params(args).vals
	if len(keys) > 0 {
		vals := ret
		ret = map[string]string{}
		for _, key := range keys {
			if val, ok := vals[key]; ok {
				ret[key] = val
			}
		}
	}
	return ret
}

func GetStrings(args []string, key string, otherwise ...string) string {
	return Params(args).Get(key, otherwise...)
}
