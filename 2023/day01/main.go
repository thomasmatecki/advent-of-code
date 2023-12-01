package main

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data, _ = os.ReadFile("01.txt")

var digitExpr = regexp.MustCompile(`(\d)`)

func FirstLast(chars []byte) (byte, byte) {
	var first, last byte

	for i := 0; ; i++ {
		if digitExpr.Match(chars[i : i+1]) {
			first = chars[i]
			break
		}
	}

	for i := len(chars); ; i-- {
		if digitExpr.Match(chars[i-1 : i]) {
			last = chars[i-1]
			break
		}
	}

	return first, last
}

func PartOne() {
	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		chars := []byte(line)
		if len(chars) == 0 {
			break
		}
		first, last := FirstLast(chars)
		n, _ := strconv.Atoi(string([]byte{first, last}))
		total += n
	}

	println("Part One:", total)
}

var numstrs = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func matchWord(chars []byte, i int) byte {
	for idx, word := range numstrs {
		j := i + len(word)
		if j <= len(chars) {
			substr := chars[i:j]
			cmpstr := []byte(word)
			if bytes.Equal(substr, cmpstr) {
				return byte(idx + 49)
			}
		}
	}

	return 0x0
}

func FirstLastDigitOrWord(chars []byte) (byte, byte) {
	var first, last byte

	for i := 0; ; i++ {
		if digitExpr.Match(chars[i : i+1]) {
			first = chars[i]
			break
		}

		if char := matchWord(chars, i); char > 0x0 {
			first = char
			break
		}

	}

	for i := len(chars); ; i-- {
		if digitExpr.Match(chars[i-1 : i]) {
			last = chars[i-1]
			break
		}

		if char := matchWord(chars, i-1); char > 0x0 {
			last = char
			break
		}
	}

	return first, last
}

func PartTwo() {

	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		chars := []byte(line)
		if len(chars) == 0 {
			break
		}
		first, last := FirstLastDigitOrWord(chars)

		n, _ := strconv.Atoi(string([]byte{first, last}))
		total += n
	}
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
