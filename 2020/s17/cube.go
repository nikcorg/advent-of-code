package s17

import (
	"context"
	"sync"
)

func cube(ctx context.Context, pos Position, neighbours []Position, initialState bool, world *World, spinup *sync.WaitGroup) {
	defer spinup.Done()

	events := world.Events()
	isactive := initialState

	var (
		evt *WorldEvent
		ok  bool
	)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case evt, ok = <-events:
				if !ok {
					return
				}
			}

			switch evt {
			default:
				activeNeighbours := 0
				for _, n := range neighbours {
					s := world.StateAt(n)
					if s == activeConwayCube {
						activeNeighbours++
					}
				}

				if isactive && activeNeighbours != 2 && activeNeighbours != 3 {
					isactive = false
					world.AlterStateAt(pos, inactiveConwayCube)
				} else if !isactive && activeNeighbours == 3 {
					world.AlterStateAt(pos, activeConwayCube)
					isactive = true
				}

				evt.Done()
			}
		}
	}()
}
