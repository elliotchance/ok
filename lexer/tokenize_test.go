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
				{lexer.TokenEOF, "", false},
			},
		},
		"func": {
			str: "func",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"main": {
			str: "main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"Main": {
			str: "Main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "Main", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"space": {
			str: " ",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false},
			},
		},
		"func-main": {
			str: "func main",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false},
				{lexer.TokenIdentifier, "main", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"paren-open": {
			str: "(",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"paren-close": {
			str: ")",
			expected: []lexer.Token{
				{lexer.TokenParenClose, ")", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"curly-open": {
			str: "{",
			expected: []lexer.Token{
				{lexer.TokenCurlyOpen, "{", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"curly-close": {
			str: "}",
			expected: []lexer.Token{
				{lexer.TokenCurlyClose, "}", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"newline": {
			str: "\n",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false},
			},
		},
		"paren-open-close": {
			str: "()",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenParenClose, ")", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"main-open-param": {
			str: "main(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main", false},
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"func-open-curly": {
			str: "func{",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false},
				{lexer.TokenCurlyOpen, "{", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"func-space-open-curly": {
			str: "func {",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false},
				{lexer.TokenCurlyOpen, "{", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"unterminated-string": {
			str: `"`,
			err: errors.New("unterminated literal, did not find closing \""),
		},
		"func-unterminated-string": {
			str: `func "`,
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false},
			},
			err: errors.New("unterminated literal, did not find closing \""),
		},
		"empty-string": {
			str: `""`,
			expected: []lexer.Token{
				{lexer.TokenString, "", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"non-empty-string": {
			str: `"foo"`,
			expected: []lexer.Token{
				{lexer.TokenString, "foo", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"non-empty-unterminated-string": {
			str: `"foo`,
			err: errors.New("unterminated literal, did not find closing \""),
		},
		"empty-comment": {
			str: `//`,
			expected: []lexer.Token{
				{lexer.TokenComment, "", false},
				{lexer.TokenEOF, "", false},
			},
			comments: []*ast.Comment{
				{},
			},
		},
		"single-comment": {
			str: `// hi`,
			expected: []lexer.Token{
				{lexer.TokenComment, " hi", false},
				{lexer.TokenEOF, "", false},
			},
			comments: []*ast.Comment{
				{Comment: " hi"},
			},
		},
		"comment-after-code": {
			str: `word //hello`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "word", false},
				{lexer.TokenComment, "hello", false},
				{lexer.TokenEOF, "", false},
			},
			comments: []*ast.Comment{
				{Comment: "hello"},
			},
		},
		"code-after-comment": {
			str: "// hello\nworld",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello", false},
				{lexer.TokenIdentifier, "world", false},
				{lexer.TokenEOF, "", false},
			},
			comments: []*ast.Comment{
				{Comment: " hello"},
			},
		},
		"number-0": {
			str: `0`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-1": {
			str: `1`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "1", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-2": {
			str: `2`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "2", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-3": {
			str: `3`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "3", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-4": {
			str: `4`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "4", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-5": {
			str: `5`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "5", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-6": {
			str: `6`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "6", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-7": {
			str: `7`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "7", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-8": {
			str: `8`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "8", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-9": {
			str: `9`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "9", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-10": {
			str: `10`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "10", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-7676673479634790000000000": {
			str: `7676673479634790000000000`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "7676673479634790000000000", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-0045": {
			str: `0045`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0045", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-0.12": {
			str: `0.12`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0.12", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-0.1200": {
			str: `0.1200`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0.1200", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"true": {
			str: `true`,
			expected: []lexer.Token{
				{lexer.TokenBool, "true", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"false": {
			str: `false`,
			expected: []lexer.Token{
				{lexer.TokenBool, "false", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"ascii-char": {
			str: `'a'`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, "a", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"zero-length-char": {
			str: `''`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, "", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"multi-byte-char": {
			str: `'😃'`,
			expected: []lexer.Token{
				{lexer.TokenCharacter, "😃", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"unterminated-char": {
			str: `'a`,
			err: errors.New("unterminated literal, did not find closing '"),
		},
		"ascii-data": {
			str: "`a`",
			expected: []lexer.Token{
				{lexer.TokenData, "a", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"zero-length-data": {
			str: "``",
			expected: []lexer.Token{
				{lexer.TokenData, "", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"multi-byte-data": {
			str: "`😃`",
			expected: []lexer.Token{
				{lexer.TokenData, "😃", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"unterminated-data": {
			str: "`a",
			err: errors.New("unterminated literal, did not find closing `"),
		},
		"number-open": {
			str: `0(`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "0", false},
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"plus": {
			str: `+`,
			expected: []lexer.Token{
				{lexer.TokenPlus, "+", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"minus": {
			str: `-`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"times": {
			str: `*`,
			expected: []lexer.Token{
				{lexer.TokenTimes, "*", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"divide": {
			str: `/`,
			expected: []lexer.Token{
				{lexer.TokenDivide, "/", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"remainder": {
			str: `%`,
			expected: []lexer.Token{
				{lexer.TokenRemainder, "%", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-negative-integer": {
			str: `-3`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false},
				{lexer.TokenNumber, "3", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"number-negative-decimal": {
			str: `-3.20`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false},
				{lexer.TokenNumber, "3.20", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"and": {
			str: `and`,
			expected: []lexer.Token{
				{lexer.TokenAnd, "and", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"or": {
			str: `or`,
			expected: []lexer.Token{
				{lexer.TokenOr, "or", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"not": {
			str: `not`,
			expected: []lexer.Token{
				{lexer.TokenNot, "not", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"==": {
			str: `==`,
			expected: []lexer.Token{
				{lexer.TokenEqual, "==", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"!=": {
			str: `!=`,
			expected: []lexer.Token{
				{lexer.TokenNotEqual, "!=", false},
				{lexer.TokenEOF, "", false},
			},
		},
		">": {
			str: `>`,
			expected: []lexer.Token{
				{lexer.TokenGreaterThan, ">", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"<": {
			str: `<`,
			expected: []lexer.Token{
				{lexer.TokenLessThan, "<", false},
				{lexer.TokenEOF, "", false},
			},
		},
		">=": {
			str: `>=`,
			expected: []lexer.Token{
				{lexer.TokenGreaterThanEqual, ">=", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"<=": {
			str: `<=`,
			expected: []lexer.Token{
				{lexer.TokenLessThanEqual, "<=", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"bool==bool": {
			str: `true==false`,
			expected: []lexer.Token{
				{lexer.TokenBool, "true", false},
				{lexer.TokenEqual, "==", false},
				{lexer.TokenBool, "false", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"comma": {
			str: `,`,
			expected: []lexer.Token{
				{lexer.TokenComma, ",", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"assign": {
			str: `=`,
			expected: []lexer.Token{
				{lexer.TokenAssign, "=", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"identifier-with-underscore": {
			str: `foo_bar`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "foo_bar", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"identifier-with-digits": {
			str: `foo123bar`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "foo123bar", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"identifier-cannot-start-with-number": {
			str: `1foo`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "1", false},
				{lexer.TokenIdentifier, "foo", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"carriage-return": {
			str: "\r",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false},
			},
		},
		"tab": {
			str: "\t",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false},
			},
		},
		"is-end-of-line-1": {
			str: "a = 1\nprint(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false},
				{lexer.TokenAssign, "=", false},
				{lexer.TokenNumber, "1", true},
				{lexer.TokenIdentifier, "print", false},
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"is-end-of-line-2": {
			str: "a = 1 \n print(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false},
				{lexer.TokenAssign, "=", false},
				{lexer.TokenNumber, "1", true},
				{lexer.TokenIdentifier, "print", false},
				{lexer.TokenParenOpen, "(", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"is-end-of-line-3": {
			str: "a = true\n",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false},
				{lexer.TokenAssign, "=", false},
				{lexer.TokenBool, "true", true},
				{lexer.TokenEOF, "", false},
			},
		},
		"for": {
			str: `for`,
			expected: []lexer.Token{
				{lexer.TokenFor, "for", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"continue": {
			str: `continue`,
			expected: []lexer.Token{
				{lexer.TokenContinue, "continue", false},
				{lexer.TokenEOF, "", false},
			},
		},
		"break": {
			str: `break`,
			expected: []lexer.Token{
				{lexer.TokenBreak, "break", false},
				{lexer.TokenEOF, "", false},
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
