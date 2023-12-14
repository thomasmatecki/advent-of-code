package main

import (
	"os"
	"regexp"
	"sort"
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

func (schematic Schematic) rowNumsOverlap(row, col int) (nums []int) {
	prevRowPartNums := schematic.partNumMatchesByRow[row]
	// Smallest index of a number ending equal or after one before the col.
	colIdx := sort.Search(len(prevRowPartNums), func(i int) bool {
		return prevRowPartNums[i][1] >= col
	})

	if colIdx == len(prevRowPartNums) {
		return
	}

	partNumMatch := prevRowPartNums[colIdx]

	if partNumMatch[0] <= col+1 {
		str := schematic.rows[row][partNumMatch[0]:partNumMatch[1]]
		val, _ := strconv.Atoi(str)
		nums = append(nums, val)
	}

	if nextColIdx := colIdx + 1; nextColIdx < len(prevRowPartNums) {
		partNumMatch = prevRowPartNums[nextColIdx]

		if partNumMatch[0] <= col+1 && partNumMatch[1] >= col {
			str := schematic.rows[row][partNumMatch[0]:partNumMatch[1]]
			val, _ := strconv.Atoi(str)
			nums = append(nums, val)
		}

	}

	return
}

func (schematic Schematic) adjacentNums(row, col int) (nums []int) {
	if row > 0 {
		nums = append(nums, schematic.rowNumsOverlap(row-1, col)...)
	}
	nums = append(nums, schematic.rowNumsOverlap(row, col)...)
	if row+1 < len(schematic.rows) {
		nums = append(nums, schematic.rowNumsOverlap(row+1, col)...)
	}
	return
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
				total += nums[0] * nums[1]
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
	schematic := NewSchematic(data)
	total := schematic.ProductGearRatios()

	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
