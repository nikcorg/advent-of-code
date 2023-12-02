package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	intrange "nikc.org/aoc2022/04/numrange"
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
	fmt.Fprint(out, "=====[ Day 04 ]=====\n")
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

func solveFirst(input string) (int, error) {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)

	total := 0

	for s.Scan() {
		splits := strings.Split(s.Text(), ",")
		if len(splits) != 2 {
			return 0, errors.New("invalid input")
		}

		ar, err := intrange.FromString(splits[0], strconv.Atoi)
		if err != nil {
			return 0, err
		}

		br, err := intrange.FromString(splits[1], strconv.Atoi)
		if err != nil {
			return 0, err
		}

		if ar.Contains(br) || br.Contains(ar) {
			total++
		}
	}

	return total, nil
}

func solveSecond(i string) (int, error) {
	s := bufio.NewScanner(strings.NewReader(i))
	s.Split(bufio.ScanLines)

	total := 0

	for s.Scan() {
		splits := strings.Split(s.Text(), ",")
		if len(splits) != 2 {
			return 0, errors.New("invalid input")
		}

		ar, err := intrange.FromString(splits[0], strconv.Atoi)
		if err != nil {
			return 0, err
		}

		br, err := intrange.FromString(splits[1], strconv.Atoi)
		if err != nil {
			return 0, err
		}

		if ar.Overlaps(br) {
			total++
		}
	}

	return total, nil
}
