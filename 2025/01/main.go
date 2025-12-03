package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2025/util"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := run(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func run(out io.Writer, input string) error {
	first := solveFirst(input)
	second := solveSecond(input)

	fmt.Fprint(out, "=====[ Day 01 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) []int {
	s := bufio.NewScanner(strings.NewReader(i))
	ms := []int{}
	for s.Scan() {
		l := s.Text()
		switch l[0] {
		case 'L':
			ms = append(ms, util.Atoi(l[1:])*-1)
		case 'R':
			ms = append(ms, util.Atoi(l[1:]))
		}
	}
	return ms
}

func solveFirst(i string) int {
	ms := parseInput(i)
	dial := 50
	zs := 0

	for _, m := range ms {
		nextDial := util.Mod(dial+m, 100)

		if nextDial == 0 {
			zs++
		}

		dial = nextDial
	}

	return zs
}

func solveSecond(i string) int {
	ms := parseInput(i)
	startedOn := 50
	zs := 0

	for _, m := range ms {
		// if the motion is more than a full rotation, we count and deduct the full rotations before
		// applying the remaining motion
		zs += util.Abs(m) / 100
		m = m - (m / 100 * 100)
		endedOn := startedOn + m

		switch {
		// count an exact zero-hit, unless the remaining motion is 0
		case endedOn == 0 && m != 0:
			zs++

		// count a zero-hit, if the remaining motion overflows
		case endedOn > 99:
			zs++

		// count a zero-hit, if the remaining motion underflows, unless we started at 0
		case endedOn < 0 && startedOn > 0:
			zs++
		}

		startedOn = util.Mod(endedOn, 100)
	}

	return zs
}
