package s11

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	const input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	out := new(bytes.Buffer)
	solver := New(context.Background(), out)

	tests := []struct {
		solve    func(io.Reader) error
		expected string
		input    string
	}{
		{solver.SolveFirst, "solution: 37\n", input},
		{solver.SolveSecond, "solution: 26\n", input},
	}

	for _, test := range tests {
		inp := strings.NewReader(test.input)

		assert.Nil(t, test.solve(inp))
		assert.Equal(t, test.expected, out.String())

		out.Reset()
		t.Log("---")
	}
}

func TestMapTileSecond(t *testing.T) {
	input1 := strings.TrimSpace(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	output1 := strings.TrimSpace(`
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`)

	output2 := strings.TrimSpace(`
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`)

	output3 := strings.TrimSpace(`
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`)

	m := &floormap{}
	m.FromString(input1)

	assert.Equal(t, input1, m.String(), "initial input")

	var dc bool

	after1, dc := step(m, mapTileSecond)
	assert.Equal(t, output1, after1.String(), "after first step")
	assert.True(t, dc)

	after2, dc := step(after1, mapTileSecond)
	assert.Equal(t, output2, after2.String(), "after second step")
	assert.True(t, dc)

	after3, dc := step(after2, mapTileSecond)
	assert.Equal(t, output3, after3.String(), "after third step")
	assert.True(t, dc)
}
