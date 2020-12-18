package s17

import (
	"context"
	"sync"
)

const activeConwayCube = "#"
const inactiveConwayCube = "."

type Update struct {
	turn       int
	pos        Position
	state      string
	wasApplied bool
}

type World struct {
	ctx               context.Context
	turn              int
	currentState      map[string]string
	stateUpdates      []*Update
	coordinateFactory func(Position) []Position
	cubes             []*cube
	updates           chan *Update
}

func newWorld(ctx context.Context, cf func(Position) []Position) *World {
	w := World{}
	w.ctx = ctx
	w.currentState = make(map[string]string)
	w.coordinateFactory = cf
	w.updates = make(chan *Update)

	go w.collectUpdates()

	return &w
}

func (w *World) collectUpdates() {
	for {
		select {
		case update := <-w.updates:
			w.stateUpdates = append(w.stateUpdates, update)
		}
	}
}

func (w *World) NextTurn() {
	w.turn++
	w.stateUpdates = []*Update{}

	wg := sync.WaitGroup{}

	for _, c := range w.cubes {
		wg.Add(1)
		go c.Update(w, w.updates, &wg)
	}

	wg.Wait()

	// due to the concurrency, despite not using buffered channels, an extra
	// update can sometimes sneak in while they're already being applied, so
	// we need to guard against not having processed them all
	applied := 0
	for applied < len(w.stateUpdates) {
		applied += w.EndTurn()
	}
}

func (w *World) EndTurn() int {
	indirectUpdates := map[string]*Update{} // using a map with `x,y,z` as key, as it will avoid duplicate updates
	appliedUpdates := 0
	// fmt.Println("committing", len(w.stateUpdates), "updates from", len(w.cubes), "workers")
	for _, u := range w.stateUpdates {
		if u.wasApplied {
			continue
		}

		k := u.pos.String()
		w.currentState[k] = u.state

		// spin up a cube on the spot, if it's the 0th turn
		if u.turn == 0 {
			w.cubes = append(w.cubes, &cube{u.pos, w.coordinateFactory(u.pos), u.state == activeConwayCube})
		} else {
			appliedUpdates++
			u.wasApplied = true
		}

		// expand onto neighbouring coordinates
		for _, c := range w.coordinateFactory(u.pos) {
			k := c.String()
			indirectUpdates[k] = &Update{w.turn, c, inactiveConwayCube, false}
		}
	}

	for k, u := range indirectUpdates {
		if _, ok := w.currentState[k]; !ok {
			w.cubes = append(w.cubes, &cube{u.pos, w.coordinateFactory(u.pos), u.state == activeConwayCube})
			w.currentState[k] = inactiveConwayCube
		}
	}

	// fmt.Println(len(w.cubes), "at the end of turn", w.turn)
	return appliedUpdates
}

func (w *World) Update(pos Position, newState string) *Update {
	return &Update{w.turn, pos, newState, false}
}

func (w *World) AlterStateAt(pos Position, newState string) {
	w.stateUpdates = append(w.stateUpdates, w.Update(pos, newState))
}

func (w *World) StateAt(pos Position) string {
	k := pos.String()

	if s, ok := w.currentState[k]; ok {
		return s
	}

	return inactiveConwayCube
}

func (w *World) ActiveCubes() int {
	actives := 0

	for _, v := range w.currentState {
		if v == activeConwayCube {
			actives++
		}
	}

	return actives
}
