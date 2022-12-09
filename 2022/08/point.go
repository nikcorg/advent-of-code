package main

import "math"

type Point struct {
	x, y int
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Equals(q Point) bool {
	return p.x == q.x && p.y == q.y
}

func (p Point) DistanceTo(q Point) float64 {
	distX := math.Abs(float64(p.x - q.x))
	distY := math.Abs(float64(p.y - q.y))
	return math.Sqrt(math.Pow(distX, 2) + math.Pow(distY, 2))
}

func pointGenerator(start Point, translateBy Point, until Point) <-chan Point {
	at := start
	c := make(chan Point)

	go func() {
		for {
			c <- at
			at = at.Add(translateBy)
			if at.Equals(until) {
				break
			}
		}

		close(c)
	}()

	return c
}
