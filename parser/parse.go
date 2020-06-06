package parser

import (
	"fmt"
	"ok/lexer"
)

// ParseString parses source code and returns the AST for the file.
func ParseString(s string) (*File, error) {
	f := &File{}
	var err error
	f.Tokens, f.Comments, err = lexer.TokenizeString(s, lexer.Options{
		IncludeComments: false,
	})
	if err != nil {
		at := "at the start"
		if len(f.Tokens) > 0 {
			at = "after " + f.Tokens[len(f.Tokens)-1].String()
		}

		return nil, fmt.Errorf("unterminated string found %s", at)
	}

	var offset int
	var consumedNothing bool
	f.Root, offset, consumedNothing, err = consumeFunc(f, offset)
	if consumedNothing {
		// This is OK. It means there was no function in the file.
	} else if err != nil {
		// Something went wrong consuming the function.
		return nil, err
	}

	if f.Tokens[offset].Kind != lexer.TokenEOF {
		return nil, fmt.Errorf("found extra %s at the end of the file",
			f.Tokens[offset])
	}

	return f, nil
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
