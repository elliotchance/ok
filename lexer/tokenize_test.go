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
			err: errors.New("unterminated literal, did not find closing \""),
		},
		"func-unterminated-string": {
			str: `func "`,
			expected: []lexer.Token{
				{lexer.TokenFunc, "func"},
			},
			err: errors.New("unterminated literal, did not find closing \""),
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
			err: errors.New("unterminated literal, did not find closing \""),
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
		"number-0": {
			str: `0`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-1": {
			str: `1`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "1"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-2": {
			str: `2`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "2"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-3": {
			str: `3`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "3"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-4": {
			str: `4`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "4"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-5": {
			str: `5`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "5"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-6": {
			str: `6`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "6"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-7": {
			str: `7`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "7"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-8": {
			str: `8`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "8"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-9": {
			str: `9`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "9"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-10": {
			str: `10`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "10"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-7676673479634790000000000": {
			str: `7676673479634790000000000`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "7676673479634790000000000"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-0045": {
			str: `0045`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0045"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-0.12": {
			str: `0.12`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0.12"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-0.1200": {
			str: `0.1200`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0.1200"},
				{lexer.TokenEOF, ""},
			},
		},
		"true": {
			str: `true`,
			expected: []lexer.Token{
				{lexer.TokenBool, "true"},
				{lexer.TokenEOF, ""},
			},
		},
		"false": {
			str: `false`,
			expected: []lexer.Token{
				{lexer.TokenBool, "false"},
				{lexer.TokenEOF, ""},
			},
		},
		"ascii-char": {
			str: `'a'`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, "a"},
				{lexer.TokenEOF, ""},
			},
		},
		"zero-length-char": {
			str: `''`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, ""},
				{lexer.TokenEOF, ""},
			},
		},
		"multi-byte-char": {
			str: `'😃'`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, "😃"},
				{lexer.TokenEOF, ""},
			},
		},
		"unterminated-char": {
			str: `'a`,
			err: errors.New("unterminated literal, did not find closing '"),
		},
		"ascii-data": {
			str: "`a`",
			expected: []lexer.Token{
				{lexer.TokenData, "a"},
				{lexer.TokenEOF, ""},
			},
		},
		"zero-length-data": {
			str: "``",
			expected: []lexer.Token{
				{lexer.TokenData, ""},
				{lexer.TokenEOF, ""},
			},
		},
		"multi-byte-data": {
			str: "`😃`",
			expected: []lexer.Token{
				{lexer.TokenData, "😃"},
				{lexer.TokenEOF, ""},
			},
		},
		"unterminated-data": {
			str: "`a",
			err: errors.New("unterminated literal, did not find closing `"),
		},
		"number-open": {
			str: `0(`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0"},
				{lexer.TokenParenOpen, "("},
				{lexer.TokenEOF, ""},
			},
		},
		"plus": {
			str: `+`,
			expected: []lexer.Token{
				{lexer.TokenPlus, "+"},
				{lexer.TokenEOF, ""},
			},
		},
		"minus": {
			str: `-`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-"},
				{lexer.TokenEOF, ""},
			},
		},
		"times": {
			str: `*`,
			expected: []lexer.Token{
				{lexer.TokenTimes, "*"},
				{lexer.TokenEOF, ""},
			},
		},
		"divide": {
			str: `/`,
			expected: []lexer.Token{
				{lexer.TokenDivide, "/"},
				{lexer.TokenEOF, ""},
			},
		},
		"remainder": {
			str: `%`,
			expected: []lexer.Token{
				{lexer.TokenRemainder, "%"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-negative-integer": {
			str: `-3`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-"},
				{lexer.TokenNumber, "3"},
				{lexer.TokenEOF, ""},
			},
		},
		"number-negative-decimal": {
			str: `-3.20`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-"},
				{lexer.TokenNumber, "3.20"},
				{lexer.TokenEOF, ""},
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
