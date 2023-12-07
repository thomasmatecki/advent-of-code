package main

import (
	"os"
	"strings"
	"testing"
)

func TestSrcIndex(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")

	rangeMaps := initRangeMaps(strings.Split(string(data), "\n"))

	if rangeMaps[2].sourceIdx(0) != 0 {
		t.Fatalf(`Fail`)
	}

	if rangeMaps[2].sourceIdx(2) != 0 {
		t.Fatalf(`Fail`)
	}
	if rangeMaps[2].sourceIdx(8) != 1 {
		t.Fatalf(`Fail`)
	}

	if rangeMaps[2].sourceIdx(12) != 2 {
		t.Fatalf(`Fail`)
	}

}

func TestTranslate(t *testing.T) {
	//var data, _ = os.ReadFile("test01.txt")

	//	rangeMaps := initRangeMaps(strings.Split(string(data), "\n"))

	//	if rangeMaps[2].translate(0) != 42 {
	//		t.Fatalf(`Fail`)
	//	}
	//
	//	if rangeMaps[2].translate(1) != 43 {
	//		t.Fatalf(`Fail`)
	//	}
	//
	// //2 44
	// //3 45
	// //4 46
	// //5 47
	//
	//	if rangeMaps[2].translate(6) != 48 {
	//		t.Fatalf(`Fail`)
	//	}
	//
	//	if rangeMaps[2].translate(7) != 57 {
	//		t.Fatalf(`Fail`)
	//	}
	//
	//	if rangeMaps[2].translate(10) != 60 {
	//		t.Fatalf(`Fail`)
	//	}
	//
	//	if rangeMaps[2].translate(11) != 0 {
	//		t.Fatalf(`Fail`)
	//	}
}

func TestTraverseRangeMaps01(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")
	lines := strings.Split(string(data), "\n")
	rangeMaps := initRangeMaps(lines)
	seeds := []int{55}
	dests := traverseMaps(seeds, &rangeMaps)

	if dests[0] != 86 {
		t.Fatalf(`Fail`)
	}
}

func TestTraverseRangeMaps(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")
	lines := strings.Split(string(data), "\n")
	seeds := initSeeds(lines[0])
	rangeMaps := initRangeMaps(lines)
	expected := []int{82, 43, 86, 35}

	for i, dest := range traverseMaps(seeds, &rangeMaps) {
		if expected[i] != dest {
			t.Fatalf(`Fail`)
		}
	}
}

func TestSeedRanges(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")
	lines := strings.Split(string(data), "\n")
	seedRanges := initSeeds(lines[0])
	rangeMaps := initRangeMaps(lines)
	seeds := initSeedsFromRanges(seedRanges, 1)
	dests := traverseMaps(seeds, &rangeMaps)
	println(dests)

}
