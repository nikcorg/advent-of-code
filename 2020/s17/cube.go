package s17

type cube struct {
	pos        Position
	neighbours []Position
	isactive   bool
}

func (c *cube) Update(world *World) *Update {
	activeNeighbours := 0
	for _, n := range c.neighbours {
		s := world.StateAt(n)
		if s == activeConwayCube {
			activeNeighbours++
		}

		if activeNeighbours > 3 {
			break
		}
	}

	if c.isactive && activeNeighbours != 2 && activeNeighbours != 3 {
		c.isactive = false
		return world.Update(c.pos, inactiveConwayCube)
	} else if !c.isactive && activeNeighbours == 3 {
		c.isactive = true
		return world.Update(c.pos, activeConwayCube)
	}

	return nil
}
