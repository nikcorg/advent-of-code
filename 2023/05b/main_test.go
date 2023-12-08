package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	type intRange struct {
//		From, To int
//	}
func TestRangeOverlap(t *testing.T) {
	cases := []struct {
		Left, Right, Overlap intRange
		Overlaps, Contained  bool
	}{
		{intRange{1, 5}, intRange{2, 7}, intRange{2, 5}, true, false},
		{intRange{2, 7}, intRange{1, 5}, intRange{2, 5}, true, false},
		{intRange{2, 7}, intRange{3, 4}, intRange{3, 4}, true, true},
		{intRange{3, 4}, intRange{2, 7}, intRange{3, 4}, true, false},
		{intRange{3, 4}, intRange{5, 7}, intRange{}, false, false},
	}

	for _, c := range cases {
		if c.Overlaps {
			overlap, ok := c.Left.GetOverlap(c.Right)
			assert.True(t, ok)
			assert.Equal(t, c.Overlap, overlap, "%+v overlaps with %+v", c.Left, c.Right)
		}

		assert.Equal(t, c.Contained, c.Left.Contains(c.Right))
	}
}

func TestRangeDifference(t *testing.T) {
	cases := []struct {
		Left, Right intRange
		Difference  []intRange
		Overlaps    bool
	}{
		{intRange{1, 5}, intRange{2, 7}, []intRange{{6, 7}}, true},
		{intRange{2, 7}, intRange{1, 5}, []intRange{{1, 1}}, true},
		{intRange{3, 4}, intRange{2, 7}, []intRange{{2, 2}, {5, 7}}, true},
		// no difference, because left fully contains r
		{intRange{2, 7}, intRange{3, 4}, []intRange{}, false},
	}

	for _, c := range cases {
		diff, ok := c.Left.GetDifference(c.Right)
		if !c.Overlaps {
			assert.False(t, ok)
			continue
		}

		assert.True(t, ok)
		assert.Equal(t, c.Difference, diff)
	}
}

//	type transform struct {
//		Range intRange
//		Mod   int
//	}
func TestTransform(t *testing.T) {

}

//	type transformer struct {
//		Name       string
//		Transforms []transform
//	}
func TestTransformer(t *testing.T) {

}
