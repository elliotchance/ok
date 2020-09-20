package parser

import (
	"errors"

	"github.com/elliotchance/ok/lexer"
)

func consumeOneOf(p *Parser, offset int, tokens []string) (lexer.Token, int, error) {
	for _, ty := range tokens {
		if t := p.tokens[offset]; t.Kind == ty {
			return t, offset + 1, nil
		}
	}

	return lexer.Token{}, offset, errors.New("no token found")
}
