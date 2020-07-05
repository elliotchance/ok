package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeInterpolate(parser *Parser, offset int) (*ast.Interpolate, int, error) {
	var err error
	offset, err = consume(parser.File, offset, []string{lexer.TokenInterpolateStart})
	if err != nil {
		return nil, offset, err
	}

	interpolate := &ast.Interpolate{}

	for {
		if parser.File.Tokens[offset].Kind == lexer.TokenInterpolateEnd {
			break
		}

		var group *ast.Group
		group, offset, err = consumeGroup(parser, offset)
		if err == nil {
			interpolate.Parts = append(interpolate.Parts, group)
			continue
		}

		// TODO(elliot): This is a bit lazy, because we can only find a string
		//  literal here.
		var literal *ast.Literal
		literal, offset, err = consumeLiteral(parser, offset)
		if err == nil {
			interpolate.Parts = append(interpolate.Parts, literal)
			continue
		}
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenInterpolateEnd})
	if err != nil {
		return nil, offset, err
	}

	return interpolate, offset, nil
}
