package main

import (
	"bufio"
	"errors"

	"nikc.org/aoc2022/12/dijkstra"
)

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
