// Copyright (c) 2021 Jing-Ying Chen. Subject to the MIT License.

package bdd

import (
	"fmt"
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

func Split(args string, seps ...string) params {
	sep := ","
	if len(seps) > 0 {
		sep = seps[0]
	}

	vals := map[string]string{}
	for _, arg := range strings.Split(args, sep) {
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

func (args params) Ensure(vals map[string]string) error {
	for k, v := range args.vals {
		if val, ok := vals[k]; !ok || val != v {
			return fmt.Errorf("%s is %s, expected %s", k, val, v)
		}
	}
	return nil
}

func ParamString(args []string, key string, otherwise ...string) string {
	return Params(args).Get(key, otherwise...)
}

// Get the first field only
func ParamField(args []string, key string, otherwise ...string) string {
	ret := Params(args).Get(key, otherwise...)
	if len(ret) > 0 {
		ret = strings.Fields(ret)[0]
	}
	return ret
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

func GetString(args []string, key string, otherwise ...string) string {
	return Params(args).Get(key, otherwise...)
}

func GetInt(args []string, key string, otherwise ...int) int {
	return Params(args).GetInt(key, otherwise...)
}

func Ensure(vals map[string]string, args ...string) error {
	return Params(args).Ensure(vals)
}

func EnsureNonEmpty(vals map[string]string, args ...string) error {
	for _, arg := range args {
		if val, ok := vals[arg]; !ok || val == "" {
			return fmt.Errorf("%s is empty", arg)
		}
	}
	return nil
}
