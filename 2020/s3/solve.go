package s3

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"sync"

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
	defer func() { cancel() }()

	lineInput := make(chan *linestream.Line, 0)
	solver := getSolver(ctx, part, linestream.SkipEmpty(lineInput))

	// Defer initialising the stream, so that all listeners have
	// time to bind themselves to the Muxxer used by the second solver.
	// Delayed signup to the Muxxer might lose messages, as it does
	// not have internal buffering.
	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	solution := <-solveStream(solver)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type partSolver = <-chan int

func getSolver(ctx context.Context, part int, in linestream.ReadOnlyLineChan) partSolver {
	muxxer := linestream.NewMuxxer(in)

	switch part {
	case 1:
		return solveSlope(3, 1, muxxer.Recv())

	case 3:
		return solveSlope(1, 2, muxxer.Recv())

	case 2:
		return multiSolve([]partSolver{
			solveSlope(1, 1, muxxer.Recv()),
			solveSlope(3, 1, muxxer.Recv()),
			solveSlope(5, 1, muxxer.Recv()),
			solveSlope(7, 1, muxxer.Recv()),
			solveSlope(1, 2, muxxer.Recv()),
		})
	}

	panic(fmt.Errorf("invalid part: %d", part))
}

const tree = '#'

func isCollision(line string, x int) int {
	pos := x % len(line)
	collided := line[pos] == tree

	if collided {
		return 1
	}

	return 0
}

func multiSolve(inputs []<-chan int) <-chan int {
	out := make(chan int)

	var results []int = make([]int, len(inputs))
	var workers int = len(inputs)

	wg := &sync.WaitGroup{}
	wg.Add(workers)

	for n, stream := range inputs {
		n := n
		stream := stream
		go func() {
			defer wg.Done()
			for {
				select {
				case v, ok := <-stream:
					if !ok {
						return
					}

					results[n] = v
				}
			}
		}()
	}

	go func() {
		defer close(out)

		wg.Wait()

		product := 1
		for _, r := range results {
			product *= r
		}
		out <- product
	}()

	return out
}

func solveStream(in <-chan int) <-chan int {
	out := make(chan int)

	var result int
	go func() {
		defer func() {
			out <- result
			close(out)
		}()

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}

				result = v
			}
		}
	}()

	return out
}

func solveSlope(slopeHoriz, slopeVert int, in linestream.ReadOnlyLineChan) <-chan int {
	out := make(chan int)

	collisions := 0
	xpos := 0
	ypos := 0

	go func() {
		defer func() {
			out <- collisions
			close(out)
		}()

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}

				if ypos%slopeVert == 0 {
					collisions += isCollision(v.Content(), xpos)
					xpos += slopeHoriz
				}
				ypos++
			}
		}
	}()

	return out
}
