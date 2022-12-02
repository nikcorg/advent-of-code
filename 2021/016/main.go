package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	if result, err := solveFirst(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 1: %d\n", result))
	}

	if result, err := solveSecond(parsed); err != nil {
		return err
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("solution 2: %d\n", result))
	}

	return nil
}

type voidT struct{}

var void = struct{}{}

type packet struct {
	Version int
	Type    int
	Value   uint64
}

type puzzleInput []packet

var hex2bin = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

var bin2int = map[string]int{
	// three-bit sequence
	"001": 1,
	"010": 2,
	"011": 3,
	"100": 4,
	"101": 5,
	"110": 6,
	"111": 7,

	// four-bit sequence
	"0001": 1,
	"0010": 2,
	"0011": 3,
	"0100": 4,
	"0101": 5,
	"0110": 6,
	"0111": 7,
	"1000": 8,
	"1001": 9,
	"1010": 10,
	"1011": 11,
	"1100": 12,
	"1101": 13,
	"1110": 14,
	"1111": 15,
}

func parseInput(raw string) (puzzleInput, error) {
	bits := make([]string, len(raw))

	for i, c := range strings.Split(raw, "") {
		bits[i] = hex2bin[c]
	}

	return bitsToPackets(strings.Join(bits, "")), nil
}

func parseLiteral(bits string) (uint64, string) {
	var (
		rest   string = bits[0:]
		chunks []string
	)

	for {
		chunk := rest[0:5]
		rest = rest[5:]

		chunks = append(chunks, chunk[1:])

		if chunk[0] == '0' {
			break
		}
	}

	val, _ := strconv.ParseUint(strings.Join(chunks, ""), 2, 64)

	// FIXME: need to strip out padding from rest string
	bitslen := len(chunks) * 5

	return val, rest
}

func parseOperator(bits string) (voidT, string) {
	var rest string

	return void, rest
}

func bitsToPackets(bits string) []packet {
	p := packet{}

	p.Version = bin2int[bits[0:3]]
	p.Type = bin2int[bits[3:6]]

	var rest string

	switch p.Type {
	case 4:
		p.Value, rest = parseLiteral(bits[7:])

	default:
		_, rest = parseOperator(bits[7:])
	}

	return append([]packet{p}, bitsToPackets(rest)...)
}

func solveFirst(input puzzleInput) (int, error) {
	return 0, errors.New("not implemented")
}

func solveSecond(input puzzleInput) (int, error) {
	return 0, errors.New("not implemented")
}
