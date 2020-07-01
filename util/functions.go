package util

import "unicode"

// IsPublic returns true if a function name is visible outside of a package.
func IsPublic(funcName string) bool {
	return len(funcName) > 0 && unicode.IsUpper(rune(funcName[0]))
}
