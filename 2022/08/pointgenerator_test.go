package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2022/util"
)

func TestPointGenerator(t *testing.T) {
	cases := []struct {
		steps       []util.Point
		startsAt    util.Point
		translateBy util.Point
		endsAt      util.Point
	}{
		{[]util.Point{util.NewPoint(0, 0), util.NewPoint(1, 1), util.NewPoint(2, 2)},
			util.NewPoint(0, 0), util.NewPoint(1, 1), util.NewPoint(3, 3)},
		{[]util.Point{util.NewPoint(3, 3), util.NewPoint(2, 2), util.NewPoint(1, 1)},
			util.NewPoint(3, 3), util.NewPoint(-1, -1), util.NewPoint(0, 0)},
		{[]util.Point{util.NewPoint(0, 0), util.NewPoint(1, 0), util.NewPoint(2, 0)},
			util.NewPoint(0, 0), util.NewPoint(1, 0), util.NewPoint(3, 0)},
		{[]util.Point{util.NewPoint(0, 0), util.NewPoint(0, 1), util.NewPoint(0, 2)},
			util.NewPoint(0, 0), util.NewPoint(0, 1), util.NewPoint(0, 3)},
		{[]util.Point{util.NewPoint(3, 3), util.NewPoint(2, 3), util.NewPoint(1, 3)},
			util.NewPoint(3, 3), util.NewPoint(-1, 0), util.NewPoint(0, 3)},
	}

	for _, tc := range cases {
		steps := []util.Point{}

		for step := range pointGenerator(tc.startsAt, tc.translateBy, tc.endsAt) {
			steps = append(steps, step)
		}

		assert.Equal(t, tc.steps, steps)
	}
}
