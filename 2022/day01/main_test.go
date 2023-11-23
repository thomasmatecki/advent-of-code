package main

import (
	"testing"
)

func TestPass(t *testing.T) {
	if 1 == 2 {
		t.Fatalf(`Faile`)
	}
}

func TestAddOne(t *testing.T) {
	if AddOne(2) != 3 {
		t.Fatalf(`Faile`)
	}
}
