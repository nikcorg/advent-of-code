package s5

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
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

	lineInput := make(linestream.LineChan, 10)
	linestream.New(ctx, bufio.NewReader(inputFile), lineInput)

	filteredInput := make(linestream.LineChan, 10)
	linestream.SkipEmpty(lineInput, filteredInput)

	solution := <-getSolver(part)(filteredInput)

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(inp linestream.LineChan) <-chan int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst

	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(inp linestream.LineChan) <-chan int {
	out := make(chan int, 0)

	maxID := 0
	candidateIDs := make(chan int, 10)

	go func() {
		defer close(out)
		for candidateID := range candidateIDs {
			maxID = int(math.Max(float64(maxID), float64(candidateID)))
		}
		out <- maxID
	}()

	go func() {
		defer close(candidateIDs)
		wg := sync.WaitGroup{}

		for r := range inp {
			cipher := r.Content()
			wg.Add(1)
			go func() {
				defer wg.Done()
				candidateIDs <- idForCipher(cipher)
			}()
		}

		wg.Wait()
	}()

	return out
}

const iterThreshold = 10

func solveSecond(inp linestream.LineChan) <-chan int {
	out := make(chan int)

	candidateIDs := make(chan int, 10) // buffered, so that it won't block the cipher calc
	allIDs := []int{}

	// accumulate a sorted list of IDs and finally scan for gap in sequence
	go func() {
		defer close(out)
		for candidateID := range candidateIDs {
			allIDs = binsert(allIDs, candidateID)
		}

		for i, id := range allIDs[1:] {
			prev := allIDs[i] // i starts from zero, so no need to subtract 1
			if prev+2 == id {
				out <- id - 1
				return
			}
		}
	}()

	// calculcate IDs from ciphers
	go func() {
		defer close(candidateIDs)
		wg := sync.WaitGroup{}

		for r := range inp {
			cipher := r.Content()
			wg.Add(1)
			go func() {
				defer wg.Done()
				candidateIDs <- idForCipher(cipher)
			}()
		}

		wg.Wait()
	}()

	return out
}

func idForCipher(cipher string) int {
	row := cipher[0:7]
	seat := cipher[7:]

	rowNum := 0
	seatNum := 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		rowNum = binsearch(cmpRow, 127, row)
	}()
	go func() {
		defer wg.Done()
		seatNum = binsearch(cmpSeat, 7, seat)
	}()

	wg.Wait()

	id := rowNum*8 + seatNum

	return id
}

type ord = string

const (
	hi ord = "higher"
	lo     = "lo"
)

func cmpRow(x rune) ord {
	switch x {
	case 'B':
		return hi
	case 'F':
		return lo
	}
	panic(fmt.Errorf("invalid input %v", x))
}

func cmpSeat(x rune) ord {
	switch x {
	case 'R':
		return hi
	case 'L':
		return lo
	}
	panic(fmt.Errorf("invalid input %v", x))
}

func binsearch(cmp func(rune) ord, max int, input string) int {
	cmin := 0
	cmax := max
	last := input[len(input)-1]

	for _, i := range input {
		switch cmp(i) {
		case hi:
			cmin = cmax - (cmax-cmin)/2
		case lo:
			cmax = cmin + (cmax-cmin)/2
		default:
			panic(fmt.Errorf("broken compare function"))
		}
	}

	switch cmp(rune(last)) {
	case hi:
		return cmax
	case lo:
		return cmin
	default:
		panic(fmt.Errorf("broken compare function"))
	}
}
