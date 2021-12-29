package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	//go:embed input.txt
	input string
)

func main() {
	if err := mainWithErr(input); err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
	}
}

func mainWithErr(input string) error {
	parsed, err := parseInput(input)
	if err != nil {
		return err
	}

	solveStart := time.Now()
	if result, err := solveFirst(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 1: %d (duration: %v)\n", result, time.Since(solveStart)))
	}

	solveStart = time.Now()
	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d (duration: %v)\n", result, time.Since(solveStart)))
	}

	return nil
}

type rule struct {
	Insert  rune
	SplitBy *regexp.Regexp
}

type bipolymer [2]byte

func (bip bipolymer) String() string {
	return string([]byte{bip[0], bip[1]})
}

type puzzleInput struct {
	Rules map[bipolymer]rule
	Base  []byte
}

var (
	errUnderflow = errors.New("underflow error")
	reSplitRule  = regexp.MustCompile(`^(\w+) -> (\w+)$`)
)

func parseInput(raw string) (puzzleInput, error) {
	rows := strings.Split(raw, "\n")

	pi := puzzleInput{}
	pi.Rules = map[bipolymer]rule{}
	pi.Base = []byte(rows[0])

	for _, row := range rows[2:] {
		matches := reSplitRule.FindStringSubmatch(row)
		pi.Rules[bipolymer{matches[1][0], matches[1][1]}] = rule{
			Insert: rune(matches[2][0]),
		}
	}

	return pi, nil
}

func splitPolymer(polymer []byte) [][]byte {
	polymers := make([][]byte, len(polymer)-1)
	for i := range polymers {
		polymers[i] = polymer[i : i+2]
	}
	return polymers
}

func solveParallel4(pctx context.Context, polymer []byte, rules map[bipolymer]rule, rounds int) (uint64, error) {
	var (
		occs     = map[byte]uint64{}
		bips     = map[bipolymer]uint64{}
		polymers = splitPolymer(polymer)
	)

	for _, e := range polymer {
		occs[e]++
	}

	for _, p := range polymers {
		bips[bipolymer{p[0], p[1]}]++
	}

	for r := 0; r < rounds; r++ {
		bipDeltas := map[bipolymer]int64{}
		for bip := range bips {
			bipOccs := bips[bip]
			rule := rules[bip]
			newElem := byte(rule.Insert)
			newBipLeft := bipolymer{bip[0], newElem}
			newBipRight := bipolymer{newElem, bip[1]}

			occs[newElem] += bipOccs

			bipDeltas[newBipLeft] += int64(bipOccs)
			bipDeltas[newBipRight] += int64(bipOccs)
			bipDeltas[bip] -= int64(bipOccs)
		}

		for bip, delta := range bipDeltas {
			if nextOccs := int64(bips[bip]) + delta; nextOccs < 0 {
				return 0, errUnderflow
			} else {
				bips[bip] = uint64(nextOccs)
			}
		}
	}

	var (
		maxOcc = uint64(0)
		minOcc = ^maxOcc
	)

	for _, es := range occs {
		if es < minOcc {
			minOcc = es
		}
		if es > maxOcc {
			maxOcc = es
		}
	}

	return maxOcc - minOcc, nil
}

func solveFirst(input puzzleInput) (uint64, error) {
	p, err := solveParallel4(context.Background(), input.Base, input.Rules, 10)
	if err != nil {
		return 0, nil
	}
	return p, nil
}

func solveSecond(input puzzleInput) (uint64, error) {
	p, err := solveParallel4(context.Background(), input.Base, input.Rules, 40)
	if err != nil {
		return 0, nil
	}
	return p, nil
}
