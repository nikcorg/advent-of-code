package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"

	util "nikc.org/aoc2023/util"
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

	return solveFirst([][]int{{util.MustParseInt(time), util.MustParseInt(record)}})
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
			times = util.Fmap(util.MustParseInt, ws.Split(line[2], -1))
		case "Distance":
			records = util.Fmap(util.MustParseInt, ws.Split(line[2], -1))
		}
	}

	return util.Zip(times, records)
}
