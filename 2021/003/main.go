package main

import (
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(input string) error {
	bits, readings, err := parseInput(input)
	if err != nil {
		return nil
	}

	if solution, err := solveFirst(bits, readings); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))
	}

	if solution, err := solveSecond(bits, readings); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))
	}

	return nil
}

func parseInput(raw string) (uint, []int, error) {
	var (
		bits uint
		rows = []int{}
	)

	for _, row := range strings.Split(raw, "\n") {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		if bits == 0 {
			bits = uint(len(row))
		}

		intval, err := strconv.ParseInt(row, 2, 64)
		if err != nil {
			return 0, nil, err
		}

		rows = append(rows, int(intval))
	}

	return bits, rows, nil
}

func onesCount(bits uint, readings []int) []int {
	var onesCounts []int = make([]int, bits)

	for _, reading := range readings {
		var bit uint
		for bit = 0; bit < bits; bit++ {
			if reading&(1<<bit) > 0 {
				onesCounts[bits-bit-1]++
			}
		}
	}

	return onesCounts
}

func solveFirst(bits uint, readings []int) (int, error) {
	var (
		onesCounts             []int = onesCount(bits, readings)
		threshold                    = len(readings) / 2
		gammaRate, epsilonRate int   = 0, 0
	)

	for bit := range onesCounts {
		ones := onesCounts[bit]
		// because bit zero from the view of the problem is actually the
		// left-most bit, we need to reverse our bitmask
		bitmask := (1 << (bits - uint(bit) - 1))
		if ones > threshold {
			gammaRate |= bitmask
		} else {
			epsilonRate |= bitmask
		}
	}

	return gammaRate * epsilonRate, nil
}

func solveSecond(bits uint, readings []int) (int, error) {
	o2r, co2r := 0, 0
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		o2r = o2Rating(0, bits, readings)
		wg.Done()
	}()

	go func() {
		co2r = co2ScrubberRating(0, bits, readings)
		wg.Done()
	}()

	wg.Wait()

	return o2r * co2r, nil
}

func keep(vals []int, bitmask int) []int {
	next := []int{}
	for _, v := range vals {
		if v&bitmask == 0 {
			continue
		}
		next = append(next, v)
	}
	return next
}

func discard(vals []int, bitmask int) []int {
	next := []int{}
	for _, v := range vals {
		if v&bitmask > 0 {
			continue
		}
		next = append(next, v)
	}
	return next
}

func o2Rating(bit, bits uint, readings []int) int {
	if len(readings) == 1 {
		return readings[0]
	}

	threshold := int(math.Ceil(float64(len(readings)) / 2))
	onesCounts := onesCount(bits, readings)

	var nextReadings []int
	if onesCounts[bit] >= threshold {
		// because bit zero from the view of the problem is actually the
		// left-most bit, we need to reverse our bitmask
		nextReadings = keep(readings, 1<<(bits-bit-1))
	} else {
		// because bit zero from the view of the problem is actually the
		// left-most bit, we need to reverse our bitmask
		nextReadings = discard(readings, 1<<(bits-bit-1))
	}

	return o2Rating(bit+1, bits, nextReadings)
}

func co2ScrubberRating(bit, bits uint, readings []int) int {
	if len(readings) == 1 {
		return readings[0]
	}

	threshold := int(math.Ceil(float64(len(readings)) / 2))
	onesCounts := onesCount(bits, readings)

	var nextReadings []int
	if onesCounts[bit] >= threshold {
		// because bit zero from the view of the problem is actually the
		// left-most bit, we need to reverse our bitmask
		nextReadings = discard(readings, 1<<(bits-bit-1))
	} else {
		// because bit zero from the view of the problem is actually the
		// left-most bit, we need to reverse our bitmask
		nextReadings = keep(readings, 1<<(bits-bit-1))
	}

	return co2ScrubberRating(bit+1, bits, nextReadings)
}
