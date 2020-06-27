package lexer_test

import (
	"testing"

	"github.com/elliotchance/ok/lexer"
	"github.com/stretchr/testify/assert"
)

func TestToken_String(t *testing.T) {
	for testName, test := range map[string]struct {
		kind     string
		value    string
		expected string
	}{
		"EOF":         {lexer.TokenEOF, "", "end of file"},
		"paren-open":  {lexer.TokenParenOpen, "(", "'('"},
		"paren-close": {lexer.TokenParenClose, ")", "')'"},
		"identifier":  {lexer.TokenIdentifier, "main", `"main"`},
	} {
		t.Run(testName, func(t *testing.T) {
			actual := lexer.Token{Kind: test.kind, Value: test.value}.String()
			assert.Equal(t, test.expected, actual)
		})
	}
}
