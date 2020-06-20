package parser

import (
	"ok/lexer"
	"strings"
)

func consumeType(parser *Parser, offset int) (string, int, error) {
	originalOffset := offset
	ty := ""

	for {
		switch {
		case parser.File.Tokens[offset].Kind == lexer.TokenSquareOpen &&
			parser.File.Tokens[offset+1].Kind == lexer.TokenSquareClose:
			offset += 2
			ty += "[]"

		case parser.File.Tokens[offset].Kind == lexer.TokenCurlyOpen &&
			parser.File.Tokens[offset+1].Kind == lexer.TokenCurlyClose:
			offset += 2
			ty += "{}"

		default:
			goto done
		}
	}
done:

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

func consumeTypes(parser *Parser, offset int) ([]string, int, error) {
	originalOffset := offset
	var types []string
	var err error

	if parser.File.Tokens[offset].Kind == lexer.TokenParenOpen {
		offset++ // skip "("

		for {
			var ty string
			ty, offset, err = consumeType(parser, offset)
			if err != nil {
				return nil, originalOffset, err
			}

			types = append(types, ty)

			if parser.File.Tokens[offset].Kind != lexer.TokenComma {
				break
			} else {
				offset++ // skip ","
			}
		}

		offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
		if err != nil {
			return nil, originalOffset, err
		}
	} else {
		var ty string
		ty, offset, err = consumeType(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		types = []string{ty}
	}

	return types, offset, nil
}
