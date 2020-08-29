// This file ensures that the standard library is generated (as Go code) so it
// can be bundled with any binaries that might need it.

package vm

//go:generate go run github.com/elliotchance/ok/cmd/lib-gen
