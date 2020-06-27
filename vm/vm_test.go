package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/vm"
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
