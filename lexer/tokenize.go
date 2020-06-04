package lexer

import (
	"errors"
	"strings"
)

// TokenizeString returns a slice of tokens from the provided str.
func TokenizeString(str string) ([]Token, error) {
	var tokens []Token
	var word string
	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		var token Token
		found := false

		switch c {
		case '"':
			token := Token{TokenString, ""}
			i++
			terminated := false
			for ; i < len(str); i++ {
				if str[i] == '"' {
					terminated = true
					break
				}
				token.Value += string(str[i])
			}

			if !terminated {
				return tokens, errors.New("unterminated string")
			}

			tokens = append(tokens, token)
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

	return tokens, nil
}

func tokenWord(word string) Token {
	word = strings.TrimSpace(word)

	switch word {
	case "":
		return Token{TokenEOF, ""}

	case "func":
		return Token{TokenFunc, word}
	}

	return Token{TokenIdentifier, word}
}
