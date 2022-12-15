package main

import (
	"bufio"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"nikc.org/aoc2022/util"
	"nikc.org/aoc2022/util/set"
)

var (
	//go:embed input.txt
	input string
	track = flag.Bool("t", false, "measure execution times")
)

func timeTrack(start time.Time, name string) {
	if *track {
		elapsed := time.Since(start)
		fmt.Fprintf(os.Stderr, "%s took %s\n", name, elapsed)
	}
}

func main() {
	flag.Parse()

	if err := mainWithErr(os.Stdout, input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(out io.Writer, input string) error {
	fmt.Fprint(out, "=====[ Day 15 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", solveFirst(input, 2000000))
	return nil
}

func solveFirst(input string, exploreY int) int {
	defer timeTrack(time.Now(), "solveFirst")

	// impossible locations
	locs := set.New[util.Point]()
	c, mut := make(chan util.Point), sync.RWMutex{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		mut.Lock()
		defer mut.Unlock()

		for {
			select {
			case <-ctx.Done():
				return

			case p := <-c:
				locs.Add(p)
			}
		}
	}()

	wg := sync.WaitGroup{}
	explore := func(p util.Point, motion util.Point, ref util.Point, refDist int) {
		defer wg.Done()
		for d := p.ManhattanDistance(ref); d < refDist; p, d = p.Add(motion), p.ManhattanDistance(ref) {
			c <- p
		}
	}

	m := parseInput(bufio.NewScanner(strings.NewReader(input)))
	motions := []util.Point{util.NewPoint(-1, 0), util.NewPoint(1, 0)}

	for sensor, beacon := range m {
		dist := sensor.ManhattanDistance(beacon)
		start := util.NewPoint(sensor.X, exploreY)

		wg.Add(len(motions))
		for _, m := range motions {
			go explore(start, m, sensor, dist)
		}
	}

	wg.Wait()
	cancel()

	// remove any beacon and sensor locations from the set of impossible locations
	mut.Lock()
	for s, b := range m {
		locs.Remove(b)
		locs.Remove(s)
	}
	mut.Unlock()

	mut.RLock()
	defer mut.RUnlock()

	return locs.Size()
}

func parseInput(s *bufio.Scanner) map[util.Point]util.Point {
	defer timeTrack(time.Now(), "parseInput")

	xyPair := regexp.MustCompile(`at x=(-?\d+), y=(-?\d+)`)

	m := map[util.Point]util.Point{}
	for s.Scan() {
		matches := xyPair.FindAllStringSubmatch(s.Text(), 2)
		if matches == nil || len(matches) != 2 {
			panic(fmt.Errorf("unexpected input: %s", s.Text()))
		}

		sensor := util.NewPoint(util.MustAtoi(matches[0][1]), util.MustAtoi(matches[0][2]))
		beacon := util.NewPoint(util.MustAtoi(matches[1][1]), util.MustAtoi(matches[1][2]))

		m[sensor] = beacon
	}

	return m
}
