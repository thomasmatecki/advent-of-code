package main

import (
	"testing"
)

func TestExample0(t *testing.T) {

	seqIdx := FindStartSequence([]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 4)
	if seqIdx != 5 {
		t.Fatalf(`Fail %d != 5`, seqIdx)
	}

}
func TestExample1(t *testing.T) {

	seqIdx := FindStartSequence([]byte("nppdvjthqldpwncqszvftbrmjlhg"), 4)
	if seqIdx != 6 {
		t.Fatalf(`Fail %d != 6`, seqIdx)
	}

}
