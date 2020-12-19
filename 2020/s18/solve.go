package s18

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/nikcorg/aoc2020/utils"
	"github.com/nikcorg/aoc2020/utils/linestream"
)

const bufSize = 1

type Solver struct {
	ctx context.Context
	out io.Writer
}

func New(ctx context.Context, out io.Writer) *Solver {
	return &Solver{ctx, out}
}

func (s *Solver) SolveFirst(inp io.Reader) error {
	return s.Solve(1, inp)
}

func (s *Solver) SolveSecond(inp io.Reader) error {
	return s.Solve(2, inp)
}

func (s *Solver) Solve(part int, inp io.Reader) error {
	lineInput := make(linestream.LineChan, bufSize)
	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	solution := solve(linestream.SkipEmpty(lineInput), getSolver(part))

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func([]*token) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return evalFirst
	case 2:
		return evalSecond
	}
	panic(fmt.Errorf("invalid part %d", part))
}

func solve(inp linestream.ReadOnlyLineChan, eval solver) int {
	total := 0
	for line := range inp {
		tokens := tokenise(strings.ReplaceAll(line.Content(), " ", ""))
		total += eval(tokens)
	}
	return total
}

var stmtTokeniser = regexp.MustCompile(`(\d+|[+*()])`)

type tokenKind int

const (
	mul tokenKind = iota + 1
	add
	val
	lhb // left hand bracket = (
	rhb // right hand bracket = )
)

type token struct {
	kind   tokenKind
	intVal int
}

func (t *token) String() string {
	switch t.kind {
	case mul:
		return "*"
	case add:
		return "+"
	case val:
		return fmt.Sprintf("%d", t.intVal)
	case lhb:
		return "("
	case rhb:
		return ")"
	}

	return fmt.Sprintf("token(%v)", t.kind)
}

func tokenise(line string) []*token {
	tokens := []*token{}

	stringTokens := stmtTokeniser.FindAllStringSubmatch(line, -1)

	for _, tok := range stringTokens {
		var t *token

		switch tok[1] {
		case "*":
			t = &token{mul, 0}
		case "+":
			t = &token{add, 0}
		case "(":
			t = &token{lhb, 0}
		case ")":
			t = &token{rhb, 0}
		default:
			t = &token{val, utils.MustAtoi(tok[1])}
		}

		tokens = append(tokens, t)
	}

	return tokens
}

func evalFirst(toks []*token) int {
	s := 0
	stacks := [][]*token{{}}

	for n := 0; n < len(toks); n++ {
		tok := toks[n]
		switch tok.kind {
		case lhb:
			s++
			stacks = append(stacks, []*token{})
		case rhb:
			tok = &token{val, evalFirst(stacks[s])}
			stacks = stacks[0 : len(stacks)-1]
			s--
			fallthrough
		default:
			stacks[s] = append(stacks[s], tok)
		}
	}

	if len(stacks) > 1 {
		panic(fmt.Errorf("expected a flat stack: %+v", stacks))
	}

	stack := stacks[0]

	if len(stack) == 0 {
		return 0
	}

	total := stack[0].intVal

	for n := 2; n < len(stack); n += 2 {
		op := stack[n-1]
		rhs := stack[n]

		switch op.kind {
		case add:
			total += rhs.intVal
		case mul:
			total *= rhs.intVal
		}
	}

	return total
}

func evalSecond(toks []*token) int {
	s := 0
	stacks := [][]*token{{}}

	// eval parenthesised expressions
	for n := 0; n < len(toks); n++ {
		tok := toks[n]
		switch tok.kind {
		case lhb:
			s++
			stacks = append(stacks, []*token{})
		case rhb:
			tok = &token{val, evalSecond(stacks[s])}
			stacks = stacks[0 : len(stacks)-1]
			s--
			fallthrough
		default:
			stacks[s] = append(stacks[s], tok)
		}
	}

	if len(stacks) > 1 {
		panic(fmt.Errorf("expected a flat stack: %+v", stacks))
	}

	stack := stacks[0]

	if len(stack) == 0 {
		return 0
	}

	// reduce additions to values
	end := len(stack)

	for n := 2; n < end; {
		lhs := stack[n-2]
		op := stack[n-1]
		rhs := stack[n]

		switch op.kind {
		case add:
			// omstart!
			stack = append(append(stack[0:n-2], &token{val, lhs.intVal + rhs.intVal}), stack[n+1:]...)
			end = len(stack)
			n = 2
		default:
			n += 2
		}
	}

	if len(stack) == 1 {
		return stack[0].intVal
	}

	// multiplications
	total := stack[0].intVal
	for n := 2; n < len(stack); n += 2 {
		op := stack[n-1]
		lhs := stack[n]
		switch op.kind {
		case mul:
			total *= lhs.intVal
		default:
			panic(fmt.Errorf("expected %+v to be mul", op))
		}
	}

	return total
}
