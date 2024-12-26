package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

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

	fmt.Fprint(out, "=====[ Day 10 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

var errOutOfBounds = errors.New("out of bounds")

type trailMap struct {
	heads         []util.Point
	grid          []int
	height, width int
}

func (m trailMap) At(p util.Point) (int, error) {
	if p.X < 0 || p.Y < 0 || p.X >= m.width || p.Y >= m.height {
		return 0, errOutOfBounds
	}

	return m.grid[m.width*p.Y+p.X], nil
}

var (
	up    = util.NewPoint(0, -1)
	right = util.NewPoint(1, 0)
	down  = util.NewPoint(0, 1)
	left  = util.NewPoint(-1, 0)
)

func parseInput(i string) trailMap {
	s := bufio.NewScanner(strings.NewReader(i))

	m := trailMap{}

	for s.Scan() {
		line := s.Text()
		if m.width == 0 {
			m.width = len(line)
		}

		for x, c := range line {
			if c == '0' {
				m.heads = append(m.heads, util.NewPoint(x, m.height))
			}
			m.grid = append(m.grid, util.MustAtoi(c))
		}

		m.height++
	}

	return m
}

func solveFirst(i string) int {
	m := parseInput(i)
	tot := 0

	for _, h := range m.heads {
		tot += startFrom(m, h)
	}

	return tot
}

func solveSecond(i string) int {
	m := parseInput(i)
	tot := make(chan int)

	ch := make(chan []util.Point)
	wg := sync.WaitGroup{}

	go func() {
		n := 0
		for range ch {
			n++
		}
		tot <- n
	}()

	for _, h := range m.heads {
		wg.Add(1)
		go func() {
			defer wg.Done()
			traverse(m, h, []util.Point{}, ch)
		}()
	}

	wg.Wait()
	close(ch)

	return <-tot
}

func startFrom(m trailMap, p util.Point) int {
	// fmt.Println("start from", p)
	e := util.NewSet[util.Point]()
	_ = visit(m, 1, p.Add(up), e) +
		visit(m, 1, p.Add(right), e) +
		visit(m, 1, p.Add(down), e) +
		visit(m, 1, p.Add(left), e)

	return len(e)
}

// visit could compose a map of the highest number that can be reached by
// traversing onwards from that square, which would optimise the map traversal
// a lot, by memoizing all trails from that square.
// this does not account for a trail being able to reach more than one trailend.

func visit(m trailMap, expect int, p util.Point, trailends util.Set[util.Point]) int {
	v, err := m.At(p)

	if err != nil {
		return 0
	} else if v != expect {
		return 0
	} else if v == 9 {
		trailends.Add(p)
		return 1
	}

	return util.FoldL(func(acc int, p util.Point) int {
		return acc + visit(m, v+1, p, trailends)
	}, 0, []util.Point{
		p.Add(up), p.Add(right), p.Add(down), p.Add(left),
	})
}

func traverse(m trailMap, next util.Point, path []util.Point, paths chan []util.Point) {
	v, err := m.At(next)
	if err != nil {
		return
	}

	if v == 9 {
		paths <- append(path, next)
		return
	}

	forks := util.Reject(func(next0 util.Point) bool {
		// valid next hops
		v0, err := m.At(next0)
		return err != nil || v0-v != 1
	}, []util.Point{
		next.Add(up), next.Add(right), next.Add(down), next.Add(left),
	})

	for _, fork := range forks {
		traverse(m, fork, append(util.CloneSlice(path), fork), paths)
	}
}
