package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustAtoi(t *testing.T) {
	assert.Equal(t, 42, MustAtoi("42"))

	assert.Panics(t, func() { MustAtoi("beep boop") })
}
