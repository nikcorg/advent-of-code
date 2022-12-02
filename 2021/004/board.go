package main

import (
	"sort"
)

const width = 5

type board struct {
	Nums           []int
	Check          map[int]struct{}
	Hits           []point
	hitsWereSorted bool
}

func (b *board) At(p point) int {
	return b.Nums[p.X+p.Y*width]
}

func (b *board) SetAt(p point, v int) {
	b.Nums[p.X+p.Y*width] = v
}

func (b *board) HitAt(p point) bool {
	for _, hp := range b.Hits {
		if hp.Equals(p) {
			return true
		}
	}
	return false
}

func (b *board) Finalise() {
	b.Check = make(map[int]struct{})
	for _, n := range b.Nums {
		b.Check[n] = struct{}{}
	}
}

func (b *board) Mark(n int) {
	if _, ok := b.Check[n]; !ok {
		return
	}

	for p, bn := range b.Nums {
		if bn == n {
			y := p / width
			x := p - y*width
			pt := point{x, y}

			// FIXME: sort on insert
			b.hitsWereSorted = false
			b.Hits = append(b.Hits, pt)
		}
	}

}

func (b *board) Winner() bool {
	if !b.hitsWereSorted {
		sort.SliceStable(b.Hits, func(i, j int) bool {
			return b.Hits[i].Y*width+b.Hits[i].X < b.Hits[j].Y*width+b.Hits[j].X
		})
		b.hitsWereSorted = true
	}

	return b.columnsHit() || b.rowsHit()
}

func (b *board) columnsHit() bool {
	for c := 0; c < width; c++ {
		var (
			seq   = 0
			nextY = 0
		)

		for _, p := range b.Hits {
			if p.X == c && p.Y == nextY {
				seq++
				nextY++
			}

			if seq == width {
				return true
			}
		}
	}

	return false
}

func (b *board) rowsHit() bool {
	for r := 0; r < width; r++ {
		var (
			seq   = 0
			nextX = 0
		)

		for _, p := range b.Hits {
			if p.Y == r && p.X == nextX {
				seq++
				nextX++
			}

			if seq == width {
				return true
			}
		}
	}

	return false
}
