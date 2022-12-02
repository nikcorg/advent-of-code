package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	l0 := newline(point{2, 9}, point{2, 3})
	assert.Equal(t, point{2, 3}, l0.From)
	assert.Equal(t, point{2, 9}, l0.To)

	l1 := newline(point{2, 2}, point{2, 8})
	l2 := newline(point{1, 3}, point{4, 3})
	p := l1.Intersection(l2)
	assert.Equal(t, []point{point{2, 3}}, p)

	l3 := newline(point{4, 2}, point{4, 8})
	p = l1.Intersection(l3)
	assert.Len(t, p, 0)
}

func TestLine2(t *testing.T) {
	l0 := newline(point{0, 0}, point{3, 0})
	l1 := newline(point{0, 0}, point{5, 0})

	expected := []point{
		point{0, 0},
		point{1, 0},
		point{2, 0},
		point{3, 0},
	}

	overlap := l0.Intersection(l1)

	assert.Equal(t, expected, overlap)
	assert.Equal(t, overlap, l1.Intersection(l0))
}
