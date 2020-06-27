package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	for testName, test := range arithmeticTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			result := number.Add(left, right)
			assert.Equal(t, test.add, result.String())
		})
	}
}
