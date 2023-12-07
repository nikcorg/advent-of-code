package main

import (
	"bytes"
	_ "embed"
	"math/rand"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestSolveFirst(t *testing.T) {
	plays := parseInput(testInput)
	solution := solveFirst(plays)
	assert.Equal(t, 6440, solution)
}

func TestSolveSecond(t *testing.T) {
	plays := parseInput(testInput)
	solution := solveSecond(plays)
	assert.Equal(t, 5905, solution)
}

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}

func TestParseInput(t *testing.T) {
	gs := parseInput(testInput)
	expect := []play{
		{[]string{"3", "2", "T", "3", "K"}, onePair, 765},
		{[]string{"T", "5", "5", "J", "5"}, threeOfAKind, 684},
		{[]string{"K", "K", "6", "7", "7"}, twoPair, 28},
		{[]string{"K", "T", "J", "J", "T"}, twoPair, 220},
		{[]string{"Q", "Q", "Q", "J", "A"}, threeOfAKind, 483},
	}

	assert.Equal(t, expect, gs)
}

func TestSortHand(t *testing.T) {
	faces := []string{"*", "2", "3", "4", "5", "6", "7", "9", "T", "J", "Q", "K", "A"}

	for i := 0; i < 1; i++ {
		shuffled := slices.Clone(faces)
		for i := range shuffled {
			swapIdx := rand.Intn(len(faces))
			shuffled[i], shuffled[swapIdx] = shuffled[swapIdx], shuffled[i]
		}

		sorted := sortHand(shuffled)

		assert.Equal(t, faces, sorted)
	}
}

func TestLess(t *testing.T) {
	cases := []struct {
		HandA, HandB string
		Exp          bool
	}{
		{"33332", "2AAAA", false},
		{"2AAAA", "33332", true},
		{"77888", "77788", false},
		{"77788", "77888", true},
	}

	for _, c := range cases {
		assert.Equal(t, c.Exp, less(strings.Split(c.HandA, ""), strings.Split(c.HandB, "")), "compare %s to %s", c.HandA, c.HandB)
	}
}

func TestHandKind(t *testing.T) {
	cases := []struct {
		Hand     string
		Expected int
	}{
		// jokers
		{"*****", fiveOfAKind},
		{"****A", fiveOfAKind},
		{"***AA", fiveOfAKind},
		{"**AAA", fiveOfAKind},
		{"*AAAA", fiveOfAKind},
		{"***KA", fourOfAKind},
		{"**KKA", fourOfAKind},
		{"*KKKA", fourOfAKind},
		{"*JJKA", threeOfAKind},
		{"*JJAA", fullHouse},
		{"*2345", onePair},

		// no jokers
		{"AAAAA", fiveOfAKind},
		{"2AAAA", fourOfAKind},
		{"A2AAA", fourOfAKind},
		{"AA2AA", fourOfAKind},
		{"AAA2A", fourOfAKind},
		{"AAAA2", fourOfAKind},
		{"AA222", fullHouse},
		{"AAA22", fullHouse},
		{"A2AA2", fullHouse},
		{"AJAA2", threeOfAKind},
		{"2J2A2", threeOfAKind},
		{"A23A2", twoPair},
		{"A23A3", twoPair},
		{"A53A2", onePair},
		{"K53A2", highCard},
	}

	for _, c := range cases {
		v := handKind(strings.Split(c.Hand, ""))
		assert.Equal(t, c.Expected, v, "hand=%s", c.Hand)
	}
}
