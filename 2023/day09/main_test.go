package main

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

var testData, _ = os.ReadFile("test01.txt")
var testValues [][]int

func init() {
	testLines := bytes.Split(testData, []byte("\n"))
	testValues = make([][]int, len(testLines))
	for idx, line := range testLines {
		for _, valueStr := range bytes.Split(line, []byte(" ")) {
			value, _ := strconv.Atoi(string(valueStr))
			testValues[idx] = append(testValues[idx], value)
		}
	}
}
func TestNextVal(t *testing.T) {
	if nextVal(testValues[0]) != 18 {
		t.Fatalf(`Fail`)
	}

	if nextVal(testValues[1]) != 28 {
		t.Fatalf(`Fail`)
	}

	if nextVal(testValues[2]) != 68 {
		t.Fatalf(`Fail`)
	}

}
