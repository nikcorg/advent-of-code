package main

import (
	"bufio"
	"context"
	"strings"
	"sync"

	"nikc.org/aoc2022/util/stack"
)

func solveFirst(input string) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := getTreeMap(bufio.NewScanner(strings.NewReader(input)))
	wg := sync.WaitGroup{}

	visible := map[Point]struct{}{}
	rchan := make(chan Point)

	go func() {
		for {
			select {
			case p := <-rchan:
				visible[p] = struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()

	maxX := m.Width() - 1
	maxY := m.Height() - 1

	type job struct {
		points <-chan Point
	}

	jobs := stack.New[*job]()

	for y := 0; y < m.Height(); y++ {
		jobs.Push(
			&job{pointGenerator(Point{0, y}, Point{1, 0}, Point{maxX, y})},
			&job{pointGenerator(Point{maxX, y}, Point{-1, 0}, Point{0, y})},
		)
	}

	for x := 0; x < m.Width(); x++ {
		jobs.Push(
			&job{pointGenerator(Point{x, 0}, Point{0, 1}, Point{x, maxY})},
			&job{pointGenerator(Point{x, maxY}, Point{0, -1}, Point{x, 0})},
		)
	}

	for n := 0; n < 42; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				job := jobs.Pop()
				if job == nil {
					return
				}
				probeFromEdge(m, job.points, rchan)
			}
		}()
	}

	wg.Wait()

	return len(visible), nil
}

func probeFromEdge(m *treeMap, points <-chan Point, r chan<- Point) {
	var (
		pt             Point
		lastVisibleHgt int
	)

	pt = <-points
	lastVisibleHgt, _ = m.At(pt)

	// the starting point is always visible by definition
	r <- pt

	for {
		if nextPt, ok := <-points; !ok {
			// stop if there are no more points
			return
		} else {
			pt = nextPt
		}

		if nextHgt, err := m.At(pt); err != nil {
			// exit if we're off the grid
			return
		} else if nextHgt > lastVisibleHgt {
			lastVisibleHgt = nextHgt
			r <- pt
		}

		// nothing more can be seen beyond a max height tall tree
		if lastVisibleHgt == maxTreeHeight {
			return
		}
	}
}
