package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"nikc.org/aoc2023/util"
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
	setup := game{Red: 12, Green: 13, Blue: 14}
	rounds := parseInput(strings.Split(input, "\n"))
	first := solveFirst(setup, rounds)
	second := solveSecond(rounds)

	fmt.Fprint(out, "=====[ Day 02 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type round struct {
	ID     int
	Reveal []game
}

type game struct {
	Red, Green, Blue int
}

var (
	rID     = regexp.MustCompile(`Game (\d+):`)
	rSample = regexp.MustCompile(`(\d+) (red|green|blue)`)
)

func parseInput(input []string) []round {
	rs := make([]round, len(input))
	for idx, l := range input {
		id := rID.FindStringSubmatch(l)[1]
		samples := strings.Split(strings.SplitN(l, ":", 2)[1], ";")

		rs[idx] = round{ID: util.MustParseInt(id)}

		for _, s := range samples {
			g := game{}
			for _, x := range strings.Split(s, ",") {
				for _, sample := range rSample.FindAllStringSubmatch(x, -1) {
					switch sample[2] {
					case "red":
						g.Red = util.MustParseInt(sample[1])
					case "green":
						g.Green = util.MustParseInt(sample[1])
					case "blue":
						g.Blue = util.MustParseInt(sample[1])
					}
				}
			}
			rs[idx].Reveal = append(rs[idx].Reveal, g)
		}

	}

	return rs
}

func solveFirst(g game, rounds []round) int {
	t := 0
	for _, r := range rounds {
		valid := true
		for _, s := range r.Reveal {
			if s.Green > g.Green || s.Red > g.Red || s.Blue > g.Blue {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		t += r.ID
	}
	return t
}

func solveSecond(rounds []round) int {
	t := 0

	for _, r := range rounds {
		mins := game{Red: 0, Green: 0, Blue: 0}

		for _, s := range r.Reveal {
			mins.Red = max(mins.Red, s.Red)
			mins.Green = max(mins.Green, s.Green)
			mins.Blue = max(mins.Blue, s.Blue)
		}
		t += mins.Red * mins.Green * mins.Blue
	}

	return t
}
