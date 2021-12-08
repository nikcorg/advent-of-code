package s2

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/nikcorg/aoc2020/utils/linestream"
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
	defer func() { cancel() }()

	lineInput := make(chan *linestream.Line, bufSize)
	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	filteredInput := linestream.SkipEmpty(lineInput)

	passwords := make(chan *passwordCandidate)
	convStream(filteredInput, passwords)

	solution := solve(getValidator(part), passwords)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

func getValidator(part int) validator {
	switch part {
	case 1:
		return validateFirst
	case 2:
		return validateSecond
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

func validateFirst(pc *passwordCandidate) bool {
	replacer := regexp.MustCompile(fmt.Sprintf(`[^%s]`, pc.seek))
	filtered := replacer.ReplaceAllString(pc.match, "")
	matches := len(filtered)
	return pc.min <= matches && matches <= pc.max
}

func validateSecond(pc *passwordCandidate) bool {
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

type validator func(*passwordCandidate) bool

func solve(valid validator, in chan *passwordCandidate) int {
	validTotal := 0

	for pc := range in {
		if valid(pc) {
			validTotal++
		}
	}

	return validTotal
}

func mustConv(in string) int {
	v, err := strconv.Atoi(in)

	if err != nil {
		panic(err)
	}

	return v
}

func convStream(in linestream.ReadOnlyLineChan, out chan *passwordCandidate) {
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
					out <- &passwordCandidate{
						min:   mustConv(matches[1]),
						max:   mustConv(matches[2]),
						seek:  matches[3],
						match: matches[4],
					}
				}
			}
		}
	}()
}
