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

	if !broadcaster.network.IsZeroed() {
		t.Fatalf("Not Zeroed")
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

func TestExecute04(t *testing.T) {
	broadcaster := InitBroadcaster("test04.txt")

	for i := 0; i < 20000; i++ {
		broadcaster.Broadcast(false)
		if !broadcaster.network.Get("mf").IsZeroed() {
			break
		}
	}
}
