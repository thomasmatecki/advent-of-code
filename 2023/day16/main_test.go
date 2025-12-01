package main

import (
	"bytes"
	"os"
	"testing"
)

func TestStep01(t *testing.T) {
	input, _ := os.ReadFile("test01.txt")
	var grid Grid = bytes.Split(input, []byte("\n"))
	beam := Beam{East, Tile{0, 0}, &grid}
	nextBeams := beam.Step()
	if len(nextBeams) != 2 {
		t.Fatalf("Fail")
	}
}

func TestCountEnergized(t *testing.T) {
	input, _ := os.ReadFile("test01.txt")
	var grid Grid = bytes.Split(input, []byte("\n"))
	energized := CountEnergized(&grid)
	if energized != 46 {
		t.Fatalf("Fail")
	}
}
