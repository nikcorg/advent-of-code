package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strings"

	"nikc.org/aoc2023/util"
)

var (
	//go:embed input.txt
	input string
)

const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	plays := parseInput(input)
	first := solveFirst(plays)
	second := solveSecond(plays)

	fmt.Fprint(out, "=====[ Day 07 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(plays []play) int {
	plays0 := slices.Clone(plays)
	sort.SliceStable(plays0, func(a, b int) bool {
		kindA := handKind(plays0[a].Hand)
		kindB := handKind(plays0[b].Hand)

		if kindA == kindB {
			return less(plays0[a].Hand, plays0[b].Hand)
		}

		return kindA < kindB
	})

	tot := 0
	for i, p := range plays0 {
		tot += (i + 1) * p.Bid
	}
	return tot
}

func solveSecond(plays []play) int {
	plays0 := util.Fmap(func(p play) play {
		p0 := p
		p0.Hand = util.Fmap(func(c string) string {
			if c == "J" {
				return "*"
			}
			return c
		}, p.Hand)
		return p0
	}, plays)

	sort.SliceStable(plays0, func(a, b int) bool {
		kindA := handKind(plays0[a].Hand)
		kindB := handKind(plays0[b].Hand)

		if kindA == kindB {
			return less(plays0[a].Hand, plays0[b].Hand)
		}

		return kindA < kindB
	})

	tot := 0
	for i, p := range plays0 {
		tot += (i + 1) * p.Bid
	}
	return tot
}

func sortHand(hand []string) []string {
	sortedHand := slices.Clone(hand)
	slices.SortStableFunc(sortedHand, func(a, b string) int {
		return cardValue(a) - cardValue(b)
	})
	return sortedHand
}

func less(a, b []string) bool {
	for _, p := range util.Zip(a, b) {
		if cardValue(p[0]) == cardValue(p[1]) {
			continue
		}

		return cardValue(p[0]) < cardValue(p[1])
	}

	return false
}

func handKind(hand []string) int {
	sortedHand := sortHand(hand)

	jokers := 0
	for _, c := range sortedHand {
		if c == "*" {
			jokers++
		}
	}

	switch {
	// ***** or ****A
	case jokers == 5 || jokers == 4:
		return fiveOfAKind
	// ***AA
	case jokers == 3 && sortedHand[3] == sortedHand[4]:
		return fiveOfAKind
	// ***AB
	case jokers == 3:
		return fourOfAKind
	// **AAA
	case jokers == 2 && sortedHand[2] == sortedHand[3] && sortedHand[2] == sortedHand[4]:
		return fiveOfAKind

	// **AAB or **ABB
	case jokers == 2 && (sortedHand[2] == sortedHand[3] || sortedHand[3] == sortedHand[4]):
		return fourOfAKind

	// **ABC
	case jokers == 2:
		return threeOfAKind

	// *AAAA
	case jokers == 1 && sortedHand[1] == sortedHand[2] && sortedHand[2] == sortedHand[3] && sortedHand[2] == sortedHand[4]:
		return fiveOfAKind

	// *AAAB or *ABBB
	case jokers == 1 && ((sortedHand[1] == sortedHand[2] && sortedHand[1] == sortedHand[3]) || (sortedHand[2] == sortedHand[3] && sortedHand[2] == sortedHand[4])):
		return fourOfAKind

	// *AABB
	case jokers == 1 && sortedHand[1] == sortedHand[2] && sortedHand[3] == sortedHand[4]:
		return fullHouse

	// *AABC or *ABBC or *ABCC
	case jokers == 1 && (sortedHand[1] == sortedHand[2] || sortedHand[2] == sortedHand[3] || sortedHand[3] == sortedHand[4]):
		return threeOfAKind
	// *ABCD
	case jokers == 1:
		return onePair

	// AAAA
	case sortedHand[0] == sortedHand[1] && sortedHand[0] == sortedHand[2] && sortedHand[0] == sortedHand[3] && sortedHand[0] == sortedHand[4]:
		return fiveOfAKind

	// AAAAB
	case sortedHand[0] == sortedHand[1] && sortedHand[0] == sortedHand[2] && sortedHand[0] == sortedHand[3]:
		return fourOfAKind

	// ABBBB
	case sortedHand[1] == sortedHand[2] && sortedHand[1] == sortedHand[3] && sortedHand[1] == sortedHand[4]:
		return fourOfAKind

	// AAABB
	case sortedHand[0] == sortedHand[1] && sortedHand[0] == sortedHand[2] && sortedHand[3] == sortedHand[4]:
		return fullHouse

	// AABBB
	case sortedHand[0] == sortedHand[1] && sortedHand[2] == sortedHand[3] && sortedHand[2] == sortedHand[4]:
		return fullHouse

	// AAABB
	case sortedHand[0] == sortedHand[1] && sortedHand[0] == sortedHand[2]:
		return threeOfAKind

	// ABBBC
	case sortedHand[1] == sortedHand[2] && sortedHand[1] == sortedHand[3]:
		return threeOfAKind

	// ABCCC
	case sortedHand[2] == sortedHand[3] && sortedHand[2] == sortedHand[4]:
		return threeOfAKind

	// AABBC or AABCC
	case sortedHand[0] == sortedHand[1] && (sortedHand[2] == sortedHand[3] || sortedHand[3] == sortedHand[4]):
		return twoPair

	// ABBCC
	case sortedHand[1] == sortedHand[2] && sortedHand[3] == sortedHand[4]:
		return twoPair

	// AABCD or ABBCD or ABCCD or ABCDD
	case sortedHand[0] == sortedHand[1] || sortedHand[1] == sortedHand[2] || sortedHand[2] == sortedHand[3] || sortedHand[3] == sortedHand[4]:
		return onePair

	// ABCDE
	default:
		return highCard
	}
}

func cardValue(s string) int {
	switch s {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	case "*":
		return 1
	}

	panic(fmt.Errorf("unknown card: %s", s))
}

type play struct {
	Hand      []string
	Kind, Bid int
}

func parseInput(input string) []play {
	s := bufio.NewScanner(strings.NewReader(input))
	ps := []play{}

	for s.Scan() {
		l := strings.Split(s.Text(), " ")

		hand := strings.Split(l[0], "")
		bid := util.ParseInt(l[1])

		ps = append(ps, play{hand, handKind(hand), bid})
	}

	return ps
}
