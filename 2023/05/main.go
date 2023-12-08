package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"

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
	seeds, maps := parseInput(input)
	first := solveFirst(seeds, maps)
	second := solveSecond(seeds, maps)

	fmt.Fprint(out, "=====[ Day 05 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(seeds []int, maps []sourceToDestMap) int {
	low := -1

	for _, s := range seeds {
		loc := s
		for _, m := range maps {
			loc = m.Ranges.MapSource(loc)
		}
		if low == -1 || loc < low {
			low = loc
		}
	}

	return low
}

func solveSecond(seeds []int, maps []sourceToDestMap) int {
	cap := make(chan struct{}, 300)
	wg := sync.WaitGroup{}
	low := atomic.Int32{}
	low.Store(-1)

	// Brute force! Slow and inefficient as anything you ever saw, but it got the job done.
	for _, sr := range chunkBy(seeds, 2) {
		// reserve slot
		cap <- struct{}{}
		wg.Add(1)
		go func(l, h int) {
			defer wg.Done()
			defer func() {
				// release slot
				<-cap
			}()

			for s := l; s < h; s++ {
				loc := s
				for _, m := range maps {
					loc = m.Ranges.MapSource(loc)
				}

				if v := low.Load(); v == -1 || loc < int(v) {
					low.Store(int32(loc))
				}
			}
		}(sr[0], sr[0]+sr[1])
	}
	wg.Wait()
	return int(low.Load())
}

type listOfRanges []mapRange

func (l listOfRanges) MapSource(n int) int {
	for _, r := range l {
		if next, ok := r.MapSource(n); ok {
			return next
		}
	}
	return n
}

type mapRange struct {
	Width, DestFrom, SourceFrom int
}

func (mr mapRange) MapSource(n int) (int, bool) {
	if n < mr.SourceFrom || mr.SourceFrom+mr.Width-1 < n {
		return n, false
	}

	diff := mr.DestFrom - mr.SourceFrom
	next := n + diff

	// fmt.Printf("%d<=%d<%d -> (%d-%d)=%d -> %d+%d=%d\n",
	// 	mr.SourceFrom, n, mr.SourceFrom+mr.Width, mr.DestFrom, mr.SourceFrom, diff, diff, n, next)

	return next, true
}

type sourceToDestMap struct {
	Name   string
	Ranges listOfRanges
}

var (
	mapLine = regexp.MustCompile(`^([-a-z]+) map:`)
)

func parseInput(input string) ([]int, []sourceToDestMap) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()

	seedLine := scanner.Text()
	if !strings.HasPrefix(seedLine, "seeds:") {
		panic(errors.New("seed line not found"))
	}

	seeds := util.Fmap(util.MustParseInt, strings.Split(seedLine[7:], " "))
	maps := []sourceToDestMap{}

	// skip empty line following seed line
	scanner.Scan()

	expectMapLine := true
	tempMap := sourceToDestMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			maps = append(maps, tempMap)
			tempMap = sourceToDestMap{}
			expectMapLine = true
			continue
		}

		if expectMapLine {
			tempMap.Name = mapLine.FindStringSubmatch(line)[1]
			expectMapLine = false
			continue
		}

		boundaries := util.Fmap(util.MustParseInt, strings.Split(line, " "))
		mr := mapRange{
			DestFrom:   boundaries[0],
			SourceFrom: boundaries[1],
			Width:      boundaries[2],
		}
		tempMap.Ranges = append(tempMap.Ranges, mr)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	maps = append(maps, tempMap)

	return seeds, maps
}

// source https://gist.github.com/mustafaturan/7a29e8251a7369645fb6c2965f8c2daf
func chunkBy[X any](items []X, chunkSize int) (chunks [][]X) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
