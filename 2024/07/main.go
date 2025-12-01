package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"slices"
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

	fmt.Fprint(out, "=====[ Day 07 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) [][]int {
	s := bufio.NewScanner(strings.NewReader(i))
	lines := [][]int{}
	for s.Scan() {
		line := s.Text()
		splits := strings.Split(line, ":")
		terms := strings.Fields(splits[1])
		nums := util.Fmap(util.MustAtoi, append([]string{splits[0]}, terms...))
		lines = append(lines, nums)
	}
	return lines
}

func solveFirst(i string) int {
	lines := parseInput(i)
	tot := 0

	for _, line := range lines {
		if !solves(line[0], line[1:], []op{add, mul}) {
			continue
		}
		tot += line[0]
	}

	return tot
}

func solveSecond(i string) int {
	lines := parseInput(i)
	tot := 0

	for _, line := range lines {
		if !solves(line[0], line[1:], []op{add, mul, mrg}) {
			continue
		}
		tot += line[0]
	}

	return tot
}

type op func(a, b int) int

var (
	add op = func(a, b int) int { return a + b }
	mul op = func(a, b int) int { return a * b }
	mrg op = func(a, b int) int { return util.MustAtoi(fmt.Sprintf("%d%d", b, a)) }
)

func solves(eq int, terms []int, ops []op) bool {
	terms = util.CloneSlice(terms)
	slices.Reverse(terms)
	results := solve(terms, ops)

	for _, r := range results {
		if r == eq {
			return true
		}
	}

	return false
}

func solve(terms []int, ops []op) []int {
	switch len(terms) {
	case 1:
		return terms[0:1]
	case 0:
		return []int{0}
	default:
	}

	term := terms[0]

	return util.FlatMap(func(f op) []int {
		return util.Fmap(func(a int) int {
			return f(term, a)
		}, solve(terms[1:], ops))
	}, ops)
}
