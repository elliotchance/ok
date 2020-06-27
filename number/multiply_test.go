package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
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
