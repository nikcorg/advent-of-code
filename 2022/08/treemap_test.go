package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"nikc.org/aoc2022/util"
)

func TestTreeMap(t *testing.T) {
	m := getTreeMap(bufio.NewScanner(strings.NewReader(testInput)))

	testcases := []struct {
		p          util.Point
		expected   int
		shouldFail bool
	}{
		{util.Point{X: 0, Y: 0}, 3, false},
		{util.Point{X: 5, Y: 0}, 0, true},
		{util.Point{X: 4, Y: 0}, 3, false},
		{util.Point{X: 4, Y: 4}, 0, false},
		{util.Point{X: 0, Y: 4}, 3, false},
		{util.Point{X: 4, Y: 5}, 0, true},
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
