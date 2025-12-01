package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
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

	fmt.Fprint(out, "=====[ Day 04 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) (map[util.Point]rune, int) {
	s := bufio.NewScanner(strings.NewReader(i))
	m, w, x, y := map[util.Point]rune{}, 0, 0, 0

	for s.Scan() {
		for _, c := range s.Text() {
			m[util.NewPoint(x, y)] = c
			x++
		}
		if y == 0 {
			w = x
		}
		y++
		x = 0
	}

	return m, w
}

func solveFirst(i string) int {
	m, _ := parseInput(i)

	tot := 0

	for p, c := range m {
		if c == 'X' {
			// x y
			tests := []bool{
				traverseFrom(p, m, -1, -1),
				traverseFrom(p, m, 0, -1),
				traverseFrom(p, m, 1, -1),
				traverseFrom(p, m, 1, 0),
				traverseFrom(p, m, 1, 1),
				traverseFrom(p, m, 0, 1),
				traverseFrom(p, m, -1, 1),
				traverseFrom(p, m, -1, 0),
			}

			for _, b := range tests {
				if b {
					tot++
				}
			}
		}
	}

	return tot
}

func solveSecond(i string) int {
	m, _ := parseInput(i)
	tot := 0

	for p, c := range m {
		if c == 'A' && sampleFrom(p, m) {
			tot++
		}
	}

	return tot
}

func sampleFrom(p util.Point, m map[util.Point]rune) bool {
	// M.S
	// .p.
	// M.S
	if m[p.Translate(-1, -1)] == 'M' && m[p.Translate(1, -1)] == 'S' &&
		m[p.Translate(-1, 1)] == 'M' && m[p.Translate(1, 1)] == 'S' {
		return true
	}

	// S.M
	// .p.
	// S.M
	if m[p.Translate(-1, -1)] == 'S' && m[p.Translate(1, -1)] == 'M' &&
		m[p.Translate(-1, 1)] == 'S' && m[p.Translate(1, 1)] == 'M' {
		return true
	}

	// S.S
	// .p.
	// M.M
	if m[p.Translate(-1, -1)] == 'S' && m[p.Translate(1, -1)] == 'S' &&
		m[p.Translate(-1, 1)] == 'M' && m[p.Translate(1, 1)] == 'M' {
		return true
	}

	// M.M
	// .p.
	// S.S
	if m[p.Translate(-1, -1)] == 'M' && m[p.Translate(1, -1)] == 'M' &&
		m[p.Translate(-1, 1)] == 'S' && m[p.Translate(1, 1)] == 'S' {
		return true
	}

	return false
}

func traverseFrom(p util.Point, m map[util.Point]rune, dx, dy int) bool {
	next := []rune{'M', 'A', 'S'}

	for _, c := range next {
		p = p.Translate(dx, dy)
		if pc, ok := m[p]; !ok || pc != c {
			return false
		}
	}

	return true
}
