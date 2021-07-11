package s10

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

const input2 = `16
10
15
5
1
11
7
19
6
12
4
`

func TestSolve(t *testing.T) {
	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}{
		{solver.SolveFirst, "solution: 220\n", input},
		{solver.SolveFirst, "solution: 35\n", input2},
		{solver.SolveSecond, "solution: 8\n", input2},
		{solver.SolveSecond, "solution: 19208\n", input},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
		t.Log("---")
	}
}
