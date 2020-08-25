package util

import (
	"sort"
	"strings"
)

func Interface(in map[string]string) string {
	var keys []string
	for key := range in {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	s := "{ "
	for i, key := range keys {
		if i > 0 {
			s += "; "
		}

		if strings.HasPrefix(in[key], "func(") {
			s += key + strings.TrimPrefix(in[key], "func")
		} else {
			s += key + " " + in[key]
		}
	}

	return s + " }"
}
