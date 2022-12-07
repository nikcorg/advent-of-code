package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"nikc.org/aoc2022/util"
	"nikc.org/aoc2022/util/stack"
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
	fmt.Fprint(out, "=====[ Day 05 ]=====\n")
	var (
		first, second string
		err           error
	)

	if first, err = solveFirst(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "first: %s\n", first)

	if second, err = solveSecond(input); err != nil {
		return err
	}

	fmt.Fprintf(out, "second: %s\n", second)

	return nil
}

func solveFirst(i string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(i))
	scanner.Split(bufio.ScanLines)

	stacks := parseStacks(scanner)
	digits := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		moves := util.Fmap(util.MustAtoi, digits.FindAllString(scanner.Text(), 3))

		n := moves[0]
		from := moves[1] - 1
		to := moves[2] - 1

		popped := util.CopySlice(stacks[from].PopN(n))
		util.ReverseSlice(popped)
		stacks[to].Push(popped...)
	}

	r := ""
	for _, s := range stacks {
		r += string(s.Peek())
	}
	return r, nil
}

func solveSecond(i string) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(i))
	scanner.Split(bufio.ScanLines)

	stacks := parseStacks(scanner)
	digits := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		moves := util.Fmap(util.MustAtoi, digits.FindAllString(scanner.Text(), 3))

		n := moves[0]
		from := moves[1] - 1
		to := moves[2] - 1

		popped := util.CopySlice(stacks[from].PopN(n))
		stacks[to].Push(popped...)
	}

	r := ""
	for _, s := range stacks {
		r += string(s.Peek())
	}
	return r, nil
}

func parseStacks(scanner *bufio.Scanner) []stack.Stack[byte] {
	b := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		b = append(b, scanner.Text())
	}

	stacksLine := b[len(b)-1]
	n := len(regexp.MustCompile(`\d+`).FindAllString(stacksLine, -1))

	stacks := make([]stack.Stack[byte], n)
	for x := 0; x < n; x++ {
		stacks[x] = stack.New[byte]()
	}

	for _, l := range util.ReverseSlice(b[0 : len(b)-1]) {
		for x := 0; x < n; x++ {
			p := x*4 + 1
			if p > len(l) {
				break
			}

			crate := l[p]
			if crate != ' ' {
				stacks[x].Push(crate)
			}
		}
	}

	return stacks
}
