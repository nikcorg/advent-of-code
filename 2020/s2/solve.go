package s2

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

type Solver struct {
	Ctx context.Context
	Inp string
}

func New(ctx context.Context, inputFile string) *Solver {
	return &Solver{ctx, inputFile}
}

func (s *Solver) Solve(part int) error {
	ctx, cancel := context.WithCancel(s.Ctx)
	defer func() { cancel() }()

	inputFile, err := os.Open(s.Inp)
	if err != nil {
		return err
	}
	defer func() { inputFile.Close() }()

	solution := <-solveStream(getSolver(part), convStream(linestream.SkipEmpty(linestream.WithDoneSignalling(cancel, linestream.New(ctx, bufio.NewReader(inputFile))))))

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part: %d", part))
}

func stringify(in []int) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = fmt.Sprintf("%d", v)
	}
	return out
}

func splitResult(xs []int) ([]int, int) {
	last := len(xs) - 1
	solution := xs[last]
	rest := xs[0:last]

	return rest, solution
}

func solveFirst(pc passwordCandidate) bool {
	replacer := regexp.MustCompile(fmt.Sprintf(`[^%s]`, pc.seek))
	filtered := replacer.ReplaceAllString(pc.match, "")
	matches := len(filtered)
	return pc.min <= matches && matches <= pc.max
}

func solveSecond(pc passwordCandidate) bool {
	leftMatch := string(pc.match[pc.min-1]) == pc.seek
	rightMatch := string(pc.match[pc.max-1]) == pc.seek
	return (leftMatch || rightMatch) && !(rightMatch && leftMatch)
}

type passwordCandidate struct {
	min   int
	max   int
	seek  string
	match string
}

type solver func(passwordCandidate) bool

func solveStream(solve solver, in chan passwordCandidate) chan int {
	out := make(chan int, 1)
	validTotal := 0

	go func() {
		defer func() {
			out <- validTotal
			close(out)
		}()

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				if solve(v) {
					validTotal++
				}
			}
		}
	}()

	return out
}

func mustConv(in string) int {
	v, err := strconv.Atoi(in)

	if err != nil {
		panic(err)
	}

	return v
}

func convStream(in linestream.LineChan) chan passwordCandidate {
	out := make(chan passwordCandidate)

	splitter := regexp.MustCompile(`^(\d+)-(\d+) (.): (.*)$`)

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				if v != nil {
					matches := splitter.FindStringSubmatch(v.Content())
					out <- passwordCandidate{
						min:   mustConv(matches[1]),
						max:   mustConv(matches[2]),
						seek:  matches[3],
						match: matches[4],
					}
				}
			}
		}
	}()

	return out
}
