package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed input.txt
	input string
)

type Tvoid struct{}

var void = struct{}{}

type point struct{ X, Y int }

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

	startTime := time.Now()
	if result, err := solveFirst(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 1: %d (%v)\n", result, time.Since(startTime)))
	}

	startTime = time.Now()
	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d (%v)\n", result, time.Since(startTime)))
	}

	return nil
}

var (
	errOutOfBounds = errors.New("out of bounds")
)

type puzzleInput struct {
	SizeMultiplier int
	Width          int
	Height         int
	Nodes          []int
}

func (pi puzzleInput) MaxWidth() int {
	if pi.SizeMultiplier == 0 {
		return pi.Width
	}

	return pi.Width * pi.SizeMultiplier
}

func (pi puzzleInput) MaxHeight() int {
	if pi.SizeMultiplier == 0 {
		return pi.Height
	}

	return pi.Height * pi.SizeMultiplier
}

func (pi puzzleInput) At(p point) (int, error) {
	if p.X < 0 || p.X >= pi.MaxWidth() || p.Y < 0 || p.Y >= pi.MaxHeight() {
		return 0, errOutOfBounds
	}

	valueAdj := 0

	if p.X >= pi.Width {
		xOverflow := p.X / pi.Width
		valueAdj += xOverflow
		p.X -= xOverflow * pi.Width
	}

	if p.Y >= pi.Height {
		yOverflow := p.Y / pi.Height
		valueAdj += yOverflow
		p.Y -= yOverflow * pi.Height
	}

	node := p.Y*pi.Width + p.X
	val := pi.Nodes[node] + valueAdj

	for val > 9 {
		val -= 9
	}

	return val, nil
}

func (pi *puzzleInput) Grow(mul int) {
	pi.SizeMultiplier = mul
}

func (pi *puzzleInput) Set(p point, v int) {
	pi.Nodes[p.Y*pi.Width+p.X] = v
}

func mustParseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parseInput(raw string) (puzzleInput, error) {
	rows := strings.Split(raw, "\n")
	pi := puzzleInput{
		Width:  len(rows[0]),
		Height: len(rows),
		Nodes:  fmap(mustParseInt, foldl(concat[string], []string{}, fmap(split(""), rows))),
	}
	return pi, nil
}

func solveFirst(input puzzleInput) (int, error) {
	start := point{0, 0}
	end := point{input.Width - 1, input.Height - 1}

	path, err := Dijkstra(input.Width, input.Height, start, end, input.At)
	if err != nil {
		return 0, err
	}

	cost := 0
	for _, p := range path {
		c, _ := input.At(p)

		if p == start {
			continue
		}
		cost += c
	}

	return cost, nil
}

func solveSecond(input puzzleInput) (int, error) {
	input.Grow(5)

	start := point{0, 0}
	end := point{input.MaxWidth() - 1, input.MaxHeight() - 1}

	path, err := Dijkstra(input.MaxWidth(), input.MaxHeight(), start, end, input.At)
	if err != nil {
		return 0, err
	}

	cost := 0
	for _, p := range path {
		c, _ := input.At(p)

		if p == start {
			continue
		}
		cost += c
	}

	return cost, nil
}
