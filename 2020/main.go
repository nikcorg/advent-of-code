package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/nikcorg/aoc2020/s1"
	"github.com/nikcorg/aoc2020/s2"
	"github.com/nikcorg/aoc2020/s3"
	"github.com/nikcorg/aoc2020/s4"
	"github.com/nikcorg/aoc2020/s5"
	"github.com/nikcorg/aoc2020/s6"
	"github.com/nikcorg/aoc2020/s7"
)

const solved = 3
const usage = `usage: %s <puzzle> <input>
- puzzle is a number between 1-%d
- input is the path to the file`

func showUsage(exitCode int) {
	io.WriteString(os.Stderr, fmt.Sprintf(usage, path.Base(os.Args[0]), solved))
	os.Exit(exitCode)
}

func main() {
	if len(os.Args) < 2 {
		showUsage(exUsage)
	}

	inputFile := os.Args[2]
	puzzle, part, err := parsePuzzleArg(os.Args[1])

	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		showUsage(exUsage)
	}

	sol := getSolver(context.Background(), puzzle, inputFile)

	if err := sol.Solve(part); err != nil {
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Solver makes problems go away
type Solver interface {
	Solve(part int) error
}

func getSolver(ctx context.Context, puzzle int, inputFilename string) Solver {
	switch puzzle {
	case 1:
		return s1.New(ctx, inputFilename)
	case 2:
		return s2.New(ctx, inputFilename)
	case 3:
		return s3.New(ctx, inputFilename)
	case 4:
		return s4.New(ctx, inputFilename)
	case 5:
		return s5.New(ctx, inputFilename)
	case 6:
		return s6.New(ctx, inputFilename)
	case 7:
		return s7.New(ctx, inputFilename)
	default:
		io.WriteString(os.Stderr, fmt.Sprintf("unknown puzzle: %d", puzzle))
		showUsage(exUsage)
	}

	return nil
}

func parsePuzzleArg(candidate string) (int, int, error) {
	parts := strings.Split(candidate, ".")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("expected puzzle to be n.1 or n.2, got: %s", candidate)
	}

	var (
		puzzle, part int
		err          error
	)

	puzzle, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid puzzle: %s", parts[0])
	}

	part, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid part: %s", parts[1])
	}

	return puzzle, part, nil
}
