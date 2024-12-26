package util

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (p Point) Translate(dx, dy int) Point {
	return NewPoint(p.X+dx, p.Y+dy)
}

func (p Point) Equals(p0 Point) bool {
	return p.X == p0.X && p.Y == p0.Y
}

func (p Point) Add(p0 Point) Point {
	return p.Translate(p0.X, p0.Y)
}
