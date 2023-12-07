package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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
	games := parseInput(input)
	first := solveFirst(games)
	second := solveSecond(games)

	fmt.Fprint(out, "=====[ Day 06 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(games [][]int) int {
	wg := sync.WaitGroup{}
	mut := sync.Mutex{}
	tot := 1

	for _, g := range games {
		wg.Add(1)
		go func(time, record int) {
			defer wg.Done()
			n := 0
			for i := 0; i < time; i++ {
				mmps := i
				travel := mmps * (time - i)
				if travel > record {
					n++
				}
			}

			mut.Lock()
			defer mut.Unlock()
			tot *= n
		}(g[0], g[1])
	}

	wg.Wait()

	return tot
}

func solveSecond(games [][]int) int {
	var (
		time, record string
	)

	for _, g := range games {
		time += fmt.Sprintf("%d", g[0])
		record += fmt.Sprintf("%d", g[1])
	}

	return solveFirst([][]int{{parseInt(time), parseInt(record)}})
}

func parseInput(input string) [][]int {
	splitter := regexp.MustCompile(`(Time|Distance):\s+(.+)$`)
	ws := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(strings.NewReader(input))

	var times, records []int

	for scanner.Scan() {
		line := splitter.FindStringSubmatch(scanner.Text())
		switch line[1] {
		case "Time":
			times = fmap(parseInt, ws.Split(line[2], -1))
		case "Distance":
			records = fmap(parseInt, ws.Split(line[2], -1))
		}
	}

	return zip(times, records)
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func fmap[X, Y any](f func(X) Y, xs []X) []Y {
	ys := make([]Y, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func zip[T any](xs, ys []T) [][]T {
	l := min(len(xs), len(ys))
	zs := make([][]T, l)

	for l > 0 {
		zs[l-1] = []T{xs[l-1], ys[l-1]}
		l--
	}

	return zs
}
