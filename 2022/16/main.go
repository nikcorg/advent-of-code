package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"nikc.org/aoc2022/util"
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
	fmt.Fprint(out, "=====[ Day 16 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", solveFirst(input))
	return nil
}

type valve struct {
	Label    string
	FlowRate int
	Next     []string
}

func parseInput(s *bufio.Scanner) (string, map[string]*valve) {
	match := regexp.MustCompile(`([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z,\s]+)$`)
	vs := map[string]*valve{}
	first := ""
	for s.Scan() {
		ms := match.FindStringSubmatch(s.Text())
		if ms == nil {
			panic(fmt.Errorf("unexpected input: %s", s.Text()))
		}

		v := valve{
			Label:    ms[1],
			FlowRate: util.MustAtoi(ms[2]),
			Next:     strings.Split(ms[3], ", "),
		}

		vs[v.Label] = &v

		if first == "" {
			first = v.Label
		}

	}

	return first, vs
}

func solveFirst(input string) int {
	root, vs := parseInput(bufio.NewScanner(strings.NewReader(input)))
	_ = vs

	curr := vs[root]

	return 0
}
