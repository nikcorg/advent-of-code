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

	fmt.Fprint(out, "=====[ Day 06 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type Vector [2]int

type AreaMap struct {
	width, height int
	objects       util.Set[util.Point]
	gLoc          util.Point
	gDir          []Vector
}

var (
	// up, right, down, left
	dirs = []Vector{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
)

func (m AreaMap) IsBlocked(p util.Point) bool {
	return m.objects.Has(p)
}

func (m AreaMap) IsWithin(p util.Point) bool {
	inRange := p.X < 0 || p.X >= m.width || p.Y < 0 || p.Y >= m.height

	return !inRange
}

func (m *AreaMap) GStep() util.Point {
	next := m.gLoc.Translate(m.gDir[0][0], m.gDir[0][1])

	if m.IsBlocked(next) {
		m.gDir = append(m.gDir[1:], m.gDir[0])
	} else {
		m.gLoc = next
	}

	return m.gLoc
}

func parseInput(i string) AreaMap {
	s := bufio.NewScanner(strings.NewReader(i))

	m := AreaMap{0, 0, util.NewSet[util.Point](), util.NewPoint(-1, -1), dirs}

	for s.Scan() {
		line := s.Text()
		if m.width == 0 {
			m.width = len(line)
		}
		m.height++

		for i, c := range line {
			switch c {
			case '#':
				m.objects.Add(util.NewPoint(i, m.height-1))

			case '^':
				m.gLoc = util.NewPoint(i, m.height-1)
			}
		}
	}

	return m
}

func solveFirst(i string) int {
	m := parseInput(i)
	ps := util.NewSet(m.gLoc)

	debug := os.Getenv("DEBUG") != ""

	for {
		next := m.GStep()
		if !m.IsWithin(next) {
			break
		}
		ps.Add(next)

	}

	if debug {
		m.Dump(ps)
	}

	return ps.Size()
}

func solveSecond(i string) int {
	return 6
}

func (m AreaMap) Dump(visited util.Set[util.Point]) {
	var p util.Point
	fmt.Println("\n-[ step ]-------------------------")
	for y := 0; y < m.height; y++ {
		p.Y = y
		for x := 0; x < m.width; x++ {
			p.X = x

			switch {
			case m.IsBlocked(p):
				fmt.Print("#")

			case p.Equals(m.gLoc):
				dumpGuard(m.gDir[0])

			case visited.Has(p):
				fmt.Print("X")

			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func dumpGuard(dir [2]int) {
	switch dir {
	case [2]int{0, -1}:
		fmt.Print("^")
	case [2]int{1, 0}:
		fmt.Print(">")
	case [2]int{0, 1}:
		fmt.Print("v")
	case [2]int{-1, 0}:
		fmt.Print("<")
	}
}
