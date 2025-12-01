package main

import (
	"bytes"
	"os"
	"strconv"
)

type Grid struct {
	blocks        [][]int
	height, width int
}

type Direction int

const (
	North Direction = iota
	East  Direction = iota
	South Direction = iota
	West  Direction = iota
)

type Pos struct {
	i, j int
}
type Path struct {
	p         Pos
	stepCount int
	tCost     int
	d         Direction
}

type Edge struct {
	i, j int
	d    Direction
}

func Input(filename string) (g Grid) {
	input, _ := os.ReadFile(filename)

	for idx, row := range bytes.Split(input, []byte{'\n'}) {
		if len(row) == 0 {
			continue
		}
		g.blocks = append(g.blocks, []int{})
		for _, c := range row {
			val, _ := strconv.Atoi(string(c))
			g.blocks[idx] = append(g.blocks[idx], val)
		}
		g.width = len(g.blocks[idx])
	}
	g.height = len(g.blocks)

	return
}

func (g *Grid) Neighbors(p Pos) (edges []Edge) {
	if p.i+1 < g.height {
		edges = append(edges, Edge{p.i + 1, p.j, North})
	}
	if p.i-1 >= 0 {
		edges = append(edges, Edge{p.i - 1, p.j, South})
	}
	if p.j+1 < g.width {
		edges = append(edges, Edge{p.i, p.j + 1, East})
	}
	if p.j-1 >= 0 {
		edges = append(edges, Edge{p.i, p.j - 1, West})
	}

	return
}

func PartOne() {
	grid := Input("17.txt")
	visited := make(map[Edge]int)
	q := []Path{
		{Pos{0, 0}, 0, 0, East},
		{Pos{0, 0}, 0, 0, South},
	}

	for len(q) > 0 {
		current := q[0]
		if current.p.i == grid.height-1 && current.p.j == grid.width-1 {
			q = q[1:]
			continue
		}
		newPaths := []Path{}
		for _, edge := range grid.Neighbors(q[0].p) {
			if current.stepCount == 3 && edge.d == current.d {
				continue
			}
			prevCost, found := visited[edge]
			stepCost := grid.blocks[edge.i][edge.j]
			nextCost := current.tCost + stepCost
			var stepCount int
			if edge.d == current.d {
				stepCount = current.stepCount + 1
			} else {
				stepCount = 1
			}

			path := Path{
				p:         Pos{edge.i, edge.j},
				d:         edge.d,
				stepCount: stepCount,
				tCost:     nextCost,
			}
			if !found || nextCost < prevCost {
				visited[edge] = nextCost
				newPaths = append(newPaths, path)
			}
		}
		q = append(q[1:], newPaths...)
	}

	println("Part One:")
}

func PartTwo() {
	println("Part Two:")
}

func main() {
	PartOne()
	PartTwo()
}
