package main

import (
	"os"
	"strconv"
	"strings"
)

func Sum(ls []int) (s int) {
	for _, l := range ls {
		s += l
	}
	return
}

func FindMatches(l int, conditionStr string) [][]int {
	target := strings.Repeat("#", l)

	matches := new([][]int)
	for i := l; i <= len(conditionStr); i++ {
		if strings.Replace(conditionStr[i-l:i], "?", "#", -1) != target {
			continue
		}
		if i-l-1 >= 0 && conditionStr[i-l-1] == '#' {
			continue
		}
		if i < len(conditionStr) && conditionStr[i] == '#' {
			continue
		}

		*matches = append(*matches, []int{i - l, i})
	}

	return *matches

}

func RecurCount(conditionStr string, contigLens []int, contigSum int) int {
	if contigSum > len(conditionStr) {
		return 0
	}
	if len(contigLens) == 0 {
		if strings.Contains(conditionStr, "#") {
			return 0
		} else {
			return 1
		}
	}
	var substr string

	total := 0
	matches := FindMatches(contigLens[0], conditionStr)
	for _, match := range matches {
		if strings.Contains(conditionStr[:match[0]], "#") {
			break // As soon we are passed the first #, the matching no longer works
		}

		remainingContigSum := contigSum - match[1]

		if match[1] == len(conditionStr) {
			if len(contigLens) == 1 {
				total += 1
			} else {
				continue
			}
		} else {
			substr = "." + conditionStr[match[1]+1:]
			if len(substr) > remainingContigSum { // Does this ever not happen?
				c := RecurCount(substr, contigLens[1:], remainingContigSum)
				total += c
			}
		}
	}
	return total
}

func ParseRow(row string, reps int) (conditionStr string, contigLens []int) {
	splitRow := strings.Split(row, " ")
	conditionStrVal, contiGroupStr := splitRow[0], splitRow[1]
	vStrs := strings.Split(contiGroupStr, ",")
	contigLensVals := make([]int, len(vStrs))
	for i, vStr := range vStrs {
		contigLensVals[i], _ = strconv.Atoi(string(vStr))
	}

	for i := 0; i < reps; i++ {
		contigLens = append(contigLens, contigLensVals...)
		conditionStr += conditionStrVal
	}
	return
}

func CountFile(filename string, reps int) int {
	var input, _ = os.ReadFile(filename)
	var rows = strings.Split(string(input), "\n")
	total := 0

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		conditionStr, contigLens := ParseRow(row, reps)
		count := RecurCount(conditionStr, contigLens, Sum(contigLens))

		total += count
	}
	return total
}

func PartOne() {
	total := CountFile("12.txt", 1)
	println("Part One:", total)
}

func PartTwo() {
	total := CountFile("12.txt", 5)
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
