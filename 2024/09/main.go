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
	typ, id, length int
}

func (b *fileblock) Slice() fileblock {
	b0 := fileblock{b.typ, b.id, 1}
	b.length--
	return b0
}

func (b fileblock) String() string {
	switch b.typ {
	case space:
		return fmt.Sprintf(`[space len=%d]`, b.length)

	default:
		return fmt.Sprintf(`[id=%d len=%d]`, b.id, b.length)
	}
}

type fs []fileblock

func (f fs) String() string {
	s := ""
	for _, b := range f {
		switch b.typ {
		case space:
			s += strings.Repeat(".", b.length)
		default:
			s += strings.Repeat(string(fmt.Sprintf("%d", b.id)[0]), b.length)
		}
	}
	return s
}

func parseInput(i string) fs {
	s := bufio.NewScanner(strings.NewReader(i))
	fm := fs{}
	id := 0

	for s.Scan() {
		for i, x := range strings.Split(s.Text(), "") {
			bs := util.MustAtoi(x)
			if i%2 == 0 {
				fm = append(fm, fileblock{typ: file, id: id, length: bs})
				id++
			} else {
				fm = append(fm, fileblock{typ: space, length: bs})
			}
		}
	}

	return fm
}

func solveFirst(i string) int {
	dense := parseInput(i)
	sparse := fs{}
	l, r := 0, len(dense)-1

	for l < r {
		if dense[l].typ != space {
			for dense[l].length > 0 {
				sparse = append(sparse, dense[l].Slice())
			}
			l++
			continue
		}

		for dense[l].length > 0 {
			for dense[r].typ == space || dense[r].length == 0 {
				r--
			}
			sparse = append(sparse, dense[r].Slice())
			dense[l].length--
		}

		l++
	}

	// drain final right block
	for dense[r].length > 0 {
		sparse = append(sparse, dense[r].Slice())
	}

	tot := 0

	for i, b := range sparse {
		if b.typ == space {
			continue
		}
		tot += b.id * i
	}

	return tot
}

func solveSecond(in string) int {
	dense := parseInput(in)
	tot := 0
	l, r := 0, len(dense)-1

	for l < r {
		for dense[r].typ == space {
			r--
		}

		for l < r {
			if dense[l].typ != space || dense[l].length < dense[r].length {
				l++
			} else {
				break
			}
		}

		if l >= r {
			l = 0
			r--
			continue
		}

		// empty block size matches file block size = simple swap
		if dense[r].length == dense[l].length {
			dense[l], dense[r] = dense[r], dense[l]
		} else {
			// file block is smaller than empty space = split space, splice and swap
			dense[l].length -= dense[r].length
			tmp := dense[r]
			dense[r] = fileblock{typ: space, length: tmp.length}
			dense = append(dense[0:l], append(fs{tmp}, dense[l:]...)...)
		}

		r--
		l = 0
	}

	i := 0
	for _, b := range dense {
		if b.typ == space {
			i += b.length
			continue
		}
		for b.length > 0 {
			tot += b.id * i
			b.length--
			i++
		}
	}

	return tot
}
