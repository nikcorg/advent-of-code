package s17

import (
	"errors"
	"fmt"
)

type Position []int

var errUnsupportedCoordinatePosition = errors.New("unsupported coordinate position")

func (p *Position) String() string {
	pp := *p
	switch len(pp) {
	case 4:
		return fmt.Sprintf("%d,%d,%d,%d", pp[0], pp[1], pp[2], pp[3])
	case 3:
		return fmt.Sprintf("%d,%d,%d", pp[0], pp[1], pp[2])
	}

	panic(errUnsupportedCoordinatePosition)
}
