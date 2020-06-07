package number_test

import (
	"ok/number"
	"testing"

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
