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
	parsed, err := parseInput(input)
	if err != nil {
		return err
	}

	if result, err := solveFirst(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 1: %d\n", result))
	}

	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d\n", result))
	}

	return nil
}

type fish int
type puzzleInput []fish

func parseInput(raw string) (puzzleInput, error) {
	nums := []fish{}
	for _, x := range strings.Split(raw, ",") {
		num, _ := strconv.Atoi(x)
		nums = append(nums, fish(num))
	}
	return nums, nil
}

func solveFirst(input puzzleInput) (int, error) {
	daysRemaining := 80
	sol := simulate(input, daysRemaining)
	return sol, nil
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func solveSecond(input puzzleInput) (int, error) {
	var (
		sol int = 0
		wg  sync.WaitGroup
	)

	daysRemaining := 256
	chunkSize := len(input) / min(16, len(input))
	chunkStart := 0

	for chunkStart+chunkSize < len(input) {
		chunk := input[chunkStart : chunkStart+chunkSize]
		wg.Add(1)
		go func(chunk []fish, daysRemaining int) {
			sol += simulate(chunk, daysRemaining)
			wg.Done()
		}(chunk, daysRemaining)
		chunkStart += chunkSize
	}

	sol += simulate(input[chunkStart:], daysRemaining)

	wg.Wait()

	return sol, nil
}

func simulate(inp []fish, days int) int {
	schoal := map[fish]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	for _, f := range inp {
		schoal[f]++
	}

	for days > 0 {
		spawn := schoal[0]

		for g := 1; g <= 8; g++ {
			schoal[fish(g-1)] = schoal[fish(g)]
		}

		schoal[6] += spawn
		schoal[8] = spawn

		days--
	}

	tot := 0
	for _, n := range schoal {
		tot += n
	}

	return tot
}
