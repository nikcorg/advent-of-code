package s12

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"regexp"

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

	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	cmdStream := make(chan *Instruction, bufSize)
	convStream(linestream.SkipEmpty(lineInput), cmdStream)

	solveStream := getSolver(part)
	solution := solveStream(cmdStream)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(<-chan *Instruction) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func convStream(inp linestream.ReadOnlyLineChan, out chan<- *Instruction) {
	splitInstruction := regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

	go func() {
		defer close(out)
		for line := range inp {
			matches := splitInstruction.FindStringSubmatch(line.Content())

			if matches == nil {
				fmt.Println("bad instruction", line)
				continue
			}

			units := utils.MustAtoi(matches[2])

			switch matches[1] {
			case "N":
				out <- &Instruction{moveNorth, units}
			case "S":
				out <- &Instruction{moveSouth, units}
			case "E":
				out <- &Instruction{moveEast, units}
			case "W":
				out <- &Instruction{moveWest, units}
			case "L":
				out <- &Instruction{rotLeft, units}
			case "R":
				out <- &Instruction{rotRight, units}
			case "F":
				out <- &Instruction{moveForward, units}
			}
		}
	}()
}

func solveFirst(inp <-chan *Instruction) int {
	ship := &Vessel{
		facing: east,
		Point: &Point{
			X: 0,
			Y: 0,
		},
	}

	for ins := range inp {
		switch ins.cmd {
		case moveForward:
			ship.Forward(ins.units)
		case rotLeft:
			fallthrough
		case rotRight:
			ship.Rotate(ins.cmd, ins.units)
		case moveNorth:
			ship.Move(north, ins.units)
		case moveSouth:
			ship.Move(south, ins.units)
		case moveEast:
			ship.Move(east, ins.units)
		case moveWest:
			ship.Move(west, ins.units)
		}
	}

	return utils.Abs(ship.Y) + utils.Abs(ship.X)
}

func solveSecond(inp <-chan *Instruction) int {
	ship := &Vessel{
		facing: east,
		Point:  &Point{0, 0},
	}

	wp := &Point{-1, 10}

	for ins := range inp {
		switch ins.cmd {
		case moveForward:
			dX, dY := wp.Diff(ship.Point)
			ship.Point.Translate(dX*ins.units, dY*ins.units)
			wp.Translate(dX*ins.units, dY*ins.units)

		case rotRight:
			fallthrough
		case rotLeft:
			wp.Rotate(ship.Point, ins.cmd, ins.units)

		default:
			wp.Move(ins.cmd, ins.units)
		}
	}

	return utils.Abs(ship.Y) + utils.Abs(ship.X)
}
