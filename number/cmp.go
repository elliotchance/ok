package number

import "github.com/cockroachdb/apd/v2"

// Cmp compares x and y and returns (regardless of the precision):
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func Cmp(a, b *apd.Decimal) int {
	return a.Cmp(b)
}
