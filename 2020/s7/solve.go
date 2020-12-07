package s7

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"sync"

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
	ctx, cancel := context.WithCancel(s.ctx)
	defer func() { cancel() }()

	lineInput := make(linestream.LineChan, bufSize)
	linestream.New(ctx, bufio.NewReader(inp), lineInput)

	containmentRules := make(chan *containmentRule, bufSize)
	convStream(linestream.SkipEmpty(lineInput), containmentRules)

	solution := <-solveStream(getSolver(part), containmentRules)

	io.WriteString(os.Stdout, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type ruleset map[string]*containmentRule
type solver func(ruleset) int

type containmentRule struct {
	colour      string
	contains    map[string]int
	containedBy map[string]bool
}

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst

	case 2:
		return solveSecond
	}
	panic(fmt.Errorf("unknown part %d", part))
}

const colourToSolve = "shiny gold"

func solveFirst(rs ruleset) int {
	total := 0

	colourHits := make(chan string, 1)

	go func() {
		colours := map[string]bool{}

		for h := range colourHits {
			if _, ok := colours[h]; !ok {
				total++
			}
			colours[h] = true
		}
	}()

	wg := sync.WaitGroup{}

	var traverse func(r *containmentRule)

	traverse = func(r *containmentRule) {
		defer wg.Done()

		for c := range r.containedBy {
			cr := rs[c]
			colourHits <- cr.colour
			wg.Add(1)
			go traverse(cr)
		}
	}

	wg.Add(1)
	go traverse(rs[colourToSolve])

	wg.Wait()
	close(colourHits)

	return total
}

func solveSecond(rs ruleset) int {
	total := 0

	bags := make(chan int, 1)
	go func() {
		for cnt := range bags {
			total += cnt
		}
	}()

	wg := sync.WaitGroup{}

	var traverse func(*containmentRule, int)
	traverse = func(r *containmentRule, multiplier int) {
		defer wg.Done()
		for colour, num := range r.contains {
			bags <- num * multiplier

			wg.Add(1)
			go traverse(rs[colour], multiplier*num)
		}
	}

	wg.Add(1)
	go traverse(rs[colourToSolve], 1)
	wg.Wait()

	close(bags)

	return total
}

func solveStream(solve solver, rules chan *containmentRule) chan int {
	out := make(chan int)

	allRules := ruleset{}

	go func() {
		defer close(out)

		for ir := range rules {
			// fmt.Println(*ir)

			if r, ok := allRules[ir.colour]; ok {
				// a rule can exist before it's defined, if the colour is contained by another rule defined earlier
				// we copy the `containedBy` from the existing stub rule into the incoming rule
				ir.containedBy = r.containedBy
			}

			allRules[ir.colour] = ir

			// patch the ruleset to reference this rule, for any colours contained by this rule
			for contains := range ir.contains {
				if r, ok := allRules[contains]; !ok {
					// the colour rule doesn't exist yet, create a provisional rule
					pr := &containmentRule{}
					pr.containedBy = map[string]bool{ir.colour: true}
					allRules[contains] = pr
				} else {
					// the colour rule exists, patch the rule's containedBy map
					r.containedBy[ir.colour] = true
				}
			}
		}

		out <- solve(allRules)
	}()

	return out
}

func convStream(inp linestream.ReadOnlyLineChan, out chan<- *containmentRule) {
	ruleSplitter := regexp.MustCompile(`^(.*) bags contain (.*)\.$`)
	containsSplitter := regexp.MustCompile(`(?:(?:, )?((\d+) (\D*?)) bags?)`)

	go func() {
		defer close(out)

		for rulePhrase := range inp {
			ruleParts := ruleSplitter.FindStringSubmatch(rulePhrase.Content())

			if ruleParts == nil {
				panic(fmt.Errorf("unable to match rule: %s", rulePhrase.Content()))
			}

			colour := ruleParts[1]
			contains := ruleParts[2]

			rule := &containmentRule{colour, map[string]int{}, map[string]bool{}}

			for _, cr := range containsSplitter.FindAllStringSubmatch(contains, -1) {
				containedColour := cr[3]
				containedAmount := mustAtoi(cr[2])
				rule.contains[containedColour] = containedAmount
			}

			out <- rule
		}
	}()
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
