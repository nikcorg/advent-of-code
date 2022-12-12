package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"nikc.org/aoc2022/12/dijkstra"
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
	fmt.Fprint(out, "=====[ Day 12 ]=====\n")

	var (
		first, second int
		err           error
	)

	if first, err = solveFirst(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %d\n", first)

	if second, err = solveSecond(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

var errImpossible = errors.New("impossible")

func traversalCost(m *elevationMap) func(dijkstra.Point, dijkstra.Point) (int, error) {
	return func(from dijkstra.Point, to dijkstra.Point) (int, error) {
		goFrom, err := m.At(from)
		if err != nil {
			return 0, err
		} else if goFrom == 'S' {
			goFrom = 'a'
		} else if goFrom == 'E' {
			goFrom = 'a'
		}

		goTo, err := m.At(to)
		if err != nil {
			return 0, err
		} else if goTo == 'E' {
			goTo = 'z'
		} else if goTo == 'S' {
			goTo = 'a'
		}

		if goFrom < goTo-1 {
			return 0, fmt.Errorf("%w: can't go from %s to %s", errImpossible, string(goFrom), string(goTo))
		}

		// Because the route involves drops and Dijkstra can't handle negative costs,
		// we need to offset the cost by 26 to avoid going below zero.
		cost := 26 + int(goFrom) - int(goTo)

		return cost, nil
	}
}

func solveFirst(input string) (int, error) {
	m, err := newMap(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		return 0, err
	}

	path, err := dijkstra.Dijkstra(m.Width(), m.Height(), m.Start(), m.End(), m.Points(), traversalCost(m))

	if err != nil {
		return 0, err
	}

	return len(path), nil
}

func solveSecond(input string) (int, error) {
	m, err := newMap(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		return 0, err
	}

	locs := []dijkstra.Point{}
	for p := range m.Points() {
		if v, _ := m.At(p); v == 'a' {
			locs = append(locs, p)
		}
	}

	// What we should do here is to retain the cost from square to square map to keep the
	// recalculation costs down, but who cares about a few CPU cycles, right? At least not
	// until it becomes necessary for completion within a reasonable time...
	shortestPathLen := math.MaxInt
	for _, p := range locs {
		path, err := dijkstra.Dijkstra(m.Width(), m.Height(), p, m.End(), m.Points(), traversalCost(m))
		if err != nil {
			// some starting points will not have a path
			continue
		}
		if len(path) < shortestPathLen {
			shortestPathLen = len(path)
		}
	}

	return shortestPathLen, nil
}

type elevationMap struct {
	width      int
	elevations []byte
	start      dijkstra.Point
	end        dijkstra.Point
}

var (
	errOutOfBounds        = errors.New("out of bounds")
	errStartOrEndNotFound = errors.New("start or end not found")
	errEmptyInput         = errors.New("empty input")
)

func (m *elevationMap) Height() int {
	return len(m.elevations) / m.width
}
func (m *elevationMap) Width() int {
	return m.width
}

func (m *elevationMap) Start() dijkstra.Point {
	return m.start
}

func (m *elevationMap) End() dijkstra.Point {
	return m.end
}

func (m *elevationMap) Points() <-chan dijkstra.Point {
	maxX, maxY := m.width, len(m.elevations)/m.width

	c := make(chan dijkstra.Point)

	go func() {
		defer close(c)
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				c <- dijkstra.NewPoint(x, y)
			}
		}
	}()

	return c
}

func (m *elevationMap) At(p dijkstra.Point) (byte, error) {
	if p.X < 0 || p.X >= m.width || p.Y < 0 {
		return byte(0), errOutOfBounds
	}

	i := p.X + p.Y*m.width

	if i < 0 || i >= len(m.elevations) {
		return byte(0), errOutOfBounds
	}

	return m.elevations[i], nil
}

func newMap(s *bufio.Scanner) (*elevationMap, error) {
	m := &elevationMap{}

	if firstScan := s.Scan(); !firstScan && s.Err() != nil {
		return nil, s.Err()
	} else if !firstScan {
		return nil, errEmptyInput
	}

	m.width = len(s.Text())
	m.elevations = []byte(s.Text())

	foundStart := false
	foundEnd := false
	scanOffset := 0

	for s.Scan() {
		m.elevations = append(m.elevations, []byte(s.Text())...)

		if !foundStart || !foundEnd {
			for i, c := range m.elevations[scanOffset:] {
				if c == 'S' {
					m.start = dijkstra.NewPoint((scanOffset+i)%m.width, (scanOffset+i)/m.width)
					foundStart = true
				} else if c == 'E' {
					m.end = dijkstra.NewPoint((scanOffset+i)%m.width, (scanOffset+i)/m.width)
					foundEnd = true
				}

				if foundStart && foundEnd {
					break
				}
			}
			scanOffset = len(m.elevations)
		}
	}

	if !foundStart || !foundEnd {
		return nil, errStartOrEndNotFound
	}

	return m, nil
}
