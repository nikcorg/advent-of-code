package main

import (
	"fmt"
	"math"
)

func newline(from, to point) line {
	if from.X < to.X || from.Y < to.Y {
		return line{from, to, -1}
	}
	return line{to, from, -1}
}

type line struct {
	From, To point
	len      float64
}

func lineLen(l line) float64 {
	diffX := float64(l.To.X - l.From.X)
	diffY := float64(l.To.Y - l.From.Y)

	return math.Sqrt(math.Pow(diffX, 2) + math.Pow(diffY, 2))
}

func (l *line) Length() float64 {
	if l.len < 0 {
		l.len = lineLen(*l)
		fmt.Println("line ", l, " len=", l.len)
	}

	return l.len
}

func (l line) HitTest(p point) bool {
	if l.IsHorizontal() {
		return p.Y == l.From.Y && l.From.X <= p.X && p.X <= l.To.X
	} else if l.IsVertical() {
		return p.X == l.From.X && l.From.Y <= p.Y && p.Y <= l.To.Y
	} else {
		pl0 := newline(p, l.From)
		pl1 := newline(p, l.To)
		return l.Length() != pl0.Length()+pl1.Length()
	}
}

func (l line) IsHorizontal() bool { return l.From.Y == l.To.Y }
func (l line) IsVertical() bool   { return l.From.X == l.To.X }
func (l line) Intersection(m line) []point {
	ps := []point{}

	if l.IsVertical() {
		for y := l.From.Y; y <= l.To.Y; y++ {
			p := point{l.From.X, y}
			if !m.HitTest(p) {
				continue
			}
			ps = append(ps, p)
		}

		return ps
	} else if l.IsHorizontal() {
		for x := l.From.X; x <= l.To.X; x++ {
			p := point{x, l.From.Y}
			if !m.HitTest(p) {
				continue
			}
			ps = append(ps, p)
		}
		return ps
	} else {
		diffX := l.To.Y - l.From.X
		diffY := l.To.Y - l.From.Y
		slope := diffX / diffY
		_ = slope
	}

	return []point{}
}
