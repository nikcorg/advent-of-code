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
	turn    int
	x, y, z int
	state   string
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
	started      bool
	ctx          context.Context
	turn         int
	currentState map[string]string
	stateUpdates []*Update
	updates      chan *Update
	stateMutex   sync.Mutex
	events       *EventMuxxer
	eventSource  chan *WorldEvent
	cubeWorkers  int
	turnWG       *WorldEvent
}

func newWorld(ctx context.Context) *World {
	w := World{}
	w.ctx = ctx
	w.currentState = make(map[string]string)
	w.eventSource = make(chan *WorldEvent)
	w.events = NewMuxxer(w.eventSource)
	w.updates = make(chan *Update, 0)

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

	for _, u := range w.stateUpdates {
		if u.turn != w.turn {
			panic(errors.New("turn mismatch in update"))
		}

		k := fmt.Sprintf("%d,%d,%d", u.x, u.y, u.z)
		w.currentState[k] = u.state

		// spin up a cube on the spot, if it's the 0th turn
		if u.turn == 0 {
			go cube(w.ctx, u.x, u.y, u.z, u.state == activeConwayCube, w, &spinups)
			spinups.Add(1)
			w.cubeWorkers++
		}

		// expand onto neighbouring coordinates
		for _, c := range surroundingXYZCoords(u.x, u.y, u.z) {
			k := fmt.Sprintf("%d,%d,%d", c[0], c[1], c[2])
			indirectUpdates[k] = &Update{w.turn, c[0], c[1], c[2], inactiveConwayCube}
		}
	}

	w.stateUpdates = []*Update{}

	for k, u := range indirectUpdates {
		if _, ok := w.currentState[k]; !ok {
			go cube(w.ctx, u.x, u.y, u.z, false, w, &spinups)
			spinups.Add(1)
			w.currentState[k] = inactiveConwayCube
			w.cubeWorkers++
		}
	}

	spinups.Wait()
	w.turn++
}

func (w *World) AlterStateAt(x, y, z int, newState string) {
	w.updates <- &Update{w.turn, x, y, z, newState}
}

func (w *World) StateAt(x, y, z int) string {
	k := fmt.Sprintf("%d,%d,%d", x, y, z)

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
