package s12

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `
F10
N3
F7
R90
F11
	`

	const input2 = `
F10
R180
F20
L180
F20
R90
R90
F20
L90
L90
F10
`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}{
		{solver.SolveFirst, "solution: 25\n", input},
		{solver.SolveSecond, "solution: 286\n", input},
		{solver.SolveSecond, "solution: 0\n", input2},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
