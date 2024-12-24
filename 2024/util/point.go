package util

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) Translate(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) Equals(p0 Point) bool {
	return p.X == p0.X && p.Y == p0.Y
}
