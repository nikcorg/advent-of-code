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

const lastMention = 0
const penultimateMention = 1
const empty = 0

type mentionRecord [2]int

func solve(init string, stopAfter int) int {
	mentionLog := make([]mentionRecord, 38061056) // "cheating" to avoid having to expand later
	lastSpoken := 0
	turn := 1

	for _, n := range strings.Split(init, ",") {
		lastSpoken = utils.MustAtoi(n)
		mentionLog[lastSpoken] = mentionRecord{turn, empty}
		turn++
	}

	for ; turn <= stopAfter; turn++ {
		mentions := mentionLog[lastSpoken]

		if mentions[penultimateMention] != empty {
			last := mentions[lastMention]
			penultimate := mentions[penultimateMention]
			mentionDiff := last - penultimate

			lastSpoken = mentionDiff
		} else {
			lastSpoken = 0
		}

		// expand log when necessary
		// if cap(mentionLog) < lastSpoken {
		// 	nextCap := cap(mentionLog) * 2
		// 	for nextCap < lastSpoken {
		// 		nextCap *= 2
		// 	}
		// 	expandedLog := make([]mentionRecord, nextCap)
		// 	copy(expandedLog, mentionLog)
		// 	mentionLog = expandedLog
		// }

		if lsMentions := mentionLog[lastSpoken]; lsMentions[lastMention] != empty {
			lsMentions[penultimateMention], lsMentions[lastMention] = lsMentions[lastMention], turn
			mentionLog[lastSpoken] = lsMentions
		} else {
			mentionLog[lastSpoken] = mentionRecord{turn, empty}
		}
	}

	return lastSpoken
}
