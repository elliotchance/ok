package util

// StringSliceMap creates a new slice with a function applied.
func StringSliceMap(ss []string, fn func(string) string) []string {
	ss2 := make([]string, len(ss))
	for i, v := range ss {
		ss2[i] = fn(v)
	}

	return ss2
}
