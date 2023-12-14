package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRecurCount00(t *testing.T) {
	input, _ := os.ReadFile("test01.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	if FindReflection(pattern, 0) != 4 {
		t.Fatalf(`Fail`)
	}
}
func TestRecurCount01(t *testing.T) {
	input, _ := os.ReadFile("test02.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	if FindReflection(pattern, 0) != -1 {
		t.Fatalf(`Fail`)
	}
}

func TestRecurCount02(t *testing.T) {
	input, _ := os.ReadFile("test03.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	if FindReflection(pattern, 0) != 6 {
		t.Fatalf(`Fail`)
	}
}

func TestTranspose(t *testing.T) {
	input, _ := os.ReadFile("test01.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	transposed := Transpose(pattern)
	if len(transposed) != 9 {
		t.Fatalf(`Fail`)
	}

	if len(transposed[0]) != 7 {
		t.Fatalf(`Fail`)
	}
}

func TestTransposedCount00(t *testing.T) {
	input, _ := os.ReadFile("test01.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	transposed := Transpose(pattern)

	if FindReflection(transposed, 0) != -1 {
		t.Fatalf(`Fail`)
	}
}

func TestTransposedCount01(t *testing.T) {
	input, _ := os.ReadFile("test02.txt")
	pattern := bytes.Split(input, []byte{'\n'})
	transposed := Transpose(pattern)

	if FindReflection(transposed, 0) != 5 {
		t.Fatalf(`Fail`)
	}
}
