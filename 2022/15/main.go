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
	"nikc.org/aoc2022/util/stack"
)

var (
	//go:embed input.txt
	input string
	track = flag.Bool("t", false, "measure execution times")
	dump  = flag.Bool("d", false, "dump triangles")
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
	fmt.Fprintf(out, "second: %d\n", solveSecond(input, 0, 4_000_000))
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

var (
	downAndLeft  = util.NewPoint(-1, 1)
	downAndRight = util.NewPoint(1, 1)
	upAndRight   = util.NewPoint(1, -1)
	upAndLeft    = util.NewPoint(-1, -1)
)

// While this solution works, it doesn't work every time. If let run, it finds more than one
// location; 5 with the test data, and 3 with the actual data. I'm not sure if the problme lies
// with my solution or the data, most likely the former. I'm miffed about the non-deterministic
// nature of my solution, but guessing got me the second star. And it's at least reasonably quick,
// particularly compared to the first part. I'm slightly tempted to try and optimise it.
func solveSecond(input string, minXY, maxXY int) int {
	defer timeTrack(time.Now(), "solveSecond")

	m := parseInput(bufio.NewScanner(strings.NewReader(input)))

	type seekerJob struct {
		from   util.Point
		to     util.Point
		motion util.Point
	}

	triangles := []Triangle{}
	seekers := stack.New[*seekerJob]()

	for s, b := range m {
		d := s.ManhattanDistance(b)

		triangles = append(triangles,
			// Top left
			Triangle{s, s.Add(util.NewPoint(-d, 0)), s.Add(util.NewPoint(0, -d))},
			// Bottom left
			Triangle{s, s.Add(util.NewPoint(-d, 0)), s.Add(util.NewPoint(0, d))},
			// // Top right
			Triangle{s, s.Add(util.NewPoint(d, 0)), s.Add(util.NewPoint(0, -d))},
			// // Bottom right
			Triangle{s, s.Add(util.NewPoint(d, 0)), s.Add(util.NewPoint(0, d))},
		)

		r := d + 1

		seekers.Push(
			// from above to the left
			&seekerJob{s.Add(util.NewPoint(0, r*-1)), s.Add(util.NewPoint(r*-1, 0)), downAndLeft},
			// from the left to below
			&seekerJob{s.Add(util.NewPoint(r*-1, 0)), s.Add(util.NewPoint(0, r)), downAndRight},
			// from below to the right
			&seekerJob{s.Add(util.NewPoint(0, r)), s.Add(util.NewPoint(r, 0)), upAndRight},
			// from the right to above
			&seekerJob{s.Add(util.NewPoint(r, 0)), s.Add(util.NewPoint(0, r*-1)), upAndLeft},
		)

	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		lostBeacon util.Point
		c          = make(chan util.Point)
	)

	explore := func(ctx context.Context, job *seekerJob) {
		p := job.from

		for !p.Equals(job.to) {
			select {
			case <-ctx.Done():
				return
			default:
			}

			// Hit test triangles when the point is within the search area
			if p.X >= minXY && p.Y >= minXY && p.X <= maxXY && p.Y <= maxXY {
				hit := false
				for _, t := range triangles {
					if t.Contains(p) {
						hit = true
						break
					}
				}
				if !hit {
					c <- p
					return
				}
			}

			// Translate location
			p = p.Add(job.motion)
		}
	}

	for i := 0; i < util.Min(23, seekers.Size()); i++ {
		go func() {
			for {
				job := seekers.Pop()
				if job == nil {
					return
				}
				explore(ctx, job)
			}
		}()
	}

	lostBeacon = <-c

	if *dump {
		dumpTriangles(triangles, lostBeacon, minXY, maxXY)
	}

	return lostBeacon.X*4_000_000 + lostBeacon.Y
}

func dumpTriangles(ts []Triangle, b util.Point, minXY, maxXY int) {
	defer timeTrack(time.Now(), "dump triangles")

	scale := 0.001

	f, err := os.Create("triangles.js")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, `const c = document.createElement("canvas");`)
	fmt.Fprintln(f)
	fmt.Fprintf(f, `c.setAttribute("width", %d); c.setAttribute("height", %d);`, int(float64(maxXY)*scale), int(float64(maxXY)*scale))
	fmt.Fprintln(f)
	fmt.Fprintln(f, `document.body.appendChild(c);`)
	fmt.Fprintf(f, `const ctx = c.getContext("2d");`)
	fmt.Fprintln(f)

	for _, tg := range ts {
		fmt.Fprintln(f, "ctx.fillStyle=`rgba(${Math.floor(Math.random()*255)},${Math.floor(Math.random()*255)},${Math.floor(Math.random()*255)},0.2)`")
		fmt.Fprint(f, "ctx.beginPath();")
		fmt.Fprintf(f, "ctx.moveTo(%d, %d);", int(float64(tg.p1.X)*scale), int(float64(tg.p1.Y)*scale))
		fmt.Fprintf(f, "ctx.lineTo(%d, %d);", int(float64(tg.p2.X)*scale), int(float64(tg.p2.Y)*scale))
		fmt.Fprintf(f, "ctx.lineTo(%d, %d);", int(float64(tg.p3.X)*scale), int(float64(tg.p3.Y)*scale))
		fmt.Fprintf(f, "ctx.closePath();")
		fmt.Fprintln(f, "ctx.fill();")
	}

	fmt.Fprintln(f, `ctx.fillStyle="black";`)
	fmt.Fprintf(f, `ctx.arc(%d, %d, %d, 0, 2*Math.PI);`, int(float64(b.X)*scale), int(float64(b.Y)*scale), 20)
	fmt.Fprintln(f)
}
