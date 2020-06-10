package number_test

import (
	"ok/number"
	"testing"

	"github.com/stretchr/testify/assert"
)

var arithmeticTests = map[string]struct {
	left, right                 string
	add, subtract               string
	multiply, divide, remainder string
}{
	"zero": {
		"0", "0",
		"0", "0",

		// Must skip the divide by zero. There is a separate test for this in
		// TestDivide/TestRemainder.
		"0", "SKIP", "SKIP",
	},
	"negative-zero": {
		"0", "-0",
		"0", "0",

		// Must skip the divide by zero. There is a separate test for this in
		// TestDivide/TestRemainder.
		"0", "SKIP", "SKIP",
	},
	"positive-ints": {
		"456", "123",
		"579", "333",
		"56088", "4", "87",
	},
	"mixed-ints": {
		"-456", "123",
		"-333", "-579",
		"-56088", "-4", "-87",
	},
	"same-precision": {
		"45.69", "1.23",
		"46.92", "44.46",
		"56.20", "37.15", "0.18",
	},
	"precision-greater-left": {
		"45.69", "1.231",
		"46.921", "44.459",
		"56.244", "37.116", "0.143",
	},
	"precision-greater-right": {
		"45.692", "1.28",
		"46.972", "44.412",

		// 45.692 * 1.28 = 58.48576. Make sure rounding is 58.486.
		// 45.692 / 1.28 = 35.696875. Make sure rounding is 35.697.
		"58.486", "35.697", "0.892",
	},
	"precision-retained": {
		"34.000", "1.200",
		"35.200", "32.800",
		"40.800", "28.333", "0.400",
	},
}

var newNumberTests = map[string]struct {
	ratString string
	precision int
	str       string
	isZero    bool
}{
	"0":           {"0/1", 0, "0", true},
	"1":           {"1/1", 0, "1", false},
	"100":         {"100/1", 0, "100", false},
	"1.23":        {"123/100", 2, "1.23", false},
	"0.5":         {"1/2", 1, "0.5", false},
	"12.3400":     {"617/50", 4, "12.3400", false},
	"-45.1":       {"-451/10", 1, "-45.1", false},
	"00123456.10": {"1234561/10", 2, "123456.10", false},
	"0.000":       {"0/1", 3, "0.000", true},
	"-0":          {"0/1", 0, "0", true},
	"-0.0":        {"0/1", 1, "0.0", true},
}

var comparisonTests = map[string]struct {
	left, right string
	cmp         int
}{
	"zero": {
		"0", "0",
		0,
	},
	"negative-zero": {
		"0", "-0",
		0,
	},
	"positive-ints": {
		"456", "123",
		1,
	},
	"mixed-ints": {
		"-456", "123",
		-1,
	},
	"same-precision": {
		"45.69", "45.69",
		0,
	},
	"precision-greater-left": {
		"45.690", "45.69",
		0,
	},
	"precision-greater-right": {
		"45", "45.0000",
		0,
	},
}

func TestNewNumber(t *testing.T) {
	for n, test := range newNumberTests {
		t.Run(n, func(t *testing.T) {
			actual := number.NewNumber(n)
			assert.Equal(t, test.ratString, actual.Rat.String())
			assert.Equal(t, test.precision, actual.Precision)
		})
	}
}

func TestNumber_String(t *testing.T) {
	for n, test := range newNumberTests {
		t.Run(n, func(t *testing.T) {
			actual := number.NewNumber(n).String()
			assert.Equal(t, test.str, actual)
		})
	}
}

func TestNumber_IsZero(t *testing.T) {
	for n, test := range newNumberTests {
		t.Run(n, func(t *testing.T) {
			actual := number.NewNumber(n).IsZero()
			assert.Equal(t, test.isZero, actual)
		})
	}
}
