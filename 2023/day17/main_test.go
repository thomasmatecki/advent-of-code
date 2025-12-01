package main

import (
	"fmt"
	"testing"
)

func TestStep01(t *testing.T) {
	input := Input("test01.txt")
	fmt.Println(input)

	if 11111111111111 != 2 {
		t.Fatalf("Fail")
	}
}
