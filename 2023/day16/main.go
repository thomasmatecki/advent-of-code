package main

import (
	"bytes"
	"os"
	"slices"
)

type Tile struct {
	i, j int
}

type Grid [][]byte

type Direction int

const (
	North Direction = iota
	East  Direction = iota
	South Direction = iota
	West  Direction = iota
)

type Beam struct {
	direction Direction
	tile      Tile
	grid      *Grid
}

func (beam Beam) TileVal() byte {
	return (*beam.grid)[beam.tile.i][beam.tile.j]
}

func (beam Beam) Step() []Beam {
	switch beam.direction {
	case East:
		beam.tile.j += 1
	case West:
		beam.tile.j -= 1
	case North:
		beam.tile.i -= 1
	case South:
		beam.tile.i += 1
	}

	if beam.tile.i < 0 ||
		beam.tile.i >= len(*beam.grid) ||
		beam.tile.j < 0 ||
		beam.tile.j >= len((*beam.grid)[0]) {
		return []Beam{}
	}

	tileVal := beam.TileVal()
	switch tileVal {

	case '-':
		if beam.direction == North || beam.direction == South {
			return []Beam{
				{East, Tile{beam.tile.i, beam.tile.j}, beam.grid},
				{West, Tile{beam.tile.i, beam.tile.j}, beam.grid},
			}
		}
	case '|':
		if beam.direction == East || beam.direction == West {
			return []Beam{
				{North, Tile{beam.tile.i, beam.tile.j}, beam.grid},
				{South, Tile{beam.tile.i, beam.tile.j}, beam.grid},
			}
		}
	case '\\':
		switch beam.direction {
		case East:
			beam.direction = South
		case North:
			beam.direction = West
		case South:
			beam.direction = East
		case West:
			beam.direction = North
		}
	case '/':
		switch beam.direction {
		case East:
			beam.direction = North
		case North:
			beam.direction = East
		case South:
			beam.direction = West
		case West:
			beam.direction = South
		}
	case '.':
	}
	return []Beam{beam}
}

func CountEnergized(grid *Grid, initial Beam) (energizedTiles int) {
	visited := make(map[Tile][]Direction)
	beams := []Beam{initial}

	for len(beams) > 0 {
		beam := beams[0]
		nextSteps := beam.Step()
		unVisitedNextSteps := []Beam{}

		for _, beam := range nextSteps {
			directions := (visited)[beam.tile]
			if !slices.Contains(directions, beam.direction) {
				(visited)[beam.tile] = append((visited)[beam.tile], beam.direction)
				unVisitedNextSteps = append(unVisitedNextSteps, beam)
			}
		}

		beams = append(beams[1:], unVisitedNextSteps...)
	}

	return len(visited)
}

func PartOne() {
	input, _ := os.ReadFile("16.txt")
	var grid Grid = bytes.Split(input, []byte("\n"))
	total := CountEnergized(&grid, Beam{East, Tile{0, -1}, &grid})
	println("Part One:", total)
}

func PartTwo() {
	input, _ := os.ReadFile("16.txt")
	var grid Grid = bytes.Split(input, []byte("\n"))
	width := len(grid[0])
	length := len(grid[0])
	maxEnergized := 0

	for i := 0; i < width; i++ {
		maxEnergized = max(maxEnergized, CountEnergized(&grid, Beam{East, Tile{i, -1}, &grid}))
		maxEnergized = max(maxEnergized, CountEnergized(&grid, Beam{West, Tile{i, width}, &grid}))
	}
	for j := 0; j < length; j++ {
		maxEnergized = max(maxEnergized, CountEnergized(&grid, Beam{South, Tile{-1, j}, &grid}))
		maxEnergized = max(maxEnergized, CountEnergized(&grid, Beam{North, Tile{length, j}, &grid}))
	}
	println("Part Two:", maxEnergized)
}

func main() {
	PartOne()
	PartTwo()
}
