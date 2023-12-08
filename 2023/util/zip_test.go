package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	as := []int{1, 3, 5}
	bs := []int{2, 4, 6, 7}

	cs := Zip(as, bs)

	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, cs)
}
