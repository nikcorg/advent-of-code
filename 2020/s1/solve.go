package s1

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

type void struct{}
type partSolver func(x int, xs map[int]void) ([]int, bool)

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

	input := make(chan *linestream.Line, 0)
	linestream.New(ctx, bufio.NewReader(inp), input)
	multiplicands, product := splitResult(<-solveStream(getSolver(part), convStream(linestream.SkipEmpty(input))))

	io.WriteString(s.out, fmt.Sprintf("solution: %s=%d\n", strings.Join(stringify(multiplicands), "*"), product))

	return nil
}

func getSolver(part int) partSolver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	default:
		panic(fmt.Errorf("invalid part: %d", part))
	}
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

func solveSecond(x int, xs map[int]void) ([]int, bool) {
	n := 2020 - x

	for k := range xs {
		m := n - k

		if m < 0 {
			continue
		}

		if _, ok := xs[m]; ok {
			// sort the numbers before returning for a stable result
			if x > m {
				x, m = m, x
			}
			if m > k {
				m, k = k, m
			}

			return []int{x, m, k, x * m * k}, true
		}
	}

	return []int{}, false
}

func solveFirst(x int, xs map[int]void) ([]int, bool) {
	k := 2020 - x

	if _, ok := xs[k]; ok {
		if x > k {
			x, k = k, x
		}
		return []int{x, k, x * k}, true
	}

	return []int{}, false
}

func solveStream(solve partSolver, in chan int) chan []int {
	out := make(chan []int, 1)

	var inputs = make(map[int]void)
	var Void = struct{}{}

	go func() {
		defer close(out)

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}

				if sol, ok := solve(v, inputs); ok {
					out <- sol
					return
				}

				inputs[v] = Void
			}
		}
	}()

	return out
}

func convStream(in linestream.ReadOnlyLineChan) chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			intval, err := strconv.Atoi(v.Content())
			if err != nil {
				panic(fmt.Errorf("error converting %v to int: %v", v, err))
			}
			out <- intval
		}
	}()

	return out
}
