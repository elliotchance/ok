package parser

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
)

// ParseString parses source code and returns the AST for the file.
func ParseString(s string) *Parser {
	parser := &Parser{
		File: &File{},
	}
	parser.File.Funcs = map[string]*ast.Func{}

	var err error
	parser.File.Tokens, parser.File.Comments, err = lexer.TokenizeString(s, lexer.Options{
		IncludeComments: false,
	})
	if err != nil {
		at := "at the start"
		if len(parser.File.Tokens) > 0 {
			at = "after " + parser.File.Tokens[len(parser.File.Tokens)-1].String()
		}

		parser.Errors = append(parser.Errors,
			fmt.Errorf("unterminated string found %s", at))

		return parser
	}

	var offset int
	for {
		switch parser.File.Tokens[offset].Kind {
		case lexer.TokenFunc:
			// TODO(elliot): Check for already declared functions.
			var fn *ast.Func
			fn, offset, err = consumeFunc(parser, offset)
			if err != nil {
				parser.Errors = append(parser.Errors, err)

				goto done
			}
			parser.File.Funcs[fn.Name] = fn

		case lexer.TokenTest:
			var t *ast.Test
			t, offset, err = consumeTest(parser, offset)
			if err != nil {
				parser.Errors = append(parser.Errors, err)

				goto done
			}
			parser.File.Tests = append(parser.File.Tests, t)

		case lexer.TokenEOF:
			goto done

		default:
			parser.Errors = append(parser.Errors,
				fmt.Errorf("found extra %s at the end of the file",
					parser.File.Tokens[offset]))

			return parser
		}
	}
done:

	return parser
}

func consume(f *File, offset int, expected []string) (int, error) {
	for i, t := range expected {
		tok := f.Tokens[offset+i]

		if t != tok.Kind {
			var after string
			if offset+i-1 >= 0 {
				after = f.Tokens[offset+i-1].Kind
			}

			return offset + i, newTokenMismatch(expected[i], after, tok.Kind)
		}
	}

	return offset + len(expected), nil
}
