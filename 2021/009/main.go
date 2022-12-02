package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(input string) error {
	parsed, err := parseInput(input)
	if err != nil {
		return err
	}

	if result, err := solveFirst(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 1: %d\n", result))
	}

	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d\n", result))
	}

	return nil
}

type point struct{ X, Y int }

type puzzleInput struct {
	height int
	width  int
	tiles  []int
}

var (
	errOutOfBounds = errors.New("out of bounds")
)

func (pi puzzleInput) At(p point) (int, error) {
	if p.X < 0 || p.Y < 0 || p.X >= pi.width || p.Y >= pi.height {
		return 0, errOutOfBounds
	}

	tile := p.Y*pi.width + p.X

	if tile < 0 || tile >= len(pi.tiles) {
		return 0, errOutOfBounds
	}

	return pi.tiles[tile], nil
}

func parseInput(raw string) (puzzleInput, error) {
	pi := puzzleInput{}

	for _, row := range strings.Split(raw, "\n") {
		if row == "" {
			break
		}

		tiles := fmap(strings.Split(row, ""), mustAtoi)
		if pi.width == 0 {
			pi.width = len(tiles)
		}
		pi.height++
		pi.tiles = append(pi.tiles, tiles...)
	}

	return pi, nil
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func add1(a int) int {
	return a + 1
}

func fmap[T any, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

func foldl[T any, U any](in []T, init U, f func(U, T) U) U {
	acc := init

	for _, v := range in {
		acc = f(acc, v)
	}

	return acc
}

func tail[T any](xs []T, n int) []T {
	s := len(xs) - n
	return xs[s:]
}

func solveFirst(input puzzleInput) (int, error) {
	lows := []int{}

	for y := 0; y < input.height; y++ {
		for x := 0; x < input.width; x++ {
			at, _ := input.At(point{x, y})
			threshold := 4
			score := 0
			check := []point{
				{x - 1, y},
				{x, y - 1},
				{x + 1, y},
				{x, y + 1},
			}

			for _, p := range check {
				if v, err := input.At(p); err != nil {
					threshold--
				} else if err == nil && v > at {
					score++
				}
			}

			if score == threshold {
				lows = append(lows, at)
			}
		}
	}

	return foldl(fmap(lows, add1), 0, sum), nil
}

func solveSecond(input puzzleInput) (int, error) {
	basins := []int{}

	for y := 0; y < input.height; y++ {
		for x := 0; x < input.width; x++ {
			at, _ := input.At(point{x, y})
			threshold := 4
			score := 0
			check := []point{
				{x - 1, y},
				{x, y - 1},
				{x + 1, y},
				{x, y + 1},
			}

			for _, p := range check {
				if v, err := input.At(p); err != nil {
					threshold--
				} else if err == nil && v > at {
					score++
				}
			}

			if score == threshold {
				visits := make(map[point]struct{})
				size := len(visit(visits, input, point{x, y}))
				basins = append(basins, size)
			}
		}
	}

	sort.Slice(basins, func(a, b int) bool {
		return basins[a] < basins[b]
	})

	return foldl(tail(basins, 3), 1, mul), nil
}

func visit(visited map[point]struct{}, pi puzzleInput, p point) map[point]struct{} {
	if _, ok := visited[p]; ok {
		return visited
	} else if v, err := pi.At(p); err == nil && v < 9 {
		visited[p] = struct{}{}

		visit(visited, pi, point{p.X - 1, p.Y})
		visit(visited, pi, point{p.X + 1, p.Y})
		visit(visited, pi, point{p.X, p.Y - 1})
		visit(visited, pi, point{p.X, p.Y + 1})
	}

	return visited
}
