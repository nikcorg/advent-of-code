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
	enqueuedUpdates   []*Update
	coordinateFactory func(Position) []Position
	cubes             []*cube
}

func newWorld(ctx context.Context, cf func(Position) []Position) *World {
	w := World{}
	w.ctx = ctx
	w.currentState = make(map[string]string)
	w.coordinateFactory = cf

	return &w
}

func (w *World) NextTurn() {
	w.turn++
	w.enqueuedUpdates = []*Update{}
	updates := make(chan *Update)

	go func() {
		defer close(updates)
		wg := sync.WaitGroup{}

		for _, c := range w.cubes {
			wg.Add(1)
			go c.Update(w, updates, &wg)
		}

		wg.Wait()
	}()

	pendingUpdates := []*Update{}

	for update := range updates {
		pendingUpdates = append(pendingUpdates, update)
	}

	for _, u := range pendingUpdates {
		w.applyUpdate(u)
	}
}

func (w *World) applyUpdate(u *Update) {
	indirectUpdates := map[string]*Update{} // using a map with `x,y,z` as key, as it will avoid duplicate updates
	k := u.pos.String()
	w.currentState[k] = u.state

	// spin up a cube on the spot, if it's the 0th turn
	if u.turn == 0 {
		w.cubes = append(w.cubes, &cube{u.pos, w.coordinateFactory(u.pos), u.state == activeConwayCube})
	} else {
		u.wasApplied = true
	}

	// expand onto neighbouring coordinates
	for _, c := range w.coordinateFactory(u.pos) {
		k := c.String()
		indirectUpdates[k] = &Update{w.turn, c, inactiveConwayCube, false}
	}

	for k, u := range indirectUpdates {
		if _, ok := w.currentState[k]; !ok {
			w.cubes = append(w.cubes, &cube{u.pos, w.coordinateFactory(u.pos), u.state == activeConwayCube})
			w.currentState[k] = inactiveConwayCube
		}
	}

}

func (w *World) EndTurn() {
	for _, u := range w.enqueuedUpdates {
		w.applyUpdate(u)
	}
}

func (w *World) Update(pos Position, newState string) *Update {
	return &Update{w.turn, pos, newState, false}
}

func (w *World) AlterStateAt(pos Position, newState string) {
	w.enqueuedUpdates = append(w.enqueuedUpdates, w.Update(pos, newState))
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
