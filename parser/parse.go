package parser

import (
	"io/ioutil"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// ParseString parses source code from a string. The fileName is used in error
// messages, the file does not have to exist.
func (parser *Parser) ParseString(s string, fileName string) {
	defer func() {
		err := parser.resolveInterfaces()
		if err != nil {
			parser.appendErrorAt("", err.Error())
		}
	}()

	var err error
	options := lexer.Options{
		IncludeComments: false,
	}
	var comments []*ast.Comment
	parser.tokens, comments, err = lexer.TokenizeString(s, options, fileName)
	parser.comments = append(parser.comments, comments...)

	if err != nil {
		at := "at the start"
		if len(parser.tokens) > 0 {
			at = "after " + parser.tokens[len(parser.tokens)-1].String()
		}

		parser.appendErrorf(nil, "unterminated string found %s", at)

		return
	}

	var offset int
	for {
		switch parser.tokens[offset].Kind {
		case lexer.TokenIdentifier:
			var name string
			var value *ast.Literal
			name, value, offset, err = consumeConstant(parser, offset)
			if err != nil {
				parser.appendErrorAt(parser.pos(offset), err.Error())

				goto done
			}

			// TODO(elliot): Check for redefinition.
			parser.Constants[name] = value

		case lexer.TokenFunc:
			// TODO(elliot): Check for already declared functions.
			// TODO(elliot): Function cannot be anonymous here.
			var fn *ast.Func
			fn, offset, _, err = consumeFunc(parser, offset)
			if err != nil {
				parser.appendErrorAt(parser.pos(offset), err.Error())

				goto done
			}

			// TODO(elliot): Remove this. We need to index by real name for
			//  resolving types. But this also means that types can't exist
			//  beyond the root level.
			parser.funcs[fn.Name] = fn

			parser.funcs[fn.UniqueName] = fn

		case lexer.TokenTest:
			var t *ast.Test
			t, offset, err = consumeTest(parser, offset)
			if err != nil {
				parser.appendError(t, err.Error())

				goto done
			}
			parser.tests = append(parser.tests, t)

		case lexer.TokenImport:
			var imp *ast.Import
			imp, offset, err = consumeImport(parser, offset)
			if err != nil {
				parser.appendError(imp, err.Error())

				goto done
			}
			parser.imports[imp.VariableName] = imp.PackageName

		case lexer.TokenEOF:
			goto done

		default:
			parser.appendErrorf(nil,
				"found extra %s at the end of the file",
				parser.tokens[offset])

			return
		}
	}
done:

	return
}

func consume(p *Parser, offset int, expected []string) (int, error) {
	originalOffset := offset

	for i, t := range expected {
		tok := p.tokens[offset+i]

		if t != tok.Kind {
			var after string
			if offset+i-1 >= 0 {
				after = p.tokens[offset+i-1].Kind
			}

			return originalOffset, newTokenMismatch(expected[i], after, tok.Kind)
		}
	}

	return offset + len(expected), nil
}

// ParseFile will read and parse the file. If the file cannot be read the error
// will be appended to the parser errors.
func (parser *Parser) ParseFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		parser.appendError(nil, err.Error())
		return
	}

	parser.ParseString(string(data), fileName)
}

// ParseDirectory will read and parse all ".ok" files in a directory (not
// recursive). This is analogous to parsing a package. If any of the files (or
// the directory itself) cannot be read the error will be appended to the parser
// errors.
//
// If includeTests = true, then ".okt" files will also be included.
func (parser *Parser) ParseDirectory(dirPath string, includeTests bool) {
	fileNames, err := getAllOKFilesInPath(dirPath, includeTests)
	if err != nil {
		parser.appendError(nil, err.Error())
		return
	}

	for _, fileName := range fileNames {
		parser.ParseFile(fileName)
	}
}
