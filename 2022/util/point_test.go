package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceTo(t *testing.T) {
	cases := []struct {
		from, to Point
		dist     int
	}{
		// We only need horizontal and vertical vector lengths
		{Point{0, 0}, Point{3, 0}, 3},
		{Point{3, 0}, Point{0, 0}, 3},
		{Point{0, 3}, Point{0, 0}, 3},
		{Point{0, 0}, Point{0, 3}, 3},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.dist, int(tc.from.DistanceTo(tc.to)))
	}
}
