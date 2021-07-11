package s9

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

const bufSize = 1

type Solver struct {
	ctx      context.Context
	out      io.Writer
	preamble int
}

func New(ctx context.Context, out io.Writer, preamble int) *Solver {
	return &Solver{ctx, out, preamble}
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

	numbers := make(chan int, bufSize)
	go convStream(linestream.SkipEmpty(lineInput), numbers)

	solve := getSolver(part)
	solution := <-solve(numbers, s.preamble)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))
	return nil
}

type solver func(<-chan int, int) <-chan int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst

	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("invalid part %d", part))
}

func _solveFirst(inp []int, trg int) int {
	for i, n := range inp {
		for _, m := range inp[i+1:] {
			if n+m == trg {
				return trg
			}
		}
	}

	return -1
}

func _max(ns []int) int {
	var r int = ns[0]

	for _, n := range ns[1:] {
		if n > r {
			r = n
		}
	}

	return r
}

func _min(ns []int) int {
	var r int = ns[0]

	for _, n := range ns[1:] {
		if n < r {
			r = n
		}
	}

	return r
}

func _solveSecond(inp []int, trg int) int {
	for i, n := range inp {
		if i+1 >= len(inp) {
			return -1
		}

		tail := inp[i+1:]
		total := n

		for j, m := range tail {
			total += m
			if total == trg {
				rng := inp[i : i+j+2]
				from := _min(rng)
				to := _max(rng)

				return from + to
			}
		}
	}

	return -1
}

func solveSecond(inp <-chan int, preamble int) <-chan int {
	out := make(chan int)
	dup := make(chan int)
	trg := 0

	ctx, cancel := context.WithCancel(context.Background())

	// FIXME: restructure the solvers to run both parts in one go
	// it's silly to solve the first part again, but with the current
	// structure of the program, the solutions are run independently
	go func() {
		defer cancel()
		trg = <-solveFirst(dup, preamble)
	}()

	go func() {
		nums := []int{}

		for num := range inp {
			nums = append(nums, num)

			select {
			case <-ctx.Done():
				res := _solveSecond(nums, trg)
				if res > 0 {
					out <- res
					return
				}

			case dup <- num:
			}
		}
	}()

	return out
}

func solveFirst(inp <-chan int, preamble int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		nums := []int{}
		for num := range inp {
			if len(nums) < preamble {
				nums = append(nums, num)
				continue
			}

			if res := _solveFirst(nums, num); res != num {
				out <- num
				return
			}

			nums = append(nums[1:], num)
		}
	}()

	return out
}

func convStream(inp linestream.ReadOnlyLineChan, out chan<- int) {
	defer close(out)

	for line := range inp {
		num, err := strconv.Atoi(line.Content())
		if err != nil {
			panic(err)
		}
		out <- num
	}
}
