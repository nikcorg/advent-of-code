package s8

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"

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

	lineInput := make(linestream.LineChan, bufSize)
	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	instructions := make(chan instruction, bufSize)
	convStream(linestream.SkipEmpty(lineInput), instructions)

	solution := <-solveStream(getSolver(part), instructions)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type instruction struct {
	op      string
	arg     int
	visited bool
}

type solver func([]instruction) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("unsolved part %d", part))
}

func convStream(inp linestream.ReadOnlyLineChan, out chan<- instruction) {
	lineSplitter := regexp.MustCompile(`^(acc|jmp|nop) ([-+]\d+)$`)

	go func() {
		defer close(out)
		for line := range inp {
			parts := lineSplitter.FindStringSubmatch(line.Content())

			if parts == nil {
				io.WriteString(os.Stderr, fmt.Sprintf("invalid instruction: %s\n", line.Content()))
				continue
			}

			op := parts[1]
			rawArg := parts[2]

			arg, err := strconv.Atoi(rawArg)
			if err != nil {
				io.WriteString(os.Stderr, fmt.Sprintf("invalid argument: %s\n", rawArg))
				continue
			}

			out <- instruction{op, arg, false}
		}
	}()
}

func solveStream(solve solver, inp <-chan instruction) <-chan int {
	out := make(chan int)

	go func() {
		program := []instruction{}
		for ins := range inp {
			program = append(program, ins)
		}

		out <- solve(program)
	}()

	return out
}

func runProgram(program []instruction, errorOnRevisit bool) (int, error) {
	acc := 0
	i := 0

	for {
		if i >= len(program) {
			break
		} else if program[i].visited {
			if errorOnRevisit {
				return 0, errors.New("infinite loop")
			}
			break
		}

		program[i].visited = true

		switch program[i].op {
		case "jmp":
			i += program[i].arg
		case "acc":
			acc += program[i].arg
			fallthrough
		case "nop":
			fallthrough
		default:
			i++
		}
	}

	return acc, nil
}

func solveFirst(program []instruction) int {
	acc, err := runProgram(program, false)

	if err != nil {
		panic(err)
	}

	return acc
}

func solveSecond(program []instruction) int {
	res := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	programs := make(chan []instruction, 10)

	attempt := func(prog []instruction) {
		acc, err := runProgram(prog, true)

		if err != nil {
			return
		}

		res <- acc
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case prog := <-programs:
				go attempt(prog)
			}
		}
	}()

	for i, line := range program {
		select {
		case <-ctx.Done():
			break
		default:
			switch line.op {
			case "jmp":
				alt := make([]instruction, len(program))
				copy(alt, program)
				alt[i] = instruction{"nop", line.arg, false}

				programs <- alt

			case "nop":
				alt := make([]instruction, len(program))
				copy(alt, program)
				alt[i] = instruction{"jmp", line.arg, false}

				programs <- alt
			}
		}
	}

	select {
	case <-ctx.Done():
		panic(errors.New("failed to solve before deadline"))
	case sol := <-res:
		return sol
	}
}
