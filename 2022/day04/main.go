package main

import (
	"os"
)

func PartOne() {
	data, _ := os.ReadFile("04.txt")
	total := 0

}

func PartTwo() {
	data, _ := os.ReadFile("04.txt")
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
