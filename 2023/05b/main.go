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
	seeds1, seeds2, maps := parseInput(input)
	first := solveFirst(seeds1, maps)
	second := solveSecond(seeds2, maps)

	fmt.Fprint(out, "=====[ Day 05 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

func solveFirst(seeds []intRange, tfrs []transformer) int {
	return 0
}

func solveSecond(seeds []intRange, tfrs []transformer) int {
	return 0
}

func newIntRange(from, width int) intRange {
	return intRange{from, from + width - 1}
}

type intRange struct {
	From, To int
}

func (r intRange) Overlaps(s intRange) bool {
	return r.Equals(s) || s.Contains(r) || (r.From <= s.From && s.From <= r.To) || (r.From <= s.To && s.To <= r.To)
}

func (r intRange) Contains(s intRange) bool {
	return r.From <= s.From && s.To <= r.To
}

func (r intRange) GetOverlap(s intRange) (intRange, bool) {
	if !r.Overlaps(s) {
		return intRange{}, false
	}

	from := max(s.From, r.From)
	to := min(s.To, r.To)

	return intRange{from, to}, true
}

func (r intRange) Equals(s intRange) bool {
	return r.From == s.From && r.To == s.To
}

func (r intRange) GetDifference(s intRange) ([]intRange, bool) {
	if r.Equals(s) || !r.Overlaps(s) || r.Contains(s) {
		return []intRange{}, false
	}

	diffs := []intRange{}

	// is there a lower difference?
	if s.From < r.From {
		diffs = append(diffs, intRange{s.From, r.From - 1})
	}

	// is there an upper difference?
	if s.To > r.To {
		diffs = append(diffs, intRange{r.To + 1, s.To})
	}

	return diffs, true
}

func newTransform(r intRange, amt int) transform {
	return transform{r, amt}
}

type transform struct {
	Range intRange
	Amt   int
}

type transformer struct {
	Name       string
	Transforms []transform
}

func (tf transformer) AddTransform(t transform) {
	tf.Transforms = append(tf.Transforms, t)
}

var (
	mapLine = regexp.MustCompile(`^([-a-z]+) map:`)
)

func parseInput(input string) ([]intRange, []intRange, []transformer) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()

	seedLine := scanner.Text()
	if !strings.HasPrefix(seedLine, "seeds:") {
		panic(errors.New("seed line not found"))
	}
	seeds := strings.Split(seedLine[7:], " ")

	// in part 1 the seeds are just numbers, i.e. ranges with length=1
	seeds1 := util.Fmap(func(s string) intRange {
		return newIntRange(util.MustParseInt(s), 1)
	}, seeds)

	// in part 2 everything is a range
	seeds2 := util.Fmap(func(s []string) intRange {
		return newIntRange(util.MustParseInt(s[0]), util.MustParseInt(s[1]))
	}, util.ChunkBy(seeds, 2))

	tfrs := []transformer{}

	// skip empty line following seed line
	scanner.Scan()

	expectMapLine := true
	tfr := transformer{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			tfrs = append(tfrs, tfr)
			tfr = transformer{}
			expectMapLine = true
			continue
		}

		if expectMapLine {
			tfr.Name = mapLine.FindStringSubmatch(line)[1]
			expectMapLine = false
			continue
		}

		boundaries := util.Fmap(util.MustParseInt, strings.Split(line, " "))
		tfr.AddTransform(
			newTransform(
				newIntRange(boundaries[1], boundaries[2]),
				boundaries[1]-boundaries[0]))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	tfrs = append(tfrs, tfr)

	return seeds1, seeds2, tfrs
}
