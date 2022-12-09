package main

import (
	"bufio"
	"errors"
	"strings"

	"nikc.org/aoc2022/util"
)

var (
	errOutOfBounds = errors.New("out of bounds")
)

type treeMap struct {
	width int
	grid  []int
}

func (m *treeMap) Width() int {
	return m.width
}

func (m *treeMap) Height() int {
	return len(m.grid) / m.width
}

func (m *treeMap) At(p Point) (int, error) {
	idx := p.y*m.width + p.x

	if p.x >= m.width || idx < 0 || idx >= len(m.grid) {
		return 0, errOutOfBounds
	}

	return m.grid[idx], nil
}

func getTreeMap(s *bufio.Scanner) *treeMap {
	m := treeMap{0, []int{}}

	for s.Scan() {
		if m.width == 0 {
			m.width = len(s.Text())
		}

		m.grid = append(m.grid, util.Fmap(util.MustAtoi, strings.Split(s.Text(), ""))...)
	}

	return &m
}
