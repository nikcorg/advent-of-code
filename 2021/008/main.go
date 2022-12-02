package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
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

type puzzleInput interface{}

func parseInput(raw string) (puzzleInput, error) {
	return nil, errors.New("not implemented")
}

func solveFirst(input puzzleInput) (int, error) {
	return 0, errors.New("not implemented")
}

func solveSecond(input puzzleInput) (int, error) {
	return 0, errors.New("not implemented")
}
