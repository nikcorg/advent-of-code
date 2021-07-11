package s12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaypointMove(t *testing.T) {
	tests := []struct {
		init     *Point
		update   *Instruction
		expected *Point
	}{
		{&Point{-1, 10}, &Instruction{moveNorth, 3}, &Point{-4, 10}},
		{&Point{-1, 10}, &Instruction{moveSouth, 3}, &Point{2, 10}},
		{&Point{-1, 10}, &Instruction{moveEast, 3}, &Point{-1, 13}},
		{&Point{-1, 10}, &Instruction{moveWest, 3}, &Point{-1, 7}},
	}

	for n, test := range tests {
		p := test.init
		p.Move(test.update.cmd, test.update.units)
		assert.Truef(t, test.expected.Equal(p), "test %d", n)
	}
}

func TestWaypointRotate(t *testing.T) {
	tests := []struct {
		ref      *Point
		init     *Point
		update   *Instruction
		expected *Point
	}{
		// Counter-clockwise
		{&Point{0, 0}, &Point{-1, 10}, &Instruction{rotLeft, 90}, &Point{-10, -1}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotLeft, 90}, &Point{-2, 2}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotLeft, 180}, &Point{0, -2}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotLeft, 270}, &Point{4, 0}},

		// Clockwise
		{&Point{0, 0}, &Point{-1, 10}, &Instruction{rotRight, 90}, &Point{10, 1}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotRight, 270}, &Point{-2, 2}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotRight, 180}, &Point{0, -2}},
		{&Point{1, 1}, &Point{2, 4}, &Instruction{rotRight, 90}, &Point{4, 0}},
	}

	for n, test := range tests {
		p := test.init

		p.Rotate(test.ref, test.update.cmd, test.update.units)

		assert.Truef(t, test.expected.Equal(p), "test %d, expected %+v, got %+v", n+1, test.expected, p)
	}
}
