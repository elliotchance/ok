package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// ParseString parses source code and returns the AST for the file.
func ParseString(s string, fileName string) *Parser {
	parser := &Parser{
		File:       &File{},
		finalizers: map[string][]*ast.Finally{},
	}
	parser.File.Funcs = map[string]*ast.Func{}
	parser.File.Imports = map[string]string{}

	var err error
	options := lexer.Options{
		IncludeComments: false,
	}
	parser.File.Tokens, parser.File.Comments, err = lexer.TokenizeString(s,
		options, fileName)
	if err != nil {
		at := "at the start"
		if len(parser.File.Tokens) > 0 {
			at = "after " + parser.File.Tokens[len(parser.File.Tokens)-1].String()
		}

		parser.AppendErrorf(nil, "unterminated string found %s", at)

		return parser
	}

	var offset int
	for {
		switch parser.File.Tokens[offset].Kind {
		case lexer.TokenFunc:
			// TODO(elliot): Check for already declared functions.
			// TODO(elliot): Function cannot be anonymous here.
			var fn *ast.Func
			fn, offset, _, err = consumeFunc(parser, offset)
			if err != nil {
				parser.AppendErrorAt(parser.File.Pos(offset), err.Error())

				goto done
			}
			parser.File.Funcs[fn.Name] = fn

		case lexer.TokenTest:
			var t *ast.Test
			t, offset, err = consumeTest(parser, offset)
			if err != nil {
				parser.AppendError(t, err.Error())

				goto done
			}
			parser.File.Tests = append(parser.File.Tests, t)

		case lexer.TokenImport:
			var imp *ast.Import
			imp, offset, err = consumeImport(parser, offset)
			if err != nil {
				parser.AppendError(imp, err.Error())

				goto done
			}
			parser.File.Imports[imp.PackageName] = imp.PackageName

		case lexer.TokenEOF:
			goto done

		default:
			parser.AppendErrorf(nil,
				"found extra %s at the end of the file",
				parser.File.Tokens[offset])

			return parser
		}
	}
done:

	return parser
}

func consume(f *File, offset int, expected []string) (int, error) {
	originalOffset := offset

	for i, t := range expected {
		tok := f.Tokens[offset+i]

		if t != tok.Kind {
			var after string
			if offset+i-1 >= 0 {
				after = f.Tokens[offset+i-1].Kind
			}

			return originalOffset, newTokenMismatch(expected[i], after, tok.Kind)
		}
	}

	return offset + len(expected), nil
}
