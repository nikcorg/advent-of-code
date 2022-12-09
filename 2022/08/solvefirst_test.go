package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveFirst(t *testing.T) {
	v, err := solveFirst(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 21, v)
}
