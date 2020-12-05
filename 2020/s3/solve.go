package s3

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

type Solver struct {
	Ctx context.Context
	Inp string
}

func New(ctx context.Context, inputFilename string) *Solver {
	return &Solver{ctx, inputFilename}
}

func (s *Solver) Solve(part int) error {
	ctx, cancel := context.WithCancel(s.Ctx)
	defer func() { cancel() }()

	inputFile, err := os.Open(s.Inp)
	if err != nil {
		return err
	}
	defer func() { inputFile.Close() }()

	lineInput := make(chan *linestream.Line, 0)
	linestream.New(ctx, bufio.NewReader(inputFile), lineInput)

	filteredInput := make(chan *linestream.Line, 0)
	linestream.SkipEmpty(lineInput, filteredInput)

	solution := <-solveStream(getSolver(ctx, part, filteredInput))

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type partSolver = <-chan int

func getSolver(ctx context.Context, part int, in linestream.LineChan) partSolver {
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
					// fmt.Printf("worker %d: %d\n", n, v)
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
