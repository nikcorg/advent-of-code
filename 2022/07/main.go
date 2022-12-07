package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"nikc.org/aoc2022/util/stack"
)

var (
	//go:embed input.txt
	input string
)

const (
	totalCap, requiredCap uint = 70_000_000, 30_000_000
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	fmt.Fprint(out, "=====[ Day 07 ]=====\n")

	var (
		first, second uint
		err           error
	)

	if first, err = solveFirst(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %d\n", first)

	if second, err = solveSecond(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(input string) (uint, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	sizes, _, err := scanFS(s)
	if err != nil {
		return 0, err
	}

	var total uint

	for _, size := range sizes {
		if size <= 1e5 {
			total += size
		}
	}

	return total, nil
}

func solveSecond(input string) (uint, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	sizes, used, err := scanFS(s)
	if err != nil {
		return 0, err
	}

	var (
		minCandidate uint = math.MaxUint
		threshold    uint = requiredCap - (totalCap - used)
	)

	for _, size := range sizes {
		if size >= threshold && size < minCandidate {
			minCandidate = size
		}
	}

	return minCandidate, nil
}

func scanFS(s *bufio.Scanner) (map[string]uint, uint, error) {
	var (
		sizes = make(map[string]uint)
		dirs  = stack.New[string]()
	)

	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()

		// Parse the terminal output. Any lines not starting with a '$'
		// is a directory listing and the only ones of those that need
		// any attention are those that start with a number. The first
		// three characters are enough to distinguish between the two
		// commands and uninteresting lines in listings.
		switch line[:3] {
		case "dir": // no-op

		case "$ c": // $ cd <dir>
			targetDir := line[5:]
			switch targetDir {
			case "/":
				dirs.Clear()

			case "..":
				dirs.Pop()

			default:
				// Push the full path, to avoid collisions in
				// in the sizes table
				dirs.Push(fmt.Sprintf("%s/%s", dirs.Peek(), targetDir))
			}

		case "$ l": // $ ls
			// Reset the dir size when starting a listing. This is
			// a potential optimisation (skip relisting an already
			// listed dir), but I don't care about microseconds.
			sizes[dirs.Peek()] = 0

		default:
			parts := strings.SplitN(line, " ", 2)
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, 0, err
			}

			// Everything is below the root
			sizes[""] += uint(size)
			dirs.Each(func(d string) {
				sizes[d] += uint(size)
			})
		}
	}

	return sizes, sizes[""], nil
}
