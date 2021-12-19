package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	xs := []point{
		point{1, 2},
		point{2, 2},
		point{3, 2},
		point{2, 2},
	}

	uxs := uniq(xs, func(p point) string { return p.String() })

	expected := []point{
		point{1, 2},
		point{2, 2},
		point{3, 2},
	}

	assert.Equal(t, expected, uxs)
}

func TestFilter(t *testing.T) {
	xs := []point{
		point{1, 2},
		point{2, 2},
		point{3, 2},
		point{4, 2},
	}

	uxs := filter(xs, func(p point) bool { return p.X%2 == 0 })

	expected := []point{
		point{2, 2},
		point{4, 2},
	}

	assert.Equal(t, expected, uxs)
}
