package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2023/util"
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
	seqs := parseInput(input)
	first := solveFirst(seqs)
	second := solveSecond(seqs)

	fmt.Fprint(out, "=====[ Day 09 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(seqs [][]int) int {
	tot := 0

	for _, seq := range seqs {
		tot += solveSeqFwd(seq)
	}

	return tot
}

func solveSecond(seqs [][]int) int {
	tot := 0

	for _, seq := range seqs {
		nextN := solveSeqRev(seq)
		tot += nextN
	}

	return tot
}

func solveSeqFwd(seq []int) int {
	next := []int{}
	prev := seq[0]
	zeroSeq := true

	for _, n := range seq[1:] {
		nextN := n - prev
		next = append(next, nextN)
		prev = n

		if zeroSeq && nextN != 0 {
			zeroSeq = false
		}
	}

	if zeroSeq {
		return last(seq)
	}

	return last(seq) + solveSeqFwd(next)
}

func solveSeqRev(seq []int) int {
	next := []int{}
	prev := seq[0]
	zeroSeq := true

	for _, n := range seq[1:] {
		nextN := n - prev
		next = append(next, nextN)
		prev = n

		if zeroSeq && nextN != 0 {
			zeroSeq = false
		}
	}

	var nextN int
	if !zeroSeq {
		nextN = first(seq) - solveSeqRev(next)
	} else {
		nextN = first(seq)
	}

	return nextN
}

func parseInput(i string) [][]int {
	s := bufio.NewScanner(strings.NewReader(i))

	out := [][]int{}

	for s.Scan() {
		l := util.Fmap(util.MustParseInt, strings.Split(s.Text(), " "))
		out = append(out, l)
	}

	return out
}

func first[T any](xs []T) T {
	return xs[0]
}

func last[T any](xs []T) T {
	return xs[len(xs)-1]
}
