package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"

	"nikc.org/aoc2022/util/set"
)

var (
	//go:embed input.txt
	input string

	errBoundaryNotFound       = errors.New("packet boundary not found")
	errStartOfPacketNotFound  = errors.New("start of packet not found")
	errStartOfMessageNotFound = errors.New("start of message not found")
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	var (
		first, second int
		err           error
	)

	fmt.Fprint(out, "=====[ Day 06 ]=====\n")

	if first, err = solveFirst([]byte(input)); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %d\n", first)

	if second, err = solveSecond([]byte(input)); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(d []byte) (int, error) {
	pos, err := findPacketBoundary(4, d)
	if err != nil {
		return 0, errStartOfPacketNotFound
	}

	return pos, nil
}

func solveSecond(d []byte) (int, error) {
	pos, err := findPacketBoundary(14, d)
	if err != nil {
		return 0, errStartOfMessageNotFound
	}

	return pos, nil
}

func findPacketBoundary(size int, d []byte) (int, error) {
	if len(d) >= size {
		for n := 0; n < len(d)-size; n++ {
			if set.New(d[n:n+size]...).Size() == size {
				return n + size, nil
			}
		}
	}

	return 0, errBoundaryNotFound
}
