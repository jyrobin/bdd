package bdd

import "strings"

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
	} else if len(otherwise) > 0 {
		return otherwise[0]
	}
	return ""
}
