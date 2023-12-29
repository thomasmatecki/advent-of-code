package main

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Input map[byte]int

type Rng struct {
	lbnd int
	ubnd int
}

func (rng *Rng) lclamp(val int) {
	rng.lbnd = max(rng.lbnd, val)
}
func (rng *Rng) uclamp(val int) {
	rng.ubnd = min(rng.ubnd, val)
}

type Space map[byte]*Rng

func (space *Space) Size() int {
	count := 1
	for _, c := range []byte("xmas") {
		count *= (*space)[c].ubnd - (*space)[c].lbnd
	}
	return count
}

func (space *Space) Partiion(cond *Conditional) Space {
	otherSpace := Space{}

	for _, c := range []byte("xmas") {
		rng := *(*space)[c]
		otherSpace[c] = &rng
	}

	if cond.comp == IfGreaterThan {
		(*space)[cond.key].lclamp(cond.val)
		(otherSpace)[cond.key].uclamp(cond.val)
	} else {
		(*space)[cond.key].uclamp(cond.val - 1)
		(otherSpace)[cond.key].lclamp(cond.val - 1)
	}

	return otherSpace
}

func (in *Input) Sum() (sum int) {
	for _, v := range *in {
		sum += v
	}
	return
}

type Rule interface {
	Applies(input *Input) bool
	Accepted(input *Input) bool
	Count(space *Space) (int, Space)
}

type Rules []Rule

func (rs *Rules) Accepted(input *Input) bool {
	for _, rule := range *rs {
		if rule.Applies(input) {
			return rule.Accepted(input)
		}
	}
	panic("no result!")
}
func (rs *Rules) Count(initialSpace *Space) (totalCount int) {
	space := initialSpace

	for _, rule := range *rs {
		count, nextSpace := rule.Count(space)
		totalCount += count
		if nextSpace == nil {
			break
		} else {
			space = &nextSpace
		}
	}
	return
}

type Lookup map[string]Rules
type Comparator byte

const (
	IfGreaterThan Comparator = iota
	IfLessThan    Comparator = iota
)

type Terminal struct {
	accept bool
}

/*
 *
 */
func (t Terminal) Applies(input *Input) bool {
	return true
}
func (t Terminal) Accepted(input *Input) bool {
	return t.accept
}
func (t Terminal) Count(space *Space) (int, Space) {
	if t.accept {
		return space.Size(), nil
	} else {
		return 0, nil
	}
}

/*
 *
 */
type Goto struct {
	dest   string
	lookup *Lookup
}

func (g Goto) Applies(input *Input) bool {
	return true
}
func (g Goto) Accepted(input *Input) bool {
	rules := (*g.lookup)[g.dest]
	return rules.Accepted(input)
}
func (g Goto) Count(space *Space) (int, Space) {
	rules := (*g.lookup)[g.dest]
	return rules.Count(space), nil
}

/*
 *
 */
type Conditional struct {
	key  byte
	comp Comparator
	val  int
}

func NewConditional(match []string) Conditional {
	var (
		comp Comparator
	)

	if match[2][0] == '>' {
		comp = IfGreaterThan
	} else {
		comp = IfLessThan
	}

	val, _ := strconv.Atoi(match[3])

	return Conditional{
		key:  match[1][0],
		comp: comp,
		val:  val,
	}
}

func (c *Conditional) Eval(input *Input) bool {
	compVal := (*input)[c.key]
	if c.comp == IfGreaterThan {
		return compVal > c.val
	} else {
		return compVal < c.val
	}
}

/**
 *
 */
type ConditionalGoto struct {
	cond Conditional
	goTo Goto
}

func (cg ConditionalGoto) Applies(input *Input) bool {
	return cg.cond.Eval(input)
}

func (cg ConditionalGoto) Accepted(input *Input) bool {
	return cg.goTo.Accepted(input)
}

func (cg ConditionalGoto) Count(space *Space) (int, Space) {
	otherSpace := space.Partiion(&cg.cond)
	goToCount, _ := cg.goTo.Count(space)
	return goToCount, otherSpace
}

/**
 */
type ConditionalTerminal struct {
	cond   Conditional
	accept bool
}

func (ct ConditionalTerminal) Applies(input *Input) bool {
	return ct.cond.Eval(input)
}
func (ct ConditionalTerminal) Accepted(input *Input) bool {
	return ct.accept
}
func (ct ConditionalTerminal) Count(space *Space) (int, Space) {
	otherSpace := space.Partiion(&ct.cond)
	if ct.accept {
		return space.Size(), otherSpace
	} else {
		return 0, otherSpace
	}
}

var (
	workflowExpr             = regexp.MustCompile(`(\w+)\{(.+)\}`)
	gotoLabelExpr            = regexp.MustCompile(`(\w+)`)
	conditionalGotoLabelExpr = regexp.MustCompile(`(\w)([<>])(\d+):([a-z]+)`)
	conditionalTerminalExpr  = regexp.MustCompile(`(\w)([<>])(\d+):([AR])`)
	valAssignmentExpr        = regexp.MustCompile(`(\w)=(\d+)`)
)

func ParseLastRule(s string, lookup *Lookup) Rule {
	if s == "A" {
		return Terminal{true}
	} else if s == "R" {
		return Terminal{false}
	} else if label := gotoLabelExpr.FindStringSubmatch(s); label != nil {
		return Goto{label[0], lookup}
	}
	panic("no!")
}

func ParseRule(s string, lookup *Lookup) Rule {
	if match := conditionalGotoLabelExpr.FindStringSubmatch(s); match != nil {

		return ConditionalGoto{
			cond: NewConditional(match),
			goTo: Goto{
				dest:   match[4],
				lookup: lookup,
			},
		}

	} else if match := conditionalTerminalExpr.FindStringSubmatch(s); match != nil {
		var accept bool

		if match[4] == "A" {
			accept = true
		} else {
			accept = false
		}

		return ConditionalTerminal{
			cond:   NewConditional(match),
			accept: accept,
		}
	}
	panic("no!")
}

func (lookup *Lookup) Add(workflowStr []byte) {
	match := workflowExpr.FindSubmatch(workflowStr)
	rules := []Rule{}

	rulesStr := strings.Split(string(match[2]), ",")

	for _, ruleStr := range rulesStr[:len(rulesStr)-1] {
		rules = append(rules, ParseRule(ruleStr, lookup))
	}
	lastRule := ParseLastRule(rulesStr[len(rulesStr)-1], lookup)
	rules = append(rules, lastRule)

	(*lookup)[string(match[1])] = rules
}

func BuildWorkflows(workflowsStrs []byte) Lookup {
	workflows := make(Lookup)

	for _, workflowStr := range bytes.Split(workflowsStrs, []byte("\n")) {
		workflows.Add(workflowStr)
	}

	return workflows
}

func BuildInputs(inputStrs []byte) []Input {
	lines := bytes.Split(bytes.Trim(inputStrs, "\n"), []byte("\n"))
	inputs := make([]Input, len(lines))

	for idx, inputStr := range lines {
		trimmed := inputStr[1 : len(inputStr)-1]
		valStrs := bytes.Split(trimmed, []byte{','})
		inputs[idx] = make(Input)
		input := &inputs[idx]
		for _, valStr := range valStrs {
			match := valAssignmentExpr.FindSubmatch(valStr)
			key := match[1][0]
			val, _ := strconv.Atoi(string(match[2]))
			(*input)[key] = val
		}
	}
	return inputs
}

func Exec(filename string) (total int) {
	lines, _ := os.ReadFile(filename)
	split := bytes.Split(lines, []byte("\n\n"))
	workflowsStrs, inputStrs := split[0], split[1]
	workflows := BuildWorkflows(workflowsStrs)
	inputs := BuildInputs(inputStrs)
	start := workflows["in"]
	for _, input := range inputs {
		if start.Accepted(&input) {
			total += input.Sum()
		}
	}
	return
}

func PartOne() {
	println("Part One:", Exec("19.txt"))
}

func NewSpace(factor int) *Space {
	return &Space{
		'x': &Rng{0, factor},
		'm': &Rng{0, factor},
		'a': &Rng{0, factor},
		's': &Rng{0, factor},
	}
}

func Explore(filename string) (total int) {
	space := NewSpace(4000)
	lines, _ := os.ReadFile(filename)
	split := bytes.Split(lines, []byte("\n\n"))
	workflowsStrs, _ := split[0], split[1]
	workflows := BuildWorkflows(workflowsStrs)
	start := workflows["in"]
	return start.Count(space)
}

func PartTwo() {
	println("Part Two:", Explore("19.txt"))
}

func main() {
	PartOne()
	PartTwo()
}
