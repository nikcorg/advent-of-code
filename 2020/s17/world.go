package s17

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const activeConwayCube = "#"
const inactiveConwayCube = "."

type Update struct {
	turn  int
	pos   Position
	state string
}

type WorldEvent struct {
	wg  sync.WaitGroup
	n   int
	sem sync.Mutex
}

func (e *WorldEvent) Add(delta int) {
	e.sem.Lock()
	defer e.sem.Unlock()
	e.wg.Add(delta)
	e.n += delta
}

func (e *WorldEvent) Done() {
	e.sem.Lock()
	defer e.sem.Unlock()
	e.wg.Done()
	e.n--
	// fmt.Println("Done event received", e.n, "remaining")
}

func (e *WorldEvent) Wait() {
	e.wg.Wait()
}

type World struct {
	started           bool
	ctx               context.Context
	turn              int
	currentState      map[string]string
	stateUpdates      []*Update
	updates           chan *Update
	stateMutex        sync.Mutex
	events            *EventMuxxer
	eventSource       chan *WorldEvent
	cubeWorkers       int
	turnWG            *WorldEvent
	coordinateFactory func(Position) []Position
}

func newWorld(ctx context.Context, cf func(Position) []Position) *World {
	w := World{}
	w.ctx = ctx
	w.currentState = make(map[string]string)
	w.eventSource = make(chan *WorldEvent)
	w.events = NewMuxxer(w.eventSource)
	w.updates = make(chan *Update, 1)
	w.coordinateFactory = cf
	go w.receiveUpdates()

	return &w
}

func (w *World) receiveUpdates() {
	for {
		select {
		case <-w.ctx.Done():
			return

		case u := <-w.updates:
			w.stateUpdates = append(w.stateUpdates, u)
		}
	}
}

func (w *World) Events() <-chan *WorldEvent {
	return w.events.Recv()
}

func (w *World) NextTurn() {
	w.StartTurn()
	w.turnWG.Wait()
	w.EndTurn()
}

func (w *World) StartTurn() {
	w.stateMutex.Lock()
	defer w.stateMutex.Unlock()

	w.started = true
	w.turnWG = &WorldEvent{}
	w.turnWG.Add(w.cubeWorkers)

	w.eventSource <- w.turnWG
}

func (w *World) EndTurn() {
	w.stateMutex.Lock()
	defer w.stateMutex.Unlock()

	spinups := sync.WaitGroup{}
	indirectUpdates := map[string]*Update{} // using a map with `x,y,z` as key, as it will avoid duplicate updates

	fmt.Println("committing", len(w.stateUpdates), "updates from", w.cubeWorkers, "workers", len(w.updates))

	for _, u := range w.stateUpdates {
		if u.turn != w.turn {
			panic(errors.New("turn mismatch in update"))
		}

		k := u.pos.String()
		w.currentState[k] = u.state

		// spin up a cube on the spot, if it's the 0th turn
		if u.turn == 0 {
			spinups.Add(1)
			go cube(w.ctx, u.pos, w.coordinateFactory(u.pos), u.state == activeConwayCube, w, &spinups)
			w.cubeWorkers++
		}

		// expand onto neighbouring coordinates
		for _, c := range w.coordinateFactory(u.pos) {
			k := c.String()
			indirectUpdates[k] = &Update{w.turn, c, inactiveConwayCube}
		}
	}

	w.stateUpdates = []*Update{}

	for k, u := range indirectUpdates {
		if _, ok := w.currentState[k]; !ok {
			spinups.Add(1)
			go cube(w.ctx, u.pos, w.coordinateFactory(u.pos), false, w, &spinups)
			w.currentState[k] = inactiveConwayCube
			w.cubeWorkers++
		}
	}

	spinups.Wait()
	fmt.Println("ended turn", w.turn, "with", w.cubeWorkers, "workers")
	w.turn++
}

func (w *World) AlterStateAt(pos Position, newState string) {
	w.updates <- &Update{w.turn, pos, newState}
}

func (w *World) StateAt(pos Position) string {
	k := pos.String() // fmt.Sprintf("%d,%d,%d", x, y, z)

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
