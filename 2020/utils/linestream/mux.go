package linestream

type Muxxer struct {
	source ReadOnlyLineChan
	chans  []WriteOnlyLineChan
}

func NewMuxxer(source ReadOnlyLineChan) *Muxxer {
	r := &Muxxer{source: source}

	go r.listen()

	return r
}

func (m *Muxxer) cleanup() {
	for _, c := range m.chans {
		close(c)
	}
}

func (m *Muxxer) listen() {
	defer m.cleanup()

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
	for _, c := range m.chans {
		c <- v
	}
}

func (m *Muxxer) Recv() ReadOnlyLineChan {
	out := make(LineChan)
	m.chans = append(m.chans, out)
	return out
}
