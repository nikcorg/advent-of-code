package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2022/util"
)

func TestTriangle(t *testing.T) {
	x := Triangle{util.NewPoint(0, 0), util.NewPoint(5, 0), util.NewPoint(0, 5)}

	for _, c := range []struct {
		p   util.Point
		exp bool
	}{
		{util.NewPoint(0, 0), true},
		{util.NewPoint(0, 1), true},
		{util.NewPoint(-1, 0), false},
		{util.NewPoint(0, -1), false},

		// on the hypotenuse
		{util.NewPoint(0, 5), true},
		{util.NewPoint(1, 4), true},
		{util.NewPoint(2, 3), true},
		{util.NewPoint(3, 2), true},
		{util.NewPoint(4, 1), true},
		{util.NewPoint(5, 0), true},

		// one step beyond
		{util.NewPoint(0, 6), false},
		{util.NewPoint(1, 5), false},
		{util.NewPoint(2, 4), false},
		{util.NewPoint(3, 3), false},
		{util.NewPoint(4, 2), false},
		{util.NewPoint(5, 1), false},
		{util.NewPoint(6, 0), false},
	} {
		assert.Equal(t, c.exp, x.Contains(c.p))
	}
}
