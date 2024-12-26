package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointAdd(t *testing.T) {
	assert.Equal(t, Point{3, 4}, Point{3, 3}.Add(Point{0, 1}))
}
