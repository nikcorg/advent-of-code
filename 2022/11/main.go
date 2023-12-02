package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
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
	fmt.Fprint(out, "=====[ Day 11 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", solveFirst(input))
	fmt.Fprintf(out, "second: %d\n", solveSecond(input))
	return nil
}

type operand int

const (
	multiply operand = iota
	add
)

func solveFirst(input string) int {
	ids, monkeys := parseMonkeys(bufio.NewScanner(strings.NewReader(input)))

	for n := 0; n < 20; n++ {
		for _, id := range ids {
			m := monkeys[id]

			for _, i := range m.items {
				worryLevel := m.Inspect(i)
				worryLevel = worryLevel / 3
				next := m.PassTo(worryLevel)

				monkeys[next].Receive(worryLevel)
			}

			m.items = []int64{}
		}
	}

	inspections := sort.IntSlice{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}

	sort.Sort(sort.Reverse(inspections))

	return inspections[0] * inspections[1]
}

func solveSecond(input string) int {
	ids, monkeys := parseMonkeys(bufio.NewScanner(strings.NewReader(input)))

	// This solution required looking for spoilers on r/adventofcode
	// Thanks go to u/_smallconfusion for ridding me of my huge confusion.
	// https://www.reddit.com/r/adventofcode/comments/zih7gf/comment/izr79go/?utm_source=share&utm_medium=web2x&context=3
	mod := int64(0)
	for _, m := range monkeys {
		if mod == 0 {
			mod = int64(m.testDivisor)
		} else {
			mod *= int64(m.testDivisor)
		}
	}

	for n := 0; n < 1e4; n++ {
		for _, id := range ids {
			m := monkeys[id]

			for _, i := range m.items {
				worryLevel := m.Inspect(i)
				next := m.PassTo(worryLevel)

				monkeys[next].Receive(worryLevel % mod)
			}

			m.items = []int64{}
		}
	}

	inspections := sort.IntSlice{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}

	sort.Sort(sort.Reverse(inspections))

	return inspections[0] * inspections[1]
}

type monkey struct {
	id          int
	items       []int64
	op          func(int64) int64
	testDivisor int
	nextIfTrue  int
	nextIfFalse int
	inspections int
}

func (m *monkey) Receive(i int64) {
	m.items = append(m.items, i)
}

func (m *monkey) PassTo(i int64) int {
	if i%int64(m.testDivisor) == 0 {
		return m.nextIfTrue
	} else {
		return m.nextIfFalse
	}
}

func (m *monkey) Inspect(i int64) int64 {
	m.inspections++
	return m.op(i)
}

func parseMonkeys(s *bufio.Scanner) ([]int, map[int]*monkey) {
	monkeys := map[int]*monkey{}

	matchOp := regexp.MustCompile(`new = old (\+|\*) (\d+|old)`)

	var (
		ids = []int{}
		m   *monkey
	)

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}
		switch line[0:4] {
		case "Monk": // Monkey <N>
			m = &monkey{id: util.MustAtoi(s.Text()[7:8])}
			ids = append(ids, m.id)
			monkeys[m.id] = m
		case "Star": // Starting items:
			m.items = util.Fmap(parseInt64, strings.Split(line[16:], ", "))
		case "Oper": // Operation:
			matches := matchOp.FindStringSubmatch(line)
			switch matches[1] {
			case "*":
				if matches[2] == "old" {
					m.op = square
				} else {
					m.op = multiplier(parseInt64(matches[2]))
				}
			case "+":
				if matches[2] == "old" {
					m.op = double
				} else {
					m.op = adder(parseInt64(matches[2]))
				}
			}
		case "Test": // Test: divisible by
			m.testDivisor = util.MustAtoi(line[19:])
		case "If t": // If true: throw to monkey
			m.nextIfTrue = util.MustAtoi(line[25:])
		case "If f": // If false: throw to monkey
			m.nextIfFalse = util.MustAtoi(line[26:])
		}
	}

	return ids, monkeys
}

func parseInt64(s string) int64 {
	return int64(util.MustAtoi(s))
}

func double(n int64) int64 {
	return n + n
}

func square(n int64) int64 {
	return n * n
}

func multiplier(n int64) func(int64) int64 {
	return func(m int64) int64 {
		return n * m
	}
}

func adder(n int64) func(int64) int64 {
	return func(m int64) int64 {
		return n + m
	}
}
