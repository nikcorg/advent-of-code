package s16

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
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
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	lineInput := make(linestream.LineChan, bufSize)

	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	solve := getSolver(part)
	solution := solve(lineInput)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(linestream.ReadOnlyLineChan) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

const headerOwnTix = "your ticket:"
const headerNearbyTix = "nearby tickets:"

func solveFirst(inp linestream.ReadOnlyLineChan) int {
	var (
		tv = getTicketValidator(inp)
	)

	for line := range inp {
		// Empty line separates own ticket from the nearby tickets block
		if line.Content() == "" {
			break
		}
	}

	solution := 0

	// All Nearby tickets until the end now
	for line := range linestream.SkipEmpty(inp) {
		if line.Content() == headerNearbyTix {
			continue
		}

		tix := newTicketFromCSVString(line.Content())
		for _, v := range tv.InvalidValues(&tix) {
			solution += v
		}
	}

	return solution
}

func solveSecond(inp linestream.ReadOnlyLineChan) int {
	var (
		tv        = getTicketValidator(inp)
		ownTicket Ticket
	)

	for line := range inp {
		// Empty line separates own ticket from the nearby tickets block
		if line.Content() == "" {
			break
		} else if line.Content() == headerOwnTix {
			continue
		}

		ownTicket = newTicketFromCSVString(line.Content())
	}

	var (
		numFields          = len(ownTicket)
		unidentifiedFields = numFields
		identifiedFields   = make([]uint, unidentifiedFields, unidentifiedFields)
		fieldValues        = make([][]int, numFields, numFields)
	)

	// Remaining lines are nearby tickets, we can read until the end
	for line := range linestream.SkipEmpty(inp) {
		if line.Content() == headerNearbyTix {
			continue
		}

		tix := newTicketFromCSVString(line.Content())
		if !tv.Validates(&tix) {
			continue
		}

		for f := 0; f < numFields; f++ {
			if fieldValues[f] == nil {
				fieldValues[f] = []int{tix[f]}
			} else {
				fieldValues[f] = append(fieldValues[f], tix[f])
			}
		}
	}

	var (
		outcomes = make([]uint, numFields)
	)

	for fid := 0; fid < numFields; fid++ {
		var outcome uint = 0

		for n, val := range fieldValues[fid] {
			matches := tv.MatchFields(val)
			if n == 0 {
				outcome = matches
			} else {
				// Use bitwise AND to keep only common matches
				outcome = outcome & matches
			}
		}

		outcomes[fid] = outcome
	}

	var (
		// the mask to clear the bits for the fields which have been matched
		sieveMask uint = 0
	)

	for unidentifiedFields > 0 {
		for fid, outcome := range outcomes {
			// solved fields are zeroed out
			if outcome == 0 {
				continue
			}

			// apply the sieveMask to the column's field mask
			outcome = outcome &^ sieveMask

			// find the first field with only 1 high bit
			if bits.OnesCount32(uint32(outcome)) == 1 {
				// update the discovered fields mask
				sieveMask = sieveMask | outcome

				// zero the outcome for this field so it can be ignored, we can't remove it,
				// as we need the mask indexes to match columns indexes
				outcomes[fid] = 0

				// store the field's bitmask in the column slot, so it can be mapped later.
				// a smarter programmer would match columns and fields here and quit early
				// when the necessary fields have been found, but i'm not one of them...
				identifiedFields[fid] = outcome

				// take on down, pass it around...
				unidentifiedFields--
			}
		}
	}

	if unidentifiedFields > 0 {
		panic(fmt.Errorf("unidentified fields remaining: %d", unidentifiedFields))
	}

	var (
		solution = 1
	)

	for fid, fieldMask := range identifiedFields {
		var (
			f   *FieldConfiguration
			val int
		)

		// match the column to a field
		for b := 0; b <= numFields; b++ {
			bitMask := uint(math.Pow(2, float64(b)))
			if bitMask&fieldMask != 0 {
				f = tv.fields[b]
				val = ownTicket[fid]
				break
			}
		}

		// fmt.Printf("fid=%2d, mask=%020s, idx=%2d, name=%20s, value=%d\n", fid, strconv.FormatInt(int64(fieldMask), 2), idx, f.name, val)

		if strings.HasPrefix(f.name, "departure") {
			solution *= val
		}
	}

	return solution
}

func getTicketValidator(inp linestream.ReadOnlyLineChan) *TicketValidator {
	tv := &TicketValidator{}

	for line := range inp {
		// Empty line finishes the configuration block
		if line.Content() == "" {
			break
		}

		tv.AddField(newFieldConfigurationFromString(line.Content()))
	}

	return tv
}

func newTicketFromCSVString(csv string) Ticket {
	vs := strings.Split(csv, ",")
	t := make(Ticket, len(vs), len(vs))
	for n, v := range vs {
		t[n] = utils.MustAtoi(v)
	}
	return t
}
