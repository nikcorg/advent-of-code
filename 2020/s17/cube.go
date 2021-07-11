package s17

import "sync"

type cube struct {
	pos        Position
	neighbours []Position
	isactive   bool
}

func (c *cube) Update(world *World, updates chan<- *Update, wg *sync.WaitGroup) {
	defer wg.Done()

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
		updates <- world.Update(c.pos, inactiveConwayCube)
	} else if !c.isactive && activeNeighbours == 3 {
		c.isactive = true
		updates <- world.Update(c.pos, activeConwayCube)
	}
}
