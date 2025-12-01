package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

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

	fmt.Fprint(out, "=====[ Day 05 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) (map[int]map[int]struct{}, [][]int) {
	s := bufio.NewScanner(strings.NewReader(i))

	order := map[int]map[int]struct{}{}
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		instr := util.Fmap(util.MustAtoi, strings.Split(s.Text(), "|"))
		page := instr[0]
		before := instr[1]

		if befores, ok := order[page]; ok {
			befores[before] = struct{}{}
		} else {
			order[page] = map[int]struct{}{before: struct{}{}}
		}
	}

	runs := [][]int{}
	for s.Scan() {
		run := util.Fmap(util.MustAtoi, strings.Split(s.Text(), ","))
		runs = append(runs, run)
	}

	return order, runs
}

func solveFirst(i string) int {
	rules, runs := parseInput(i)

	validRuns := [][]int{}

	for _, run := range runs {
		if isValidRun(run, rules) {
			validRuns = append(validRuns, run)
		}
	}

	tot := 0

	for _, run := range validRuns {
		tot += run[len(run)/2]
	}

	return tot
}

func solveSecond(i string) int {
	rules, runs := parseInput(i)

	invalidRuns := [][]int{}

	for _, run := range runs {
		if !isValidRun(run, rules) {
			invalidRuns = append(invalidRuns, run)
		}
	}

	tot := 0

	for _, run := range invalidRuns {
		healed := healInvalidRun(run, rules)

		tot += healed[len(healed)/2]
	}

	return tot
}

func isValidRun(pages []int, rules map[int]map[int]struct{}) bool {
	printed := []int{pages[0]}

	for _, page := range pages[1:] {
		rule, ok := rules[page]
		if !ok {
			printed = append(printed, page)
			continue
		}

		for _, pn := range printed {
			if _, ok := rule[pn]; ok {
				return false
			}
		}

		printed = append(printed, page)
	}

	return true
}

func healInvalidRun(pages []int, rules map[int]map[int]struct{}) []int {
	printed := []int{pages[0]}

	for _, page := range pages[1:] {
		rule, ok := rules[page]
		if !ok {
			printed = append(printed, page)
			continue
		}

		spliced := false
		for i, pn := range printed {
			if _, ok := rule[pn]; ok {
				if i == 0 {
					printed = append([]int{page}, printed...)
				} else {
					before, after := util.CloneSlice(printed[0:i]), util.CloneSlice(printed[i:])
					printed = append(append(before, page), after...)
				}
				spliced = true
				break
			}

		}

		if spliced {
			continue
		}
		printed = append(printed, page)
	}

	return printed
}
