package lexer

import (
	"fmt"
	"ok/ast"
	"strings"
)

// TokenizeString returns a slice of tokens from the provided str.
func TokenizeString(str string, options Options) ([]Token, []*ast.Comment, error) {
	var tokens []Token
	var comments []*ast.Comment
	var word string

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		var token Token
		found := false

		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			token := Token{TokenNumber, ""}
			for ; i < len(runes) && isDecimalCharacter(runes[i]); i++ {
				token.Value += string(runes[i])
			}

			i--
			tokens = append(tokens, token)
			continue

		case '/':
			token.Kind = TokenComment
			i += 2
			for ; i < len(runes) && runes[i] != '\n'; i++ {
				token.Value += string(runes[i])
			}

			comments = append(comments, &ast.Comment{
				Comment: token.Value,
			})
			if options.IncludeComments {
				tokens = append(tokens, token)
			}
			continue

		case '\'', '"', '`':
			value, newI, err := readQuotedLiteral(runes, i, c)
			if err != nil {
				// It's important that we return the tokens up until now so that
				// the parser has to the context to say where the error
				// occurred.
				return tokens, comments, err
			}

			i = newI
			tokens = append(tokens, Token{tokenKindForQuote(c), value})
			continue

		case ' ', '\n':
			found = true

		case '(', ')', '{', '}':
			found = true
			token = Token{string(c), string(c)}
		}

		if found && word != "" {
			tokens = append(tokens, tokenWord(word))
			word = ""
		}

		if token.Kind != "" {
			tokens = append(tokens, token)
		}

		if !found {
			word += string(c)
		}
	}

	token := tokenWord(word)
	if token.Kind != TokenEOF {
		tokens = append(tokens, token)
	}

	tokens = append(tokens, Token{TokenEOF, ""})

	return tokens, comments, nil
}

func readQuotedLiteral(str []rune, i int, quote rune) (string, int, error) {
	i++
	terminated := false
	value := ""
	for ; i < len(str); i++ {
		if str[i] == quote {
			terminated = true
			break
		}

		value += string(str[i])
	}

	if !terminated {
		return "", i, fmt.Errorf(
			"unterminated literal, did not find closing %c", quote)
	}

	return value, i, nil
}

func tokenWord(word string) Token {
	word = strings.TrimSpace(word)

	switch word {
	case "":
		return Token{TokenEOF, ""}

	case "true", "false":
		return Token{TokenBool, word}

	case "func":
		return Token{TokenFunc, word}
	}

	return Token{TokenIdentifier, word}
}

func isDecimalCharacter(c rune) bool {
	return (c >= '0' && c <= '9') || c == '.'
}

func tokenKindForQuote(quote rune) (kind string) {
	switch quote {
	case '\'':
		kind = TokenCharacter

	case '"':
		kind = TokenString

	case '`':
		kind = TokenData
	}

	return
}
