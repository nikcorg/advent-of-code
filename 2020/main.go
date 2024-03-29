package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nikcorg/aoc2020/s1"
	"github.com/nikcorg/aoc2020/s10"
	"github.com/nikcorg/aoc2020/s11"
	"github.com/nikcorg/aoc2020/s12"
	"github.com/nikcorg/aoc2020/s13"
	"github.com/nikcorg/aoc2020/s14"
	"github.com/nikcorg/aoc2020/s15"
	"github.com/nikcorg/aoc2020/s16"
	"github.com/nikcorg/aoc2020/s17"
	"github.com/nikcorg/aoc2020/s18"
	"github.com/nikcorg/aoc2020/s19"
	"github.com/nikcorg/aoc2020/s2"
	"github.com/nikcorg/aoc2020/s3"
	"github.com/nikcorg/aoc2020/s4"
	"github.com/nikcorg/aoc2020/s5"
	"github.com/nikcorg/aoc2020/s6"
	"github.com/nikcorg/aoc2020/s7"
	"github.com/nikcorg/aoc2020/s8"
	"github.com/nikcorg/aoc2020/s9"
	"github.com/nikcorg/aoc2020/utils"
)

const solved = 19
const inputDir = "_inputs"

type SolverFunc func(io.Reader) error

func runPuzzle(ctx context.Context, sol SolverFunc, fileName string) error {
	var err error

	inputFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() { inputFile.Close() }()

	return sol(inputFile)
}

func parseArgs() (int, int) {
	var startAt, stopAt int = 0, solved

	if len(os.Args) > 1 {

		firstArg := os.Args[1]

		if firstArg[0] == '-' {
			startAt = 0
			stopAt = utils.MustAtoi(firstArg[1:])
		} else if firstArg[len(firstArg)-1] == '-' {
			startAt = utils.MustAtoi(firstArg[0 : len(firstArg)-1])
			stopAt = solved
		} else {
			stopAt = utils.MustAtoi(firstArg)
			startAt = stopAt - 1
		}
	}

	return startAt, stopAt
}

func main() {
	ctx := context.Background()
	start := time.Now()

	startAt, stopAt := parseArgs()

	for puzzle := startAt; puzzle < stopAt; puzzle++ {
		solver := getSolver(ctx, os.Stdout, puzzle+1)

		inputFile := fmt.Sprintf("%s/%d.txt", inputDir, puzzle+1)

		io.WriteString(os.Stdout, fmt.Sprintf("%d.1: ", puzzle+1))
		startFirst := time.Now()
		if err := runPuzzle(ctx, solver.SolveFirst, inputFile); err != nil {
			io.WriteString(os.Stderr, err.Error())
		} else {
			durationFirst := time.Since(startFirst)
			io.WriteString(os.Stdout, fmt.Sprintf("duration=%v\n", durationFirst))
		}

		io.WriteString(os.Stdout, fmt.Sprintf("%d.2: ", puzzle+1))
		startSecond := time.Now()
		if err := runPuzzle(ctx, solver.SolveSecond, inputFile); err != nil {
			io.WriteString(os.Stderr, err.Error())
		} else {

			durationSecond := time.Since(startSecond)
			io.WriteString(os.Stdout, fmt.Sprintf("duration=%v\n", durationSecond))
		}
	}

	io.WriteString(os.Stdout, fmt.Sprintf("total duration %v\n", time.Since(start)))
}

// Solver makes problems go away
type Solver interface {
	SolveFirst(inp io.Reader) error
	SolveSecond(inp io.Reader) error
	Solve(part int, inp io.Reader) error
}

func getSolver(ctx context.Context, out io.Writer, puzzle int) Solver {
	switch puzzle {
	case 1:
		return s1.New(ctx, out)
	case 2:
		return s2.New(ctx, out)
	case 3:
		return s3.New(ctx, out)
	case 4:
		return s4.New(ctx, out)
	case 5:
		return s5.New(ctx, out)
	case 6:
		return s6.New(ctx, out)
	case 7:
		return s7.New(ctx, out)
	case 8:
		return s8.New(ctx, out)
	case 9:
		return s9.New(ctx, out, 25)
	case 10:
		return s10.New(ctx, out)
	case 11:
		return s11.New(ctx, out)
	case 12:
		return s12.New(ctx, out)
	case 13:
		return s13.New(ctx, out)
	case 14:
		return s14.New(ctx, out)
	case 15:
		return s15.New(ctx, out)
	case 16:
		return s16.New(ctx, out)
	case 17:
		return s17.New(ctx, out, 6)
	case 18:
		return s18.New(ctx, out)
	case 19:
		return s19.New(ctx, out)
	default:
		io.WriteString(os.Stderr, fmt.Sprintf("unknown puzzle: %d\n", puzzle))
		os.Exit(1)
	}

	return nil
}
