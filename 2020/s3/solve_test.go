package s3

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/nikcorg/aoc2020/utils/linestream"
	"github.com/stretchr/testify/assert"
)

const input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestSolver(t *testing.T) {
	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
	}{
		{solver.SolveFirst, "solution: 7\n"},
		{solver.SolveSecond, "solution: 336\n"},
	}

	for _, test := range tests {
		inp := strings.NewReader(input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
	}
}

func TestSolveSlope(t *testing.T) {
	tests := []struct {
		slopeX, slopeY, expected int
	}{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	for i, test := range tests {
		inp := make(linestream.LineChan)

		linestream.New(context.Background(), bufio.NewReader(strings.NewReader(input)), inp)
		actual := <-solveSlope(test.slopeX, test.slopeY, linestream.SkipEmpty(inp))

		assert.Equalf(t, test.expected, actual, "test %d, expected %d, got %d", i, test.expected, actual)
	}
}
