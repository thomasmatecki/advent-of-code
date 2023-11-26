package main

import (
	"os"
	"regexp"
	"strings"
)

func comp(play, beats, loses byte) int {
	if play == beats {
		return 6
	} else if play == loses {
		return 0
	} else {
		return 3
	}
}

func score(opp, self byte) int {

	if self == 'X' {
		return 1 + comp(opp, 'C', 'B')
	} else if self == 'Y' {
		return 2 + comp(opp, 'A', 'C')
	} else {
		return 3 + comp(opp, 'B', 'A')
	}
}

type OppOut struct {
	Opp byte
	Out byte
}

var oppOutcome = map[OppOut]byte{
	{'A', 'Z'}: 'Y',
	{'B', 'Z'}: 'Z',
	{'C', 'Z'}: 'X',

	{'A', 'Y'}: 'X',
	{'B', 'Y'}: 'Y',
	{'C', 'Y'}: 'Z',

	{'A', 'X'}: 'Z',
	{'B', 'X'}: 'X',
	{'C', 'X'}: 'Y',
}

var r = regexp.MustCompile(`(\w) (\w)`)

func PartOne() {
	data, _ := os.ReadFile("02.txt")
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			matches := r.FindStringSubmatch(line)
			total += score(matches[1][0], matches[2][0])
		}
	}

	println("Part One:", total)
}

func PartTwo() {
	data, _ := os.ReadFile("02.txt")
	total := 0
	for _, line := range strings.Split(string(data), "\n") {
		if len(line) > 0 {
			matches := r.FindStringSubmatch(line)
			opp := matches[1][0]
			out := matches[2][0]
			self := oppOutcome[OppOut{opp, out}]
			total += score(opp, self)
		}
	}
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
