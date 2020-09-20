package parser

import (
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

var typeTokens = []string{
	lexer.TokenAny,
	lexer.TokenBool,
	lexer.TokenChar,
	lexer.TokenData,
	lexer.TokenNumber,
	lexer.TokenString,
}

func consumeType(parser *Parser, offset int) (ty *types.Type, _ int, _ error) {
	originalOffset := offset

	for {
		switch {
		case parser.tokens[offset].Kind == lexer.TokenSquareOpen &&
			parser.tokens[offset+1].Kind == lexer.TokenSquareClose:
			offset += 2
			defer func() {
				if ty != nil {
					ty = ty.ToArray()
				}
			}()

		case parser.tokens[offset].Kind == lexer.TokenCurlyOpen &&
			parser.tokens[offset+1].Kind == lexer.TokenCurlyClose:
			offset += 2
			defer func() {
				if ty != nil {
					ty = ty.ToMap()
				}
			}()

		default:
			goto done
		}
	}
done:

	// A function?
	var err error
	offset, err = consume(parser, offset, []string{lexer.TokenFunc})
	if err == nil {
		fn := &ast.Func{}
		var args []*types.Type
		args, offset, err = consumeTypes(parser, offset, true)
		if err != nil {
			return nil, originalOffset, err
		}

		for _, ty := range args {
			fn.Arguments = append(fn.Arguments, &ast.Argument{Type: ty})
		}

		// Returns is optional, so don't error if it wasn't consumed.
		fn.Returns, offset, _ = consumeTypes(parser, offset, false)

		return types.NewFunc(args, fn.Returns), offset, nil
	}

	var t lexer.Token
	t, offset, err = consumeOneOf(parser, offset, typeTokens)
	if err != nil {
		// Any identifier is also valid.
		var ident *ast.Identifier
		ident, offset, err = consumeIdentifier(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		t.Kind = ident.Name
	}

	return types.TypeFromString(strings.Split(t.Kind, " ")[0]), offset, nil
}

func consumeTypes(parser *Parser, offset int, allowEmpty bool) ([]*types.Type, int, error) {
	originalOffset := offset
	var tys []*types.Type
	var err error

	if parser.tokens[offset].Kind == lexer.TokenParenOpen {
		offset++ // skip "("

		if allowEmpty &&
			parser.tokens[offset].Kind == lexer.TokenParenClose {
			goto close
		}

		for {
			var ty *types.Type
			ty, offset, err = consumeType(parser, offset)
			if err != nil {
				return nil, originalOffset, err
			}

			tys = append(tys, ty)

			if parser.tokens[offset].Kind != lexer.TokenComma {
				break
			} else {
				offset++ // skip ","
			}
		}

	close:
		offset, err = consume(parser, offset, []string{lexer.TokenParenClose})
		if err != nil {
			return nil, originalOffset, err
		}
	} else {
		var ty *types.Type
		ty, offset, err = consumeType(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		tys = []*types.Type{ty}
	}

	return tys, offset, nil
}
