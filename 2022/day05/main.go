package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var crateExpr = regexp.MustCompile(`[(\w)]`)
var moveExpr = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
var data, _ = os.ReadFile("05.txt")

type CrateNode struct {
	val  byte
	next *CrateNode
}

type CrateStack struct {
	head *CrateNode
	tail *CrateNode
}

func (stack *CrateStack) Append(val byte) {
	if stack.head == nil {
		stack.head = &CrateNode{
			val:  val,
			next: stack.head,
		}
		stack.tail = stack.head
	} else {
		stack.tail.next = &CrateNode{
			val:  val,
			next: nil,
		}
		stack.tail = stack.tail.next
	}
}

func (stack *CrateStack) Pop() byte {
	node := stack.head

	if stack.head == stack.tail {
		stack.tail = nil
		stack.head = nil
	} else {
		stack.head = stack.head.next
	}

	return node.val
}

func (stack *CrateStack) Push(val byte) {
	if stack.head == nil {
		stack.head = &CrateNode{
			val:  val,
			next: stack.head,
		}
		stack.tail = stack.head
	} else {
		stack.head = &CrateNode{
			val:  val,
			next: stack.head,
		}
	}
}

/*
I want a pretty output in the debug window, but this doesn't work!
*/
func (stack CrateStack) String() string {
	node := stack.head
	var crates []byte
	for node != nil {
		crates = append(crates, node.val)
		node = node.next
	}
	return string(crates)
}

func initStacks(lines []string) [9]CrateStack {
	crateStacks := new([9]CrateStack)

	for _, line := range lines[0:8] {
		for i := 0; i < 9; i++ {
			j := i * 4
			chunk := line[j : j+3]
			match := crateExpr.FindStringSubmatch(chunk)
			if match != nil {
				crateStacks[i].Append(match[0][0])
			}
		}
	}
	return *crateStacks

}

type Move struct {
	count, from, to int
}

func parseMove(line string) Move {

	match := moveExpr.FindStringSubmatch(line)

	count, _ := strconv.Atoi(match[1])
	from, _ := strconv.Atoi(match[2])
	to, _ := strconv.Atoi(match[3])
	return Move{
		count: count,
		from:  from,
		to:    to,
	}
}

func PartOne() {
	lines := strings.Split(string(data), "\n")
	crateStacks := initStacks(lines)

	for _, line := range lines[10:] {
		if len(line) == 0 {
			break
		}

		move := parseMove(line)

		for i := 0; i < move.count; i++ {
			crateStacks[move.to-1].Push(crateStacks[move.from-1].Pop())
		}
	}

	heads := new([9]byte)

	for idx, stack := range crateStacks {
		heads[idx] = stack.head.val
	}

	println("Part One:", string(heads[:]))
}

func PartTwo() {

	lines := strings.Split(string(data), "\n")
	crateStacks := initStacks(lines)

	for _, line := range lines[10:] {
		match := moveExpr.FindStringSubmatch(line)
		if len(match) == 0 {
			break
		}
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])
		moves := make([]byte, count)

		for i := 0; i < count; i++ {
			moves[i] = crateStacks[from-1].Pop()
		}

		for i := count - 1; i >= 0; i-- {
			crateStacks[to-1].Push(moves[i])
		}
	}

	heads := new([9]byte)

	for idx, stack := range crateStacks {
		heads[idx] = stack.head.val
	}

	println("Part Two:", string(heads[:]))

}

func main() {
	PartOne()
	PartTwo()
}
