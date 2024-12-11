package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"nikc.org/aoc2024/util"
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
	first := solveFirst(input)
	second := solveSecond(input)

	fmt.Fprint(out, "=====[ Day 03 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(i string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllString(i, -1)
	tot := 0

	for _, m := range matches {
		sm := re.FindStringSubmatch(m)
		if len(sm) != 3 {
			panic(errors.New("invalid match"))
		}

		nums := util.Fmap(util.MustAtoi, sm[1:])
		tot += nums[0] * nums[1]
	}

	return tot
}

func solveSecond(i string) int {
	doop := regexp.MustCompile(`do\(\)`)
	donotop := regexp.MustCompile(`don't\(\)`)

	i0 := ""

	for {
		disable := donotop.FindStringIndex(i)
		if disable == nil {
			// append the remaining input and exit
			i0 += i
			break
		}

		// copy from input until don't()
		i0 += i[0:disable[0]]

		// strip input until after don't()
		i = i[disable[1]:]

		enable := doop.FindStringIndex(i)
		if enable == nil {
			// if no more enables are found, we're done
			break
		}

		// strip input until after do()
		i = i[enable[1]:]
	}

	return solveFirst(i0)
}
