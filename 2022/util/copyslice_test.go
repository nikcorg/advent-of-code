package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopySlice(t *testing.T) {
	// this is a pointless test purely for coverage
	xs := []int{1, 2, 3, 4, 5}
	ys := CopySlice(xs)

	assert.Equal(t, xs, ys)
}
