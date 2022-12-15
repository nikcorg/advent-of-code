package main

import "nikc.org/aoc2022/util"

type Triangle struct {
	p1, p2, p3 util.Point
}

// Triangle hit test sourced from https://www.gamedev.net/forums/topic.asp?topic_id=295943
func (t *Triangle) Contains(p util.Point) bool {
	b1 := sign(p, t.p1, t.p2) < 0.0
	b2 := sign(p, t.p2, t.p3) < 0.0
	b3 := sign(p, t.p3, t.p1) < 0.0
	return ((b1 == b2) && (b2 == b3))
}

func sign(p1, p2, p3 util.Point) float64 {
	return float64(p1.X-p3.X)*float64(p2.Y-p3.Y) - float64(p2.X-p3.X)*float64(p1.Y-p3.Y)
}
