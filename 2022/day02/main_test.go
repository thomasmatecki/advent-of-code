package main

import (
	"testing"
)

/*
*
A X Rock
B Y Paper
C Z Scissors
*/
func TestScoreWin(t *testing.T) {
	if score('A', 'Y') != 8 {
		t.Fatalf(`Fail`)
	}

	if score('B', 'Z') != 9 {
		t.Fatalf(`Fail`)
	}

	if score('C', 'X') != 7 {
		t.Fatalf(`Fail`)
	}
}

func TestScoreLose(t *testing.T) {
	if score('B', 'X') != 1 {
		t.Fatalf(`Fail`)
	}
}

func TestScoreDraw(t *testing.T) {
	if score('A', 'X') != 4 {
		t.Fatalf(`Fail`)
	}

	if score('B', 'Y') != 5 {
		t.Fatalf(`Fail`)
	}

	if score('C', 'Z') != 6 {
		t.Fatalf(`Fail`)
	}

}
