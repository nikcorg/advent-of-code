package s4

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/nikcorg/aoc2020/utils/linestream"
)

type Solver struct {
	Ctx context.Context
	Inp string
}

func New(ctx context.Context, inputFile string) *Solver {
	return &Solver{ctx, inputFile}
}

func (s *Solver) Solve(part int) error {
	ctx, cancel := context.WithCancel(s.Ctx)
	defer func() { cancel() }()

	inputFile, err := os.Open(s.Inp)
	if err != nil {
		return err
	}
	defer func() { inputFile.Close() }()

	lineInput := make(linestream.LineChan, 0)
	linestream.New(ctx, bufio.NewReader(inputFile), lineInput)

	solution := <-solveStream(getSolver(part), convStream(lineInput))

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(*passport, int) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveSecond
	}

	panic(fmt.Errorf("invalid part: %d", part))
}

func solveSecond(p *passport, validPassports int) int {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range mandatory {
		fv, isset := p.fields[f]
		if !isset {
			return validPassports
		}

		switch f {
		case "byr":
			if !validNumber(1920, 2002, fv) {
				return validPassports
			}

		case "iyr":
			if !validNumber(2010, 2020, fv) {
				return validPassports
			}

		case "eyr":
			if !validNumber(2020, 2030, fv) {
				return validPassports
			}

		case "hgt":
			if !validHeight(fv) {
				return validPassports
			}

		case "hcl":
			if !hexColourMatcher.MatchString(fv) {
				return validPassports
			}

		case "ecl":
			if !eyeColourMatcher.MatchString(fv) {
				return validPassports
			}

		case "pid":
			if !pidMatcher.MatchString(fv) {
				return validPassports
			}
		}
	}
	return validPassports + 1
}

var hexColourMatcher = regexp.MustCompile(`^#[\da-f]{6}$`)
var eyeColourMatcher = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var heightMatcher = regexp.MustCompile(`^(\d+)(in|cm)$`)
var pidMatcher = regexp.MustCompile(`^\d{9}$`)

func validHeight(hv string) bool {
	height := heightMatcher.FindStringSubmatch(hv)

	if len(height) == 0 {
		return false
	}

	switch height[2] {
	case "in":
		return validNumber(59, 76, height[1])

	case "cm":
		return validNumber(150, 193, height[1])
	}

	return false
}

func validNumber(min, max int, year string) bool {
	conv, err := strconv.Atoi(year)
	if err != nil {
		return false
	}

	if conv < min || conv > max {
		return false
	}

	return true
}

func solveFirst(p *passport, validPassports int) int {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range mandatory {
		if _, isset := p.fields[f]; !isset {
			return validPassports
		}
	}
	return validPassports + 1
}

func solveStream(solve solver, inp chan *passport) chan int {
	out := make(chan int)

	result := 0

	go func() {
		defer close(out)

		for {
			select {
			case p, ok := <-inp:
				if !ok {
					out <- result
					return
				}

				result = solve(p, result)
			}
		}

	}()

	return out
}

type passport struct {
	fields map[string]string
}

func newPassport() *passport {
	return &passport{map[string]string{}}
}

func (p *passport) Valid() bool {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, f := range mandatory {
		if _, isset := p.fields[f]; !isset {
			return false
		}
	}
	return true
}

func (p *passport) SetField(field, value string) {
	p.fields[field] = value
}

func (p *passport) SetFields(fs []kv) {
	for _, f := range fs {
		p.SetField(f.k, f.v)
	}
}

func convStream(inp chan *linestream.Line) chan *passport {
	out := make(chan *passport)

	go func() {
		defer close(out)

		var p *passport = newPassport()

		shouldSend := false

		for {
			select {
			case v, ok := <-inp:
				if !ok {
					if shouldSend {
						out <- p
					}
					return
				}

				if v.Content() == "" {
					out <- p
					p = newPassport()
					shouldSend = false
					continue
				}

				p.SetFields(splitValues(v.Content()))
				shouldSend = true
			}
		}
	}()

	return out
}

type kv struct {
	k string
	v string
}

func splitValues(raw string) []kv {
	var kvs []kv

	for _, rawPair := range strings.Split(raw, " ") {
		rawKV := strings.Split(rawPair, ":")
		kvs = append(kvs, kv{rawKV[0], rawKV[1]})
	}

	return kvs
}
