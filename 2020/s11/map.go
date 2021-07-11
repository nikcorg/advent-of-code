package s11

import (
	"errors"
	"strings"
)

type tileKind int

const (
	unset tileKind = iota
	floorTile
	emptySeat
	occupiedSeat
	observer
)

var (
	errInvalidCoordinate = errors.New("invalid coordinate")
)

type floormap struct {
	tiles []tileKind
	width int
}

func (f *floormap) Init(w, h int) {
	f.width = w
	f.tiles = make([]tileKind, h*w)
}

func (f *floormap) SetTileAt(x, y int, v tileKind) {
	pos := y*f.width + x
	f.tiles[pos] = v
}

func (f *floormap) TileAt(x, y int) (tileKind, error) {
	if x < 0 || x >= f.Width() || y < 0 || y >= f.Height() {
		return unset, errInvalidCoordinate
	}

	pos := y*f.width + x

	if pos < 0 || pos >= len(f.tiles) {
		return unset, errInvalidCoordinate
	}

	return f.tiles[pos], nil
}

func (f *floormap) Occupied(x, y int) (bool, error) {
	t, err := f.TileAt(x, y)

	if err != nil {
		return false, err
	}

	return t == occupiedSeat, nil
}

func (f *floormap) Height() int {
	return len(f.tiles) / f.width
}

func (f *floormap) Width() int {
	return f.width
}

func (f *floormap) OccupiedAdjacent(x, y int) int {
	startX := x - 1
	startY := y - 1
	endX := x + 1
	endY := y + 1

	occupied := 0

	for checkY := startY; checkY <= endY; checkY++ {
		for checkX := startX; checkX <= endX; checkX++ {
			// Skip self
			if x == checkX && y == checkY {
				continue
			}

			isOccupied, err := f.Occupied(checkX, checkY)

			if err != nil && err == errInvalidCoordinate {
				continue
			} else if err != nil && err != errInvalidCoordinate {
				panic(err)
			} else if isOccupied {
				occupied++
			}
		}
	}

	return occupied
}

func (f *floormap) OccupiedVisibleFrom(x, y int) int {
	occupied := 0
	nearestSeatMap := map[string]tileKind{}
	done := false
	satisfied := 0

	for dist := 1; !done && satisfied < 8; dist++ {
		// check all directions clockwise
		checks := []struct {
			dir string
			x   int
			y   int
		}{
			{"N", x, y - dist},
			{"NE", x + dist, y - dist},
			{"E", x + dist, y},
			{"SE", x + dist, y + dist},
			{"S", x, y + dist},
			{"SW", x - dist, y + dist},
			{"W", x - dist, y},
			{"NW", x - dist, y - dist},
		}

		// stay positive!
		done = true

		for _, check := range checks {
			tile, err := f.TileAt(check.x, check.y)

			if err != nil && err == errInvalidCoordinate {
				continue
			} else if err != nil && err != errInvalidCoordinate {
				panic(err)
			}

			// at least one coordinate is still within bounds
			done = false

			if _, ok := nearestSeatMap[check.dir]; !ok && (tile == occupiedSeat || tile == emptySeat) {
				satisfied++
				switch tile {
				case occupiedSeat:
					occupied++
					fallthrough
				default:
					nearestSeatMap[check.dir] = tile
				}
			}
		}
	}

	return occupied
}

func (f *floormap) FromString(str string) {
	for n, l := range strings.Split(str, "\n") {
		if n == 0 {
			f.width = len(l)
		}

		xs := make([]tileKind, f.width)

		for m, c := range strings.Split(l, "") {
			switch c {
			case "L":
				xs[m] = emptySeat
			case "#":
				xs[m] = occupiedSeat
			default:
				xs[m] = floorTile
			}
		}

		f.tiles = append(f.tiles, xs...)
	}
}

func (f *floormap) String() string {
	s := ""

	for p := 0; p < len(f.tiles); p++ {
		if p > 0 && p%f.Width() == 0 {
			s += "\n"
		}
		switch f.tiles[p] {
		case emptySeat:
			s += "L"
		case occupiedSeat:
			s += "#"
		case observer:
			s += "X"
		default:
			s += "."
		}
	}

	return s
}
