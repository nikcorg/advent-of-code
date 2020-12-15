package s15

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
	inputs := [][]string{
		{`0,3,6`, `436`, `175594`},
		{`1,3,2`, `1`, `2578`},
		{`2,1,3`, `10`, `3544142`},
		{`1,2,3`, `27`, `261214`},
		{`2,3,1`, `78`, `6895259`},
		{`3,2,1`, `438`, `18`},
		{`3,1,2`, `1836`, `362`},
	}

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)
	type testCase struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}
	tests := []testCase{}

	for _, inp := range inputs {
		t.Logf("expect %s from %s", inp[1], inp[0])
		tests = append(tests, testCase{solver.SolveFirst, fmt.Sprintf("solution: %s\n", inp[1]), inp[0]})
		tests = append(tests, testCase{solver.SolveSecond, fmt.Sprintf("solution: %s\n", inp[2]), inp[0]})
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}
