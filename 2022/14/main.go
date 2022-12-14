package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2022/util"
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
	fmt.Fprint(out, "=====[ Day 14 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", solveFirst(input))
	fmt.Fprintf(out, "second: %d\n", solveSecond(input))
	return nil
}

func solveFirst(input string) int {
	m := createMap(bufio.NewScanner(strings.NewReader(input)))

	isFree := func(p util.Point) bool { return m.At(p) == EMPTY }
	grains := 0

	for {
		s := newSand(sandSpawn)

		for s.Drop(isFree) && m.WithinBounds(s.XY()) {
		}

		if !m.WithinBounds(s.XY()) {
			break
		}

		m.Set(s.XY(), SAND)
		grains++
	}

	return grains
}

func solveSecond(input string) int {
	m := createMap(bufio.NewScanner(strings.NewReader(input)))
	m.SetInfinite(true)

	floorY := m.MaxXY().Y + 2
	isFree := func(p util.Point) bool { return m.At(p) == EMPTY && p.Y < floorY }
	grains := 0

	for m.At(sandSpawn) == EMPTY {
		s := newSand(sandSpawn)

		for s.Drop(isFree) {
		}

		m.Set(s.XY(), SAND)
		grains++
	}

	return grains
}

var (
	sandSpawn    = util.NewPoint(500, 0)
	down         = util.NewPoint(0, 1)
	downAndLeft  = util.NewPoint(-1, 1)
	downAndRight = util.NewPoint(1, 1)
	motions      = []util.Point{down, downAndLeft, downAndRight}
)

type sand struct {
	p util.Point
}

func newSand(p util.Point) *sand {
	return &sand{p: p}
}

func (s *sand) XY() util.Point {
	return s.p
}

func (s *sand) Drop(isFree func(util.Point) bool) bool {
	for _, motion := range motions {
		if next := s.p.Add(motion); isFree(next) {
			s.p = next
			return true
		}
	}

	return false
}

func createMap(s *bufio.Scanner) *caveMap {
	m := newCaveMap()

	for s.Scan() {
		turtles := util.Fmap(func(s string) util.Point {
			xy := strings.Split(s, ",")
			return util.NewPoint(util.MustAtoi(xy[0]), util.MustAtoi(xy[1]))
		}, strings.Split(s.Text(), " -> "))

		for _, fromTo := range zip(turtles, turtles[1:]) {
			from := fromTo[0]
			to := fromTo[1]

			if from.X == to.X { // Draw vertically
				for y := min(from.Y, to.Y); y < max(from.Y, to.Y)+1; y++ {
					m.Set(util.NewPoint(from.X, y), WALL)
				}
			} else if from.Y == to.Y { // Draw horizontally
				for x := min(from.X, to.X); x < max(from.X, to.X)+1; x++ {
					m.Set(util.NewPoint(x, from.Y), WALL)
				}
			}
		}
	}

	return m
}

type tile byte

const (
	EMPTY = tile('.')
	SAND  = tile('+')
	WALL  = tile('#')
)

func newCaveMap() *caveMap {
	return &caveMap{tiles: map[util.Point]tile{}, floorY: -1, infinite: false}
}

type caveMap struct {
	tiles    map[util.Point]tile
	minXY    *util.Point
	maxXY    *util.Point
	floorY   int
	infinite bool
}

func (m *caveMap) maybeUpdateMinMax(p util.Point) {
	if m.minXY == nil {
		p := p
		m.minXY = &p
	} else if m.minXY.X > p.X || m.minXY.Y > p.Y {
		m.minXY.X = min(m.minXY.X, p.X)
		m.minXY.Y = min(m.minXY.Y, p.Y)
	}

	if m.maxXY == nil {
		p := p
		m.maxXY = &p
	} else if m.maxXY.X < p.X || m.maxXY.Y < p.Y {
		m.maxXY.X = max(m.maxXY.X, p.X)
		m.maxXY.Y = max(m.maxXY.Y, p.Y)
	}
}

func (m *caveMap) SetInfinite(i bool) {
	m.infinite = i
}

func (m *caveMap) SetFloorY(y int) {
	m.floorY = y
}

func (m *caveMap) WithinBounds(p util.Point) bool {
	if m.infinite && m.floorY > -1 {
		return p.Y < m.floorY
	} else if m.infinite {
		return true
	}

	minXY, maxXY := m.MinXY(), m.MaxXY()
	// Not checking Y-underflow on purpose
	return minXY.X <= p.X && p.X <= maxXY.X && p.Y <= maxXY.Y
}

func (m *caveMap) MinXY() util.Point {
	if m.minXY == nil {
		for p := range m.tiles {
			m.maybeUpdateMinMax(p)
		}
	}
	return *m.minXY
}

func (m *caveMap) MaxXY() util.Point {
	if m.maxXY == nil {
		for p := range m.tiles {
			m.maybeUpdateMinMax(p)
		}
	}
	return *m.maxXY
}

func (m *caveMap) Set(p util.Point, b tile) {
	m.maybeUpdateMinMax(p)
	m.tiles[p] = b
}

func (m *caveMap) At(p util.Point) tile {
	x, ok := m.tiles[p]
	if !ok {
		return EMPTY
	}

	return x
}

func zip[T any](as, bs []T) [][2]T {
	cnt := min(len(as), len(bs))
	cs := make([][2]T, cnt)
	for i := 0; i < cnt; i++ {
		cs[i] = [2]T{as[i], bs[i]}
	}
	return cs
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
