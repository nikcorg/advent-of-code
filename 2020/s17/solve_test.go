package s17

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `.#.
..#
###
`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out, 6)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
		offset   uint64
	}{
		{solver.SolveFirst, "solution: 112\n", input, 0},
		// {solver.SolveSecond, "solution: 1\n", input2, 0},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
