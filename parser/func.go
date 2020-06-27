package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeFunc(parser *Parser, offset int) (*ast.Func, int, error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenFunc, lexer.TokenIdentifier,
		lexer.TokenParenOpen,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	fn := &ast.Func{
		Name: parser.File.Tokens[offset-2].Value,
	}

	var args []*ast.Argument
	args, offset, err = consumeArguments(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}
	fn.Arguments = args

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	if parser.File.Tokens[offset].Kind != lexer.TokenCurlyOpen {
		fn.Returns, offset, err = consumeTypes(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}
	}

	fn.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return fn, offset, nil
}

func consumeArguments(parser *Parser, offset int) ([]*ast.Argument, int, error) {
	originalOffset := offset

	// Catch no arguments.
	if parser.File.Tokens[offset].Kind == lexer.TokenParenClose {
		return nil, offset, nil
	}

	var args []*ast.Argument
	for {
		var args2 []*ast.Argument
		var err error
		args2, offset, err = consumeArgument(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		args = append(args, args2...)

		if parser.File.Tokens[offset].Kind != lexer.TokenComma {
			break
		} else {
			offset++ // skip ","
		}
	}

	return args, offset, nil
}

func consumeArgument(parser *Parser, offset int) ([]*ast.Argument, int, error) {
	originalOffset := offset
	var err error

	var names []string
	for {
		offset, err = consume(parser.File, offset, []string{lexer.TokenIdentifier})
		if err != nil {
			return nil, originalOffset, err
		}

		names = append(names, parser.File.Tokens[offset-1].Value)

		if parser.File.Tokens[offset].Kind != lexer.TokenComma {
			break
		} else {
			offset++ // skip ","
		}
	}

	var ty string
	ty, offset, err = consumeType(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	var args []*ast.Argument
	for _, name := range names {
		args = append(args, &ast.Argument{
			Name: name,
			Type: ty,
		})
	}

	return args, offset, nil
}
