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

type partSolver func(x int, xs []int) ([]int, bool)

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
	multiplicands, product := splitResult(<-solveStream(getSolver(part), convStream(input)))

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

func solveSecond(x int, xs []int) ([]int, bool) {
	for i, n := range xs {
		for _, m := range xs[i+1:] {
			if x+n+m == 2020 {
				return []int{x, n, m, x * n * m}, true
			}
		}
	}

	return []int{}, false
}

func solveFirst(x int, xs []int) ([]int, bool) {
	for _, n := range xs {
		if n+x == 2020 {
			return []int{n, x, n * x}, true
		}
	}

	return []int{}, false
}

func solveStream(solve partSolver, in chan int) chan []int {
	out := make(chan []int, 1)

	var inputs []int

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

				inputs = append(inputs, v)
			}
		}
	}()

	return out
}

func convStream(in linestream.LineChan) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}

				intval, err := strconv.Atoi(v.Content())
				if err != nil {
					panic(fmt.Errorf("error converting %v to int: %v", v, err))
				}
				out <- intval
			}
		}
	}()

	return out
}
