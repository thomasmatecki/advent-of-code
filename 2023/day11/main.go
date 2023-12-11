package main

import (
	"bytes"
	"os"
)

type Pos struct {
	i, j int
}

var input, _ = os.ReadFile("11.txt")
var newLine = []byte("\n")
var image = bytes.Split(input, newLine)

type Universe struct {
	image                    *[][]byte
	rowMeasures, colMeasures []int
}

func (universe *Universe) SumDistances() int {

	visited := make(map[Pos]bool)
	total := 0
	for i, row := range *universe.image {
		for j := range row {
			if (*universe.image)[i][j] == '#' {
				g0 := Pos{i, j}
				for g1 := range visited {
					total += universe.Distance(g0, g1)
				}
				visited[g0] = true
			}
		}
	}
	return total
}

func (universe *Universe) Distance(p0, p1 Pos) int {
	return scalarDistance(p0.i, p1.i, &universe.rowMeasures) + scalarDistance(p0.j, p1.j, &universe.colMeasures)
}

func scalarDistance(k0, k1 int, measures *[]int) (d int) {
	var fr, to int
	if k1 > k0 {
		fr, to = k0, k1
	} else {
		fr, to = k1, k0
	}
	for k := fr; k < to; k++ {
		d += (*measures)[k]
	}
	return
}

func InitUniverse(image *[][]byte, expansionScale int) *Universe {
	u := new(Universe)

	u.rowMeasures = make([]int, len(*image))
	u.colMeasures = make([]int, len((*image)[0]))
	u.image = image

	for i := range u.rowMeasures {
		u.rowMeasures[i] = expansionScale
	}

	for j := range u.colMeasures {
		u.colMeasures[j] = expansionScale
	}

	for i, row := range *u.image {
		for j := range row {
			if (*u.image)[i][j] == '#' {
				u.rowMeasures[i] = 1
				u.colMeasures[j] = 1
			}
		}
	}

	return u
}

func PartOne() {
	universe := InitUniverse(&image, 2)
	println("Part One", universe.SumDistances())
}

func PartTwo() {
	universe := InitUniverse(&image, 1000000)
	println("Part Two", universe.SumDistances())
}

func main() {
	PartOne()
	PartTwo()
}
