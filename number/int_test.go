package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

var intTests = map[string]int64{
	"0":                0,
	"1234":             1234,
	"-100":             -100,
	"123.4":            123,
	"123.8":            123,
	"1234567890123456": 1234567890123456,
}

func TestInt64(t *testing.T) {
	for n, expected := range intTests {
		t.Run(n, func(t *testing.T) {
			num := number.NewNumber(n)
			assert.Equal(t, expected, number.Int64(num))
		})
	}
}

func TestInt(t *testing.T) {
	for n, expected := range intTests {
		t.Run(n, func(t *testing.T) {
			num := number.NewNumber(n)
			assert.Equal(t, int(expected), number.Int(num))
		})
	}
}
