package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeKeyValues(parser *Parser, offset int) ([]*ast.KeyValue, int, error) {
	var keyValues []*ast.KeyValue

	for {
		var keyValue *ast.KeyValue
		var err error
		keyValue, offset, err = consumeKeyValue(parser, offset)
		if err != nil {
			return nil, offset, err
		}

		keyValues = append(keyValues, keyValue)

		if parser.File.Tokens[offset].Kind == lexer.TokenComma {
			offset++
		} else {
			break
		}
	}

	return keyValues, offset, nil
}

func consumeKeyValue(parser *Parser, offset int) (*ast.KeyValue, int, error) {
	var err error
	node := &ast.KeyValue{}

	node.Key, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err != nil {
		return nil, offset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenColon})
	if err != nil {
		return nil, offset, err
	}

	node.Value, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
