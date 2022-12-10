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

	fmt.Fprintf(out, "first: %d\n", solveFirst(input))
	fmt.Fprintf(out, "second:\n%s\n", solveSecond(input))

	return nil
}

func solveFirst(input string) int {
	observed := 0
	c := NewComputer(input)
	c.OnTick(func(tick int) {
		if tick >= 20 && tick <= 220 && (tick-20)%40 == 0 {
			observed += tick * c.X()
		}
	})
	c.Run()

	return observed
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

	c := NewComputer(input)
	c.OnTick(func(tick int) {
		ypos := floor((tick - 1) / 40)
		xpos := (tick - 1) % 40

		if c.X()-1 > xpos || c.X()+1 < xpos {
			display[ypos][xpos] = '.'
		} else {
			display[ypos][xpos] = '#'
		}
	})
	c.Run()

	return strings.Join(util.Fmap(func(bs []byte) string { return string(bs) }, display), "\n")
}

func floor(x int) int {
	return int(math.Floor(float64(x)))
}
