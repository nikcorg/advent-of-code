package s14

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/nikcorg/aoc2020/utils"
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
	lineInput := make(linestream.LineChan, bufSize)
	cmdStream := make(chan *Operation, bufSize)
	solveStream := getSolver(part)

	convStream(linestream.SkipEmpty(lineInput), cmdStream)
	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	solution := solveStream(cmdStream)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type Cmd int

const (
	mask Cmd = iota + 1
	write
)

func (c Cmd) String() string {
	switch c {
	case mask:
		return "mask"
	case write:
		return "write"
	default:
		panic(fmt.Errorf("invalid op: %d", c))
	}
}

type Operation struct {
	action  Cmd
	data    string
	address int
}

func convStream(inp linestream.ReadOnlyLineChan, out chan<- *Operation) {
	go func() {
		defer close(out)

		splitter := regexp.MustCompile(`^(mask|mem)(?:\[(\d+)\])? = (.*)$`)

		for line := range inp {
			matches := splitter.FindStringSubmatch(line.Content())

			if matches == nil {
				panic(fmt.Errorf("invalid input %s", line.Content()))
			}

			var nextOp *Operation
			switch matches[1] {
			case "mask":
				nextOp = &Operation{mask, matches[3], 0}
			case "mem":
				nextOp = &Operation{
					write,
					fmt.Sprintf("%036s", strconv.FormatInt(int64(utils.MustAtoi(matches[3])), 2)),
					utils.MustAtoi(matches[2]),
				}
			default:
				panic(fmt.Errorf("unknown op: %s", matches[1]))
			}

			out <- nextOp
		}
	}()
}

type solver func(<-chan *Operation) int64

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(inp <-chan *Operation) int64 {
	var applyMask func(string) string
	mem := map[int]string{}

	prepareMask := func(mask string) func(string) string {
		transforms := map[int]string{}

		for p, c := range mask {
			switch c {
			case '0':
				fallthrough
			case '1':
				transforms[p] = string(c)
			}
		}

		return func(x string) string {
			xs := strings.Split(x, "")
			for p, b := range transforms {
				xs[p] = b
			}
			return strings.Join(xs, "")
		}
	}

	for op := range inp {
		switch op.action {
		case mask:
			applyMask = prepareMask(op.data)

		case write:
			mem[op.address] = applyMask(op.data)

		default:
			panic(errors.New("unknown action"))
		}
	}

	var total int64 = 0
	for _, v := range mem {
		total += mustParseInt(v)
	}

	return total
}

func solveSecond(inp <-chan *Operation) int64 {
	var applyMask func(string) []string
	mem := map[int]string{}

	prepareMask := func(mask string) func(string) []string {
		transforms := map[int]string{}

		for p, c := range mask {
			switch c {
			case 'X':
				fallthrough
			case '1':
				transforms[p] = string(c)
			}
		}

		var mapFloatingBits func(string, []string) []string
		mapFloatingBits = func(prefix string, xs []string) []string {
			for p, c := range xs {
				if c == "X" {
					head := xs[0:p]
					tail := xs[p+1:]
					nextPrefix := prefix + strings.Join(head, "")
					zero := mapFloatingBits(nextPrefix+"0", tail)
					one := mapFloatingBits(nextPrefix+"1", tail)
					return append(append([]string{}, zero...), one...)
				}
			}

			return []string{prefix + strings.Join(xs, "")}
		}

		return func(x string) []string {
			xs := strings.Split(x, "")
			for p, b := range transforms {
				xs[p] = b
			}

			return mapFloatingBits("", xs)
		}
	}

	for op := range inp {
		switch op.action {
		case mask:
			applyMask = prepareMask(op.data)

		case write:
			binAddress := fmt.Sprintf("%036s", strconv.FormatInt(int64(op.address), 2))
			addresses := applyMask(binAddress)
			for _, addr := range addresses {
				mem[int(mustParseInt(addr))] = op.data
			}

		default:
			panic(errors.New("unknown action"))
		}
	}

	var total int64 = 0
	for _, v := range mem {
		total += mustParseInt(v)
	}

	return total
}

func mustParseInt(s string) int64 {
	v, e := strconv.ParseInt(s, 2, 64)
	if e != nil {
		panic(e)
	}
	return v
}
