package vm

import (
	"fmt"
	"strings"
	"unicode"
)

// Register is the name of a register. At the moment variable names can also be
// used as registers, but that will be removed in the future. When that happens
// these can be refactored into an int.
type Register string

type TypeRegister string

type SymbolRegister string

// String returns either "rX" or the name of the variable. See Register for
// details.
func (r Register) String() string {
	// An empty register is sometimes used when there is no register need, like
	// how Key is optional in NextArray.
	if r == "" {
		return "_"
	}

	if unicode.IsDigit([]rune(r)[0]) {
		return "$" + string(r)
	}

	return string(r)
}

// Registers is multiple sequential registers. It might represents function
// arguments, for example.
type Registers []Register

// String returns the registers as a set, like "(r0, r1)". The values will be
// wrapped with parenthesis even if there are one or zero elements.
func (rs Registers) String() string {
	ss := make([]string, len(rs))
	for i, r := range rs {
		ss[i] = r.String()
	}

	return fmt.Sprintf("(%s)", strings.Join(ss, ", "))
}
