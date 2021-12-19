package main

import (
	"fmt"
	"strconv"
)

type point struct{ X, Y int }

func pointFromString(sx, sy string) point {
	x, _ := strconv.Atoi(sx)
	y, _ := strconv.Atoi(sy)

	return point{x, y}
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}
