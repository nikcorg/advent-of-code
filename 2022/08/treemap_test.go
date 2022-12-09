package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMap(t *testing.T) {
	m := getTreeMap(bufio.NewScanner(strings.NewReader(testInput)))

	testcases := []struct {
		p          Point
		expected   int
		shouldFail bool
	}{
		{Point{0, 0}, 3, false},
		{Point{5, 0}, 0, true},
		{Point{4, 0}, 3, false},
		{Point{4, 4}, 0, false},
		{Point{0, 4}, 3, false},
		{Point{4, 5}, 0, true},
	}

	for _, tc := range testcases {
		v, e := m.At(tc.p)

		if tc.shouldFail {
			assert.Error(t, e)
		} else {
			assert.Equal(t, tc.expected, v)
		}
	}

}
