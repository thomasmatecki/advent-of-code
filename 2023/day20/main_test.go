package main

import (
	"testing"
)

func TestExecute01(t *testing.T) {
	broadcaster := InitBroadcaster("test01.txt")
	for i := 0; i < 1000; i++ {
		broadcaster.Broadcast(false)
	}
	product := broadcaster.HiLoProduct()

	if product != 32000000 {
		t.Fatalf("Fail")
	}
}

func TestExecute02(t *testing.T) {
	broadcaster := InitBroadcaster("test02.txt")

	for i := 0; i < 1000; i++ {
		broadcaster.Broadcast(false)
	}
	product := broadcaster.HiLoProduct()
	if product != 11687500 {
		t.Fatalf("Fail")
	}
}
