package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"nikc.org/aoc2022/util"
	"nikc.org/aoc2022/util/set"
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
	var (
		first, second int
		err           error
	)

	fmt.Fprint(out, "=====[ Day 09 ]=====\n")

	if first, err = solveFirst(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %d\n", first)

	if second, err = solveSecond(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

var (
	trRight = util.NewPoint(1, 0)
	trLeft  = util.NewPoint(-1, 0)
	trUp    = util.NewPoint(0, -1)
	trDown  = util.NewPoint(0, 1)
)

func solveFirst(input string) (int, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	head := util.NewPoint(0, 0)
	tail := util.NewPoint(0, 0)
	visited := set.New(tail)

	for s.Scan() {
		args := strings.Split(s.Text(), " ")
		direction := args[0]
		amount := util.MustAtoi(args[1])

		for n := 0; n < amount; n++ {
			switch direction {
			case "R":
				head = head.Add(trRight)
			case "L":
				head = head.Add(trLeft)
			case "U":
				head = head.Add(trUp)
			case "D":
				head = head.Add(trDown)
			default:
				return 0, fmt.Errorf("unknown cmd: %s", s.Text())
			}

			distX, distY := head.DistanceX(tail), head.DistanceY(tail)

			if abs(distX) > 1 || abs(distY) > 1 {
				switch {
				case distX != 0 && distY != 0:
					tail = tail.Add(util.NewPoint(distX/abs(distX), distY/abs(distY)))

				case distX != 0:
					tail = tail.Add(util.NewPoint(distX/abs(distX), 0))

				case distY != 0:
					tail = tail.Add(util.NewPoint(0, distY/abs(distY)))
				}

			}

			visited.Add(tail)

		}
	}

	return visited.Size(), nil
}

func solveSecond(input string) (int, error) {
	s := bufio.NewScanner(strings.NewReader(input))

	var (
		knots = []util.Point{
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
			util.NewPoint(0, 0),
		}
		numKnots = len(knots)
		headKnot = 0
		tailKnot = numKnots - 1
		visited  = set.New(knots[tailKnot])
	)

	for s.Scan() {
		args := strings.Split(s.Text(), " ")
		direction := args[0]
		amount := util.MustAtoi(args[1])

		for n := 0; n < amount; n++ {
			switch direction {
			case "R":
				knots[headKnot] = knots[headKnot].Add(trRight)
			case "L":
				knots[headKnot] = knots[headKnot].Add(trLeft)
			case "U":
				knots[headKnot] = knots[headKnot].Add(trUp)
			case "D":
				knots[headKnot] = knots[headKnot].Add(trDown)
			default:
				return 0, fmt.Errorf("unknown cmd: %s", s.Text())
			}

			for i := 1; i < numKnots; i++ {
				head, tail := knots[i-1], knots[i]
				distX, distY := head.DistanceX(tail), head.DistanceY(tail)

				if abs(distX) > 1 || abs(distY) > 1 {
					switch {
					case distX != 0 && distY != 0:
						tail = tail.Add(util.NewPoint(distX/abs(distX), distY/abs(distY)))

					case distX != 0:
						tail = tail.Add(util.NewPoint(distX/abs(distX), 0))

					case distY != 0:
						tail = tail.Add(util.NewPoint(0, distY/abs(distY)))
					}
				}

				knots[i] = tail
			}

			visited.Add(knots[tailKnot])
		}
	}

	return visited.Size(), nil
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
