package lexer

import "fmt"

// Pos describes the position of a token.
type Pos struct {
	FileName                    string
	LineNumber, CharacterNumber int
}

func (pos Pos) add(characters int) Pos {
	pos.CharacterNumber += characters

	return pos
}

func (pos *Pos) nextLine() {
	pos.LineNumber++
	pos.CharacterNumber = 0
}

// String returns a human-readable position.
func (pos *Pos) String() string {
	return fmt.Sprintf("%s:%d:%d", pos.FileName, pos.LineNumber, pos.CharacterNumber)
}
