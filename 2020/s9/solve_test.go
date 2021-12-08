package s9

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

func TestSolve(t *testing.T) {
	preamble := 5

	out := new(bytes.Buffer)
	solver := New(context.Background(), out, preamble)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
	}{
		{solver.SolveFirst, "solution: 127\n"},
		{solver.SolveSecond, "solution: 62\n"},
	}

	for _, test := range tests {
		inp := strings.NewReader(input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
