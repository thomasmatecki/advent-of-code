package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data, _ = os.ReadFile("03.txt")
var partNoExpr = regexp.MustCompile(`(\d+)`)
var gearNoExpr = regexp.MustCompile(`(\*)`)

type Schematic struct {
	rows                []string
	rowWidth            int
	partNumMatchesByRow [][][]int
}

func NewSchematic(lines []byte) (schematic Schematic) {
	schematic.rows = strings.Split(string(lines), "\n")
	schematic.rowWidth = len(schematic.rows[0])
	schematic.partNumMatchesByRow = make([][][]int, len(schematic.rows))
	for i, row := range schematic.rows {
		schematic.partNumMatchesByRow[i] = partNoExpr.FindAllStringSubmatchIndex(row, -1)
	}
	return
}

func (schematic Schematic) isSymbol(row, col int) bool {
	if row < 0 || col < 0 || row >= len(schematic.rows)-1 || col > schematic.rowWidth-1 {
		return false
	}

	return schematic.rows[row][col] != '.'
}

func (schematic Schematic) adjacentToSymbol(row, from, to int) bool {
	for col := from - 1; col <= to; col++ {
		if schematic.isSymbol(row-1, col) || schematic.isSymbol(row+1, col) {
			return true
		}
	}
	if schematic.isSymbol(row, from-1) || schematic.isSymbol(row, to) {
		return true
	}
	return false
}

func (schematic Schematic) SumPartNums() (total int) {

	for row, line := range schematic.rows {
		if len(line) == 0 {
			continue
		}

		matches := partNoExpr.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			from, to := match[0], match[1]
			if schematic.adjacentToSymbol(row, from, to) {
				val, _ := strconv.Atoi(line[from:to])
				total += val
			}
		}

	}
	return

}

func (schematic Schematic) adjacentNums(row, col int) []int {
	return []int{}
}

func (schematic Schematic) ProductGearRatios() (total int) {

	for row, line := range schematic.rows {
		if len(line) == 0 {
			continue
		}
		matches := gearNoExpr.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			col := match[0]
			if nums := schematic.adjacentNums(row, col); len(nums) == 2 {
				// TODO
			}
		}

	}
	return

}

func PartOne() {
	schematic := NewSchematic(data)
	total := schematic.SumPartNums()
	println("Part One:", total)
}

func PartTwo() {

	total := 0

	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
