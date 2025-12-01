package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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

	fmt.Fprint(out, "=====[ Day 01 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) ([]int, []int) {
	s := bufio.NewScanner(strings.NewReader(i))
	lefts, rights := []int{}, []int{}
	for s.Scan() {
		sints := strings.Fields(s.Text())

		i, _ := strconv.Atoi(sints[0])
		lefts = append(lefts, i)

		i, _ = strconv.Atoi(sints[1])
		rights = append(rights, i)
	}

	return lefts, rights
}

func solveFirst(i string) int {
	lefts, rights := parseInput(i)
	slices.Sort(lefts)
	slices.Sort(rights)

	tot := 0

	for i := range lefts {
		tot += int(math.Abs(float64(lefts[i]) - float64(rights[i])))
	}

	return tot
}

func solveSecond(i string) int {
	lefts, rights := parseInput(i)

	freqs := freqMap(rights)
	tot := 0

	for _, l := range lefts {
		tot += l * freqs[l]
	}

	return tot
}

func freqMap(xs []int) map[int]int {
	fs := map[int]int{}

	for _, x := range xs {
		fs[x] += 1
	}

	return fs
}
