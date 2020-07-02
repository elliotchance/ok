package number_test

import (
	"testing"

	"github.com/elliotchance/ok/number"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	for testName, test := range map[string]string{
		"0.123": "-2.0955709236097195568",
		"12.34": "2.5128460184772417554",
	} {
		t.Run(testName, func(t *testing.T) {
			result := number.Log(number.NewNumber(testName))
			assert.Equal(t, test, number.Format(result, -1))
		})
	}
}
