package main

import (
	"bytes"
	"os"
	"testing"
)

func setup(filename string) (pointSet map[Pos]bool, testGrid Grid) {
	var testData, _ = os.ReadFile(filename)
	testGrid = bytes.Split(testData, []byte(newLine))
	pointSet = testGrid.LoopPointSet()
	return
}
func Test01OddIntersections(t *testing.T) {
	pointSet, testGrid := setup("test01.txt")
	if !OddIntersections(&testGrid, &pointSet, Pos{2, 2}) {
		t.Fatalf(`Fail`)
	}
}
func Test04OddIntersections(t *testing.T) {
	pointSet, testGrid := setup("test04.txt")
	if !OddIntersections(&testGrid, &pointSet, Pos{6, 2}) {
		t.Fatalf(`Fail`)
	}

	if !OddIntersections(&testGrid, &pointSet, Pos{6, 3}) {
		t.Fatalf(`Fail`)
	}

	if !OddIntersections(&testGrid, &pointSet, Pos{6, 7}) {
		t.Fatalf(`Fail`)
	}

	if !OddIntersections(&testGrid, &pointSet, Pos{6, 8}) {
		t.Fatalf(`Fail`)
	}

	if OddIntersections(&testGrid, &pointSet, Pos{4, 3}) {
		t.Fatalf(`Fail`)
	}

	if OddIntersections(&testGrid, &pointSet, Pos{3, 3}) {
		t.Fatalf(`Fail`)
	}

	if OddIntersections(&testGrid, &pointSet, Pos{3, 4}) {
		t.Fatalf(`Fail`)
	}

	if OddIntersections(&testGrid, &pointSet, Pos{8, 3}) {
		t.Fatalf(`Fail`)
	}
}

func Test03CountInner(t *testing.T) {
	_, testGrid := setup("test03.txt")
	innerCount := testGrid.CountInner()
	if innerCount != 10 {
		t.Fatalf(`Fail`)
	}
}

func Test04CountInner(t *testing.T) {
	_, testGrid := setup("test04.txt")
	if testGrid.CountInner() != 4 {
		t.Fatalf(`Fail`)
	}
}

func Test05OddIntersections(t *testing.T) {
	pointSet, testGrid := setup("test05.txt")
	if !OddIntersections(&testGrid, &pointSet, Pos{4, 8}) {
		t.Fatalf(`Fail`)
	}

	if OddIntersections(&testGrid, &pointSet, Pos{5, 1}) {
		t.Fatalf(`Fail`)
	}
}

func Test05CountInner(t *testing.T) {
	_, testGrid := setup("test05.txt")
	innerCount := testGrid.CountInner()
	if innerCount != 8 {
		t.Fatalf(`Fail`)
	}
}
