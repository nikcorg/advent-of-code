package s13

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	config := convStream(linestream.SkipEmpty(lineInput))

	solveStream := getSolver(part)
	solution := solveStream(config)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type configuration struct {
	earliestTime int
	buses        []int
}

func convStream(inp linestream.ReadOnlyLineChan) *configuration {
	earliestTime := utils.MustAtoi((<-inp).Content())
	buses := []int{}

	for _, l := range strings.Split((<-inp).Content(), ",") {
		if l == "x" {
			buses = append(buses, 0)
			continue
		}

		buses = append(buses, utils.MustAtoi(l))
	}

	return &configuration{earliestTime, buses}
}

type solver func(*configuration) uint64

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("invalid solver %d", part))
}

func solveFirst(inp *configuration) uint64 {
	busID := inp.buses[0]
	departure := inp.earliestTime - inp.earliestTime%busID + busID

	for _, bus := range inp.buses[1:] {
		if bus == 0 {
			continue
		}

		busDeparture := inp.earliestTime - inp.earliestTime%bus + bus
		if busDeparture < departure {
			busID = bus
			departure = busDeparture
		}
	}

	return uint64((departure - inp.earliestTime) * busID)
}

func solveSecond(inp *configuration) uint64 {
	var (
		nextIncrement    uint64   = uint64(inp.buses[0])
		refDepartureTime uint64   = 0
		hits             int      = 0
		nonZeroBuses     []uint64 = []uint64{uint64(inp.buses[0])}
	)

	for {
		if ok, nextHits := matches(refDepartureTime, inp.buses[1:]); ok {
			break
		} else if nextHits > hits {
			hits = nextHits
			nonZeroBuses = append(nonZeroBuses, uint64(inp.buses[nextHits]))
			nextIncrement = utils.LCM(nonZeroBuses[0], nonZeroBuses[1], nonZeroBuses[2:]...)
		}

		refDepartureTime += nextIncrement
	}

	return refDepartureTime
}

type InterimResult struct {
	hits          int
	nextIncrement uint64
	departureTime uint64
}

func matches(refDeparture uint64, buses []int) (bool, int) {
	var lastNonZeroOffset int = 0

	for offs, bus := range buses {
		if bus == 0 {
			continue
		}

		var (
			busID    uint64 = uint64(bus)
			timeOffs uint64 = uint64(offs + 1)
		)

		if (refDeparture+timeOffs)%busID != 0 {
			return false, lastNonZeroOffset
		}

		lastNonZeroOffset = offs + 1
	}

	return true, 0
}
