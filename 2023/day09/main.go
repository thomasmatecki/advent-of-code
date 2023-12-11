package main

import (
	"bytes"
	"os"
	"strconv"
)

var data, _ = os.ReadFile("09.txt")
var values [][]int

func init() {
	lines := bytes.Split(data, []byte("\n"))
	values = make([][]int, len(lines))
	for idx, line := range lines {
		for _, valueStr := range bytes.Split(line, []byte(" ")) {
			value, _ := strconv.Atoi(string(valueStr))
			values[idx] = append(values[idx], value)
		}
	}
}

func all(v int, is []int) bool {
	for _, i := range is {
		if v != i {
			return false
		}
	}
	return true
}

func nextRow(is []int) []int {
	js := make([]int, len(is)-1)
	for i := 0; i < len(is)-1; i++ {
		js[i] = is[i+1] - is[i]
	}
	return js
}

func nextVal(is []int) int {
	if all(0, is) {
		return 0
	} else {
		return is[len(is)-1] + nextVal(nextRow(is))
	}
}

func prevVal(is []int) int {
	if all(0, is) {
		return 0
	} else {
		return is[0] - prevVal(nextRow(is))
	}
}

func PartOne() {
	total := 0
	for _, row := range values {
		total += nextVal(row)
	}
	println("Part One:", total)
}

func PartTwo() {
	total := 0
	for _, row := range values {
		total += prevVal(row)
	}
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
