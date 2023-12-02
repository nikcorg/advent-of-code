package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFmap(t *testing.T) {
	xs := Fmap(func(x string) string { return fmt.Sprintf("hello %s", x) }, []string{"hello", "world"})

	assert.Equal(t, []string{"hello hello", "hello world"}, xs)
}
