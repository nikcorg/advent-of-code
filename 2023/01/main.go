package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
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
	inputs := strings.Split(input, "\n")
	first := solveFirst(inputs)
	second := solveSecond(inputs)

	fmt.Fprint(out, "=====[ Day 01 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(input []string) int {
	r := regexp.MustCompile(`\d`)
	tot := 0
	for _, l := range input {
		ns := r.FindAllString(l, -1)
		a, _ := strconv.Atoi(ns[0])
		b, _ := strconv.Atoi(ns[len(ns)-1])
		tot += a*10 + b
	}
	return tot
}

func reverse(s string) string {
	xs := strings.Split(s, "")
	slices.Reverse(xs)
	return strings.Join(xs, "")
}

func solveSecond(input []string) int {
	r := regexp.MustCompile(`\d|one|eno|two|owt|three|eerht|four|rouf|five|evif|six|xis|seven|neves|eight|thgie|nine|enin`)
	stoi := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	tot := 0

	for _, l := range input {
		var a, b string = r.FindString(l), reverse(r.FindString(reverse(l)))

		for n := len(l) - 1; n > 0; n-- {
			s := l[n:]

			if x := r.FindString(s); x != "" {
				b = x
				break
			}
		}

		v := 0

		if n, ok := stoi[a]; ok {
			v += n * 10
		} else {
			n, _ = strconv.Atoi(a)
			v += n * 10
		}

		if n, ok := stoi[b]; ok {
			v += n
		} else {
			n, _ = strconv.Atoi(b)
			v += n
		}

		tot += v
	}

	return tot
}
