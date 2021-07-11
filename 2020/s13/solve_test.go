package s13

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `939
7,13,x,x,59,x,31,19`

	const input2 = `000
17,x,13,19`

	const input3 = `000
67,7,59,61`

	const input4 = `000
67,x,7,59,61`

	const input5 = `000
1789,37,47,1889`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
		offset   uint64
	}{
		{solver.SolveFirst, "solution: 295\n", input, 0},
		{solver.SolveSecond, "solution: 1068781\n", input, 0},
		{solver.SolveSecond, "solution: 3417\n", input2, 0},
		{solver.SolveSecond, "solution: 754018\n", input3, 0},
		{solver.SolveSecond, "solution: 779210\n", input4, 0},
		{solver.SolveSecond, "solution: 1202161486\n", input5, 0},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
