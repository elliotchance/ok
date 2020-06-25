package vm_test

import (
	"ok/compiler"
	"ok/parser"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVM_Run(t *testing.T) {
	t.Run("call-with-args", func(t *testing.T) {
		p := parser.ParseString(`
func foo() {
    print("called foo")
}

func main() {
    foo()
    print(add(2.3, 7.90))
}

func add(a, b number) number {
    return a + b
}
`)
		require.Nil(t, p.Errors)
		f, err := compiler.CompileFile(p.File)
		require.NoError(t, err)

		m := vm.NewVM(f.Funcs, f.Tests, "pkg")
		assert.NoError(t, m.Run())
	})
}
