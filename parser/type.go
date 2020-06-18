package parser

import (
	"ok/lexer"
	"strings"
)

func consumeType(parser *Parser, offset int) (string, int, error) {
	originalOffset := offset
	ty := ""

	if parser.File.Tokens[offset].Kind == lexer.TokenSquareOpen &&
		parser.File.Tokens[offset+1].Kind == lexer.TokenSquareClose {
		offset += 2
		ty += "[]"
	}

	var err error
	var t lexer.Token
	t, offset, err = consumeOneOf(parser.File, offset, []string{
		lexer.TokenAny,
		lexer.TokenBool,
		lexer.TokenChar,
		lexer.TokenData,
		lexer.TokenNumber,
		lexer.TokenString,
	})
	if err != nil {
		return "", originalOffset, err
	}

	ty += strings.Split(t.Kind, " ")[0]

	return ty, offset, nil
}
