package main

import (
	"bytes"
	"os"
)

var data, _ = os.ReadFile("10.txt")
var grid Grid

const newLine = "\n"

type Pos struct {
	i, j int
}

type Grid [][]byte

func init() {
	grid = bytes.Split(data, []byte(newLine))
}

func (grid *Grid) Start() Pos {
	for i, row := range *grid {
		for j, val := range row {
			if val == 'S' {
				return Pos{i, j}
			}
		}
	}
	panic("No Start?")
}

func (grid *Grid) valIs(i, j int, vals ...byte) bool {
	pVal := (*grid)[i][j]
	for _, val := range vals {
		if pVal == val {
			return true
		}
	}
	return false
}

func (grid *Grid) First(p Pos) Pos {
	// TODO: There are lots of boundary positions where this doesn't work.
	if p.i != 0 && grid.valIs(p.i-1, p.j, '|', '7', 'F') {
		return Pos{p.i - 1, p.j}
	}
	if grid.valIs(p.i+1, p.j, '|', 'J', 'L') {
		return Pos{p.i + 1, p.j}
	}
	if grid.valIs(p.i, p.j-1, '-', 'L', 'F') {
		return Pos{p.i, p.j - 1}
	}
	if grid.valIs(p.i, p.j+1, '-', '7', 'J') {
		return Pos{p.i, p.j + 1}
	}
	panic("No First?")
}

func (grid *Grid) Next(c Pos, p Pos) Pos {
	/*
	 * | is a vertical pipe connecting north and south.
	 * - is a horizontal pipe connecting east and west.
	 * L is a 90-degree bend connecting north and east.
	 * J is a 90-degree bend connecting north and west.
	 * 7 is a 90-degree bend connecting south and west.
	 * F is a 90-degree bend connecting south and east.
	 */
	if grid.valIs(c.i, c.j, '|') {
		return Pos{c.i + (c.i - p.i), c.j}
	}

	if grid.valIs(c.i, c.j, '-') {
		return Pos{c.i, c.j + (c.j - p.j)}
	}

	if c.i == p.i {
		if grid.valIs(c.i, c.j, 'L', 'J') {
			return Pos{c.i - 1, c.j}
		}
		if grid.valIs(c.i, c.j, '7', 'F') {
			return Pos{c.i + 1, c.j}
		}
	} else if c.j == p.j {
		if grid.valIs(c.i, c.j, 'L', 'F') {
			return Pos{c.i, c.j + 1}
		}
		if grid.valIs(c.i, c.j, 'J', '7') {
			return Pos{c.i, c.j - 1}
		}
	}

	panic("No Next?")
}

func (grid *Grid) LoopPointSet() (pointSet map[Pos]bool) {
	pointSet = make(map[Pos]bool)
	start := grid.Start()
	pointSet[start] = true
	curr := grid.First(start)
	pointSet[curr] = true
	prev := start

	for curr != start {
		next := grid.Next(curr, prev)
		prev = curr
		curr = next
		pointSet[curr] = true
	}
	return
}

func PartOne() {
	count := len(grid.LoopPointSet())
	println("Part One:", count/2)
}

func OddIntersections(grid *Grid, loopPoints *map[Pos]bool, p Pos) bool {
	// The Even-Odd Rule: https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule
	count := 0
	for p.i > 0 {
		p.i--
		if !(*loopPoints)[p] {
			continue
		}
		if grid.valIs(p.i, p.j, '-') {
			count += 2
		}
		if grid.valIs(p.i, p.j, 'F', 'J') {
			count++
		}
		if grid.valIs(p.i, p.j, '7', 'L') {
			count--
		}
	}
	count /= 2
	if count < 0 {
		count *= -1
	}
	return (count % 2) == 1
}

func (grid *Grid) CountInner() (count int) {

	loopPoints := grid.LoopPointSet()

	for i, row := range *grid {
		for j := range row {
			p := Pos{i, j}
			if !loopPoints[p] && OddIntersections(grid, &loopPoints, p) {
				count++
			}
		}
	}
	return

}

func PartTwo() {
	total := grid.CountInner()
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
