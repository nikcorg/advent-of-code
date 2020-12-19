package s18

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}{
		{solver.SolveFirst, "solution: 71\n", "1 + 2 * 3 + 4 * 5 + 6"},
		{solver.SolveFirst, "solution: 26\n", "2 * 3 + (4 * 5)"},
		{solver.SolveFirst, "solution: 437\n", "5 + (8 * 3 + 9 + 3 * 4 * 3)"},
		{solver.SolveFirst, "solution: 12240\n", "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"},
		{solver.SolveFirst, "solution: 13632\n", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"},
		// {solver.SolveSecond, "solution: 1\n", input2, 0},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		t.Logf("expression: %s", test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
