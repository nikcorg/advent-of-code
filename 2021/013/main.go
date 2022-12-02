package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
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

	parsed, _ = parseInput(input)
	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2:\n%s\n", result))
	}

	return nil
}

type point struct{ X, Y int }

func (p point) translateX(shift int) point {
	return point{p.X + shift, p.Y}
}

func (p point) translateY(shift int) point {
	return point{p.X, p.Y + shift}
}

type fold struct {
	Axis rune
	Pos  int
}

type puzzleInput struct {
	Folds  []fold
	Dots   map[point]struct{}
	Height int
	Width  int
}

func (pi *puzzleInput) Init() {
	pi.Dots = make(map[point]struct{})
}

var (
	void   = struct{}{}
	reFold = regexp.MustCompile(`fold along (x|y)=(\d+)$`)
)

func parseInput(raw string) (puzzleInput, error) {
	pi := puzzleInput{}
	pi.Init()

	parseFolds := false
	for _, row := range strings.Split(raw, "\n") {
		if row == "" {
			parseFolds = true
			continue
		}

		if parseFolds {
			matches := reFold.FindStringSubmatch(row)
			axis := rune(matches[1][0])
			pos := mustParseInt(matches[2])

			pi.Folds = append(pi.Folds, fold{axis, pos})
		} else {
			parts := strings.Split(row, ",")
			x := mustParseInt(parts[0])
			y := mustParseInt(parts[1])
			pi.Height = max(pi.Height, y+1)
			pi.Width = max(pi.Width, x+1)
			pi.Dots[point{x, y}] = void
		}
	}

	return pi, nil
}

func printMap(input puzzleInput) string {
	out := bytes.Buffer{}

	for y := 0; y < input.Height; y++ {
		for x := 0; x < input.Width; x++ {
			if _, ok := input.Dots[point{x, y}]; ok {
				io.WriteString(&out, "#")
			} else {
				io.WriteString(&out, ".")
			}
		}
		io.WriteString(&out, "\n")
	}

	return out.String()
}

func solveFirst(input puzzleInput) (int, error) {
	input = solve(input, input.Folds[0:1])
	return len(input.Dots), nil
}

func solveSecond(input puzzleInput) (string, error) {
	input = solve(input, input.Folds)
	return printMap(input), nil
}

func solve(input puzzleInput, folds []fold) puzzleInput {
	for _, fold := range folds {
		switch fold.Axis {
		case 'x':
			input.Width = fold.Pos
			for p := range input.Dots {
				if p.X > fold.Pos {
					newPoint := p.translateX(2 * (fold.Pos - p.X))
					input.Dots[newPoint] = void
					delete(input.Dots, p)
				}
			}
		case 'y':
			input.Height = fold.Pos
			for p := range input.Dots {
				if p.Y > fold.Pos {
					newPoint := p.translateY(2 * (fold.Pos - p.Y))
					input.Dots[newPoint] = void
					delete(input.Dots, p)
				}
			}
		}
	}

	return input
}

func mustParseInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
