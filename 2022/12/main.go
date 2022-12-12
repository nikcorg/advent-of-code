package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
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

// Calculate traversal cost in a reverse perspective, to enable calculating distances starting
// at the end location.
func traversalCostRev(m *elevationMap) func(dijkstra.Point, dijkstra.Point) (int, error) {
	return func(to dijkstra.Point, from dijkstra.Point) (int, error) {
		goFrom, err := m.At(from)
		if err != nil {
			return 0, err
		}

		goTo, err := m.At(to)
		if err != nil {
			return 0, err
		}

		// Upwards only a height difference of 1 is possible to traverse
		if goFrom < goTo-1 {
			return 0, fmt.Errorf("%w: can't go from %s to %s", errImpossible, string(goFrom), string(goTo))
		}

		// Because a traversable vertice is essentially free (1 or negative), it can be a
		// constant. Reachable is the only meaningful test.
		return 1, nil
	}
}

func solveFirst(input string) (int, error) {
	m, err := newMap(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		return 0, err
	}

	path, _, err := dijkstra.FindPath(m.End(), m.Start(), traversalCostRev(m))

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

	// Run the finder algorithm once to construct the cost graph
	_, costs, err := dijkstra.FindPath(m.End(), m.Start(), traversalCostRev(m))
	if err != nil {
		return 0, err
	}

	shortestPathLen := costs[m.Start()]

	// Because the traversal cost is constant, the distance (cost) at each point of the graph
	// is also the traversal distance
	for p := range m.Points() {
		if v, _ := m.At(p); v != 'a' {
			continue
		}

		if cost, ok := costs[p]; ok && cost < shortestPathLen {
			shortestPathLen = costs[p]
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
					p := (scanOffset + i)
					m.elevations[p] = 'a'
					m.start = dijkstra.NewPoint(p%m.width, p/m.width)
					foundStart = true
				} else if c == 'E' {
					p := (scanOffset + i)
					m.elevations[p] = 'z'
					m.end = dijkstra.NewPoint(p%m.width, p/m.width)
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
