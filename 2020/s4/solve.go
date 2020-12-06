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

const bufSize = 1

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

	lineInput := make(linestream.LineChan, bufSize)
	linestream.New(ctx, bufio.NewReader(inputFile), lineInput)

	chunkedInput := make(linestream.ChunkedLineChan, bufSize)
	linestream.Chunked(lineInput, chunkedInput)

	passports := make(chan *passport, bufSize)
	convStream(chunkedInput, passports)

	solution := <-solveStream(getValidator(part), passports)

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type validator func(*passport) bool

func getValidator(part int) validator {
	switch part {
	case 1:
		return validateLax
	case 2:
		return validateStrict
	}

	panic(fmt.Errorf("invalid part: %d", part))
}

func validateStrict(p *passport) bool {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range mandatory {
		fv, isset := p.fields[f]
		if !isset {
			return false
		}

		switch f {
		case "byr":
			if !validNumber(1920, 2002, fv) {
				return false
			}

		case "iyr":
			if !validNumber(2010, 2020, fv) {
				return false
			}

		case "eyr":
			if !validNumber(2020, 2030, fv) {
				return false
			}

		case "hgt":
			if !validHeight(fv) {
				return false
			}

		case "hcl":
			if !hexColourMatcher.MatchString(fv) {
				return false
			}

		case "ecl":
			if !eyeColourMatcher.MatchString(fv) {
				return false
			}

		case "pid":
			if !pidMatcher.MatchString(fv) {
				return false
			}
		}
	}
	return true
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

func validateLax(p *passport) bool {
	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, f := range mandatory {
		if _, isset := p.fields[f]; !isset {
			return false
		}
	}

	return true
}

func solveStream(valid validator, inp chan *passport) chan int {
	out := make(chan int)

	valids := make(chan bool, 1)
	totalValid := 0

	go func() {
		defer close(out)

		for isValid := range valids {
			if isValid {
				totalValid++
			}
		}

		out <- totalValid
	}()

	go func() {
		defer close(valids)

		for p := range inp {
			valids <- valid(p)
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

func convStream(inp linestream.ReadOnlyChunkedLineChan, out chan *passport) {
	go func() {
		defer close(out)

		var p *passport = newPassport()

		for v := range inp {
			p = newPassport()
			for _, l := range v {
				p.SetFields(splitValues(l.Content()))
			}
			out <- p
		}
	}()
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
