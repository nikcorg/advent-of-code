package linestream

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipEmpty(t *testing.T) {
	const input = `line 1
line 2

line 3

line 4

line 5

line 6


`

	const expectedOutput = `line 1
line 2
line 3
line 4
line 5
line 6`
	output := make(LineChan)

	go func() {
		defer close(output)
		for n, line := range strings.Split(input, "\n") {
			output <- &Line{n, line}
		}
	}()

	expected := strings.Split(expectedOutput, "\n")

	for actual := range SkipEmpty(output) {
		assert.Equal(t, expected[0], actual.Content())
		expected = expected[1:]
	}

	assert.Empty(t, expected)
}
