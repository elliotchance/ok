package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

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
			assert.Equal(t, test.divide, result.String())
		})
	}

	t.Run("divide-by-zero", func(t *testing.T) {
		left := number.NewNumber("1")
		right := number.NewNumber("0")
		result, err := number.Divide(left, right)
		assert.EqualError(t, err, "division by zero")
		assert.Equal(t, "0", result.String())
	})
}
