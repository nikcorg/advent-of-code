package packet

import (
	"fmt"
	"math"
	"strings"
)

const (
	GreaterThan = iota - 1
	SameSame
	LessThan
)

type Packet interface {
	Compare(Packet) int
}

func NewIntPacket(d int) IntPacket {
	return IntPacket{value: d}
}

type IntPacket struct {
	value int
}

func (p IntPacket) String() string {
	return fmt.Sprintf("%d", p.value)
}

func (p IntPacket) ToList() Packet {
	return NewListPacket(p)
}

func (p IntPacket) Compare(q Packet) int {
	switch q := q.(type) {
	case ListPacket:
		return p.ToList().Compare(q)

	case IntPacket:
		switch {
		case p.value == q.value:
			return SameSame

		case p.value < q.value:
			return GreaterThan

		case p.value > q.value:
			return LessThan
		}
	}

	panic(fmt.Errorf("impossible comparison: %+v <> %+v", p, q))
}

func NewListPacket(vs ...Packet) ListPacket {
	return ListPacket{values: vs}
}

type ListPacket struct {
	values []Packet
}

func (p ListPacket) String() string {
	ss := []string{}
	for _, v := range p.values {
		ss = append(ss, v.(fmt.Stringer).String())
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ","))
}

func (p *ListPacket) Append(q Packet) {
	p.values = append(p.values, q)
}

func (p ListPacket) Compare(q Packet) int {
	switch q := q.(type) {
	case ListPacket:
		switch {
		case len(p.values) == 0 && len(q.values) == 0:
			return SameSame

		case len(p.values) == 0 && len(q.values) > 0:
			return GreaterThan

		case len(q.values) == 0:
			return LessThan
		}

		for _, pq := range zip(p.values, q.values) {
			if pq[0] == nil {
				return GreaterThan
			} else if pq[1] == nil {
				return LessThan
			}

			comp := pq[0].Compare(pq[1])
			if comp != 0 {
				return comp
			}
		}

		return SameSame

	case IntPacket:
		return p.Compare(q.ToList())
	}

	panic(fmt.Errorf("impossible comparison: %+v <> %+v", p, q))
}

func zip[T any](as, bs []T) [][2]T {
	cnt := max(len(as), len(bs))
	cs := make([][2]T, cnt)

	for i := 0; i < cnt; i++ {
		c := [2]T{}
		if i < len(as) {
			c[0] = as[i]
		}
		if i < len(bs) {
			c[1] = bs[i]
		}
		cs[i] = c
	}

	return cs
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
