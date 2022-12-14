package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"nikc.org/aoc2022/13/packet"
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
	fmt.Fprint(out, "=====[ Day 13 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", solveFirst(input))
	fmt.Fprintf(out, "second: %d\n", solveSecond(input))
	return nil
}

func solveFirst(input string) int {
	s := bufio.NewScanner(strings.NewReader(input))
	n, valids := 1, 0
	var left, right packet.Packet
	for s.Scan() {
		if s.Text() == "" {
			left, right = nil, nil
			n++
			continue
		}

		if left == nil {
			left = packet.ParseMessage(s.Text())
		} else if right == nil {
			right = packet.ParseMessage(s.Text())
		}

		if left != nil && right != nil && left.Compare(right) == packet.GreaterThan {
			valids += n
		}
	}

	return valids
}

const (
	two = "[[2]]"
	six = "[[6]]"
)

func solveSecond(input string) int {
	pks := packet.PacketSlice{packet.ParseMessage(two), packet.ParseMessage(six)}

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		pks = append(pks, packet.ParseMessage(s.Text()))
	}

	sort.Sort(pks)

	ptwo, psix := -1, -1

	for i, p := range pks {
		if p.(fmt.Stringer).String() == two {
			ptwo = i + 1
		} else if p.(fmt.Stringer).String() == six {
			psix = i + 1
		}

		if ptwo >= 0 && psix >= 0 {
			break
		}
	}

	return ptwo * psix
}
