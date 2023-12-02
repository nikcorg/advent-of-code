package intrange

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntRange(t *testing.T) {
	cases := []struct {
		a, b      NumRange[int]
		contained bool
		overlaps  bool
	}{
		{NumRange[int]{2, 4}, NumRange[int]{6, 8}, false, false},
		{NumRange[int]{2, 3}, NumRange[int]{4, 5}, false, false},
		{NumRange[int]{5, 7}, NumRange[int]{7, 9}, false, true},
		{NumRange[int]{2, 8}, NumRange[int]{3, 7}, true, true},
		{NumRange[int]{6, 6}, NumRange[int]{4, 6}, false, true},
		{NumRange[int]{2, 6}, NumRange[int]{4, 8}, false, true},

		{NumRange[int]{1, 3}, NumRange[int]{2, 5}, false, true},
		{NumRange[int]{1, 6}, NumRange[int]{4, 5}, true, true},
		{NumRange[int]{1, 5}, NumRange[int]{4, 5}, true, true},
	}

	for _, c := range cases {
		assert.Equal(t, c.a.Contains(c.b), c.contained)
		assert.Equal(t, c.a.Overlaps(c.b), c.overlaps)
	}
}

func TestNew(t *testing.T) {
	r := New(1, 2)
	assert.NotNil(t, r)
	assert.Equal(t, 1, r.Lower())
	assert.Equal(t, 2, r.Upper())
	assert.Equal(t, "1-2", r.String())
}

func TestFromString(t *testing.T) {
	cases := []struct {
		in         string
		lo, hi     int
		shouldFail bool
	}{
		{"1-2", 1, 2, false},
		{"99-10", 10, 99, false},
		{"hello world", 0, 0, true},
		{"99-poop", 0, 0, true},
		{"beep-boop", 0, 0, true},
	}

	for _, tc := range cases {
		r, err := FromString(tc.in, strconv.Atoi)
		if tc.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.lo, r.Lower())
			assert.Equal(t, tc.hi, r.Upper())
		}
	}
}
