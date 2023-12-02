package dijkstra

type Point struct {
	X, Y int
}

func (p Point) Neighbours() <-chan Point {
	c := make(chan Point)
	go func() {
		defer close(c)
		for _, p := range [4]Point{
			NewPoint(p.X, p.Y-1),
			NewPoint(p.X, p.Y+1),
			NewPoint(p.X-1, p.Y),
			NewPoint(p.X+1, p.Y),
		} {
			c <- p
		}
	}()
	return c
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}
