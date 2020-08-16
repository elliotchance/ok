package util

import (
	"regexp"
	"strings"
)

// MatchesGlob is a very rudimentary implementation of glob matching on strings.
// Right now it only supports "*" that will match zero or more characters
// anywhere within the string.
func MatchesGlob(s, glob string) bool {
	regex := "^" + strings.ReplaceAll(glob, "*", ".*") + "$"

	return regexp.MustCompile(regex).MatchString(s)
}
