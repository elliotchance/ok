package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
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
		"56088", "3.7073170731707317073", "87",
	},
	"mixed-ints": {
		"-456", "123",
		"-333", "-579",
		"-56088", "-3.7073170731707317073", "-87",
	},
	"same-precision": {
		"45.69", "1.23",
		"46.92", "44.46",
		"56.1987", "37.146341463414634146", "0.18",
	},
	"precision-greater-left": {
		"45.69", "1.231",
		"46.921", "44.459",
		"56.24439", "37.116165718927701056", "0.143",
	},
	"precision-greater-right": {
		"45.692", "1.28",
		"46.972", "44.412",
		"58.48576", "35.696875", "0.892",
	},
	"zero-padded": {
		"34.000", "1.200",
		"35.2", "32.8",
		"40.8", "28.333333333333333333", "0.4",
	},
}

func TestAdd(t *testing.T) {
	for testName, test := range arithmeticTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result := number.Add(left, right)
			assert.Equal(t, test.add, number.Format(result, -1))
		})
	}
}

func TestSubtract(t *testing.T) {
	for testName, test := range arithmeticTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result := number.Subtract(left, right)
			assert.Equal(t, test.subtract, number.Format(result, -1))
		})
	}
}

func TestMultiply(t *testing.T) {
	for testName, test := range arithmeticTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result := number.Multiply(left, right)
			assert.Equal(t, test.multiply, number.Format(result, -1))
		})
	}
}

func TestDivide(t *testing.T) {
	for testName, test := range arithmeticTests {
		if test.divide == "SKIP" {
			continue
		}

		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result, err := number.Divide(left, right)
			assert.NoError(t, err)
			assert.Equal(t, test.divide, number.Format(result, -1))
		})
	}

	t.Run("divide-by-zero", func(t *testing.T) {
		left := number.NewNumber("1")
		right := number.NewNumber("0")
		result, err := number.Divide(left, right)
		assert.EqualError(t, err, "division by zero")
		assert.Equal(t, "0", number.Format(result, -1))
	})
}

func TestRemainder(t *testing.T) {
	for testName, test := range arithmeticTests {
		if test.divide == "SKIP" {
			continue
		}

		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result, err := number.Remainder(left, right)
			assert.NoError(t, err)
			assert.Equal(t, test.remainder, number.Format(result, -1))
		})
	}

	t.Run("divide-by-zero", func(t *testing.T) {
		left := number.NewNumber("1")
		right := number.NewNumber("0")
		result, err := number.Remainder(left, right)
		assert.EqualError(t, err, "division by zero")
		assert.Equal(t, "0", number.Format(result, -1))
	})
}
