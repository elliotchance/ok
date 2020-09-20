package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeInterpolate(parser *Parser, offset int) (*ast.Interpolate, int, error) {
	originalOffset := offset
	var err error
	offset, err = consume(parser, offset, []string{lexer.TokenInterpolateStart})
	if err != nil {
		return nil, offset, err
	}

	interpolate := &ast.Interpolate{
		// Needed to ensure a mangled interpolation will be reported correctly.
		Pos: parser.pos(originalOffset),
	}

	isMangled := false
	for {
		if parser.tokens[offset].Kind == lexer.TokenInterpolateEnd {
			break
		}

		// If the interpolation is mangled we need to consume all tokens until
		// TokenInterpolateEnd so that the calling function can continue
		// parsing.
		if isMangled {
			offset++
			continue
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

		// We reach here if the interpolation is empty (like "{}"), which is not
		// valid. However, we cannot return an error otherwise the caller will
		// think this patch was not an interpolation expression.
		//
		// TODO(elliot): To properly recover from this we need to consume until TokenInterpolateEnd
		parser.appendError(interpolate,
			"empty interpolation, perhaps missing an escape")
		isMangled = true
	}

	offset, err = consume(parser, offset, []string{lexer.TokenInterpolateEnd})
	if err != nil {
		return nil, offset, err
	}

	return interpolate, offset, nil
}
