package main

import (
	_ "embed"
	"errors"
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

	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d\n", result))
	}

	return nil
}

type point struct {
	X, Y int
}

func (p point) Equals(op point) bool {
	return p.X == op.X && p.Y == op.Y
}

type puzzleInput struct {
	Boards []*board
	Nums   []int
}

func parseNumbers(unparsed string, splitter *regexp.Regexp) []int {
	nums := []int{}
	for _, s := range splitter.Split(unparsed, -1) {
		s = strings.TrimSpace(s)
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

var (
	commaSplit = regexp.MustCompile(`,`)
	wsSplit    = regexp.MustCompile(`\s+`)
)

func parseInput(raw string) (puzzleInput, error) {
	lines := strings.Split(raw, "\n")
	pi := puzzleInput{}

	// the first line is the numbers input
	pi.Nums = parseNumbers(lines[0], commaSplit)

	nums := []int{}
	for _, line := range lines[2:] {
		line = strings.TrimSpace(line)
		if line == "" {
			b := board{Nums: nums}
			b.Finalise()
			pi.Boards = append(pi.Boards, &b)
			nums = []int{}
			continue
		}

		nums = append(nums, parseNumbers(line, wsSplit)...)
	}

	return pi, nil
}

func solveFirst(input puzzleInput) (int, error) {
	for _, n := range input.Nums {
		for _, b := range input.Boards {
			b.Mark(n)
			if b.Winner() {
				for _, p := range b.Hits {
					b.SetAt(p, 0)
				}

				sum := 0
				for _, num := range b.Nums {
					sum += num
				}

				return sum * n, nil
			}
		}
	}

	return 0, errors.New("no winner")
}

func solveSecond(input puzzleInput) (int, error) {
	winners := make(map[int]struct{})
	toWin := len(input.Boards)

	for _, n := range input.Nums {
		for i, b := range input.Boards {
			if _, ok := winners[i]; ok {
				continue
			}

			b.Mark(n)
			if b.Winner() {
				winners[i] = struct{}{}
				toWin--

				// keep going until all boards have won
				if toWin > 0 {
					continue
				}

				for _, p := range b.Hits {
					b.SetAt(p, 0)
				}

				sum := 0
				for _, num := range b.Nums {
					sum += num
				}

				return sum * n, nil
			}
		}
	}

	return 0, errors.New("no winnder")
}
