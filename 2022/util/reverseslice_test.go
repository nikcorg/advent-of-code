package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSlice(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.out, ReverseSlice(tc.in))
	}
}
