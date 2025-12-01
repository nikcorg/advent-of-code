package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2024/util"
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
	first := solveFirst(input)
	second := solveSecond(input)

	fmt.Fprint(out, "=====[ Day 11 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func parseInput(i string) []int64 {
	s := bufio.NewScanner(strings.NewReader(i))
	ns := []int64{}
	for s.Scan() {
		ns = append(ns, util.Fmap(util.MustAtoi64, strings.Fields(s.Text()))...)
	}
	return ns
}

func solveFirst(i string) int {
	stones := parseInput(i)

	for b := 0; b < 25; b++ {
		n := 0
		for n < len(stones) {
			switch {
			case stones[n] == 0:
				stones[n] = 1
			case digits(stones[n])%2 == 0:
				a, b := split(stones[n])
				stones = append(stones[0:n], append([]int64{a, b}, stones[n+1:]...)...)
				n++
			default:
				stones[n] *= 2024
			}
			n++
		}
	}
	return len(stones)
}

type blinkadj [2]int64

func solveSecond(i string) int64 {
	init := parseInput(i)
	stones := map[int64]int64{}

	for _, n := range init {
		stones[n] = 1
	}

	for x := 0; x < 75; x++ {
		blink := []blinkadj{}
		for v, n := range stones {
			switch {
			case v == 0:
				blink = append(blink, blinkadj{0, -n}, blinkadj{1, n})
			case digits(v)%2 == 0:
				a, b := split(v)
				blink = append(blink, blinkadj{v, -n}, blinkadj{a, n}, blinkadj{b, n})
			default:
				blink = append(blink, blinkadj{v, -n}, blinkadj{v * 2024, n})
			}
		}

		for _, adj := range blink {
			if adj[1] == 0 {
				continue
			}
			stones[adj[0]] += adj[1]
		}
	}

	var tot int64 = 0

	for v, n := range stones {
		if n == 0 {
			continue
		}
		tot += stones[v]
	}

	return tot
}

func digits(b int64) int {
	return len(fmt.Sprintf("%d", b))
}

func split(x int64) (int64, int64) {
	s := fmt.Sprintf("%d", x)
	mid := len(s) / 2
	a, b := s[0:mid], s[mid:]

	return util.MustAtoi64(a), util.MustAtoi64(b)
}
