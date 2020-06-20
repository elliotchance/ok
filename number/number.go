package number

import (
	"math/big"
	"strconv"
	"strings"
)

// Number represents a number value.
type Number struct {
	Rat *big.Rat

	// Precision describes the number of digits after the decimal place.
	Precision int
}

// NewNumber creates a new number from a string that is well formatted.
func NewNumber(s string) *Number {
	// We do not care about the error because we expect the number to be well
	// formatted.
	rat, _ := big.NewRat(0, 1).SetString(s)

	return &Number{
		Rat:       rat,
		Precision: precision(s),
	}
}

// String returns the string representation of the number retaining the current
// precision.
func (n *Number) String() string {
	return n.Rat.FloatString(n.Precision)
}

// IsZero returns true if the value equals 0 (at any precision).
func (n *Number) IsZero() bool {
	return n.Rat.Cmp(new(big.Rat)) == 0
}

// Int64 returns the int64 value. This is only for internal use when we need the
// actual int for things like array sizes. Non-integers will be rounded to the
// nearest integer.
func (n *Number) Int64() int64 {
	i64, _ := strconv.ParseInt(n.Rat.FloatString(0), 10, 64)

	return i64
}

// Int returns the int value. This is only for internal use when we need the
// actual int for things like array lengths. Non-integers will be rounded to the
// nearest integer.
func (n *Number) Int() int {
	return int(n.Int64())
}

func precision(s string) int {
	if idx := strings.Index(s, "."); idx > 0 {
		return len(s) - idx - 1
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
