package lexer_test

import (
	"errors"
	"ok/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected []lexer.Token
		err      error
	}{
		"empty": {"", []lexer.Token{
			{lexer.TokenEOF, ""},
		}, nil},
		"func": {"func", []lexer.Token{
			{lexer.TokenFunc, "func"},
			{lexer.TokenEOF, ""},
		}, nil},
		"main": {"main", []lexer.Token{
			{lexer.TokenIdentifier, "main"},
			{lexer.TokenEOF, ""},
		}, nil},
		"Main": {"Main", []lexer.Token{
			{lexer.TokenIdentifier, "Main"},
			{lexer.TokenEOF, ""},
		}, nil},
		"space": {" ", []lexer.Token{
			{lexer.TokenEOF, ""},
		}, nil},
		"func-main": {"func main", []lexer.Token{
			{lexer.TokenFunc, "func"},
			{lexer.TokenIdentifier, "main"},
			{lexer.TokenEOF, ""},
		}, nil},
		"paren-open": {"(", []lexer.Token{
			{lexer.TokenParenOpen, "("},
			{lexer.TokenEOF, ""},
		}, nil},
		"paren-close": {")", []lexer.Token{
			{lexer.TokenParenClose, ")"},
			{lexer.TokenEOF, ""},
		}, nil},
		"curly-open": {"{", []lexer.Token{
			{lexer.TokenCurlyOpen, "{"},
			{lexer.TokenEOF, ""},
		}, nil},
		"curly-close": {"}", []lexer.Token{
			{lexer.TokenCurlyClose, "}"},
			{lexer.TokenEOF, ""},
		}, nil},
		"newline": {"\n", []lexer.Token{
			{lexer.TokenEOF, ""},
		}, nil},
		"paren-open-close": {"()", []lexer.Token{
			{lexer.TokenParenOpen, "("},
			{lexer.TokenParenClose, ")"},
			{lexer.TokenEOF, ""},
		}, nil},
		"main-open-param": {"main(", []lexer.Token{
			{lexer.TokenIdentifier, "main"},
			{lexer.TokenParenOpen, "("},
			{lexer.TokenEOF, ""},
		}, nil},
		"func-open-curly": {"func{", []lexer.Token{
			{lexer.TokenFunc, "func"},
			{lexer.TokenCurlyOpen, "{"},
			{lexer.TokenEOF, ""},
		}, nil},
		"func-space-open-curly": {"func {", []lexer.Token{
			{lexer.TokenFunc, "func"},
			{lexer.TokenCurlyOpen, "{"},
			{lexer.TokenEOF, ""},
		}, nil},
		"unterminated-string": {
			`"`,
			nil,
			errors.New("unterminated string"),
		},
		"func-unterminated-string": {
			`func "`,
			[]lexer.Token{
				{lexer.TokenFunc, "func"},
			},
			errors.New("unterminated string"),
		},
		"empty-string": {`""`, []lexer.Token{
			{lexer.TokenString, ""},
			{lexer.TokenEOF, ""},
		}, nil},
		"non-empty-string": {`"foo"`, []lexer.Token{
			{lexer.TokenString, "foo"},
			{lexer.TokenEOF, ""},
		}, nil},
		"non-empty-unterminated-string": {
			`"foo`,
			nil,
			errors.New("unterminated string"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			actual, err := lexer.TokenizeString(test.str)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
