package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	data            []byte
	packetBoundary  int
	messageBoundary int
}{
	{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7, 19},
	{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5, 23},
	{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 6, 23},
	{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10, 29},
	{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11, 26},
	{[]byte("abcdaaabbbabaabbbaabbab"), 4, -1},
	{[]byte("abbbaaabbbabaabbbaabbab"), -1, -1},
	{[]byte(""), -1, -1},
}

func TestMainWithErr(t *testing.T) {
	for _, tc := range testCases {
		err := mainWithErr(&bytes.Buffer{}, string(tc.data))
		if tc.packetBoundary < 0 || tc.messageBoundary < 0 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestSolveFirst(t *testing.T) {
	for _, tc := range testCases {
		v, err := solveFirst(tc.data)
		if tc.packetBoundary < 0 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.packetBoundary, v)
		}
	}
}

func TestSolveSecond(t *testing.T) {
	for _, tc := range testCases {
		v, err := solveSecond(tc.data)
		if tc.messageBoundary < 0 {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.messageBoundary, v)
		}
	}
}
