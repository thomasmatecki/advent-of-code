package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
var data, _ = os.ReadFile("04.txt")

func parse(line string) [4]int {
	matches := r.FindStringSubmatch(line)
	var vs [4]int
	for i := 0; i < 4; i++ {
		vs[i], _ = strconv.Atoi(matches[i+1])
	}
	return vs
}

func encloses(vs [4]int) bool {
	return (vs[2] <= vs[0] && vs[1] <= vs[3]) ||
		(vs[2] >= vs[0] && vs[1] >= vs[3])
}

func PartOne() {
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			break
		}
		vs := parse(line)
		if encloses(vs) {
			total += 1
		}
	}

	println("Part One:", total)
}

func overlaps(vs [4]int) bool {
	iv0 := vs[:2]
	iv1 := vs[2:]
	return (iv0[0] <= iv1[0] && iv1[0] <= iv0[1]) ||
		(iv0[0] <= iv1[1] && iv1[1] <= iv0[1]) ||
		(iv1[0] <= iv0[0] && iv0[0] <= iv1[1]) ||
		(iv1[0] <= iv0[1] && iv0[1] <= iv1[1])
}

func PartTwo() {
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			break
		}

		vs := parse(line)
		if overlaps(vs) {
			total += 1
		}
	}

	println("Part Two:", total)

}

func main() {
	PartOne()
	PartTwo()
}
