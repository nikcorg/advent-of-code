package s6

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

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
	ctx := context.Background()

	lineInput := make(linestream.LineChan, 1)
	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	groups := convStream(linestream.WithChunking(lineInput))

	solve := getSolver(part)
	solution := <-solve(groups)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type group struct {
	answers map[string]int
	chunk   linestream.Chunk
}

type solver func(<-chan group) <-chan int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(inp <-chan group) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		total := 0
		for group := range inp {
			keys := []string{}
			for k := range group.answers {
				keys = append(keys, k)
			}
			total += len(keys)
		}

		out <- total
	}()

	return out
}

func solveSecond(inp <-chan group) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		total := 0

		for group := range inp {
			for _, as := range group.answers {
				if as == len(group.chunk) {
					total++
				}
			}
		}

		out <- total
	}()

	return out
}

func convStream(inp linestream.ReadOnlyChunkedLineChan) <-chan group {
	out := make(chan group)

	go func() {
		defer close(out)

		for chunk := range inp {
			g := group{}
			g.answers = make(map[string]int)

			for _, line := range chunk {
				g.chunk = chunk
				for _, answer := range strings.Split(line.Content(), "") {
					if as, ok := g.answers[answer]; ok {
						g.answers[answer] = as + 1
					} else {
						g.answers[answer] = 1
					}
				}
			}

			out <- g
		}
	}()

	return out
}
