package main

import (
	"os"
	"strings"
)

func priority(c byte) int {
	if 'a' <= c && c <= 'z' {
		return int(c - 96)
	} else if 'A' <= c && c <= 'Z' {
		return int(c - 38)
	}
	panic("No!")
}

func sharedCharPriority(line string) int {
	lineLen := len(line)
	if lineLen == 0 {
		return 0
	}

	var charMap = make(map[byte]bool, 0)
	chars := []byte(line)

	for _, char := range chars[:lineLen/2] {
		charMap[char] = true
	}

	for _, char := range chars[lineLen/2:] {
		if charMap[char] {
			return priority(char)
		}
	}
	panic("No!")
}

func PartOne() {
	data, _ := os.ReadFile("03.txt")
	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		total += sharedCharPriority(line)
	}

	println("Part One:", total)
}

func maxOccurring(lines []string) byte {
	charCount := make(map[byte]int, 0)
	for _, line := range lines {
		chars := []byte(line)

		charMap := make(map[byte]bool, 0)

		for _, char := range chars {
			charMap[char] = true
		}

		for char := range charMap {
			charCount[char] += 1
			if charCount[char] == 3 {
				return char
			}
		}

	}

	panic("No!")

}

func PartTwo() {
	data, _ := os.ReadFile("03.txt")
	lines := strings.Split(string(data), "\n")
	total := 0

	for i := 0; i < len(lines)-1; i += 3 {
		chunk := lines[i : i+3]
		total += priority(maxOccurring(chunk))
	}
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
