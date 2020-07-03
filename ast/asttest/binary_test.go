package asttest_test

import (
	"testing"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/stretchr/testify/assert"
)

func TestNewBinary(t *testing.T) {
	node := asttest.NewBinary(
		asttest.NewLiteralData([]byte("foo")),
		"/",
		asttest.NewLiteralData([]byte("bar")),
	)
	assert.Equal(t, asttest.NewLiteralData([]byte("foo")), node.Left)
	assert.Equal(t, "/", node.Op)
	assert.Equal(t, asttest.NewLiteralData([]byte("bar")), node.Right)
}
