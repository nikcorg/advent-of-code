package s1

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/nikcorg/aoc2020/utils/ob"
)

const input = `1721
979
366
299
675
1456
`

func TestFirst(t *testing.T) {
	out := &ob.Capture{}
	ctx := context.Background()
	solver := New(ctx, out)

	tests := []struct {
		solver   func(io.Reader) error
		expected string
	}{
		{solver.SolveFirst, "solution: 1721*299=514579\n"},
		{solver.SolveSecond, "solution: 675*979*366=241861950\n"},
	}

	for _, test := range tests {
		rdr := strings.NewReader(input)

		assert.Nil(t, test.solver(rdr), "does not return an error")
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}