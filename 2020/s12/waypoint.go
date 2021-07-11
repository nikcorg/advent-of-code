package s12

import "fmt"

type Point struct {
	Y int
	X int
}

func (w *Point) Diff(ref *Point) (int, int) {
	return w.X - ref.X, w.Y - ref.Y
}

func (w *Point) Move(dir Command, units int) {
	switch dir {
	case moveNorth:
		w.Y -= units
	case moveSouth:
		w.Y += units
	case moveWest:
		w.X -= units
	case moveEast:
		w.X += units
	}
}

func (w *Point) Translate(x, y int) {
	w.X += x
	w.Y += y
}

func (w *Point) Rotate(ref *Point, dir Command, units int) {
	if w.Equal(ref) {
		fmt.Println("no need to rotate as reference is on waypoint")
		return
	}

	// Translate the point so it sits on the origin
	w.Translate(-ref.X, -ref.Y)

	for n := 0; n < units/90; n++ {
		switch dir {
		case rotLeft:
			w.X, w.Y = w.Y, -w.X

		case rotRight:
			w.X, w.Y = -w.Y, w.X

		default:
			panic(fmt.Errorf("invalid rotation: %v", dir))
		}
	}

	// Move the point back around the reference point
	w.Translate(ref.X, ref.Y)
}

func (w *Point) Equal(ref *Point) bool {
	return w.Y == ref.Y && w.X == ref.X
}

func (w *Point) String() string {
	return fmt.Sprintf("{NS=%d, EW=%d}", w.Y, w.X)
}
