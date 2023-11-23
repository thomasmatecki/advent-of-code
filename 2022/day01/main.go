package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func PartOne() {
	data, _ := os.ReadFile("01.txt")
	var max_total = 0
	for _, group := range strings.Split(string(data), "\n\n") {
		var total = 0
		for _, item_cals := range strings.Split(group, "\n") {
			item_cali, _ := strconv.Atoi(item_cals)
			total += item_cali
		}
		max_total = max(total, max_total)
	}
	println("Part One:", max_total)
}

func PartTwo() {
	data, _ := os.ReadFile("01.txt")
	max_totals := []int{0, 0, 0}
	for _, group := range strings.Split(string(data), "\n\n") {
		var total = 0
		for _, item_cals := range strings.Split(group, "\n") {
			item_cali, _ := strconv.Atoi(item_cals)
			total += item_cali
		}

		max_totals[0] = max(total, max_totals[0])
		slices.Sort(max_totals)
	}
	println("Part Two:", max_totals[0]+max_totals[1]+max_totals[2])
}

func main() {
	PartOne()
	PartTwo()
}
