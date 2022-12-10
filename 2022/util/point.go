package util

import "math"

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) DistanceTo(q Point) float64 {
	distX := math.Abs(float64(p.X - q.X))
	distY := math.Abs(float64(p.Y - q.Y))
	return math.Sqrt(math.Pow(distX, 2) + math.Pow(distY, 2))
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}
