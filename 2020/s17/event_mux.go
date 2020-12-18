package s17

import (
	"errors"
	"sync"
)

var (
	ErrMuxxerSourceClosed = errors.New("source channel closed")
)

type EventMuxxer struct {
	source     chan *WorldEvent
	chans      []chan *WorldEvent
	chansMutex sync.Mutex
	closed     bool
	subsWait   sync.WaitGroup
}

func NewMuxxer(source chan *WorldEvent) *EventMuxxer {
	r := &EventMuxxer{source: source, closed: false}

	go r.listen()

	return r
}

func (m *EventMuxxer) closeAll() {
	m.chansMutex.Lock()
	defer m.chansMutex.Unlock()

	m.closed = true
	for _, c := range m.chans {
		close(c)
	}
}

func (m *EventMuxxer) listen() {
	defer m.closeAll()

	for {
		select {
		case v, ok := <-m.source:
			if !ok {
				return
			}
			m.bcast(v)
		}
	}
}

func (m *EventMuxxer) bcast(v *WorldEvent) {
	if m.closed {
		return
	}

	m.subsWait.Wait()

	m.chansMutex.Lock()
	defer m.chansMutex.Unlock()

	for _, c := range m.chans {
		c <- v
	}
}

func (m *EventMuxxer) Recv() <-chan *WorldEvent {
	out := make(chan *WorldEvent, 1)
	m.subsWait.Add(1)

	go func() {
		m.chansMutex.Lock()
		defer m.chansMutex.Unlock()
		defer m.subsWait.Done()
		m.chans = append(m.chans, out)
	}()

	if m.closed {
		close(out)
	} else {
	}

	return out
}
