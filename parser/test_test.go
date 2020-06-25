package parser_test

import (
	"ok/ast"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Test
		errs     []error
	}{
		"empty": {
			str: `test "foo bar" {}`,
			expected: &ast.Test{
				Name: "foo bar",
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str)

			assertEqualErrors(t, test.errs, p.Errors)
			assert.Equal(t, []*ast.Test{test.expected}, p.File.Tests)
			assert.Len(t, p.File.Funcs, 0)
		})
	}
}
