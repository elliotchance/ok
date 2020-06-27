package parser_test

import (
	"testing"

	"github.com/elliotchance/ok/parser"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected map[string]string
	}{
		"math": {
			str: `import "math"`,
			expected: map[string]string{
				"math": "math",
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str)

			assert.Nil(t, p.Errors)
			assert.Equal(t, test.expected, p.File.Imports)
		})
	}
}
