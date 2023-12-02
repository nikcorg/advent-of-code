package main

import (
	"bufio"
	"strings"

	"nikc.org/aoc2022/util"
)

type Computer struct {
	clock    int
	x        int
	program  *bufio.Scanner
	observer func(int)
}

func NewComputer(program string) *Computer {
	return &Computer{x: 1, clock: 1, program: bufio.NewScanner(strings.NewReader(program))}
}

func (c *Computer) OnTick(f func(int)) {
	c.observer = f
}

func (c *Computer) X() int {
	return c.x
}

func (c *Computer) Run() {
	c.clock = 1

	for c.program.Scan() {
		line := c.program.Text()
		switch line[0:4] {
		case "noop":
			c.tick(1)
		case "addx":
			c.tick(2)
			c.x += util.MustAtoi(line[5:])
		}
	}
}

func (c *Computer) tick(n int) {
	for n := n; n > 0; n-- {
		c.observer(c.clock)
		c.clock++
	}
}
