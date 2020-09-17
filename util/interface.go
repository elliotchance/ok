package util

import (
	"sort"
	"strings"

	"github.com/elliotchance/ok/types"
)

func Interface(in map[string]*types.Type) string {
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

		if in[key].Kind == types.KindFunc {
			s += key + strings.TrimPrefix(in[key].String(), "func")
		} else {
			s += key + " " + in[key].String()
		}
	}

	return s + " }"
}
