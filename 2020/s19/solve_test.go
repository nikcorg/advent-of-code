package s19

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
		offset   uint64
	}{
		{solver.SolveFirst, "solution: 2\n", input, 0},
		// {solver.SolveSecond, "solution: 848\n", input, 0},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
