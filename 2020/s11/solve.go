package s11

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

const bufSize = 1

type Solver struct {
	ctx context.Context
	out io.Writer
}

func New(ctx context.Context, out io.Writer) *Solver {
	return &Solver{ctx, out}
}

func (s *Solver) SolveFirst(inp io.Reader) error {
	return s.Solve(1, inp)
}

func (s *Solver) SolveSecond(inp io.Reader) error {
	return s.Solve(2, inp)
}

func (s *Solver) Solve(part int, inp io.Reader) error {
	lineInput := make(linestream.LineChan, bufSize)

	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	fm := convStream(linestream.SkipEmpty(lineInput))
	solve := getSolver(part)
	solution := solve(fm)

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

func convStream(inp linestream.ReadOnlyLineChan) *floormap {
	fm := &floormap{}

	for line := range inp {
		if fm.width == 0 {
			fm.width = len(line.Content())
		}

		tiles := make([]tileKind, len(line.Content()))

		for n, tile := range strings.Split(line.Content(), "") {
			switch tile {
			case "L":
				tiles[n] = emptySeat
			case "#":
				tiles[n] = occupiedSeat
			default:
				tiles[n] = floorTile
			}
		}

		fm.tiles = append(fm.tiles, tiles...)
	}

	return fm
}

type solver func(*floormap) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("invalid part %d", part))
}

func solveFirst(fm *floormap) int {
	mapTile := func(prevMap *floormap, tile tileKind, x, y int) tileKind {
		switch tile {
		case emptySeat:
			if prevMap.OccupiedAdjacent(x, y) == 0 {
				return occupiedSeat
			}

		case occupiedSeat:
			if prevMap.OccupiedAdjacent(x, y) >= 4 {
				return emptySeat
			}
		}

		return tile
	}

	nextMap := stepUntilStable(fm, mapTile)
	occupied := 0

	for _, t := range nextMap.tiles {
		if t == occupiedSeat {
			occupied++
		}
	}

	return occupied
}

func mapTileSecond(prevMap *floormap, tile tileKind, x, y int) tileKind {
	switch tile {
	case occupiedSeat:
		if prevMap.OccupiedVisibleFrom(x, y) >= 5 {
			return emptySeat
		}
	case emptySeat:
		if prevMap.OccupiedVisibleFrom(x, y) == 0 {
			return occupiedSeat
		}
	}

	return tile
}

func solveSecond(fm *floormap) int {
	nextMap := stepUntilStable(fm, mapTileSecond)
	occupied := 0

	for _, t := range nextMap.tiles {
		if t == occupiedSeat {
			occupied++
		}
	}

	return occupied
}

type tileMapper func(prevMap *floormap, tile tileKind, x, y int) tileKind

func step(fm *floormap, tm tileMapper) (*floormap, bool) {
	prevMap := fm
	nextMap := &floormap{}
	nextMap.Init(prevMap.Width(), prevMap.Height())
	didChange := false

	wg := sync.WaitGroup{}

	for x := 0; x < prevMap.Width(); x++ {
		x := x
		wg.Add(1)

		// Spawn goroutine for each X-column
		go func() {
			defer wg.Done()

			for y := 0; y < prevMap.Height(); y++ {
				tile, err := prevMap.TileAt(x, y)

				if err != nil {
					if err == errInvalidCoordinate {
						continue
					}
					panic(err)
				}

				if nextTile := tm(prevMap, tile, x, y); nextTile != tile {
					nextMap.SetTileAt(x, y, nextTile)
					didChange = true
					continue
				}

				nextMap.SetTileAt(x, y, tile)
			}
		}()
	}

	wg.Wait()

	return nextMap, didChange
}

func stepUntilStable(fm *floormap, tm tileMapper) *floormap {
	nextMap, didChange := step(fm, tm)
	for didChange {
		nextMap, didChange = step(nextMap, tm)
	}
	return nextMap
}
