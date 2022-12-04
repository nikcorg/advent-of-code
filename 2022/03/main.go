package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2022/03/set"
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
	fmt.Fprint(out, "=====[ Day 03 ]=====\n")
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
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	total := 0

	_ = total
	for scanner.Scan() {
		l, r := split(scanner.Text())
		o := overlap(l, r)

		for _, c := range o {
			total += priority(byte(c))
		}
	}

	return total, nil
}

func solveSecond(input string) (int, error) {
	total := 0

	r := bufio.NewReader(strings.NewReader(input))

	var (
		a, b, c string
		err     error
	)

	for !errors.Is(io.EOF, err) {
		a, b, c, err = nextThree(r)
		if err != nil && !errors.Is(io.EOF, err) {
			return 0, err
		}

		o := overlap(overlap(a, b), overlap(a, c))

		for _, c := range o {
			total += priority(byte(c))
		}
	}

	return total, nil
}

func nextThree(r *bufio.Reader) (string, string, string, error) {
	var (
		a, b, c string
		err     error
	)

	a, err = r.ReadString('\n')

	if err == nil {
		b, err = r.ReadString('\n')
	}

	if err == nil {
		c, err = r.ReadString('\n')
	}

	return a, b, c, err
}

func priority(ch byte) int {
	switch {
	case ch >= 'a' && ch <= 'z':
		return int(ch-'a') + 1
	case ch >= 'A' && ch <= 'Z':
		return int(ch-'A') + 27
	}
	return 0
}

func overlap(left, right string) string {
	o, l, r := "", set.New([]byte(left)), set.New([]byte(right))

	if r.Size() > l.Size() {
		l, r = r, l
	}

	for x := range l {
		if _, ok := r[byte(x)]; ok {
			o += string(x)
		}
	}

	return o
}

func split(inp string) (string, string) {
	p := len(inp) / 2
	return inp[0:p], inp[p:]
}
