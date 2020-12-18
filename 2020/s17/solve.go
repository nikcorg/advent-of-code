package s17

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

const bufSize = 1

type Solver struct {
	ctx              context.Context
	out              io.Writer
	simulationCycles int
}

func New(ctx context.Context, out io.Writer, simulationCycles int) *Solver {
	return &Solver{ctx, out, simulationCycles}
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

	init, side := slurpInput(linestream.SkipEmpty(lineInput))
	solve := getSolver(part)
	solution := solve(s.ctx, init, side, s.simulationCycles)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(context.Context, []string, int, int) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(ctx context.Context, init []string, side, maxCycles int) int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	world := newWorld(ctx, surroundingXYZCoords)
	z := 0

	// set initial state
	for n, pos := range init {
		x := n % side
		y := n % (side * side) / side

		if pos == activeConwayCube {
			world.AlterStateAt(Position{x, y, z}, activeConwayCube)
		}
	}

	world.EndTurn()

	for n := 0; n < maxCycles; n++ {
		world.NextTurn()
	}

	return world.ActiveCubes()
}

func solveSecond(ctx context.Context, init []string, side, maxCycles int) int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	world := newWorld(ctx, surroundingXYZWCoords)
	z := 0
	w := 0

	// set initial state
	for n, pos := range init {
		x := n % side
		y := n % (side * side) / side

		if pos == activeConwayCube {
			world.AlterStateAt(Position{x, y, z, w}, activeConwayCube)
		}
	}

	world.EndTurn()

	for n := 0; n < maxCycles; n++ {
		world.NextTurn()
	}

	return world.ActiveCubes()
}

func surroundingXYZCoords(pos Position) []Position {
	sx := pos[0]
	sy := pos[1]
	sz := pos[2]

	n := 27
	cs := make([]Position, 0, n-1) // 26 neighbours
	offs := -1

	for i := 0; i < 27; i++ {
		x := sx + ((i % 3) + offs)
		y := sy + ((i % 9 / 3) + offs)
		z := sz + ((i / 9) + offs)

		if x == sx && y == sy && z == sz {
			continue
		}

		cs = append(cs, Position{x, y, z})
	}

	return cs
}

func surroundingXYZWCoords(pos Position) []Position {
	sw := pos[3]

	xyz := surroundingXYZCoords(pos)
	xyzw := []Position{}

	for w := 0; w < 3; w++ {
		for _, p := range xyz {
			p = append(p, sw+(w-1))
			xyzw = append(xyzw, p)
		}
	}

	sx := pos[0]
	sy := pos[1]
	sz := pos[2]
	offs := -1

	// Because XYZ omits the origin position, we need to explicitly
	// inlude the two neighbouring coordinates on the W dimension
	for i := 0; i < 3; i++ {
		w := sw + i + offs
		if w == sw {
			continue
		}

		xyzw = append(xyzw, Position{sx, sy, sz, w})
	}

	return xyzw
}

func slurpInput(inp linestream.ReadOnlyLineChan) ([]string, int) {
	var (
		init []string
		side int = -1
	)

	for line := range inp {
		if side < 0 {
			side = len(line.Content())
		}

		init = append(init, strings.Split(line.Content(), "")...)
	}

	return init, side
}
