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

type puzzleInput []int

func parseInput(raw string) (puzzleInput, error) {
	positions := []int{}

	for _, s := range strings.Split(raw, ",") {
		n, _ := strconv.Atoi(s)
		positions = append(positions, n)
	}

	return positions, nil
}

func solveFirst(input puzzleInput) (int, error) {
	maxPos := 0

	for _, pos := range input {
		maxPos = max(maxPos, pos)
	}

	var (
		mut     sync.Mutex
		wg      sync.WaitGroup
		minCost int = -1
	)

	for pos := 0; pos <= maxPos; pos++ {
		wg.Add(1)

		go func(targetPos int) {
			cost := 0

			for _, pos := range input {
				cost += abs(targetPos - pos)
			}

			if minCost == -1 || minCost > cost {
				mut.Lock()
				defer mut.Unlock()

				if minCost == -1 || minCost > cost {
					minCost = cost
				}
			}
			wg.Done()
		}(pos)
	}

	wg.Wait()

	return minCost, nil
}

func solveSecond(input puzzleInput) (int, error) {
	maxPos := 0

	for _, pos := range input {
		maxPos = max(maxPos, pos)
	}

	var (
		mut     sync.Mutex
		wg      sync.WaitGroup
		minCost int = -1
	)

	for pos := 0; pos <= maxPos; pos++ {
		wg.Add(1)

		go func(targetPos int) {
			cost := 0

			for _, pos := range input {
				cost += moveCost(abs(targetPos - pos))
			}

			if minCost == -1 || minCost > cost {
				mut.Lock()
				defer mut.Unlock()

				if minCost == -1 || minCost > cost {
					minCost = cost
				}
			}
			wg.Done()
		}(pos)
	}

	wg.Wait()

	return minCost, nil
}

var (
	costCache = map[int]int{}
	cacheMut  sync.Mutex
)

func moveCost(dist int) int {
	cacheMut.Lock()
	defer cacheMut.Unlock()

	if cost, ok := costCache[dist]; ok {
		return cost
	}

	cost := 0
	for i := dist; i > 0; i-- {
		cost += i
	}

	costCache[dist] = cost

	return cost
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
