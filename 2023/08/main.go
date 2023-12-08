package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"nikc.org/aoc2023/util"
)

var (
	//go:embed input.txt
	input string
)

func init() {
	if os.Getenv("LOG_LEVEL") == "debug" {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
	}
}

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	c, nm := parseInput(input)
	first := solveFirst(c, nm)
	c.Reset()
	second := solveSecond(c, nm)

	fmt.Fprint(out, "=====[ Day 08 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)
	return nil
}

func solveFirst(c cursor, nm nodemap) int {
	cur := "AAA"

	for steps := 1; true; steps++ {
		switch c.C() {
		case "L":
			cur = nm[cur].Left
		case "R":
			cur = nm[cur].Right
		}

		if cur == "ZZZ" {
			return steps
		}

		c.Adv()
	}

	return 0
}

func solveSecond(c cursor, nm nodemap) int {
	nodes := []string{}

	for n := range nm {
		if strings.HasSuffix(n, "A") {
			nodes = append(nodes, n)
		}
	}

	steppers, finished, cycles := len(nodes), 0, util.ParseIntOrDefault(os.Getenv("CYCLES"), 1)
	distances := make([]int, steppers)

	for steps := 1; finished < steppers*cycles; steps++ {
		for i, n := range nodes {
			switch c.C() {
			case "L":
				nodes[i] = nm[n].Left
			case "R":
				nodes[i] = nm[n].Right
			}

			if strings.HasSuffix(nodes[i], "Z") {
				finished++

				if distances[i] == 0 {
					slog.Debug("endnode reached", "ghost", i, "node", nodes[i], "steps", steps)
					distances[i] = steps
				} else {
					slog.Debug("endnode reached",
						"ghost", i, "node", nodes[i], "steps", steps,
						"identical?", steps%distances[i] == 0,
						"cycle#", steps/distances[i])
				}
			}
		}

		c.Adv()
	}

	var lcm int

	switch len(nodes) {
	case 1:
		lcm = distances[0]
	case 2:
		lcm = util.LCM(distances[0], distances[1])
	default:
		lcm = util.LCM(distances[0], distances[1], distances[2:]...)
	}

	return lcm
}

type nodemap map[string]node
type node struct {
	Left, Right string
}

func newCursor(steps []string) cursor {
	return cursor{0, steps}
}

type cursor struct {
	c     int
	steps []string
}

func (c cursor) C() string {
	return string(c.steps[c.c])
}

func (c *cursor) Reset() {
	c.c = 0
}

func (c *cursor) Adv() {
	c.c = (c.c + 1) % len(c.steps)
}

func parseInput(i string) (cursor, nodemap) {
	rNodeLine := regexp.MustCompile(`^([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)$`)
	r := bufio.NewScanner(strings.NewReader(i))

	r.Scan() // first line is the cursor path
	c := newCursor(strings.Split(r.Text(), ""))
	nm := make(nodemap)

	for r.Scan() {
		if r.Text() == "" {
			continue
		}
		n := rNodeLine.FindStringSubmatch(r.Text())
		nm[n[1]] = node{n[2], n[3]}
	}

	return c, nm
}
