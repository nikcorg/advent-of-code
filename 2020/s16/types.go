package s16

import (
	"fmt"
	"math"
	"regexp"

	"github.com/nikcorg/aoc2020/utils"
)

type Ticket []int

type Range struct {
	upper, lower int
}

func (r *Range) String() string {
	return fmt.Sprintf("%d-%d", r.lower, r.upper)
}

func (r *Range) InRange(v int) bool {
	return r.upper >= v && v >= r.lower
}

type FieldConfiguration struct {
	name        string
	validRanges []*Range
}

func (c *FieldConfiguration) String() string {
	s := fmt.Sprintf("%s: ", c.name)

	for _, r := range c.validRanges {
		s += r.String() + " "
	}

	return s
}

var fieldConfigSplitter = regexp.MustCompile(`([^:]+): (?:(\d+)-(\d+)) or (?:(\d+)-(\d+))`)

func newFieldConfigurationFromString(s string) *FieldConfiguration {
	matches := fieldConfigSplitter.FindStringSubmatch(s)

	if matches == nil {
		panic(fmt.Errorf("invalid field config: %s", s))
	}

	cfg := &FieldConfiguration{}
	cfg.name = matches[1]
	cfg.AddRange(utils.MustAtoi(matches[3]), utils.MustAtoi(matches[2]))
	cfg.AddRange(utils.MustAtoi(matches[5]), utils.MustAtoi(matches[4]))

	return cfg
}

func (c *FieldConfiguration) MatchesAny(v int) bool {
	for _, r := range c.validRanges {
		if r.InRange(v) {
			return true
		}
	}

	return false
}

func (c *FieldConfiguration) AddRange(upper, lower int) {
	c.validRanges = append(c.validRanges, &Range{upper, lower})
}

type TicketValidator struct {
	fields []*FieldConfiguration
}

func (v *TicketValidator) String() string {
	s := ""
	for _, f := range v.fields {
		s += f.String() + " "
	}
	return s
}

func (v *TicketValidator) InvalidValues(t *Ticket) []int {
	invalidValues := make([]int, 0, len(*t))

	for _, tv := range *t {
		didMatch := false
		for _, f := range v.fields {
			if f.MatchesAny(tv) {
				didMatch = true
				break
			}
		}

		if !didMatch {
			invalidValues = append(invalidValues, tv)
		}
	}

	return invalidValues
}

func (v *TicketValidator) Validates(t *Ticket) bool {
	for _, tv := range *t {
		didMatch := false
		for _, f := range v.fields {
			if f.MatchesAny(tv) {
				didMatch = true
				break
			}
		}

		if !didMatch {
			return false
		}
	}
	return true
}

// MatchFields returns a bitmask where a high bit indicates a matching field
func (v *TicketValidator) MatchFields(val int) uint {
	var matched uint = 0
	for n, f := range v.fields {
		if f.MatchesAny(val) {
			matched = matched | uint(math.Pow(2, float64(n)))
		}
	}
	return matched
}

func (v *TicketValidator) AddField(f *FieldConfiguration) {
	v.fields = append(v.fields, f)
}
