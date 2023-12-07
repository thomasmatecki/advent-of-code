package main

import (
	"os"
	"strconv"
	"strings"
)

var data, _ = os.ReadFile("04.txt")

func toNumList(s string, l int) []int {
	nums := strings.Split(s, " ")
	numi := make([]int, l)
	j := 0
	for _, c := range nums {
		num, err := strconv.Atoi(c)
		if err == nil {
			numi[j] = num
			j += 1
		}
	}

	return numi
}

func toNumSet(s string, l int) map[int]bool {
	nums := strings.Split(s, " ")
	numi := make(map[int]bool, l)
	for _, c := range nums {
		num, err := strconv.Atoi(c)
		if err == nil {
			numi[num] = true
		}
	}

	return numi
}

func CountMatches(lines []string, leadLen int) []int {
	matches := make([]int, len(lines)-1)

	for idx, line := range lines {
		if len(line) == 0 {
			continue
		}

		sliced := strings.Split(line[leadLen:], " | ")
		winning, yours := sliced[0], sliced[1]

		yourNums := toNumList(yours, 25)
		winningNums := toNumSet(winning, 10)

		for _, num := range yourNums {
			if winningNums[num] {
				matches[idx] += 1
			}
		}
	}

	return matches
}

func AggregateCounts(matches []int) (total int) {
	counts := make([]int, len(matches))

	for i := 0; i < len(counts); i++ {
		counts[i] = 1
	}

	for i := 0; i < len(counts); i++ {
		for j := 1; j <= matches[i]; j++ {
			counts[i+j] += counts[i]
		}
	}

	for i := 0; i < len(counts); i++ {
		total += counts[i]
	}

	return
}

func PartOne() {
	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		sliced := strings.Split(line[10:], " | ")
		winning, yours := sliced[0], sliced[1]
		yourNums := toNumList(yours, 25)
		winningNums := toNumSet(winning, 10)
		subtotal := 0

		for _, num := range yourNums {
			if winningNums[num] {
				if subtotal == 0 {
					subtotal = 1
				} else {
					subtotal *= 2
				}

			}
		}
		total += subtotal
	}

	println("Part One:", total)
}
func PartTwo() {
	lines := strings.Split(string(data), "\n")
	matches := CountMatches(lines, 10)
	println("Part Two:", AggregateCounts(matches))
}

func main() {
	PartOne()
	PartTwo()
}
