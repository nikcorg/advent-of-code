package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"nikc.org/aoc2024/util"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	first := solveFirst(input)
	second := solveSecond(input)

	fmt.Fprint(out, "=====[ Day 02 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) [][]int {
	s := bufio.NewScanner(strings.NewReader(i))
	lines := [][]int{}
	for s.Scan() {
		ints := util.Fmap(util.MustAtoi, strings.Fields(s.Text()))
		lines = append(lines, ints)
	}
	return lines
}

func solveFirst(i string) int {
	reps := parseInput(i)
	safe := 0

	for _, rep := range reps {
		if rep[0] > rep[1] && descending(rep, 0) {
			safe++
		} else if ascending(rep, 0) {
			safe++
		}
	}

	return safe
}

func solveSecond(i string) int {
	reps := parseInput(i)
	safe := 0

	for _, rep := range reps {
		rep := rep
		if descending(rep, 1) || ascending(rep, 1) {
			safe++
		}
	}

	return safe
}

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three
// L2 (skips):
// Now, the same rules apply as before, except if removing a single level
// from an unsafe report would make it safe, the report instead counts as safe.
func descending(xs []int, skips int) bool {
	ys, i := clone(xs), 1

	for i < len(ys) {
		l, r := ys[i-1], ys[i]
		diff := math.Abs(float64(l) - float64(r))
		if l < r || diff < 1 || diff > 3 {
			if skips > 0 {
				return descending(append(clone(ys[0:i-1]), ys[i:]...), skips-1) || descending(append(clone(ys[0:i]), ys[i+1:]...), skips-1)
			}
			return false
		}
		i++
	}

	return true
}

func ascending(xs []int, skips int) bool {
	ys, i := clone(xs), 1

	for i < len(ys) {
		l, r := ys[i-1], ys[i]
		diff := math.Abs(float64(l) - float64(r))
		if l > r || diff < 1 || diff > 3 {
			if skips > 0 {
				return ascending(append(clone(ys[0:i-1]), ys[i:]...), skips-1) || ascending(append(clone(ys[0:i]), ys[i+1:]...), skips-1)
			}
			return false
		}
		i++
	}

	return true
}

func clone[T any](xs []T) []T {
	ys := make([]T, len(xs))

	copy(ys, xs)

	return ys
}
