package main

import (
	"bytes"
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	//go:embed input_test.txt
	testInput string
)

func TestMainWithErr(t *testing.T) {
	assert.NoError(t, mainWithErr(&bytes.Buffer{}, testInput))
}
