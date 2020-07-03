package parser_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"

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
			p := parser.ParseString(test.str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, []*ast.Test{test.expected}, p.File.Tests)
			assert.Len(t, p.File.Funcs, 0)
		})
	}
}
