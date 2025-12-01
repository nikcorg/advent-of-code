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

	fmt.Fprint(out, "=====[ Day 08 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type antenna struct {
	Pos  util.Point
	Freq rune
}

func (a antenna) String() string {
	return fmt.Sprintf("%c[%d, %d]", a.Freq, a.Pos.X, a.Pos.Y)
}

type gamemap struct {
	Width, Height int
	Antennae      []antenna
}

func parseInput(i string) gamemap {
	s := bufio.NewScanner(strings.NewReader(i))
	w, h, as := 0, 0, []antenna{}
	for s.Scan() {
		line := s.Text()
		if w == 0 {
			w = len(line)
		}

		for i, c := range line {
			if c == '.' {
				continue
			}
			as = append(as, antenna{util.NewPoint(i, h), c})
		}

		h++
	}

	return gamemap{w, h, as}
}

func solveFirst(i string) int {
	m := parseInput(i)

	ps := util.NewSet[util.Point]()
	gs := util.GroupBy(func(a antenna) rune { return a.Freq }, m.Antennae)

	for _, as := range gs {
		if len(as) == 1 {
			continue
		}

		pairs := [][2]antenna{}

		for i, a := range as {
			for _, a0 := range as[i+1:] {
				pairs = append(pairs, [2]antenna{a, a0})
			}
		}

		for _, p := range pairs {
			a, b := p[0], p[1]
			dx, dy := a.Pos.X-b.Pos.X, a.Pos.Y-b.Pos.Y
			dx0, dy0 := dx*-1, dy*-1

			for _, p := range []util.Point{
				a.Pos.Translate(dx0, dy0).Translate(dx0, dy0),
				b.Pos.Translate(dx, dy).Translate(dx, dy),
			} {
				if p.X < 0 || p.X >= m.Width || p.Y < 0 || p.Y >= m.Height {
					continue
				}
				ps.Add(p)
			}
		}
	}

	return len(ps)
}

func solveSecond(i string) int {
	m := parseInput(i)

	ps := util.NewSet[util.Point]()
	gs := util.GroupBy(func(a antenna) rune { return a.Freq }, m.Antennae)

	for _, as := range gs {
		if len(as) == 1 {
			continue
		}

		pairs := [][2]antenna{}

		for i, a := range as {
			for _, a0 := range as[i+1:] {
				pairs = append(pairs, [2]antenna{a, a0})
			}
		}

		for _, p := range pairs {
			a, b := p[0], p[1]
			dx, dy := a.Pos.X-b.Pos.X, a.Pos.Y-b.Pos.Y
			dx0, dy0 := dx*-1, dy*-1

			pa := a.Pos.Translate(dx0, dy0)
			pb := b.Pos.Translate(dx, dy)

			for {
				misses := 0

				if pa.X < 0 || pa.X >= m.Width || pa.Y < 0 || pa.Y >= m.Height {
					misses++
				} else {
					ps.Add(pa)
				}

				if pb.X < 0 || pb.X >= m.Width || pb.Y < 0 || pb.Y >= m.Height {
					misses++
				} else {
					ps.Add(pb)
				}

				if misses == 2 {
					break
				}

				pa = pa.Translate(dx0, dy0)
				pb = pb.Translate(dx, dy)
			}
		}
	}

	return len(ps)
}
