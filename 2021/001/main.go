package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr() error {
	if err := solveFirst(); err != nil {
		return err
	}

	if err := solveSecond(); err != nil {
		return err
	}

	return nil
}

func solveFirst() error {
	nums, err := parseInput(input)
	if err != nil {
		return err
	}

	incs := countIncrements(nums)

	io.WriteString(os.Stdout, fmt.Sprintf("increments: %d\n", incs))

	return nil
}

func countIncrements(readings []int) int {
	incs := 0
	prev := readings[0]

	for _, next := range readings[1:] {
		if next > prev {
			incs++
		}
		prev = next
	}

	return incs
}

func solveSecond() error {
	nums, err := parseInput(input)
	if err != nil {
		return err
	}

	incs := countTripletIncrements(nums)

	io.WriteString(os.Stdout, fmt.Sprintf("increments: %d\n", incs))

	return nil
}

func countTripletIncrements(readings []int) int {
	incs := 0

	prev := readings[0] + readings[1] + readings[2]
	for n := 0; n < len(readings)-2; n++ {
		next := readings[n] + readings[n+1] + readings[n+2]

		if next > prev {
			incs++
		}

		prev = next
	}

	return incs
}

func parseInput(raw string) ([]int, error) {
	rows := strings.Split(raw, "\n")
	parsed := []int{}
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		num, err := strconv.Atoi(row)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, num)
	}

	return parsed, nil
}
