package s16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFieldConfigFromString(t *testing.T) {
	tests := []struct {
		input          string
		expectedName   string
		expectedRanges [][]int
	}{
		{`class: 0-1 or 4-19`, "class", [][]int{{0, 1}, {4, 19}}},
		{`row: 0-5 or 8-19`, "row", [][]int{{0, 5}, {8, 19}}},
		{`seat: 0-13 or 16-19`, "seat", [][]int{{0, 13}, {16, 19}}},
	}

	for _, test := range tests {
		fc := newFieldConfigurationFromString(test.input)

		assert.Equal(t, len(test.expectedRanges), len(fc.validRanges))
		assert.Equal(t, test.expectedName, fc.name)

		for n, er := range test.expectedRanges {
			assert.Equal(t, er[0], fc.validRanges[n].lower)
			assert.Equal(t, er[1], fc.validRanges[n].upper)
		}
	}
}

func TestMatchFields(t *testing.T) {
	fieldInputs := []string{
		`class: 0-1 or 4-19`,
		`row: 0-5 or 8-19`,
		`seat: 0-13 or 16-19`,
	}

	tv := &TicketValidator{}

	for _, fi := range fieldInputs {
		tv.AddField(newFieldConfigurationFromString(fi))
	}

	tests := []struct {
		input         []int
		expectedMatch uint // decimal number from composing bitmasks
		expectedField string
	}{
		{[]int{3, 15, 9}, 2, "class"},
		{[]int{9, 1, 14}, 1, "row"},
		{[]int{18, 5, 9}, 4, "seat"},
	}

	var (
		outcomes []uint
	)

	for n, test := range tests {
		var outcome uint = 0
		for x, val := range test.input {
			matches := tv.MatchFields(val)
			if x == 0 {
				outcome = matches
			} else {
				outcome = outcome & tv.MatchFields(val)
			}
		}

		if n == 0 {
			outcomes = []uint{outcome}
		} else {
			for _, o := range outcomes {
				outcome = o ^ outcome
			}
			outcomes = append(outcomes, outcome)
		}

		assert.Equal(t, uint(0), test.expectedMatch^outcome)
	}
}
