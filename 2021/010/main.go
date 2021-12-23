package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
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

type puzzleInput []string

func parseInput(raw string) (puzzleInput, error) {
	return reject(strings.Split(raw, "\n"), empty), nil
}

var (
	invChunkErrValue = map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	incLineErrValue  = map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	validPair        = map[rune]rune{
		')': '(', '>': '<', ']': '[', '}': '{',
		'(': ')', '<': '>', '[': ']', '{': '}',
	}
	errInvalidChunk   = errors.New("invalid chunk")
	errIncompleteLine = errors.New("incomplete")
)

func solveFirst(input puzzleInput) (int, error) {
	score := 0
	for _, line := range input {
		if ch, err := checkLine(line); errors.Is(err, errInvalidChunk) {
			score += invChunkErrValue[ch[0]]
		}
	}
	return score, nil
}

func solveSecond(input puzzleInput) (int, error) {
	scores := []int{}
	for _, line := range input {
		if stack, err := checkLine(line); errors.Is(err, errIncompleteLine) {
			lineScore := 0
			rstack := reverse(stack)
			for _, ch := range rstack {
				lineScore *= 5
				lineScore += incLineErrValue[validPair[ch]]
			}
			scores = append(scores, lineScore)
		}
	}
	sort.Slice(scores, func(a, b int) bool {
		return scores[a] < scores[b]
	})
	middleScore := scores[len(scores)/2]
	return middleScore, nil
}

func checkLine(line string) ([]rune, error) {
	var (
		stack = []rune{}
	)

	for _, s := range line {
		switch s {
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != validPair[s] {
				return []rune{s}, errInvalidChunk
			}
			stack = stack[0 : len(stack)-1]
		default:
			stack = append(stack, s)
		}

	}

	if len(stack) > 0 {
		return stack, errIncompleteLine
	}

	return nil, nil
}

func empty(s string) bool {
	return s == ""
}

func reject[T any](in []T, f func(T) bool) []T {
	out := []T{}
	for _, x := range in {
		if f(x) {
			continue
		}
		out = append(out, x)
	}
	return out
}

func reverse[T any](in []T) []T {
	out := make([]T, len(in))
	copy(out, in)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
