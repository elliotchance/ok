package parser

import (
	"fmt"
	"ok/lexer"
)

// ParseString parses source code and returns the AST for the file.
func ParseString(s string) *Parser {
	parser := &Parser{
		File: &File{},
	}

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
	var consumedNothing bool
	parser.File.Root, offset, consumedNothing, err = consumeFunc(parser, offset)
	if consumedNothing {
		// This is OK. It means there was no function in the file.
	} else if err != nil {
		// Something went wrong consuming the function.
		parser.Errors = append(parser.Errors, err)

		return parser
	}

	if parser.File.Tokens[offset].Kind != lexer.TokenEOF {
		parser.Errors = append(parser.Errors,
			fmt.Errorf("found extra %s at the end of the file",
				parser.File.Tokens[offset]))
	}

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
