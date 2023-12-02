package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveSecond(t *testing.T) {
	v, err := solveSecond(testInput)

	assert.NoError(t, err)
	assert.Equal(t, 8, v)
}
