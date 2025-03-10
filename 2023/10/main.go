package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	c := parseInput(input)
	first := solveFirst(c)
	second := solveSecond(c)

	fmt.Fprint(out, "=====[ Day 10 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type chart struct {
	Width int
	Tiles string
}

func (c chart) At(x, y int) byte {
	return c.Tiles[y*c.Width+x]
}

func (c chart) XY(p int) (int, int) {
	if p < 0 || p >= len(c.Tiles) {
		panic(errors.New("out of bounds"))
	}

	y := p % c.Width
	x := p - y
	return x, y
}

func (c chart) Start() (int, int) {
	return c.XY(strings.Index(c.Tiles, "S"))
}

func parseInput(i string) chart {
	s := bufio.NewScanner(strings.NewReader(i))
	w, b := -1, strings.Builder{}

	for s.Scan() {
		b.Write(s.Bytes())
		if w == -1 {
			w = b.Len()
		}
	}

	return chart{Width: w, Tiles: b.String()}
}

var lookAround = map[byte][]int{
	'N': {0, -1},
	'E': {1, 0},
	'S': {0, 1},
	'W': {-1, 0},
}

func solveFirst(i chart) int {
	x, y := i.Start()
	visits := map[string]int{
		fmt.Sprintf("%d:%d", x, y): 0,
	}
	mut := sync.Mutex{}
	max := 0

	wg := sync.WaitGroup{}
	for d, t := range lookAround {
		wg.Add(1)
		go func(m byte, x, y int, prev byte) {
			defer wg.Done()

			d := 1
			n := i.At(x, y)

			for n != 'S' {
				switch {
				case m == 'N':
					// valid next tiles: 7 I F
				case m == 'E':
					// valid next tiles - J 7
				case m == 'S':
					// valid next tiles J I L
				case m == 'W':
					// valid next tiles L - F

				default:
					// end of the line
					return
				}

				// record visit
				p := fmt.Sprintf("%d:%d", x, y)

				mut.Lock()
				d0, ok := visits[p]
				if !ok || d0 > d {
					visits[p] = d
				}
				mut.Unlock()

				// if we're crossing paths and not finding a shorter path, we can stop
				if ok && d0 < d {
					return
				}

				// lookaround goes here
				d++
			}
		}(d, x+t[0], y+t[1], i.At(x, y))
	}

	wg.Wait()

	return max
}

func solveSecond(i chart) int {
	return 0
}
