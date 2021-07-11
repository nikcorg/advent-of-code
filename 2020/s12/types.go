package s12

import (
	"fmt"
)

type CompassDirection int

func (d CompassDirection) String() string {
	switch d {
	case east:
		return "E"
	case west:
		return "W"
	case north:
		return "N"
	case south:
		return "S"
	}

	panic(fmt.Errorf("invalid direction %d", d))
}

const (
	east CompassDirection = iota
	south
	west
	north
)

type Command int

const (
	moveForward Command = iota + 1
	rotRight
	rotLeft
	moveNorth
	moveSouth
	moveEast
	moveWest
)

type Instruction struct {
	cmd   Command
	units int
}

func (i Instruction) String() string {
	s := "#"

	switch i.cmd {
	case moveForward:
		s = "F"
	case rotRight:
		s = "R"
	case rotLeft:
		s = "L"
	case moveNorth:
		s = "N"
	case moveSouth:
		s = "S"
	case moveEast:
		s = "E"
	case moveWest:
		s = "W"
	}

	return fmt.Sprintf("%s%d", s, i.units)
}
