package main

import (
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
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
	fmt.Fprint(out, "=====[ Day 10 ]=====\n")

	fmt.Printf("first: %d\n", solveFirst(input))
	fmt.Printf("second:\n%s\n", solveSecond(input))

	return nil
}

func solveFirst(input string) int {
	observeAt := 20
	observations := []int{}
	c := Computer{}
	c.Load(input, 1)
	c.OnTick(func(tick int) {
		if tick == observeAt {
			fmt.Printf("%d: %d\n", tick, c.X())
			observations = append(observations, tick*c.X())
			observeAt += 40
		}
	})
	c.Run()

	total := observations[0]
	for _, n := range observations[1:] {
		total += n
	}

	return total
}

func solveSecond(input string) string {
	display := [][]byte{
		make([]byte, 40),
		make([]byte, 40),
		make([]byte, 40),
		make([]byte, 40),
		make([]byte, 40),
		make([]byte, 40),
	}

	c := Computer{}
	c.Load(input, 1)
	c.OnTick(func(tick int) {
		ypos := floor((tick - 1) / 40)
		xpos := (tick - 1) % 40

		if c.X() == xpos || c.X() == xpos-1 || c.X() == xpos+1 {
			display[ypos][xpos] = '#'
		} else {
			display[ypos][xpos] = '.'
		}
	})
	c.Run()

	return strings.Join(util.Fmap(func(bs []byte) string { return string(bs) }, display), "\n")
}

func floor(x int) int {
	return int(math.Floor(float64(x)))
}
