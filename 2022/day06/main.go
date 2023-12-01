package main

import "os"

var input, _ = os.ReadFile("06.txt")

func FindStartSequence(data []byte, seqLen int) int {

	chars := map[byte]int8{}

	for idx, char := range data {
		if idx > seqLen-1 {
			dec := data[idx-seqLen]
			if chars[dec] == 1 {
				delete(chars, dec)
			} else {
				chars[dec] -= 1
			}
		}

		chars[char] += 1

		if len(chars) == seqLen {
			return idx + 1
		}
	}
	panic("no!")
}

func PartOne() {
	seqIdx := FindStartSequence(input, 4)
	println("Part One:", seqIdx)
}

func PartTwo() {
	seqIdx := FindStartSequence(input, 14)
	println("Part Two:", seqIdx)
}

func main() {
	PartOne()
	PartTwo()
}
