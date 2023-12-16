package main

import (
	"testing"
)

func AssertHashEq(t *testing.T, input string, expected int) {
	if Hash([]byte(input)) != expected {
		t.Fatalf(`Fail: Hash(%s) != %d`, input, expected)
	}
}

func TestHash01(t *testing.T) {
	AssertHashEq(t, "rn=1", 30)
	AssertHashEq(t, "cm-", 253)
	AssertHashEq(t, "qp=3", 97)
	AssertHashEq(t, "cm=2", 47)
	AssertHashEq(t, "qp-", 14)
	AssertHashEq(t, "pc=4", 180)
	AssertHashEq(t, "ot=9", 9)
	AssertHashEq(t, "ab=5", 197)
	AssertHashEq(t, "pc-", 48)
	AssertHashEq(t, "pc=6", 214)
	AssertHashEq(t, "ot=7", 231)
	AssertHashEq(t, "rn", 0)
	AssertHashEq(t, "cm", 0)
	AssertHashEq(t, "qp", 1)
	AssertHashEq(t, "ot", 3)
}

func TestSumHashes01(t *testing.T) {
	input := [][]byte{
		[]byte("rn=1"),
		[]byte("cm-"),
		[]byte("qp=3"),
		[]byte("cm=2"),
		[]byte("qp-"),
		[]byte("pc=4"),
		[]byte("ot=9"),
		[]byte("ab=5"),
		[]byte("pc-"),
		[]byte("pc=6"),
		[]byte("ot=7"),
	}

	if SumHashes(input) != 1320 {
		t.Fatalf("fail")
	}
}

func TestArrangeLenses01(t *testing.T) {
	input := [][]byte{
		[]byte("rn=1"),
		[]byte("cm-"),
		[]byte("qp=3"),
		[]byte("cm=2"),
		[]byte("qp-"),
		[]byte("pc=4"),
		[]byte("ot=9"),
		[]byte("ab=5"),
		[]byte("pc-"),
		[]byte("pc=6"),
		[]byte("ot=7"),
	}

	if ArrangeLenses(input) != 145 {
		t.Fatalf("fail")
	}
}
