package main

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//go:embed input_test_a.txt
	testInputA string
	//go:embed input_test_b.txt
	testInputB string
	//go:embed input_test_c.txt
	testInputC string

	inputs = []string{testInputA, testInputB, testInputC}
	prices = []int{140, 772, 1930}

	_ = inputs
	_ = prices
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInputA))
}

func TestParseInput(t *testing.T) {
	parseInput(testInputA)
}

func TestSolveFirst(t *testing.T) {
	sol := solveFirst(testInputA)
	assert.Equal(t, 0, sol)
}

func TestSolveSecond(t *testing.T) {
	sol := solveSecond(testInputA)
	assert.Equal(t, 0, sol)
}
