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
	fmt.Fprintf(out, "first: %d\n", solveFirst(input, 2_000_000))
	fmt.Fprintf(out, "second: %d\n", solveSecond(input, 0, 4_000_000))
	return nil
}

func solveFirst(input string, exploreY int) int {
	defer timeTrack(time.Now(), "solveFirst")

	m := parseInput(bufio.NewScanner(strings.NewReader(input)))

	// impossible locations
	locs := set.New[util.Point]()

	for s, b := range m {
		d := s.ManhattanDistance(b)

		// Check if this beacon's range intersects with exploreY
		if s.Y-d > exploreY || s.Y+d < exploreY {
			continue
		}

		l, r := util.NewPoint(s.X, exploreY), util.NewPoint(s.X, exploreY)
		for l.ManhattanDistance(s) <= d {
			locs.Add(l)
			locs.Add(r)
			l.X -= 1
			r.X += 1
		}
	}

	// remove any beacon and sensor locations from the set of impossible locations
	for s, b := range m {
		locs.Remove(b)
		locs.Remove(s)
	}

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

// Lucky 13
const maxSeekers = 13

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

	lostBeacon := make(chan util.Point)
	lostBeacons := make(chan util.Point, maxSeekers)

	go func() {
		for candidate := range lostBeacons {
			trulyLost := true

			// eliminate false positives
			for s, b := range m {
				d := s.ManhattanDistance(b)

				if s.ManhattanDistance(candidate) <= d {
					trulyLost = false
					break
				}
			}

			if trulyLost {
				lostBeacon <- candidate
				return
			}
		}
	}()

	explore := func(ctx context.Context, job *seekerJob) *util.Point {
		p := job.from

		for !p.Equals(job.to) {
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
					return &p
				}
			}

			// Translate location
			p = p.Add(job.motion)
		}

		return nil
	}

	for n := 0; n < util.Min(maxSeekers-1, seekers.Size()); n++ {
		go func() {
			for {
				job := seekers.Pop()
				if job == nil {
					return
				}
				p := explore(ctx, job)

				select {
				case <-ctx.Done():
					return
				default:
					if p != nil {
						lostBeacons <- *p
					}
				}
			}
		}()
	}

	lb := <-lostBeacon

	if *dump {
		dumpTriangles(m, triangles, lb, minXY, maxXY)
	}

	return lb.X*4_000_000 + lb.Y
}

func dumpTriangles(m map[util.Point]util.Point, ts []Triangle, b util.Point, minXY, maxXY int) {
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
	fmt.Fprintln(f, `ctx.fillStyle="white"; ctx.fillRect(0, 0, c.width, c.height);`)
	fmt.Fprintln(f)

	for _, tg := range ts {
		fmt.Fprint(f, "ctx.beginPath();")
		fmt.Fprintln(f, "ctx.fillStyle=`rgba(${Math.floor(Math.random()*255)},${Math.floor(Math.random()*255)},${Math.floor(Math.random()*255)},0.2)`")
		fmt.Fprintf(f, "ctx.moveTo(%d, %d);", int(float64(tg.p1.X)*scale), int(float64(tg.p1.Y)*scale))
		fmt.Fprintf(f, "ctx.lineTo(%d, %d);", int(float64(tg.p2.X)*scale), int(float64(tg.p2.Y)*scale))
		fmt.Fprintf(f, "ctx.lineTo(%d, %d);", int(float64(tg.p3.X)*scale), int(float64(tg.p3.Y)*scale))
		fmt.Fprintf(f, "ctx.lineTo(%d, %d);", int(float64(tg.p1.X)*scale), int(float64(tg.p1.Y)*scale))
		fmt.Fprintln(f, "ctx.stroke();")
		fmt.Fprintln(f, "ctx.fill();")
	}

	for s, b := range m {
		if s.X >= minXY && s.X <= maxXY && s.Y >= minXY && s.Y <= maxXY {
			fmt.Fprint(f, "ctx.beginPath();")
			fmt.Fprint(f, `ctx.fillStyle="black";`)
			fmt.Fprintf(f, `ctx.arc(%d, %d, %d, 0, 2*Math.PI);`, int(float64(s.X)*scale), int(float64(s.Y)*scale), 20)
			fmt.Fprintln(f, "ctx.fill();")
		}
		if b.X >= minXY && b.X <= maxXY && b.Y >= minXY && b.Y <= maxXY {
			fmt.Fprint(f, "ctx.beginPath();")
			fmt.Fprint(f, `ctx.fillStyle="peru";`)
			fmt.Fprintf(f, `ctx.arc(%d, %d, %d, 0, 2*Math.PI);`, int(float64(b.X)*scale), int(float64(b.Y)*scale), 20)
			fmt.Fprintln(f, "ctx.stroke();")
			fmt.Fprintln(f, "ctx.fill();")
		}
	}

	fmt.Fprint(f, "ctx.beginPath();")
	fmt.Fprint(f, `ctx.fillStyle="hotpink";`)
	fmt.Fprintf(f, `ctx.arc(%d, %d, %d, 0, 2*Math.PI);`, int(float64(b.X)*scale), int(float64(b.Y)*scale), 20)
	fmt.Fprintln(f, "ctx.stroke();")
	fmt.Fprintln(f, "ctx.fill();")
	fmt.Fprintln(f)
	fmt.Fprintln(f, `const i = document.createElement("img");`)
	fmt.Fprintln(f, `i.src = c.toDataURL();`)
	fmt.Fprintln(f, `document.body.appendChild(i);`)
	fmt.Fprintln(f)
}
