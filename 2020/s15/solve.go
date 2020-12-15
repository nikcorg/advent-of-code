package s15

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/nikcorg/aoc2020/utils"
	"github.com/nikcorg/aoc2020/utils/linestream"
)

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

	lineInput := make(chan *linestream.Line, 0)
	solver := getSolver(part)

	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	solution := solver(lineInput)

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
	init := <-inp
	return solve(init.Content(), 2020)
}

func solveSecond(inp linestream.ReadOnlyLineChan) int {
	init := <-inp
	return solve(init.Content(), 30_000_000)
}

const previousMention = 0
const penultimateMention = 1
const unmentioned = -1

func solve(init string, stopAfter int) int {
	numbers := map[int][2]int{}
	lastSpoken := 0
	turn := 1

	for _, n := range strings.Split(init, ",") {
		lastSpoken = utils.MustAtoi(n)
		numbers[lastSpoken] = [2]int{turn, unmentioned}
		turn++
	}

	for ; turn <= stopAfter; turn++ {
		mentions := numbers[lastSpoken]

		if mentions[penultimateMention] == unmentioned && mentions[previousMention]+1 == turn {
			lastSpoken = 0
		} else {
			lm := mentions[previousMention]
			pm := mentions[penultimateMention]
			mentionDiff := lm - pm

			lastSpoken = mentionDiff
		}

		if log, ok := numbers[lastSpoken]; ok {
			log[penultimateMention], log[previousMention] = log[previousMention], turn
			numbers[lastSpoken] = log
		} else {
			numbers[lastSpoken] = [2]int{turn, unmentioned}
		}
	}

	return lastSpoken
}

func last(xs []int) int {
	return xs[len(xs)-1]
}

func penultimate(xs []int) int {
	return xs[len(xs)-2]
}
