package s10

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/nikcorg/aoc2020/utils/linestream"
	"github.com/nikcorg/aoc2020/utils/slices"
)

const bufSize = 1

type Solver struct {
	ctx context.Context
	out io.Writer
}

func New(ctx context.Context, out io.Writer) *Solver {
	return &Solver{ctx, out}
}

func (s *Solver) SolveFirst(inp io.Reader) error {
	return s.Solve(1, inp)
}

func (s *Solver) SolveSecond(inp io.Reader) error {
	return s.Solve(2, inp)
}

func (s *Solver) Solve(part int, inp io.Reader) error {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	lineInput := make(linestream.LineChan, bufSize)

	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	solve := getSolver(part)
	solution := solve(linestream.SkipEmpty(lineInput))

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(linestream.ReadOnlyLineChan) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(inp linestream.ReadOnlyLineChan) int {
	adapters := slices.SortedIntSlice{0}

	for line := range inp {
		adapters = adapters.Insert(mustAtoi(line.Content()))
	}

	diff1 := 0
	diff3 := 1 // the last adapter is always +3

	for n := 1; n < len(adapters); n++ {
		prev := adapters[n-1]
		curr := adapters[n]
		diff := curr - prev

		if diff > 3 {
			panic(fmt.Errorf("impossible to go from %d to %d, diff %v", prev, curr, diff))
		}

		switch diff {
		case 1:
			diff1++
		case 3:
			diff3++
		}
	}

	return diff1 * diff3
}

func solveSecond(inp linestream.ReadOnlyLineChan) int {
	adapters := slices.SortedIntSlice{0}

	for line := range inp {
		adapters = adapters.Insert(mustAtoi(line.Content()))
	}

	adapters = adapters.Append(adapters.Last() + 3)
	solved := solveForks(adapters)

	return solved
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func solveForks(adapters slices.SortedIntSlice) int {
	solution := 1

	left := 0
	right := left + 1
	seqStart := 0

	for right < len(adapters) {
		leftVal := adapters[left]
		rightVal := adapters[right]

		if diff(leftVal, rightVal) < 3 {
			left++
			right = left + 1
			continue
		}

		sequenceLen := right - seqStart

		if sequenceLen > 5 {
			panic(fmt.Errorf("unhandled sequence length: %d", sequenceLen))
		}

		switch sequenceLen {
		case 5:
			solution *= 7

		case 4:
			solution *= 4

		case 3:
			solution *= 2
		}

		left = right
		right = left + 1
		seqStart = left
	}

	return solution
}

func mustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
