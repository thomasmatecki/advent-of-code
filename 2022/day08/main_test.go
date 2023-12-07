package main

import (
	"os"
	"strings"
	"testing"
)

func TestExample0(t *testing.T) {

	var data, _ = os.ReadFile("sample.txt")
	var input = strings.Split(string(data), "\n")
	expected := 95437
	total := SumTotals(input)
	if total != expected {
		t.Fatalf(`Fail %d != %d`, total, expected)
	}

}
