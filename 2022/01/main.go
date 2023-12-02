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
	fmt.Fprint(out, "=====[ Day 01 ]=====\n")

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
	var (
		calories    int
		maxCalories int
	)

	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		l := s.Text()

		if l == "" {
			calories = 0
		} else {
			c, err := strconv.Atoi(l)
			if err != nil {
				return 0, err
			}

			calories += c
		}

		if calories > maxCalories {
			maxCalories = calories
		}
	}

	return maxCalories, nil
}

func solveSecond(input string) (int, error) {
	var (
		calories, a, b, c int
		err               error
		l                 string
	)

	// Use a Reader instead of a Scanner, because that gives us the EOF err
	r := bufio.NewReader(strings.NewReader(input))

	for !errors.Is(io.EOF, err) {
		if l, err = r.ReadString('\n'); err != nil && !errors.Is(io.EOF, err) {
			return 0, err
		}

		l = strings.TrimSpace(l)

		// Any non-empty value should be a number
		if l != "" {
			c, err := strconv.Atoi(l)
			if err != nil {
				return 0, err
			}

			calories += c
		}

		// Update and reset on EOF and empty lines (pack delimiter)
		if l == "" || errors.Is(io.EOF, err) {
			if calories > a {
				a, b, c = calories, a, b
			} else if calories > b {
				b, c = calories, b
			} else if calories > c {
				c = calories
			}
			calories = 0
		}
	}

	return a + b + c, nil
}
