package main

import "nikc.org/aoc2022/util"

func pointGenerator(start util.Point, translateBy util.Point, until util.Point) <-chan util.Point {
	c := make(chan util.Point)

	go func() {
		for at := start; !at.Equals(until); at = at.Add(translateBy) {
			c <- at
		}

		close(c)
	}()

	return c
}
