package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	//go:embed input.txt
	input             string
	errNoIntersection = errors.New("no intersection")
	matchPoints       = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)
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

type puzzleinput []line

func parseInput(raw string) (puzzleinput, error) {
	pi := puzzleinput{}

	for _, row := range strings.Split(raw, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		matches := matchPoints.FindStringSubmatch(row)

		if matches == nil {
			return nil, fmt.Errorf("no match found in: %s", row)
		}

		pi = append(pi, newline(pointFromString(matches[1], matches[2]), pointFromString(matches[3], matches[4])))
	}

	fmt.Println(pi)

	return pi, nil
}

func solveFirst(input puzzleinput) (int, error) {
	intersects := map[point][]line{}
	lines := filter(input, func(l line) bool {
		return l.IsHorizontal() || l.IsVertical()
	})

	for _, lineA := range lines {
		for _, lineB := range lines {
			if lineA == lineB {
				continue
			}
			for _, p := range lineA.Intersection(lineB) {
				if ps, ok := intersects[p]; ok {
					intersects[p] = append(ps, lineA, lineB)
				} else {
					intersects[p] = []line{lineA, lineB}
				}
			}
		}
	}

	count := 0

	for p, ls := range intersects {
		lsu := uniq(ls, func(l line) line {
			return l
		})

		if len(lsu) > 1 {
			count++
			fmt.Println(p, " ", lsu)
		}
	}

	return count, nil
}

func solveSecond(lines puzzleinput) (int, error) {
	intersects := map[point][]line{}

	for _, lineA := range lines {
		for _, lineB := range lines {
			if lineA == lineB {
				continue
			}
			for _, p := range lineA.Intersection(lineB) {
				if ps, ok := intersects[p]; ok {
					intersects[p] = append(ps, lineA, lineB)
				} else {
					intersects[p] = []line{lineA, lineB}
				}
			}
		}
	}

	count := 0

	for p, ls := range intersects {
		lsu := uniq(ls, func(l line) line {
			return l
		})

		if len(lsu) > 1 {
			count++
			fmt.Println(p, " ", lsu)
		}
	}

	return count, nil
}
