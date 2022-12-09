package main

import (
	"bufio"
	"context"
	"strings"
	"sync"
	"time"

	"nikc.org/aoc2022/util/stack"
)

type View struct {
	from Point
	dist int
}

func solveSecond(input string) (int, error) {
	m := getTreeMap(bufio.NewScanner(strings.NewReader(input)))

	vds := map[Point]int{}

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := make(chan View, 13) // lucky 13
	rmut := sync.RWMutex{}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case vp := <-r:
				rmut.Lock()
				if ds, ok := vds[vp.from]; ok {
					vds[vp.from] = ds * vp.dist
				} else {
					vds[vp.from] = vp.dist
				}
				rmut.Unlock()
			}
		}
	}()

	maxX := m.Height() - 1
	maxY := m.Width() - 1

	type probeSpec struct {
		lookAt Point
		points <-chan Point
	}

	jobs := stack.New[*probeSpec]()

	for y := 0; y < maxY; y++ {
		jobs.Push(
			&probeSpec{Point{-1, 0}, pointGenerator(Point{0, y}, Point{1, 0}, Point{maxX, y})},
			&probeSpec{Point{1, 0}, pointGenerator(Point{maxX, y}, Point{-1, 0}, Point{0, y})},
		)
	}

	for x := 0; x < maxX; x++ {
		jobs.Push(
			&probeSpec{Point{0, -1}, pointGenerator(Point{x, 0}, Point{0, 1}, Point{x, maxY})},
			&probeSpec{Point{0, 1}, pointGenerator(Point{x, maxY}, Point{0, -1}, Point{x, 0})},
		)
	}

	for n := 0; n < 42; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				nextJob := jobs.Pop()
				if nextJob == nil {
					return
				}
				probeFromPoint(m, nextJob.lookAt, nextJob.points, r)
			}
		}()
	}

	wg.Wait()

	// If the results sink isn't empty, wait until the goroutine has a chance to drain it
	for len(r) > 0 {
		time.Sleep(time.Millisecond * 7)
	}

	rmut.RLock()
	defer rmut.RUnlock()

	best := 0
	for _, d := range vds {
		if d > best {
			best = d
		}
	}

	return best, nil
}

func probeFromPoint(m *treeMap, vec Point, points <-chan Point, r chan<- View) {
	start := <-points

	// Always starts at the edge, which has a 0 distance view
	r <- View{start, 0}

	// Keep track of where the previous tree of different heights were
	seen := map[int]Point{}

	for p := range points {
		ownHgt, _ := m.At(p)
		nbrHgt, _ := m.At(p.Add(vec))

		// Equally toll or shorter than the neighbour, can only see 1 away
		if ownHgt <= nbrHgt {
			r <- View{p, 1}
		} else {
			// Find the nearest, tallest preceding tree
			foundAt := p.DistanceTo(start)
			for h := ownHgt; h <= maxTreeHeight; h++ {
				if mp, ok := seen[h]; !ok {
					continue
				} else if d := p.DistanceTo(mp); d < foundAt {
					foundAt = d
				}
			}

			r <- View{p, int(foundAt)}
		}

		seen[ownHgt] = p
	}
}
