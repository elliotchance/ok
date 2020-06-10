package number

// Cmp compares x and y and returns (regardless of the precision):
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func Cmp(a, b *Number) int {
	return a.Rat.Cmp(b.Rat)
}
