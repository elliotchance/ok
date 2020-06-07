package number_test

import (
	"ok/number"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	for testName, test := range arithmeticTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result := number.Multiply(left, right)
			assert.Equal(t, test.multiply, result.String())
		})
	}
}
