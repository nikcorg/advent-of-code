package s11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOccupiedAdjacent(t *testing.T) {
	input := strings.TrimSpace(`
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

	input2 := strings.TrimSpace(`
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

	tests := []struct {
		input    string
		x        int
		y        int
		expected int
	}{
		{input, 0, 0, 0},
		{input2, 0, 0, 2},
		{input2, 0, 1, 3},
		{input2, 3, 3, 5},
		{input2, 3, 8, 8},
	}

	for _, test := range tests {
		m := floormap{}
		m.FromString(test.input)
		assert.Equal(t, test.expected, m.OccupiedAdjacent(test.x, test.y))
	}
}

func TestOccupiedVisibleFrom(t *testing.T) {
	input := strings.TrimSpace(`
.#####.
#.#.#.#
##...##
#..L..#
##...##
#.#.#.#
.#####.`)

	input2 := strings.TrimSpace(`
.............
.L.L.#.#.#.#.
.............`)

	input3 := strings.TrimSpace(`
.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`)

	input4 := strings.TrimSpace(`
.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`)

	input5 := strings.TrimSpace(`
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

	tests := []struct {
		input    string
		x        int
		y        int
		expected int
	}{
		{input, 3, 3, 4},
		{input2, 1, 1, 0},
		{input3, 3, 3, 0},
		{input3, 2, 3, 6},
		{input4, 3, 4, 8},
		{input5, 0, 0, 3},
		{input5, 0, 2, 5},
	}

	for n, test := range tests {
		m := floormap{}
		m.FromString(test.input)
		m.SetTileAt(test.x, test.y, observer)
		actual := m.OccupiedVisibleFrom(test.x, test.y)
		assert.Equalf(t, test.expected, actual, "test %d, expected %d, got %d", n+1, test.expected, actual)
	}
}
