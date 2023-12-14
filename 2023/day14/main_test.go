package main

import (
	"testing"
)

func TestTiltN(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Tilt(
		0, len(rocks[0]), 1,
		0, len(rocks), 1,
		false,
	)
	for _, r := range rocks {
		println(string(r))
	}
}

func TestTiltS(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Tilt(
		0, len(rocks[0]), 1,
		len(rocks)-1, -1, -1,
		false,
	)
	for _, r := range rocks {
		println(string(r))
	}
}

func TestTiltW(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Tilt(
		0, len(rocks), 1,
		0, len(rocks[0]), 1,
		true,
	)
	for _, r := range rocks {
		println(string(r))
	}
}

func TestTiltE(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Tilt(
		0, len(rocks), 1,
		len(rocks[0])-1, -1, -1,
		true,
	)
	for _, r := range rocks {
		println(string(r))
	}
}
func TestSpin01(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Spin()
	for _, r := range rocks {
		println(string(r))
	}
}
func TestSpin02(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Spin()
	rocks.Spin()
	for _, r := range rocks {
		println(string(r))
	}
}

func TestSpin03(t *testing.T) {
	rocks := Input("test01.txt")
	rocks.Spin()
	rocks.Spin()
	rocks.Spin()
	for _, r := range rocks {
		println(string(r))
	}
}
