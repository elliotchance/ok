package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

func consumeFunc(parser *Parser, offset int) (_ *ast.Func, _ int, anon bool, finalErr error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser, offset, []string{lexer.TokenFunc})
	if err != nil {
		return nil, originalOffset, anon, err
	}

	fn := &ast.Func{
		Pos:        parser.pos(originalOffset),
		UniqueName: parser.nextFunctionName(),
	}

	// The name of the function is optional. For anonymous functions we will
	// generate an internal name so they can be referenced.
	//
	// TODO(elliot): Anonymous functions should not be allowed just anywhere.

	offset, err = consume(parser, offset, []string{lexer.TokenIdentifier})
	if err == nil {
		fn.Name = parser.tokens[offset-1].Value
	} else {
		anon = true
	}

	offset, err = consume(parser, offset, []string{lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, anon, err
	}

	var args []*ast.Argument
	args, offset, err = consumeArguments(parser, offset)
	if err != nil {
		return nil, originalOffset, anon, err
	}
	fn.Arguments = args

	offset, err = consume(parser, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, anon, err
	}

	// The next "{" will be ambiguous because it could be a map or the start of
	// the block. Try to consume types first, and fallback to treating it as the
	// start of the block.
	fn.Returns, offset, err = consumeTypes(parser, offset, false)

	if parser.tokens[offset].Kind != lexer.TokenCurlyOpen && err != nil {
		fn.Returns, offset, err = consumeTypes(parser, offset, false)
		if err != nil {
			return nil, originalOffset, anon, err
		}
	}

	parser.functionNames = append(parser.functionNames, fn.Name)
	defer func() {
		parser.functionNames = parser.functionNames[:len(parser.functionNames)-1]
	}()

	fn.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, originalOffset, anon, err
	}

	return fn, offset, anon, nil
}

func consumeArguments(parser *Parser, offset int) ([]*ast.Argument, int, error) {
	originalOffset := offset

	// Catch no arguments.
	if parser.tokens[offset].Kind == lexer.TokenParenClose {
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

		if parser.tokens[offset].Kind != lexer.TokenComma {
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
		offset, err = consume(parser, offset, []string{lexer.TokenIdentifier})
		if err != nil {
			return nil, originalOffset, err
		}

		names = append(names, parser.tokens[offset-1].Value)

		if parser.tokens[offset].Kind != lexer.TokenComma {
			break
		} else {
			offset++ // skip ","
		}
	}

	var ty *types.Type
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
