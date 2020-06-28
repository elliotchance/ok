package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

var newNumberTests = map[string]struct {
	fmtNeg1, fmtZero, fmtThree string
	isZero                     bool
	wholeDigits                int
}{
	"0":           {"0", "0", "0.000", true, 1},
	"1":           {"1", "1", "1.000", false, 1},
	"100":         {"100", "100", "100.000", false, 3},
	"1.23":        {"1.23", "1", "1.230", false, 1},
	"1.2343":      {"1.2343", "1", "1.234", false, 1},
	"0.5":         {"0.5", "1", "0.500", false, 1},
	"12.8400":     {"12.84", "13", "12.840", false, 2},
	"-45.1":       {"-45.1", "-45", "-45.100", false, 2},
	"00123456.10": {"123456.1", "123456", "123456.100", false, 6},
	"0.000":       {"0", "0", "0.000", true, 1},
	"-0":          {"0", "0", "0.000", true, 1},
	"-0.0":        {"0", "0", "0.000", true, 1},
}

func TestNewNumber(t *testing.T) {
	for n := range newNumberTests {
		t.Run(n, func(t *testing.T) {
			actual := number.NewNumber(n)
			assert.NotNil(t, actual)
		})
	}
}

func TestIsZero(t *testing.T) {
	for n, test := range newNumberTests {
		t.Run(n, func(t *testing.T) {
			actual := number.IsZero(number.NewNumber(n))
			assert.Equal(t, test.isZero, actual)
		})
	}
}
