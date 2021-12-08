package s12

import "fmt"

type Vessel struct {
	facing CompassDirection
	*Point
}

func (v *Vessel) String() string {
	f := ""

	switch v.facing {
	case north:
		f = "N"
	case south:
		f = "S"
	case east:
		f = "E"
	case west:
		f = "W"
	}

	return fmt.Sprintf("facing %s at N/S %d E/W %d", f, v.Y, v.X)
}

func (v *Vessel) Forward(units int) {
	switch v.facing {
	case north:
		v.Move(north, units)
	case south:
		v.Move(south, units)
	case west:
		v.Move(west, units)
	case east:
		v.Move(east, units)
	}
}

func (v *Vessel) Rotate(dir Command, units int) {
	directionsLeft := []CompassDirection{north, west, south, east}
	directionsRight := []CompassDirection{north, east, south, west}

	switch dir {
	case rotLeft:
		// match order to current facing
		for directionsLeft[0] != v.facing {
			directionsLeft = append(directionsLeft[1:], directionsLeft[0])
		}

		// step around the compass
		for n := 0; n < units/90; n++ {
			directionsLeft = append(directionsLeft[1:], directionsLeft[0])
		}

		v.facing = directionsLeft[0]

	case rotRight:
		// match order to current facing
		for directionsRight[0] != v.facing {
			directionsRight = append(directionsRight[1:], directionsRight[0])
		}

		// step around the compass
		for n := 0; n < units/90; n++ {
			directionsRight = append(directionsRight[1:], directionsRight[0])
		}

		v.facing = directionsRight[0]
	}
}

func (v *Vessel) Move(dir CompassDirection, units int) {
	switch dir {
	case north:
		v.Y -= units
	case south:
		v.Y += units
	case west:
		v.X -= units
	case east:
		v.X += units
	}
}

func (v *Vessel) ForwardTo(point *Point, units int) {

}

func (v *Vessel) MoveTo(point *Point, units int) {

}
