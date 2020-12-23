package types

import (
	"unicode"
)

func tokenize(ty string) []string {
	var tokens []string
	word := ""
	for i := 0; i < len(ty); i++ {
		if unicode.IsSpace(rune(ty[i])) {
			if word != "" {
				tokens = append(tokens, word)
				word = ""
			}
			continue
		}

		if ty[i] == '(' || ty[i] == ')' || ty[i] == ',' ||
			ty[i] == '[' || ty[i] == ']' ||
			ty[i] == '{' || ty[i] == '}' {
			if word != "" {
				tokens = append(tokens, word)
				word = ""
			}
			tokens = append(tokens, string(ty[i]))
			continue
		}

		word += string(ty[i])
	}

	if word != "" {
		tokens = append(tokens, word)
	}

	return tokens
}

func parseType(tokens []string, offset int) (*Type, int) {
	switch tokens[offset] {
	case "func":
		return parseFunc(tokens, offset)

	case "[":
		offset += 2
		var ty *Type
		ty, offset = parseType(tokens, offset)

		return &Type{
			Kind:    KindArray,
			Element: ty,
		}, offset

	case "{":
		offset += 2
		var ty *Type
		ty, offset = parseType(tokens, offset)

		return &Type{
			Kind:    KindMap,
			Element: ty,
		}, offset
	}

	ty := &Type{
		Kind: kindFromString(tokens[offset]),
	}

	if ty.Kind == KindUnresolvedInterface {
		ty.Name = tokens[offset]
	}

	return ty, offset + 1
}

func parseFunc(tokens []string, offset int) (*Type, int) {
	ty := &Type{
		Kind: KindFunc,
	}
	offset += 2 // skip "func" "("
	for tokens[offset] != ")" {
		if tokens[offset] == "," {
			offset++
			continue
		}

		var argType *Type
		argType, offset = parseType(tokens, offset)
		ty.Arguments = append(ty.Arguments, argType)
	}

	offset++ // skip ")"

	if offset < len(tokens) && tokens[offset] != "," && tokens[offset] != ")" {
		if tokens[offset] == "(" {
			offset++ // skip "("
			for tokens[offset] != ")" {
				if tokens[offset] == "," {
					offset++
					continue
				}

				var returnType *Type
				returnType, offset = parseType(tokens, offset)
				ty.Returns = append(ty.Returns, returnType)
			}
			offset++ // skip ")"
		} else {
			var returnType *Type
			returnType, offset = parseType(tokens, offset)
			ty.Returns = []*Type{returnType}
		}
	}

	return ty, offset
}
