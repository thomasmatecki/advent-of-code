package main

import (
	"bytes"
	"os"
	"regexp"
)

type DirIter struct {
	cursor     int
	directions []byte
}

func (dirIter *DirIter) Next() byte {
	next := dirIter.directions[dirIter.cursor]
	dirIter.cursor++
	dirIter.cursor = dirIter.cursor % len(dirIter.directions)
	return next
}

type NodeId [3]byte

func EndsInZ(nodeId NodeId) bool {
	return nodeId[2] == 'Z'
}

func IsZZZ(nodeId NodeId) bool {
	return nodeId == NodeId{'Z', 'Z', 'Z'}
}

var Start = NodeId{'A', 'A', 'A'}
var End = NodeId{'Z', 'Z', 'Z'}

type NodePair struct {
	left  NodeId
	right NodeId
}

type NodeMap map[NodeId]NodePair

func NewNodeMap(lines [][]byte) *NodeMap {
	nodeMap := make(NodeMap, len(lines))
	nodeId := new(NodeId)
	nodePair := new(NodePair)
	for _, line := range lines[2:] {
		match := nodeExpr.FindSubmatch(line)
		if match == nil {
			continue
		}
		copy(nodeId[:], match[1][:3])
		copy(nodePair.left[:], match[2][:3])
		copy(nodePair.right[:], match[3][:3])
		nodeMap[*nodeId] = *nodePair
	}

	return &nodeMap
}

var data, _ = os.ReadFile("08.txt")
var nodeExpr = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)
var lines = bytes.Split(data, []byte("\n"))
var nodeMap = NewNodeMap(lines)

func traverseCount(start NodeId, until func(NodeId) bool) int {

	dirIter := DirIter{
		directions: lines[0],
		cursor:     0,
	}

	nodeId := start
	count := 0
	for !until(nodeId) {
		nodePair := (*nodeMap)[nodeId]
		if dirIter.Next() == 'L' {
			nodeId = nodePair.left
		} else {
			nodeId = nodePair.right
		}
		count++
	}

	return count

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func PartOne() {
	count := traverseCount(Start, IsZZZ)
	println("Part One", count)
}

func PartTwo() {
	total := 1

	for nodeId := range *nodeMap {
		if nodeId[2] == 'A' {
			count := traverseCount(nodeId, EndsInZ)
			total = lcm(total, count)
		}
	}

	println("Part Two", total)
}

func main() {
	PartOne()
	PartTwo()
}
