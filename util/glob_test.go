package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchesGlob(t *testing.T) {
	tests := map[string]struct {
		s, glob string
		match   bool
	}{
		"1":  {"foo", "foo", true},
		"2":  {"food", "foo", false},
		"3":  {"Foo", "foo", false},
		"4":  {"", "", true},
		"5":  {"foo", "foo*", true},
		"6":  {"food", "foo*", true},
		"7":  {"foo", "*foo", true},
		"8":  {"aafoo", "*foo", true},
		"9":  {"foo", "f*oo", true},
		"10": {"f12oo", "f*oo", true},
		"11": {"foo", "f**oo", true},
		"12": {"f12oo", "f**oo", true},
		"13": {"f12o1o", "f**oo", false},
		"14": {"f1o2o", "f*o*o", true},
		"15": {"f1oo", "f*o*o", true},
		"16": {"foo", "*", true},
		"17": {"", "*", true},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, tt.match, MatchesGlob(tt.s, tt.glob))
		})
	}
}
