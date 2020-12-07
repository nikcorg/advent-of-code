package s5

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/nikcorg/aoc2020/utils/ob"
	"github.com/stretchr/testify/assert"
)

const input = `FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`

func TestSolver(t *testing.T) {
	out := &ob.Capture{}
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
	}{
		{solver.SolveFirst, "solution: 820\n"},
	}

	for _, test := range tests {
		inp := strings.NewReader(input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
