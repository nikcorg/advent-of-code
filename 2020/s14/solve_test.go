package s14

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

	const input2 = `mask = 000000000000000000000000000000X1001X
	mem[42] = 100
	mask = 00000000000000000000000000000000X0XX
	mem[26] = 1
	`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}{
		{solver.SolveFirst, "solution: 165\n", input},
		{solver.SolveSecond, "solution: 208\n", input2},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
