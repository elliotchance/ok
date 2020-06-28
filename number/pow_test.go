package number_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

var powTests = []struct {
	a, b string
	pow  string
}{
	{"0", "4.567", "0"},
	{"1.23", "0", "1"},
	{"1.23", "4.567", "2.5739294795546728988"},
}

func TestPow(t *testing.T) {
	for _, test := range powTests {
		t.Run(fmt.Sprintf("%s,%s", test.a, test.b), func(t *testing.T) {
			a := number.NewNumber(test.a)
			b := number.NewNumber(test.b)
			result := number.Pow(a, b)
			assert.Equal(t, test.pow, number.Format(result, -1))
		})
	}
}
