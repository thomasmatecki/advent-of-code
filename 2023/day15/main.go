package main

import (
	"bytes"
	"os"
	"regexp"
)

type Lens struct {
	label       string
	focalLength int
}

type Box []Lens

func (box *Box) Insert(l Lens) Box {
	for idx, lens := range *box {
		if l.label == lens.label {
			(*box)[idx] = l
			return *box
		}
	}
	return append(*box, l)
}

func (box *Box) Remove(label string) Box {
	for idx, lens := range *box {
		if label == lens.label {
			return append((*box)[:idx], (*box)[idx+1:]...)
		}
	}
	return *box
}

func Hash(s []byte) (hash int) {
	hash = 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return
}

func SumHashes(steps [][]byte) (total int) {
	for _, step := range steps {
		total += Hash(step)
	}
	return
}

func PartOne() {
	input, _ := os.ReadFile("15.txt")
	total := SumHashes(bytes.Split(input, []byte{','}))
	println("Part One:", total)
}

var insertExpr = regexp.MustCompile(`(\w+)=(\d+)`)
var removeExpr = regexp.MustCompile(`(\w+)-`)

func ArrangeLenses(steps [][]byte) (total int) {
	boxes := new([256]Box)

	for _, step := range steps {
		match := insertExpr.FindSubmatch(step)
		if match != nil {
			idx := Hash(match[1])
			box := boxes[idx]
			lens := Lens{
				label:       string(match[1]),
				focalLength: int(match[2][0]) - 48,
			}
			boxes[idx] = box.Insert(lens)
			continue
		}

		match = removeExpr.FindSubmatch(step)
		if match != nil {
			label := match[1]
			idx := Hash(match[1])
			box := boxes[idx]
			boxes[idx] = box.Remove(string(label))
		}

	}
	for boxIdx, box := range boxes {
		for lensIdx, lens := range box {
			total += ((boxIdx + 1) * (lensIdx + 1) * lens.focalLength)
		}
	}
	return

}

func PartTwo() {
	input, _ := os.ReadFile("15.txt")
	total := ArrangeLenses(bytes.Split(input, []byte{','}))
	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
