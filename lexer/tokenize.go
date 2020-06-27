package lexer

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
)

// TokenizeString returns a slice of tokens from the provided str.
func TokenizeString(str string, options Options) ([]Token, []*ast.Comment, error) {
	var tokens []Token
	var comments []*ast.Comment
	var word string

	runes := []rune(str)
	runesLen := len(runes)
	endOfLineForNextToken := false
	for i := 0; i < runesLen; i++ {
		c := runes[i]
		var token Token
		found := false

		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// Make sure we are not in the middle of reading a word. The digit
			// would be part of the identifier.
			if word == "" {
				token.Kind = TokenNumberLiteral
				for ; i < runesLen && isDecimalCharacter(runes[i]); i++ {
					token.Value += string(runes[i])
				}

				i--
				tokens = appendToken(tokens, token, &endOfLineForNextToken)
				continue
			}

		case '/':
			if i+1 < runesLen && runes[i+1] == '/' {
				token.Kind = TokenComment
				i += 2
				for ; i < runesLen && runes[i] != '\n'; i++ {
					token.Value += string(runes[i])
				}

				comments = append(comments, &ast.Comment{
					Comment: token.Value,
				})
				if options.IncludeComments {
					tokens = appendToken(tokens, token, &endOfLineForNextToken)
				}
				continue

			} else if i+1 < runesLen && runes[i+1] == '=' {
				token.Value = TokenDivideAssign
				i++

				found = true
				token.Kind = token.Value
			} else {
				found = true
				token = NewToken(string(c), string(c))
			}

		case '\'', '"', '`':
			value, newI, err := readQuotedLiteral(runes, i, c)
			if err != nil {
				// It's important that we return the tokens up until now so that
				// the parser has to the context to say where the error
				// occurred.
				return tokens, comments, err
			}

			i = newI
			newToken := NewToken(tokenKindForQuote(c), value)
			tokens = appendToken(tokens, newToken, &endOfLineForNextToken)
			found = true

		case ' ':
			found = true

		case '\n', '\r':
			// Only set the end of line if we are not in the middle of consuming
			// a token, since we need the newline to be after the most recent
			// finished token.
			if word == "" && len(tokens) > 0 {
				tokens[len(tokens)-1].IsEndOfLine = true
			} else {
				endOfLineForNextToken = true
			}
			found = true

		case '+', '-':
			token.Value = string(c)
			if i < runesLen-1 && (runes[i+1] == c || runes[i+1] == '=') {
				token.Value += string(runes[i+1])
				i++
			}

			found = true
			token.Kind = token.Value

		case '(', ')', '[', ']', '{', '}',
			'*', '%', '=', '!', '>', '<', ',', ';', ':', '.':
			token.Value = string(c)
			if i < runesLen-1 && runes[i+1] == '=' {
				token.Value = string(c) + "="
				i++
			}
			found = true
			token.Kind = token.Value
		}

		if found && word != "" {
			tokens = appendToken(tokens, tokenWord(word), &endOfLineForNextToken)
			word = ""
		}

		if token.Kind != "" {
			tokens = appendToken(tokens, token, &endOfLineForNextToken)
		}

		if !found {
			word += string(c)
		}
	}

	token := tokenWord(word)
	if token.Kind != TokenEOF {
		tokens = append(tokens, token)
	}

	tokens = append(tokens, NewToken(TokenEOF, ""))

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
		return NewToken(TokenEOF, "")

	case "true", "false":
		return NewToken(TokenBoolLiteral, word)

	case "and", "break", "case", "continue", "else", "if", "for", "func", "not",
		"or", "switch", "any", "bool", "char", "data", "number", "string", "in",
		"return", "test", "assert", "import":
		return NewToken(word, word)
	}

	return NewToken(TokenIdentifier, word)
}

func isDecimalCharacter(c rune) bool {
	return (c >= '0' && c <= '9') || c == '.'
}

func tokenKindForQuote(quote rune) (kind string) {
	switch quote {
	case '\'':
		kind = TokenCharLiteral

	case '"':
		kind = TokenStringLiteral

	case '`':
		kind = TokenDataLiteral
	}

	return
}

func appendToken(tokens []Token, token Token, endOfLine *bool) []Token {
	if *endOfLine {
		*endOfLine = false
		token.IsEndOfLine = true
	}

	return append(tokens, token)
}
