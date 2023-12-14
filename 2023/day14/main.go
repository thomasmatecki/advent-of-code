package main

import (
	"bytes"
	"os"
	"strings"
)

var newLine = []byte{'\n'}

type Rocks [][]byte

func (rocks *Rocks) Tilt(
	oStart, oEnd, oStep int,
	nStart, nEnd, nStep int,
	transposed bool,
) {
	for o := oStart; o < oEnd; o += oStep {
		l := nStart - nStep
		for n := nStart; n != nEnd; n += nStep {
			var i, j int

			if transposed {
				i, j = o, n
			} else {
				j, i = o, n
			}

			if (*rocks)[i][j] == '#' {
				l = n
			}

			if (*rocks)[i][j] == 'O' {
				l += nStep
				if transposed {
					(*rocks)[i][l] = 'O'
				} else {
					(*rocks)[l][j] = 'O'
				}
				if l != n {
					(*rocks)[i][j] = '.'
				}
			}
		}
	}
}

func (rocks *Rocks) TiltNorth() {
	rocks.Tilt(
		0, len((*rocks)[0]), 1,
		0, len(*rocks), 1,
		false,
	)
}

func (rocks *Rocks) TiltSouth() {
	rocks.Tilt(
		0, len((*rocks)[0]), 1,
		len(*rocks)-1, -1, -1,
		false,
	)
}

func (rocks *Rocks) TiltWest() {
	rocks.Tilt(
		0, len(*rocks), 1,
		0, len((*rocks)[0]), 1,
		true,
	)
}

func (rocks *Rocks) TiltEast() {
	rocks.Tilt(
		0, len(*rocks), 1,
		len((*rocks)[0])-1, -1, -1,
		true,
	)
}
func (rocks *Rocks) Spin() {
	rocks.TiltNorth()
	rocks.TiltWest()
	rocks.TiltSouth()
	rocks.TiltEast()
}

func (rocks *Rocks) NorthLoad() (load int) {
	for i := 0; i < len(*rocks); i++ {
		for _, val := range (*rocks)[i] {
			if val == 'O' {
				load += len(*rocks) - i
			}
		}
	}
	return
}

func (rocks *Rocks) String() string {
	strs := make([]string, len(*rocks))

	for i := 0; i < len(*rocks); i++ {
		strs[i] = string((*rocks)[i])
	}
	return strings.Join(strs, "\n")
}

func Input(filename string) Rocks {
	input, _ := os.ReadFile(filename)
	return bytes.Split(bytes.Trim(input, string(newLine)), newLine)
}

func PartOne() {
	rocks := Input("14.txt")
	rocks.TiltNorth()
	println("Part One:", rocks.NorthLoad())
}

const CYCLES = 1000000000

func PartTwo() {
	rocks := Input("14.txt")
	states := make(map[string]int)
	for i := 0; i < CYCLES; i++ {
		rocks.Spin()
		if state := states[rocks.String()]; state > 0 {
			remaining := CYCLES - i
			modRemaining := remaining % (i - state)
			i = CYCLES - modRemaining
		}
		states[rocks.String()] = i
	}
	println("Part Two:", rocks.NorthLoad())
}

func main() {
	PartOne()
	PartTwo()
}
