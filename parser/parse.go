package parser

import (
	"errors"
	"fmt"
	"ok/ast"
	"ok/lexer"
)

// ParseString parses source code and returns the AST for the function.
func ParseString(s string) (*ast.Func, error) {
	if s == "" {
		return nil, errors.New("cannot parse empty string")
	}

	tokens, err := lexer.TokenizeString(s)
	if err != nil {
		at := "at the start"
		if len(tokens) > 0 {
			at = "after " + tokens[len(tokens)-1].String()
		}

		return nil, fmt.Errorf("unterminated string found %s", at)
	}

	fn, offset, err := consumeFunc(tokens, 0)
	if err != nil {
		return nil, err
	}

	if tokens[offset].Kind != lexer.TokenEOF {
		return nil, fmt.Errorf("found extra %s at the end of the file",
			tokens[offset])
	}

	return fn, nil
}

func consume(tokens []lexer.Token, offset int, expected []string) (int, error) {
	for i, t := range expected {
		// Bail out if there are not enough tokens. We can't do this at the
		// start because we needed to check one of the tokens earlier on did not
		// match to return a potentially lower value of i.
		if offset+i >= len(tokens) {
			return offset, fmt.Errorf(
				"expecting %s after %s, but found end of file",
				expected[i], tokens[offset+i-1].Kind)
		}

		if t != tokens[offset+i].Kind {
			return offset, fmt.Errorf(
				"expecting %s after %s, but found %s",
				expected[i], tokens[offset+i-1].Kind, tokens[offset+i].Kind)
		}
	}

	return offset + len(expected), nil
}
