package main

import (
	"os"
	"testing"
)

func TestExample01(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")

	schematic := NewSchematic(data)

	if schematic.SumPartNums() != 4361 {
		t.Fatalf(`Fail`)
	}
}
