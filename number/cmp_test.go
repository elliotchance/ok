package number_test

import (
	"ok/number"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmp(t *testing.T) {
	for testName, test := range comparisonTests {
		t.Run(testName, func(t *testing.T) {
			left := number.NewNumber(test.left)
			right := number.NewNumber(test.right)
			assert.Equal(t, test.cmp, number.Cmp(left, right))
		})
	}
}
