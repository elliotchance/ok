package lexer_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/stretchr/testify/assert"
)

func pos(characterNumber int) lexer.Pos {
	return pos2(1, characterNumber)
}

func pos2(lineNumber, characterNumber int) lexer.Pos {
	return lexer.Pos{
		FileName:        "a.ok",
		LineNumber:      lineNumber,
		CharacterNumber: characterNumber,
	}
}

func TestTokenizeString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected []lexer.Token
		comments []*ast.Comment
		err      error
		options  *lexer.Options
	}{
		"empty": {
			str: "",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos(1)},
			},
		},
		"func": {
			str: "func",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"main": {
			str: "main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"Main": {
			str: "Main",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "Main", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"space": {
			str: " ",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"func-main": {
			str: "func main",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos(1)},
				{lexer.TokenIdentifier, "main", false, pos(6)},
				{lexer.TokenEOF, "", false, pos(10)},
			},
		},
		"paren-open": {
			str: "(",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "(", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"paren-close": {
			str: ")",
			expected: []lexer.Token{
				{lexer.TokenParenClose, ")", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"curly-open": {
			str: "{",
			expected: []lexer.Token{
				{lexer.TokenCurlyOpen, "{", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"curly-close": {
			str: "}",
			expected: []lexer.Token{
				{lexer.TokenCurlyClose, "}", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"newline": {
			str: "\n",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos2(2, 1)},
			},
		},
		"paren-open-close": {
			str: "()",
			expected: []lexer.Token{
				{lexer.TokenParenOpen, "(", false, pos(1)},
				{lexer.TokenParenClose, ")", false, pos(2)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"main-open-param": {
			str: "main(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "main", false, pos(1)},
				{lexer.TokenParenOpen, "(", false, pos(5)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"func-open-curly": {
			str: "func{",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos(1)},
				{lexer.TokenCurlyOpen, "{", false, pos(5)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"func-space-open-curly": {
			str: "func {",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos(1)},
				{lexer.TokenCurlyOpen, "{", false, pos(6)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"unterminated-string": {
			str: `"`,
			err: errors.New("a.ok:1:1 unterminated literal, did not find closing \""),
		},
		"func-unterminated-string": {
			str: `func "`,
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos(1)},
			},
			err: errors.New("a.ok:1:6 unterminated literal, did not find closing \""),
		},
		"empty-string": {
			str: `""`,
			expected: []lexer.Token{
				{lexer.TokenStringLiteral, "", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"non-empty-string": {
			str: `"foo"`,
			expected: []lexer.Token{
				{lexer.TokenStringLiteral, "foo", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"non-empty-unterminated-string": {
			str: `"foo`,
			err: errors.New("a.ok:1:1 unterminated literal, did not find closing \""),
		},
		"empty-comment": {
			str: `//`,
			expected: []lexer.Token{
				{lexer.TokenComment, "", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
			comments: []*ast.Comment{
				{Pos: "a.ok:1:1"},
			},
		},
		"single-comment": {
			str: `// hi`,
			expected: []lexer.Token{
				{lexer.TokenComment, " hi", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
			comments: []*ast.Comment{
				{Comment: " hi", Pos: "a.ok:1:1"},
			},
		},
		"comment-after-code": {
			str: `word //hello`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "word", false, pos(1)},
				{lexer.TokenComment, "hello", false, pos(6)},
				{lexer.TokenEOF, "", false, pos(13)},
			},
			comments: []*ast.Comment{
				{Comment: "hello", Pos: "a.ok:1:6"},
			},
		},
		"comment-after-code-2": {
			str: `word //hello`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "word", true, pos(1)},
				{lexer.TokenEOF, "", false, pos(13)},
			},
			comments: []*ast.Comment{
				{Comment: "hello", Pos: "a.ok:1:6"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"code-after-comment": {
			str: "// hello\nworld",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello", false, pos(1)},
				{lexer.TokenIdentifier, "world", false, pos2(2, 1)},
				{lexer.TokenEOF, "", false, pos2(2, 6)},
			},
			comments: []*ast.Comment{
				{Comment: " hello", Pos: "a.ok:1:1"},
			},
		},
		"multiline-comment": {
			str: "// hello\n// world",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello\n world", false, pos(1)},
				{lexer.TokenEOF, "", false, pos2(2, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello\n world", Pos: "a.ok:1:1"},
			},
		},
		"multiline-comments-joined": {
			str: "// hello\n//\n// world",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello\n\n world", false, pos(1)},
				{lexer.TokenEOF, "", false, pos2(3, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello\n\n world", Pos: "a.ok:1:1"},
			},
		},
		"multiline-comments-not-joined": {
			str: "// hello\n\n// world",
			expected: []lexer.Token{
				{lexer.TokenComment, " hello", true, pos(1)},
				{lexer.TokenComment, " world", false, pos2(3, 1)},
				{lexer.TokenEOF, "", false, pos2(3, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello", Pos: "a.ok:1:1"},
				{Comment: " world", Pos: "a.ok:3:1"},
			},
		},
		"comment-attached-to-func": {
			str: "// Foo is cool\nfunc Foo(",
			expected: []lexer.Token{
				{lexer.TokenComment, " Foo is cool", false, pos(1)},
				{lexer.TokenFunc, "func", false, pos2(2, 1)},
				{lexer.TokenIdentifier, "Foo", false, pos2(2, 6)},
				// The ( is important for the test because the last token in the
				// file will not be tested for attaching a comment to a
				// function.
				{lexer.TokenParenOpen, "(", false, pos2(2, 9)},
				{lexer.TokenEOF, "", false, pos2(2, 10)},
			},
			comments: []*ast.Comment{
				{Comment: " Foo is cool", Func: "Foo", Pos: "a.ok:1:1"},
			},
		},
		"comment-not-attached-to-func": {
			str: "// Foo is cool\n\nfunc Foo(",
			expected: []lexer.Token{
				{lexer.TokenComment, " Foo is cool", true, pos(1)},
				{lexer.TokenFunc, "func", false, pos2(3, 1)},
				{lexer.TokenIdentifier, "Foo", false, pos2(3, 6)},
				// The ( is important for the test because the last token in the
				// file will not be tested for attaching a comment to a
				// function.
				{lexer.TokenParenOpen, "(", false, pos2(3, 9)},
				{lexer.TokenEOF, "", false, pos2(3, 10)},
			},
			comments: []*ast.Comment{
				{Comment: " Foo is cool", Pos: "a.ok:1:1"},
			},
		},
		"comment-attached-to-func-2": {
			str: "// Foo is cool\nfunc Foo(",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos2(2, 1)},
				{lexer.TokenIdentifier, "Foo", false, pos2(2, 6)},
				// The ( is important for the test because the last token in the
				// file will not be tested for attaching a comment to a
				// function.
				{lexer.TokenParenOpen, "(", false, pos2(2, 9)},
				{lexer.TokenEOF, "", false, pos2(2, 10)},
			},
			comments: []*ast.Comment{
				{Comment: " Foo is cool", Func: "Foo", Pos: "a.ok:1:1"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"comment-not-attached-to-func-2": {
			str: "// Foo is cool\n\nfunc Foo(",
			expected: []lexer.Token{
				{lexer.TokenFunc, "func", false, pos2(3, 1)},
				{lexer.TokenIdentifier, "Foo", false, pos2(3, 6)},
				// The ( is important for the test because the last token in the
				// file will not be tested for attaching a comment to a
				// function.
				{lexer.TokenParenOpen, "(", false, pos2(3, 9)},
				{lexer.TokenEOF, "", false, pos2(3, 10)},
			},
			comments: []*ast.Comment{
				{Comment: " Foo is cool", Pos: "a.ok:1:1"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"multiline-comments-joined-2": {
			str: "// hello\n//\n// world",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos2(3, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello\n\n world", Pos: "a.ok:1:1"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"tokens-between-comments": {
			str: "// hello\na\n// world",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", true, pos2(2, 1)},
				{lexer.TokenEOF, "", false, pos2(3, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello", Pos: "a.ok:1:1"},
				{Comment: " world", Pos: "a.ok:3:1"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"operator-between-comments": {
			str: "// hello\n}\n// world",
			expected: []lexer.Token{
				{lexer.TokenCurlyClose, "}", true, pos2(2, 1)},
				{lexer.TokenEOF, "", false, pos2(3, 9)},
			},
			comments: []*ast.Comment{
				{Comment: " hello", Pos: "a.ok:1:1"},
				{Comment: " world", Pos: "a.ok:3:1"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"operator-between-comments-2": {
			str: "// hello\n}// world",
			expected: []lexer.Token{
				{lexer.TokenCurlyClose, "}", true, pos2(2, 1)},
				{lexer.TokenEOF, "", false, pos2(2, 10)},
			},
			comments: []*ast.Comment{
				{Comment: " hello", Pos: "a.ok:1:1"},
				{Comment: " world", Pos: "a.ok:2:2"},
			},
			options: &lexer.Options{
				IncludeComments: false,
			},
		},
		"number-0": {
			str: `0`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "0", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-1": {
			str: `1`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "1", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-2": {
			str: `2`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "2", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-3": {
			str: `3`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "3", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-4": {
			str: `4`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "4", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-5": {
			str: `5`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "5", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-6": {
			str: `6`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "6", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-7": {
			str: `7`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "7", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-8": {
			str: `8`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "8", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-9": {
			str: `9`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "9", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-10": {
			str: `10`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "10", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"number-7676673479634790000000000": {
			str: `7676673479634790000000000`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "7676673479634790000000000", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(26)},
			},
		},
		"number-0045": {
			str: `0045`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "0045", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"number-0.12": {
			str: `0.12`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "0.12", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"number-0.1200": {
			str: `0.1200`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "0.1200", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"true": {
			str: `true`,
			expected: []lexer.Token{
				{lexer.TokenBoolLiteral, "true", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"false": {
			str: `false`,
			expected: []lexer.Token{
				{lexer.TokenBoolLiteral, "false", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"ascii-char": {
			str: `'a'`,
			expected: []lexer.Token{
				{lexer.TokenCharLiteral, "a", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"zero-length-char": {
			str: `''`,
			expected: []lexer.Token{
				{lexer.TokenCharLiteral, "", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"multi-byte-char": {
			str: `'😃'`,
			expected: []lexer.Token{
				{lexer.TokenCharLiteral, "😃", false, pos(1)},

				// Note that 4 for the character position counts characters, not
				// bytes.
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"unterminated-char": {
			str: `'a`,
			err: errors.New("a.ok:1:1 unterminated literal, did not find closing '"),
		},
		"ascii-data": {
			str: "`a`",
			expected: []lexer.Token{
				{lexer.TokenDataLiteral, "a", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"zero-length-data": {
			str: "``",
			expected: []lexer.Token{
				{lexer.TokenDataLiteral, "", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"multi-byte-data": {
			str: "`😃`",
			expected: []lexer.Token{
				{lexer.TokenDataLiteral, "😃", false, pos(1)},

				// Note that 4 for the character position counts characters, not
				// bytes. Even though we are dealing with a data literal here.
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"unterminated-data": {
			str: "`a",
			err: errors.New("a.ok:1:1 unterminated literal, did not find closing `"),
		},
		"number-open": {
			str: `0(`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "0", false, pos(1)},
				{lexer.TokenParenOpen, "(", false, pos(2)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"plus": {
			str: `+`,
			expected: []lexer.Token{
				{lexer.TokenPlus, "+", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"minus": {
			str: `-`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"times": {
			str: `*`,
			expected: []lexer.Token{
				{lexer.TokenTimes, "*", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"divide": {
			str: `/`,
			expected: []lexer.Token{
				{lexer.TokenDivide, "/", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"remainder": {
			str: `%`,
			expected: []lexer.Token{
				{lexer.TokenRemainder, "%", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"number-negative-integer": {
			str: `-3`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false, pos(1)},
				{lexer.TokenNumberLiteral, "3", false, pos(2)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"number-negative-decimal": {
			str: `-3.20`,
			expected: []lexer.Token{
				{lexer.TokenMinus, "-", false, pos(1)},
				{lexer.TokenNumberLiteral, "3.20", false, pos(2)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"and": {
			str: `and`,
			expected: []lexer.Token{
				{lexer.TokenAnd, "and", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"or": {
			str: `or`,
			expected: []lexer.Token{
				{lexer.TokenOr, "or", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"not": {
			str: `not`,
			expected: []lexer.Token{
				{lexer.TokenNot, "not", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"==": {
			str: `==`,
			expected: []lexer.Token{
				{lexer.TokenEqual, "==", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"!=": {
			str: `!=`,
			expected: []lexer.Token{
				{lexer.TokenNotEqual, "!=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		">": {
			str: `>`,
			expected: []lexer.Token{
				{lexer.TokenGreaterThan, ">", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"<": {
			str: `<`,
			expected: []lexer.Token{
				{lexer.TokenLessThan, "<", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		">=": {
			str: `>=`,
			expected: []lexer.Token{
				{lexer.TokenGreaterThanEqual, ">=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"<=": {
			str: `<=`,
			expected: []lexer.Token{
				{lexer.TokenLessThanEqual, "<=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"bool=bool": {
			str: `true=false`,
			expected: []lexer.Token{
				{lexer.TokenBoolLiteral, "true", false, pos(1)},
				{lexer.TokenAssign, "=", false, pos(5)},
				{lexer.TokenBoolLiteral, "false", false, pos(6)},
				{lexer.TokenEOF, "", false, pos(11)},
			},
		},
		"bool==bool": {
			str: `true==false`,
			expected: []lexer.Token{
				{lexer.TokenBoolLiteral, "true", false, pos(1)},
				{lexer.TokenEqual, "==", false, pos(5)},
				{lexer.TokenBoolLiteral, "false", false, pos(7)},
				{lexer.TokenEOF, "", false, pos(12)},
			},
		},
		"comma": {
			str: `,`,
			expected: []lexer.Token{
				{lexer.TokenComma, ",", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"assign": {
			str: `=`,
			expected: []lexer.Token{
				{lexer.TokenAssign, "=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"identifier-with-underscore": {
			str: `foo_bar`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "foo_bar", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(8)},
			},
		},
		"identifier-with-digits": {
			str: `foo123bar`,
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "foo123bar", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(10)},
			},
		},
		"identifier-cannot-start-with-number": {
			str: `1foo`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "1", false, pos(1)},
				{lexer.TokenIdentifier, "foo", false, pos(2)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"carriage-return": {
			str: "\r",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos2(2, 1)},
			},
		},
		"tab": {
			str: "\t",
			expected: []lexer.Token{
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"is-end-of-line-1": {
			str: "a = 1\nprint(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false, pos(1)},
				{lexer.TokenAssign, "=", false, pos(3)},
				{lexer.TokenNumberLiteral, "1", true, pos(5)},
				{lexer.TokenIdentifier, "print", false, pos2(2, 1)},
				{lexer.TokenParenOpen, "(", false, pos2(2, 6)},
				{lexer.TokenEOF, "", false, pos2(2, 7)},
			},
		},
		"is-end-of-line-2": {
			str: "a = 1 \n print(",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false, pos(1)},
				{lexer.TokenAssign, "=", false, pos(3)},
				{lexer.TokenNumberLiteral, "1", true, pos(5)},
				{lexer.TokenIdentifier, "print", false, pos2(2, 2)},
				{lexer.TokenParenOpen, "(", false, pos2(2, 7)},
				{lexer.TokenEOF, "", false, pos2(2, 8)},
			},
		},
		"is-end-of-line-3": {
			str: "a = true\n",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "a", false, pos(1)},
				{lexer.TokenAssign, "=", false, pos(3)},
				{lexer.TokenBoolLiteral, "true", true, pos(5)},
				{lexer.TokenEOF, "", false, pos2(2, 1)},
			},
		},
		"for": {
			str: `for`,
			expected: []lexer.Token{
				{lexer.TokenFor, "for", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"continue": {
			str: `continue`,
			expected: []lexer.Token{
				{lexer.TokenContinue, "continue", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(9)},
			},
		},
		"break": {
			str: `break`,
			expected: []lexer.Token{
				{lexer.TokenBreak, "break", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"++": {
			str: `++`,
			expected: []lexer.Token{
				{lexer.TokenIncrement, "++", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"--": {
			str: `--`,
			expected: []lexer.Token{
				{lexer.TokenDecrement, "--", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"+=": {
			str: `+=`,
			expected: []lexer.Token{
				{lexer.TokenPlusAssign, "+=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"-=": {
			str: `-=`,
			expected: []lexer.Token{
				{lexer.TokenMinusAssign, "-=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"*=": {
			str: `*=`,
			expected: []lexer.Token{
				{lexer.TokenTimesAssign, "*=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"/=": {
			str: `/=`,
			expected: []lexer.Token{
				{lexer.TokenDivideAssign, "/=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"%=": {
			str: `%=`,
			expected: []lexer.Token{
				{lexer.TokenRemainderAssign, "%=", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"2/3": {
			str: `2/3`,
			expected: []lexer.Token{
				{lexer.TokenNumberLiteral, "2", false, pos(1)},
				{lexer.TokenDivide, "/", false, pos(2)},
				{lexer.TokenNumberLiteral, "3", false, pos(3)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"if": {
			str: `if`,
			expected: []lexer.Token{
				{lexer.TokenIf, "if", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"else": {
			str: `else`,
			expected: []lexer.Token{
				{lexer.TokenElse, "else", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		";": {
			str: `;`,
			expected: []lexer.Token{
				{lexer.TokenSemiColon, ";", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"switch": {
			str: `switch`,
			expected: []lexer.Token{
				{lexer.TokenSwitch, "switch", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"case": {
			str: `case`,
			expected: []lexer.Token{
				{lexer.TokenCase, "case", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"[": {
			str: `[`,
			expected: []lexer.Token{
				{lexer.TokenSquareOpen, "[", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"]": {
			str: `]`,
			expected: []lexer.Token{
				{lexer.TokenSquareClose, "]", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"any": {
			str: `any`,
			expected: []lexer.Token{
				{lexer.TokenAny, "any", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"bool": {
			str: `bool`,
			expected: []lexer.Token{
				{lexer.TokenBool, "bool", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"char": {
			str: `char`,
			expected: []lexer.Token{
				{lexer.TokenChar, "char", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"data": {
			str: `data`,
			expected: []lexer.Token{
				{lexer.TokenData, "data", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"number": {
			str: `number`,
			expected: []lexer.Token{
				{lexer.TokenNumber, "number", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"string": {
			str: `string`,
			expected: []lexer.Token{
				{lexer.TokenString, "string", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		":": {
			str: `:`,
			expected: []lexer.Token{
				{lexer.TokenColon, ":", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"in": {
			str: `in`,
			expected: []lexer.Token{
				{lexer.TokenIn, "in", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
		"return": {
			str: `return`,
			expected: []lexer.Token{
				{lexer.TokenReturn, "return", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"test": {
			str: `test`,
			expected: []lexer.Token{
				{lexer.TokenTest, "test", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(5)},
			},
		},
		"assert": {
			str: `assert`,
			expected: []lexer.Token{
				{lexer.TokenAssert, "assert", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		"import": {
			str: `import`,
			expected: []lexer.Token{
				{lexer.TokenImport, "import", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(7)},
			},
		},
		".": {
			str: `.`,
			expected: []lexer.Token{
				{lexer.TokenDot, ".", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(2)},
			},
		},
		"string-literal-newline": {
			str: "\"math\"\n",
			expected: []lexer.Token{
				{lexer.TokenStringLiteral, "math", true, pos(1)},
				{lexer.TokenEOF, "", false, pos2(2, 1)},
			},
		},
		"string-literal-newline-func": {
			str: "\"math\"\nfunc",
			expected: []lexer.Token{
				{lexer.TokenStringLiteral, "math", true, pos(1)},
				{lexer.TokenFunc, "func", false, pos2(2, 1)},
				{lexer.TokenEOF, "", false, pos2(2, 5)},
			},
		},
		"import-func": {
			str: "import \"math\"\nfunc foo",
			expected: []lexer.Token{
				{lexer.TokenImport, "import", false, pos(1)},
				{lexer.TokenStringLiteral, "math", true, pos(8)},
				{lexer.TokenFunc, "func", false, pos2(2, 1)},
				{lexer.TokenIdentifier, "foo", false, pos2(2, 6)},
				{lexer.TokenEOF, "", false, pos2(2, 9)},
			},
		},
		"Abs": {
			str: "Abs",
			expected: []lexer.Token{
				{lexer.TokenIdentifier, "Abs", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"interpolate-1": {
			str: `"{name}"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(2)},
				{lexer.TokenIdentifier, "name", false, pos(3)},
				{lexer.TokenParenClose, ")", false, pos(7)},
				{lexer.TokenInterpolateEnd, "", false, pos(8)},
				{lexer.TokenEOF, "", false, pos(9)},
			},
		},
		"interpolate-2": {
			str: `"hi {name}"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenStringLiteral, "hi ", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(5)},
				{lexer.TokenIdentifier, "name", false, pos(6)},
				{lexer.TokenParenClose, ")", false, pos(10)},
				{lexer.TokenInterpolateEnd, "", false, pos(11)},
				{lexer.TokenEOF, "", false, pos(12)},
			},
		},
		"interpolate-3": {
			str: `"{name}foo"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(2)},
				{lexer.TokenIdentifier, "name", false, pos(3)},
				{lexer.TokenParenClose, ")", false, pos(7)},
				{lexer.TokenStringLiteral, "foo", false, pos(8)},
				{lexer.TokenInterpolateEnd, "", false, pos(11)},
				{lexer.TokenEOF, "", false, pos(12)},
			},
		},
		"interpolate-4": {
			str: `"foo{bar}baz{qux}quux"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenStringLiteral, "foo", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(5)},
				{lexer.TokenIdentifier, "bar", false, pos(6)},
				{lexer.TokenParenClose, ")", false, pos(9)},
				{lexer.TokenStringLiteral, "baz", false, pos(10)},
				{lexer.TokenParenOpen, "(", false, pos(13)},
				{lexer.TokenIdentifier, "qux", false, pos(14)},
				{lexer.TokenParenClose, ")", false, pos(17)},
				{lexer.TokenStringLiteral, "quux", false, pos(18)},
				{lexer.TokenInterpolateEnd, "", false, pos(22)},
				{lexer.TokenEOF, "", false, pos(23)},
			},
		},
		"interpolate-plus": {
			str: `"hi {name}"+`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenStringLiteral, "hi ", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(5)},
				{lexer.TokenIdentifier, "name", false, pos(6)},
				{lexer.TokenParenClose, ")", false, pos(10)},
				{lexer.TokenInterpolateEnd, "", false, pos(11)},
				{lexer.TokenPlus, "+", false, pos(12)},
				{lexer.TokenEOF, "", false, pos(13)},
			},
		},
		"interpolate-expr-1": {
			str: `"{1}"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(2)},
				{lexer.TokenNumberLiteral, "1", false, pos(3)},
				{lexer.TokenParenClose, ")", false, pos(4)},
				{lexer.TokenInterpolateEnd, "", false, pos(5)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"interpolate-expr-2": {
			str: `"{1+2}"`,
			expected: []lexer.Token{
				{lexer.TokenInterpolateStart, "", false, pos(2)},
				{lexer.TokenParenOpen, "(", false, pos(2)},
				{lexer.TokenNumberLiteral, "1", false, pos(3)},
				{lexer.TokenPlus, "+", false, pos(4)},
				{lexer.TokenNumberLiteral, "2", false, pos(5)},
				{lexer.TokenParenClose, ")", false, pos(6)},
				{lexer.TokenInterpolateEnd, "", false, pos(7)},
				{lexer.TokenEOF, "", false, pos(8)},
			},
		},
		"try": {
			str: `try`,
			expected: []lexer.Token{
				{lexer.TokenTry, "try", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(4)},
			},
		},
		"raise": {
			str: `raise`,
			expected: []lexer.Token{
				{lexer.TokenRaise, "raise", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(6)},
			},
		},
		"on": {
			str: `on`,
			expected: []lexer.Token{
				{lexer.TokenOn, "on", false, pos(1)},
				{lexer.TokenEOF, "", false, pos(3)},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			options := lexer.Options{
				IncludeComments: true,
			}
			if test.options != nil {
				options = *test.options
			}

			actual, comments, err := lexer.TokenizeString(test.str, options, "a.ok")
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
