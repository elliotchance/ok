package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

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

func TestCmp(t *testing.T) {
	for testName, test := range comparisonTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			assert.Equal(t, test.cmp, number.Cmp(left, right))
		})
	}
}
