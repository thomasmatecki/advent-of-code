package main

import (
	"bytes"
	"os"
)

func ReflectionDiffs(b0, b1 *[][]byte) (diffs int) {
	if len(*b1) != len(*b0) {
		panic("lens don't match!")
	}
	l := len(*b1)
	for i, row := range *b0 {
		for j, val := range row {
			if val != (*b1)[l-i-1][j] {
				diffs += 1
			}
		}
	}
	return

}

func FindReflection(pattern [][]byte, smudges int) int {
	for i := 1; i < len(pattern); i++ {
		l := min(i, len(pattern)-i)
		bottom := pattern[i-l : i]
		top := pattern[i : i+l]

		if ReflectionDiffs(&top, &bottom) == smudges {
			return i
		}
	}
	return -1
}

func Transpose(pattern [][]byte) (transposed [][]byte) {
	rows := len(pattern[0])
	cols := len(pattern)
	transposed = make([][]byte, rows)
	for i := range transposed {
		transposed[i] = make([]byte, cols)
	}

	for i, row := range pattern {
		for j, val := range row {
			transposed[j][i] = val
		}
	}

	return
}

func CountReflectionAxis(filename string, smudges int) (total int) {
	input, _ := os.ReadFile(filename)
	patternStrs := bytes.Split(input, []byte("\n\n"))
	for _, patternStr := range patternStrs {
		patternStr := bytes.Trim(patternStr, "\n")
		pattern := bytes.Split(patternStr, []byte("\n"))
		if hAxis := FindReflection(pattern, smudges); hAxis > 0 {
			total += (hAxis * 100)
		} else {
			transposed := Transpose(pattern)
			vAxis := FindReflection(transposed, smudges)
			if vAxis < 0 {
				panic("no!")
			}
			total += vAxis
		}
	}
	return
}

func PartOne() {
	total := CountReflectionAxis("13.txt", 0)
	println("Part One:", total)
}

func PartTwo() {
	total := CountReflectionAxis("13.txt", 1)
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
