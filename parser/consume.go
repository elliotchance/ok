package parser

import (
	"errors"
	"ok/lexer"
)

func consumeOneOf(f *File, offset int, tokens []string) (lexer.Token, int, error) {
	for _, ty := range tokens {
		if t := f.Tokens[offset]; t.Kind == ty {
			return t, offset + 1, nil
		}
	}

	return lexer.Token{}, offset, errors.New("no token found")
}
