package s19

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/nikcorg/aoc2020/utils"
	"github.com/nikcorg/aoc2020/utils/linestream"
)

const bufSize = 1

type Solver struct {
	ctx           context.Context
	out           io.Writer
	ruleOverrides map[int]string
}

func New(ctx context.Context, out io.Writer) *Solver {
	return &Solver{ctx, out, nil}
}

func (s *Solver) SolveFirst(inp io.Reader) error {
	return s.Solve(1, inp)
}

func (s *Solver) SolveSecond(inp io.Reader) error {
	s.ruleOverrides = map[int]string{
		8:  "42 | 42 8",
		11: "42 31 | 42 11 31",
	}
	return s.Solve(2, inp)
}

func (s *Solver) Solve(part int, inp io.Reader) error {
	lineInput := make(linestream.LineChan, bufSize)
	linestream.New(s.ctx, bufio.NewReader(inp), lineInput)

	rules := parseRules(lineInput, s.ruleOverrides)
	solve := getSolver(part)
	solution := solve(rules, linestream.SkipEmpty(lineInput))

	io.WriteString(s.out, fmt.Sprintf("solution: %d\n", solution))

	return nil
}

type solver func(map[int]string, linestream.ReadOnlyLineChan) int

func getSolver(part int) solver {
	switch part {
	case 1:
		return solveFirst
	case 2:
		return solveFirst
	}
	panic(fmt.Sprintf("invalid part %d", part))
}

func solveFirst(rules map[int]string, inp linestream.ReadOnlyLineChan) int {
	matches := 0

	root := constructRuleTree(0, rules[0], rules)
	for line := range inp {

		if matchLen := root.Matches(line.Content()); matchLen == len(line.Content()) {
			matches++
		}
	}

	return matches
}

/*
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
*/

type Mode int

func (m *Mode) String() string {
	switch *m {
	case matchAny:
		return "any"
	case matchAll:
		return "all"
	case literal:
		return "lit"
	}
	return fmt.Sprintf("mode(%v)", *m)
}

const (
	matchAny Mode = iota + 1
	matchAll
	literal
)

type Node struct {
	id       int
	mode     Mode
	nextNode []*Node
	value    string
}

func (n *Node) Matches(s string) int {
	var matchedLen int
	switch n.mode {
	case matchAny:
		for _, nn := range n.nextNode {
			ruleMatches := nn.Matches(s)
			if ruleMatches > 0 {
				matchedLen = ruleMatches
				break
			}
		}
	case matchAll:
		matchesAcc := 0
		for i := 0; i < len(n.nextNode); i++ {
			ruleMatches := n.nextNode[i].Matches(s[matchesAcc:])
			if ruleMatches == 0 {
				return 0
			}
			matchesAcc += ruleMatches
		}
		matchedLen += matchesAcc
	case literal:
		if string(s[0]) == n.value {
			matchedLen = 1
		}
	}

	return matchedLen
}

func (n *Node) String() string {
	switch n.mode {
	case matchAny:
		alts := []string{}
		for _, nn := range n.nextNode {
			alts = append(alts, nn.String())
		}
		return fmt.Sprintf("(%s)", strings.Join(alts, "|"))
	case matchAll:
		alts := []string{}
		for _, nn := range n.nextNode {
			alts = append(alts, nn.String())
		}
		return fmt.Sprintf("%s", strings.Join(alts, ""))
	case literal:
		return fmt.Sprintf("%s", n.value)
	}

	return fmt.Sprintf("Node(id:%d, mode:%s)", n.id, n.mode.String())
}

func parseRules(inp linestream.ReadOnlyLineChan, overrides map[int]string) map[int]string {
	rules := map[int]string{}

	for line := range inp {
		// an empty line marks the end of the rules section
		if line.Content() == "" {
			break
		}

		splits := strings.Split(line.Content(), ":")
		id := utils.MustAtoi(splits[0])
		rule := strings.TrimSpace(splits[1])

		rules[id] = rule
	}

	if overrides != nil {
		for k, v := range overrides {
			rules[k] = v
		}
	}

	return rules
}

func constructRuleTree(id int, rule string, rules map[int]string) *Node {
	node := &Node{}
	node.id = id

	if rule[0] == '"' {
		// literal node
		node.mode = literal
		node.value = string(rule[1])
	} else if paths := strings.Split(rule, " | "); len(paths) > 1 {
		// alternate paths
		node.mode = matchAny
		for _, path := range paths {
			node.nextNode = append(node.nextNode, constructRuleTree(id, strings.TrimSpace(path), rules))
		}
	} else if ids := strings.Split(rule, " "); len(ids) > 0 {
		// list of rule ids
		node.mode = matchAll
		for _, nextRule := range ids {
			nextID := utils.MustAtoi(nextRule)
			if nextID == id {
				panic(fmt.Errorf("cyclical rule discovered in rule %d", id))
			}
			node.nextNode = append(node.nextNode, constructRuleTree(nextID, rules[utils.MustAtoi(nextRule)], rules))
		}
	} else {
		fmt.Println("no idea", rule)
	}

	return node
}
