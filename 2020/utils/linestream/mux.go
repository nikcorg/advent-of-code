package linestream

import (
	"errors"
	"sync"
)

var (
	ErrMuxxerSourceClosed = errors.New("source channel closed")
)

type Muxxer struct {
	source     ReadOnlyLineChan
	chans      []WriteOnlyLineChan
	chansMutex sync.Mutex
	closed     bool
}

func NewMuxxer(source ReadOnlyLineChan) *Muxxer {
	r := &Muxxer{source: source, closed: false}

	go r.listen()

	return r
}

func (m *Muxxer) closeAll() {
	m.closed = true
	for _, c := range m.chans {
		close(c)
	}
}

func (m *Muxxer) listen() {
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

func (m *Muxxer) bcast(v *Line) {
	if m.closed {
		return
	}

	for _, c := range m.chans {
		c <- v
	}
}

func (m *Muxxer) Recv() ReadOnlyLineChan {
	out := make(LineChan)

	if m.closed {
		close(out)
	} else {
		m.chansMutex.Lock()
		defer m.chansMutex.Unlock()

		m.chans = append(m.chans, out)
	}

	return out
}
