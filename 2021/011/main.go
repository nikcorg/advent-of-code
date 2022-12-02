package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
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

	// Need to reparse for a pristine state for part 2
	parsed, _ = parseInput(input)
	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d\n", result))
	}

	return nil
}

type point struct{ X, Y int }
type puzzleInput struct {
	Grid    []int
	Height  int
	Width   int
	Flashes int
	Resets  map[point]struct{}
}

var (
	errOutOfBounds = errors.New("out of bounds")
)

func (pi *puzzleInput) WithinBounds(p point) bool {
	return 0 <= p.X && p.X < pi.Width && 0 <= p.Y && p.Y < pi.Height
}

func (pi *puzzleInput) Reset(p point) {
	pi.Resets[p] = struct{}{}
	pi.Grid[p.Y*pi.Width+p.X] = 0
}

func (pi *puzzleInput) At(p point) int {
	return pi.Grid[p.Y*pi.Width+p.X]
}

func (pi *puzzleInput) Inc(p point) int {
	if _, ok := pi.Resets[p]; !ok {
		pi.Grid[p.Y*pi.Width+p.X]++
	}
	return pi.At(p)
}

func (pi *puzzleInput) Flash(p point) {
	pi.Flashes++
	pi.Reset(p)

	neighbours := filter(pi.WithinBounds, []point{
		{p.X - 1, p.Y},
		{p.X - 1, p.Y - 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y + 1},
		{p.X, p.Y + 1},
		{p.X - 1, p.Y + 1},
	})

	for _, p := range neighbours {
		if pi.Inc(p) > 9 {
			pi.Flash(p)
		}
	}
}

func parseInput(raw string) (puzzleInput, error) {
	pi := puzzleInput{}
	rows := strings.Split(raw, "\n")
	pi.Width = len(rows[0])
	for _, row := range rows {
		pi.Grid = append(pi.Grid, fmap(mustParseInt, strings.Split(row, ""))...)
		pi.Height++
	}
	return pi, nil
}

func solveFirst(input puzzleInput) (int, error) {
	turns := 100

	for turn := 0; turn < turns; turn++ {
		// fmt.Println("turn ", turn, " flashes ", input.Flashes)
		// for y := 0; y < input.Height; y++ {
		// 	i := y * input.Width
		// 	fmt.Println(input.Grid[i : i+input.Width])
		// }

		input.Resets = map[point]struct{}{}

		for i := 0; i < input.Width*input.Height; i++ {
			y := i / input.Width
			x := i - y*input.Width
			p := point{x, y}

			if input.Inc(p) > 9 {
				input.Flash(p)
			}
		}

	}

	return input.Flashes, nil
}

func solveSecond(input puzzleInput) (int, error) {
	for turn := 0; true; turn++ {
		if len(input.Resets) == len(input.Grid) {
			return turn, nil
		}

		input.Resets = map[point]struct{}{}

		for i := 0; i < input.Width*input.Height; i++ {
			y := i / input.Width
			x := i - y*input.Width
			p := point{x, y}

			if input.Inc(p) > 9 {
				input.Flash(p)
			}
		}
	}

	return 0, errors.New("impossible")
}

func filter[T any](f func(T) bool, xs []T) []T {
	out := []T{}
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}
	return out
}

func fmap[T any, U any](f func(T) U, xs []T) []U {
	out := make([]U, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}

func mustParseInt(x string) int {
	v, _ := strconv.Atoi(x)
	return v
}
