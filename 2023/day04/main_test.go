package main

import (
	"os"
	"strings"
	"testing"
)

/*
*
A X Rock
B Y Paper
C Z Scissors
*/
func TestExample01(t *testing.T) {
	var data, _ = os.ReadFile("test01.txt")
	lines := strings.Split(string(data), "\n")
	matches := CountMatches(lines, 8)

	if AggregateCounts(matches) != 30 {
		t.Fatalf(`Fail`)
	}
}
