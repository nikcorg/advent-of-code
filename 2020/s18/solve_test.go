package s18

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		// solve
		expected []int
		input    string
	}{
		{[]int{71, 231}, "1 + 2 * 3 + 4 * 5 + 6"},
		{[]int{51, 51}, "1 + (2 * 3) + (4 * (5 + 6))"},
		{[]int{26, 46}, "2 * 3 + (4 * 5)"},
		{[]int{437, 1445}, "5 + (8 * 3 + 9 + 3 * 4 * 3)"},
		{[]int{12240, 669060}, "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"},
		{[]int{13632, 23340}, "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"},
	}

	solvers := []func(io.Reader) error{solver.SolveFirst, solver.SolveSecond}

	for _, test := range tests {
		for n, solve := range solvers {
			inp := strings.NewReader(test.input)

			assert.Nil(t, solve(inp))
			assert.Equal(t, fmt.Sprintf("solution: %d\n", test.expected[n]), out.String())

			out.Reset()
		}
	}
}
