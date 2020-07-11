package lexer

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
)

// TokenizeString returns a slice of tokens from the provided str.
func TokenizeString(str string, options Options, fileName string) ([]Token, []*ast.Comment, error) {
	var tokens []Token
	var comments []*ast.Comment
	var word string

	runes := []rune(str)
	runesLen := len(runes)

	// endOfLineForNextToken will increment for every sequential new line. The
	// number of new lines needs to be tracked to correct the pos line number
	// after the token has been resolved.
	endOfLineForNextToken := 0

	var lastComment *ast.Comment
	pos := Pos{
		FileName:        fileName,
		LineNumber:      1,
		CharacterNumber: 0,
	}

	for i := 0; i < runesLen; i++ {
		pos.CharacterNumber++
		c := runes[i]
		var token Token
		found := false

		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// Make sure we are not in the middle of reading a word. The digit
			// would be part of the identifier.
			if word == "" {
				token.Kind = TokenNumberLiteral
				token.Pos = pos
				for ; i < runesLen && isDecimalCharacter(runes[i]); i++ {
					token.Value += string(runes[i])
				}

				i--
				tokens = appendToken(tokens, token, &endOfLineForNextToken, &pos)
				pos.CharacterNumber += len(token.Value) - 1
				continue
			}

		case '/':
			if i+1 < runesLen && runes[i+1] == '/' {
				token.Kind = TokenComment
				i += 2
				hasNewLine := false
				for ; i < runesLen; i++ {
					if runes[i] == '\n' {
						hasNewLine = true
						break
					}
					token.Value += string(runes[i])
				}

				// If we are not collecting comments as tokens we should always
				// treat the most recent token as the end of the line. This is
				// important for expression parsing that needs to know the last
				// token on the line.
				if !options.IncludeComments && len(tokens) > 0 {
					tokens[len(tokens)-1].IsEndOfLine = true
				}

				// If the previous token was a comment we append to it. However,
				// if there is a new line just before this comment then it
				// cannot be joined to the previous one.
				if lastComment != nil {
					lastComment.Comment += "\n" + token.Value

					// If we are including comments as tokens we will need to
					// append to the previous token as well.
					if len(tokens) > 0 &&
						tokens[len(tokens)-1].Kind == TokenComment {
						tokens[len(tokens)-1].Value += "\n" + token.Value
					}
				} else {
					comments = append(comments, &ast.Comment{
						Comment: token.Value,
						Pos:     pos.String(),
					})
					if options.IncludeComments {
						token.Pos = pos
						tokens = appendToken(tokens, token, &endOfLineForNextToken, &pos)
					}
					lastComment = comments[len(comments)-1]
				}

				if hasNewLine {
					pos.nextLine()
				} else {
					pos.CharacterNumber += 1 + len(token.Value)
				}

				continue

			} else if i+1 < runesLen && runes[i+1] == '=' {
				token.Value = TokenDivideAssign
				i++

				found = true
				token.Kind = token.Value
				token.Pos = pos
			} else {
				found = true
				token = NewToken(string(c), string(c), pos)
			}

		case '\'', '"', '`':
			value, newI, err := readQuotedLiteral(runes, i, c, pos)
			if err != nil {
				// It's important that we return the tokens up until now so that
				// the parser has to the context to say where the error
				// occurred.
				return tokens, comments, err
			}

			i = newI

			// Only strings can be interpolated.
			if c == '"' && containsInterpolation(value) {
				interpolatedTokens, err := interpolate(value, pos, fileName)
				if err != nil {
					return tokens, comments, err
				}

				tokens = append(tokens, interpolatedTokens...)
			} else {
				newToken := NewToken(tokenKindForQuote(c), value, pos)
				tokens = appendToken(tokens, newToken, &endOfLineForNextToken, &pos)
			}
			found = true

			// Since we support multibyte characters we have to be careful we
			// move forward the number of characters (not bytes).
			pos.CharacterNumber += len([]rune(value)) + 1

		case ' ':
			found = true

		case '\n', '\r':
			// Only set the end of line if we are not in the middle of consuming
			// a token, since we need the newline to be after the most recent
			// finished token.
			if word == "" {
				if len(tokens) > 0 {
					tokens[len(tokens)-1].IsEndOfLine = true
				}
				pos.nextLine()
			} else {
				endOfLineForNextToken++
			}
			found = true
			lastComment = nil

		case '+', '-':
			token.Value = string(c)
			token.Pos = pos
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
			token.Pos = pos

			// This is to stop comments on either side of an operator from being
			// considered a continuous comment block. However, be careful that
			// we don't reset the comment if the token is `(` because that would
			// cause us to lose the attachment of a function to its comment.
			if c != '(' {
				lastComment = nil
			}
		}

		if found && word != "" {
			// If we find a function, we might have to attach the comment.
			if len(tokens) > 0 &&
				tokens[len(tokens)-1].Kind == TokenFunc &&
				lastComment != nil &&
				// This is only needed to satisfy the IDE static analyzer.
				len(comments) > 0 {
				comments[len(comments)-1].Func = word
			}

			tokens = appendToken(tokens, tokenWord(word, pos), &endOfLineForNextToken, &pos)
			word = ""
		}

		if token.Kind != "" {
			tokens = appendToken(tokens, token, &endOfLineForNextToken, &pos)
		}

		if !found {
			word += string(c)
		}
	}

	pos.CharacterNumber++

	token := tokenWord(word, pos)
	if token.Kind != TokenEOF {
		tokens = append(tokens, token)
	}

	if endOfLineForNextToken > 0 {
		pos.LineNumber += endOfLineForNextToken
		pos.CharacterNumber = 1
	}
	tokens = append(tokens, NewToken(TokenEOF, "", pos))

	return tokens, comments, nil
}

func interpolate(s string, pos Pos, fileName string) ([]Token, error) {
	pos.CharacterNumber++
	tokens := []Token{
		NewToken(TokenInterpolateStart, "", pos),
	}

	stringLiteral := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			if stringLiteral != "" {
				tokens = append(tokens, NewToken(TokenStringLiteral, stringLiteral, pos))
				pos.CharacterNumber += len(stringLiteral)
				stringLiteral = ""
			}

			// Read everything until the closing curly. We do not allow curly
			// brackets in expressions so it doesn't need to be anymore
			// complicated than this.
			exprLen := strings.Index(s[i:], "}")

			// TODO(elliot): Handle exprLen == -1
			tokens = append(tokens, NewToken(TokenParenOpen, "(", pos))
			exprTokens, _, err := TokenizeString(s[i+1:exprLen+i], Options{}, fileName)
			if err != nil {
				return nil, err
			}

			// Fix all of the offsets for the subparse.
			for i := range exprTokens {
				exprTokens[i].Pos.CharacterNumber += pos.CharacterNumber
			}

			// Remove the EOF.
			// TODO(elliot): Should this be done with an Option instead?
			exprTokens = exprTokens[:len(exprTokens)-1]

			tokens = append(tokens, exprTokens...)
			tokens = append(tokens, NewToken(TokenParenClose, ")", pos.add(exprLen)))

			pos.CharacterNumber += exprLen + 1
			i += exprLen
		} else {
			stringLiteral += string(s[i])
		}
	}

	if stringLiteral != "" {
		tokens = append(tokens, NewToken(TokenStringLiteral, stringLiteral, pos))
		pos.CharacterNumber += len(stringLiteral)
	}

	tokens = append(tokens, NewToken(TokenInterpolateEnd, "", pos))

	return tokens, nil
}

func containsInterpolation(s string) bool {
	return strings.Contains(s, "{")
}

func readQuotedLiteral(str []rune, i int, quote rune, pos Pos) (string, int, error) {
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
			"%s unterminated literal, did not find closing %c",
			pos.String(), quote)
	}

	return value, i, nil
}

func tokenWord(word string, pos Pos) Token {
	word = strings.TrimSpace(word)

	switch word {
	case "":
		return NewToken(TokenEOF, "", pos)

	case "true", "false":
		return NewToken(TokenBoolLiteral, word, pos.add(-len(word)))

	case
		// Logical
		"and", "not", "or",

		// Control flow
		"break", "case", "continue", "else", "if", "for", "switch", "in",

		// Testing
		"test", "assert",

		// Types
		"any", "bool", "char", "data", "number", "string",

		// Statements
		"func", "return", "import",

		// Errors
		"try", "raise", "on":
		return NewToken(word, word, pos.add(-len(word)))
	}

	return NewToken(TokenIdentifier, word, pos.add(-len(word)))
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

func appendToken(tokens []Token, token Token, endOfLine *int, pos *Pos) []Token {
	if *endOfLine > 0 {
		pos.LineNumber += *endOfLine
		pos.CharacterNumber = 0

		token.IsEndOfLine = true
		*endOfLine = 0
	}

	// Operators are ingested as a single unit (like a literal would look
	// ahead). However, we cannot move the position forward at that time because
	// it would affect any non-appended token. So we have to adjust it here.
	if len(token.Value) > 1 &&
		(token.Value[1] == '=' || token.Value[1] == '+' || token.Value[1] == '-') {
		pos.CharacterNumber++
	}

	return append(tokens, token)
}
