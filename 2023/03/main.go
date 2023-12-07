package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	"nikc.org/aoc2023/util"
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
	m := parseInput(input)
	first := solveFirst(m)
	second := solveSecond(m)

	fmt.Fprint(out, "=====[ Day 03 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type point struct {
	X, Y int
}

type partsMap struct {
	Width int
	Map   string
}

func (m partsMap) At(x, y int) byte {
	return m.Map[y*m.Width+x]
}

func (m partsMap) XY(n int) (int, int) {
	return n % m.Width, n / m.Width
}

func parseInput(s string) partsMap {
	m := partsMap{}
	lines := strings.Split(s, "\n")
	m.Width = len(lines[0])
	for _, l := range lines {
		m.Map += l
	}
	return m
}

var lookAround = []point{{0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}}

func solveFirst(m partsMap) int {
	tot := 0

	mut := sync.Mutex{}
	wg := sync.WaitGroup{}

	partNos := map[string]int{}

	// find all symbols that aren't digits or '.'
	syms := regexp.MustCompile(`[^.\d]`).FindAllStringIndex(m.Map, -1)
	for _, r := range syms {
		p := r[0]
		x, y := m.XY(p)

		// for each symbol, scan all the adjacent X,Y coordinates
		for _, p := range lookAround {
			wg.Add(1)
			go func(sx, y int) {
				defer wg.Done()
				// not a digit? not a part number, so nothing to do
				if !isDigit(m.At(sx, y)) {
					return
				}
				// find the start coordinate of the number...
				i, j := sx, sx
				for {
					if i-1 >= 0 && isDigit(m.At(i-1, y)) {
						i -= 1
						continue
					}
					break
				}
				// ...and the end coordinate of the number
				for {
					if j+1 < m.Width && isDigit(m.At(j+1, y)) {
						j += 1
						continue
					}
					break
				}
				s := m.Map[y*m.Width+i : y*m.Width+j+1]
				n := util.ParseInt(s)
				mut.Lock()
				defer mut.Unlock()
				// store the number using the location as key, to avoid including
				// the same part number multiple times
				partNos[fmt.Sprintf("%d:%d:%d", y, i, j)] = n
			}(x+p.X, y+p.Y)
		}
	}

	wg.Wait()

	for _, n := range partNos {
		tot += n
	}

	return tot
}

func solveSecond(m partsMap) int {
	tot := 0

	// find all '*' symbols
	for _, r := range regexp.MustCompile(`[*]`).FindAllStringIndex(m.Map, -1) {
		x, y := m.XY(r[0])

		partNos := map[string]int{}

		mut := sync.Mutex{}
		wg := sync.WaitGroup{}

		// scan all the adjacent X,Y coordinates
		for _, p := range lookAround {
			wg.Add(1)
			go func(sx, y int) {
				defer wg.Done()
				// not a digit? not a ratio, so nothing to do
				if !isDigit(m.At(sx, y)) {
					return
				}
				i, j := sx, sx
				// find the start coordinate of the number...
				for {
					if i-1 >= 0 && isDigit(m.At(i-1, y)) {
						i -= 1
						continue
					}
					break
				}
				// ...and the end coordinate of the number
				for {
					if j+1 < m.Width && isDigit(m.At(j+1, y)) {
						j += 1
						continue
					}
					break
				}
				s := m.Map[y*m.Width+i : y*m.Width+j+1]
				n := util.ParseInt(s)
				mut.Lock()
				defer mut.Unlock()
				// store the number using the location as key, to avoid including
				// the same ratio multiple times
				partNos[fmt.Sprintf("%d:%d:%d", y, i, j)] = n
			}(x+p.X, y+p.Y)
		}

		wg.Wait()

		// if there aren't exactly two adjacent numbers, carry on
		if len(partNos) != 2 {
			continue
		}

		r := 1
		// ...otherwise, multiple the gear ratio with the found numbers and add the product
		// to the final total
		for _, n := range partNos {
			r *= n
		}
		tot += r
	}

	return tot
}

func isDigit(c byte) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}
