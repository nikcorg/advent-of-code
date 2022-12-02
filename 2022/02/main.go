package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	fmt.Fprint(out, "=====[ Day 02 ]=====\n")

	var (
		first, second int
		err           error
	)

	if first, err = solveFirst(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %d\n", first)

	if second, err = solveSecond(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

const (
	rock     = 1
	paper    = 2
	scissors = 3
	lose     = 0
	draw     = 3
	win      = 6
)

func solveFirst(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		switch scanner.Text() {
		// opponent plays rock
		case "A X": // rock - rock
			score += draw + rock
		case "A Y": // rock - paper
			score += win + paper
		case "A Z": // rock - scissors
			score += lose + scissors

		// opponent plays paper
		case "B X": // paper - rock
			score += lose + rock
		case "B Y": // paper - paper
			score += draw + paper
		case "B Z": // paper - scissors
			score += win + scissors

		// opponent plays scissors
		case "C X": // scissors - rock
			score += win + rock
		case "C Y": // scissors - paper
			score += lose + paper
		case "C Z": // scissors - scissors
			score += draw + scissors
		}
	}
	return score, nil
}

func solveSecond(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		switch scanner.Text() {
		// play to lose
		case "A X": // rock - scissors
			score += lose + scissors
		case "B X": // paper - rock
			score += lose + rock
		case "C X": // scissors - paper
			score += lose + paper

		// play to draw
		case "A Y": // rock
			score += draw + rock
		case "B Y": // paper
			score += draw + paper
		case "C Y": // scissors
			score += draw + scissors

		// play to win
		case "A Z": // rock - paper
			score += win + paper
		case "B Z": // paper - scissors
			score += win + scissors
		case "C Z": // scissors - rock
			score += win + rock
		}
	}
	return score, nil
}
