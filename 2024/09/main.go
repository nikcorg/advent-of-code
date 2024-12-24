package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strings"

	"nikc.org/aoc2024/util"
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
	first := solveFirst(input)
	second := solveSecond(input)

	fmt.Fprint(out, "=====[ Day 09 ]=====\n")
	fmt.Fprintf(out, "first: %d\n", first)
	fmt.Fprintf(out, "second: %d\n", second)

	return nil
}

const (
	space int = iota
	file
)

type fileblock struct {
	typ, length int
}

func (b *fileblock) Slice() fileblock {
	b0 := fileblock{b.typ, 1}
	b.length--
	return b0
}

type fs []fileblock

func parseInput(i string) fs {
	s := bufio.NewScanner(strings.NewReader(i))
	fm := fs{}

	id := 0
	for s.Scan() {
		for i, x := range strings.Split(s.Text(), "") {
			if i%2 == 0 {
				fm = append(fm, fileblock{typ: file, length: util.MustAtoi(x)})
				id++
			} else {
				fm = append(fm, fileblock{typ: space, length: util.MustAtoi(x)})
			}
		}
	}

	return fm
}

func solveFirst(i string) int {
	// dense := []int{1, 2, 3, 4, 5}
	dense := parseInput(i)

	// the dense needs to mimic the sparse version for copying over, meaning the
	// cursor at the right side cannot move to the next space over, until it has
	// decremented the size of the file. (maybe decrement the file size until it
	// reaches 0?) the file id from the right cursor repeats for every empty space
	// at the left cursor until the right cursor shifts onto the next file block.

	sparse := ""

	l, r := 0, len(sparse)-1

	for l < r {
		if dense[l].typ == space {
			for dense[r].typ == space || dense[r].length == 0 {
				r--
			}
			dense[l] = dense[r].Slice()
		}

		l++
	}

	tot := 0

	for i, s := range dense {
		if s.typ == space {
			continue
		}
		tot += s.id * i
	}

	return tot
}

func solveSecond(i string) int {
	return 0
}
