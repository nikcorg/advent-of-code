package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	cards := parseInput(input)
	first := solveFirst(cards)
	second := solveSecond(cards)

	fmt.Fprint(out, "=====[ Day 04 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

type card struct {
	Copies  int
	ID      int
	Winners map[string]struct{}
	Draw    []string
}

func parseInput(input string) []card {
	r := regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+) \|\s+([\d\s]+)`)
	ws := regexp.MustCompile(`\s+`)
	cs := []card{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		m := r.FindStringSubmatch(scanner.Text())
		c := card{ID: parseInt(m[1]), Copies: 1}
		c.Winners = fold(func(acc map[string]struct{}, x string) map[string]struct{} {
			acc[x] = struct{}{}
			return acc
		}, map[string]struct{}{}, ws.Split(m[2], -1))
		c.Draw = ws.Split(m[3], -1)

		cs = append(cs, c)
	}

	return cs
}

func solveFirst(cs []card) int {
	tot := 0

	for _, c := range cs {
		n := 0

		for _, x := range c.Draw {
			if _, ok := c.Winners[x]; ok {
				n++
			}
		}

		if n == 0 {
			continue
		}

		points := int(math.Pow(2, float64(n-1)))
		tot += points
	}

	return tot
}

func solveSecond(cs []card) int {
	tot := 0

	for i, c := range cs {
		tot += c.Copies
		n := 0

		for _, x := range c.Draw {
			if _, ok := c.Winners[x]; ok {
				n++
			}
		}

		if n == 0 {
			continue
		}

		for n > 0 {
			cs[i+n].Copies += c.Copies
			n--
		}
	}

	return tot
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func fold[X any, TAcc any](f func(TAcc, X) TAcc, init TAcc, xs []X) TAcc {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}
