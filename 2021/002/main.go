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
	if solution, err := solveFirst(input); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))
	}

	if solution, err := solveSecond(input); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))
	}

	return nil
}

func solveFirst(input string) (int, error) {
	cmds, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	horiz, depth := 0, 0

	for _, cmd := range cmds {
		switch cmd.dir {
		case forward:
			horiz += cmd.units
		case up:
			depth -= cmd.units
		case down:
			depth += cmd.units
		}
	}

	return horiz * depth, nil
}

func solveSecond(input string) (int, error) {
	cmds, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	horiz, depth, aim := 0, 0, 0

	for _, cmd := range cmds {
		switch cmd.dir {
		case down:
			aim += cmd.units
		case up:
			aim -= cmd.units
		case forward:
			horiz += cmd.units
			depth += aim * cmd.units
		}
	}

	return horiz * depth, nil
}

type direction int

func directionFromString(s string) direction {
	switch s {
	case "forward":
		return forward
	case "down":
		return down
	case "up":
		return up
	default:
		return unset
	}
}

func unitsFromString(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

const (
	unset direction = iota
	forward
	down
	up
)

type command struct {
	dir   direction
	units int
}

func parseInput(raw string) ([]command, error) {
	cmds := []command{}
	rows := strings.Split(raw, "\n")
	for _, row := range rows {
		row = strings.TrimSpace(row)

		if row == "" {
			continue
		}

		parts := strings.Split(row, " ")
		command := command{
			dir:   directionFromString(parts[0]),
			units: unitsFromString(parts[1]),
		}
		cmds = append(cmds, command)
	}

	return cmds, nil
}
