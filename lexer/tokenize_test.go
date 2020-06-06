package lexer_test

import (
	"errors"
	"ok/ast"
	"ok/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected []lexer.Token
		comments []*ast.Comment
		err      error
	}{
		"empty": {
			str: "",
			expected: []lexer.Token{
				{lexer.TokenEOF, ""},
			},
		},
		"func": {
			str: "func",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
				{lexer.TokenEOF, ""},
			},
		},
		"main": {
			str: "main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main"},
				{lexer.TokenEOF, ""},
			},
		},
		"Main": {
			str: "Main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "Main"},
				{lexer.TokenEOF, ""},
			},
		},
		"space": {
			str: " ",
			expected: []lexer.Token{
				{lexer.TokenEOF, ""},
			},
		},
		"func-main": {
			str: "func main",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
				{lexer.TokenIdentifier, "main"},
				{lexer.TokenEOF, ""},
			},
		},
		"paren-open": {
			str: "(",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "("},
				{lexer.TokenEOF, ""},
			},
		},
		"paren-close": {
			str: ")",
			expected: []lexer.Token{
				{lexer.TokenParenClose, ")"},
				{lexer.TokenEOF, ""},
			},
		},
		"curly-open": {
			str: "{",
			expected: []lexer.Token{
				{lexer.TokenCurlyOpen, "{"},
				{lexer.TokenEOF, ""},
			},
		},
		"curly-close": {
			str: "}",
			expected: []lexer.Token{
				{lexer.TokenCurlyClose, "}"},
				{lexer.TokenEOF, ""},
			},
		},
		"newline": {
			str: "\n",
			expected: []lexer.Token{
				{lexer.TokenEOF, ""},
			},
		},
		"paren-open-close": {
			str: "()",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "("},
				{lexer.TokenParenClose, ")"},
				{lexer.TokenEOF, ""},
			},
		},
		"main-open-param": {
			str: "main(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main"},
				{lexer.TokenParenOpen, "("},
				{lexer.TokenEOF, ""},
			},
		},
		"func-open-curly": {
			str: "func{",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
				{lexer.TokenCurlyOpen, "{"},
				{lexer.TokenEOF, ""},
			},
		},
		"func-space-open-curly": {
			str: "func {",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
				{lexer.TokenCurlyOpen, "{"},
				{lexer.TokenEOF, ""},
			},
		},
		"unterminated-string": {
			str: `"`,
			err: errors.New("unterminated string"),
		},
		"func-unterminated-string": {
			str: `func "`,
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
			},
			err: errors.New("unterminated string"),
		},
		"empty-string": {
			str: `""`,
			expected: []lexer.Token{
				{lexer.TokenString, ""},
				{lexer.TokenEOF, ""},
			},
		},
		"non-empty-string": {
			str: `"foo"`,
			expected: []lexer.Token{
				{lexer.TokenString, "foo"},
				{lexer.TokenEOF, ""},
			},
		},
		"non-empty-unterminated-string": {
			str: `"foo`,
			err: errors.New("unterminated string"),
		},
		"empty-comment": {
			str: `//`,
			expected: []lexer.Token{
				{lexer.TokenComment, ""},
				{lexer.TokenEOF, ""},
			},
			comments: []*ast.Comment{
				{},
			},
		},
		"single-comment": {
			str: `// hi`,
			expected: []lexer.Token{
				{lexer.TokenComment, " hi"},
				{lexer.TokenEOF, ""},
			},
			comments: []*ast.Comment{
				{Comment: " hi"},
			},
		},
		"comment-after-code": {
			str: `word //hello`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "word"},
				{lexer.TokenComment, "hello"},
				{lexer.TokenEOF, ""},
			},
			comments: []*ast.Comment{
				{Comment: "hello"},
			},
		},
		"code-after-comment": {
			str: "// hello\nworld",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello"},
				{lexer.TokenIdentifier, "world"},
				{lexer.TokenEOF, ""},
			},
			comments: []*ast.Comment{
				{Comment: " hello"},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			options := lexer.Options{
				IncludeComments: true,
			}
			actual, comments, err := lexer.TokenizeString(test.str, options)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.comments, comments)
		})
	}
}
