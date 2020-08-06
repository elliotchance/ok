package parser

import (
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

var types = []string{
	lexer.TokenAny,
	lexer.TokenBool,
	lexer.TokenChar,
	lexer.TokenData,
	lexer.TokenNumber,
	lexer.TokenString,
}

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

	// A function?
	var err error
	offset, err = consume(parser.File, offset, []string{lexer.TokenFunc})
	if err == nil {
		fn := &ast.Func{}
		var tys []string
		tys, offset, err = consumeTypes(parser, offset, true)
		if err != nil {
			return "", originalOffset, err
		}

		for _, ty := range tys {
			fn.Arguments = append(fn.Arguments, &ast.Argument{Type: ty})
		}

		// Returns is optional, so don't error if it wasn't consumed.
		fn.Returns, offset, _ = consumeTypes(parser, offset, false)

		return ty + fn.String(), offset, nil
	}

	var t lexer.Token
	t, offset, err = consumeOneOf(parser.File, offset, types)
	if err != nil {
		// Any identifier is also valid.
		var ident *ast.Identifier
		ident, offset, err = consumeIdentifier(parser, offset)
		if err != nil {
			return "", originalOffset, err
		}

		t.Kind = ident.Name
	}

	ty += strings.Split(t.Kind, " ")[0]

	return ty, offset, nil
}

func consumeTypes(parser *Parser, offset int, allowEmpty bool) ([]string, int, error) {
	originalOffset := offset
	var types []string
	var err error

	if parser.File.Tokens[offset].Kind == lexer.TokenParenOpen {
		offset++ // skip "("

		if allowEmpty &&
			parser.File.Tokens[offset].Kind == lexer.TokenParenClose {
			goto close
		}

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

	close:
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
