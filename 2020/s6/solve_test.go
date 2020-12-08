package s6

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `abc

a
b
c

ab
ac

a
a
a
a

b
`

func TestSolver(t *testing.T) {
	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
	}{
		{solver.SolveFirst, "solution: 11\n"},
		{solver.SolveSecond, "solution: 6\n"},
	}

	for _, test := range tests {
		inp := strings.NewReader(input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
