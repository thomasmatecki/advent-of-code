package main

import (
	"testing"
)

func TestExecute01(t *testing.T) {
	ratingSum := Exec("test01.txt")
	if ratingSum != 19114 {
		t.Fatalf("Fail")
	}
}

func TestSpace_Split(t *testing.T) {

	space := Space{
		'x': &Rng{0, 4000},
		'm': &Rng{0, 4000},
		'a': &Rng{0, 4000},
		's': &Rng{0, 4000},
	}
	otherSpace := space.Partiion(&Conditional{
		'a', IfGreaterThan, 1000,
	})

	origSize := space.Size()

	if space['a'].lbnd != 1000 {
		t.Fatalf("Fail")
	}

	if space.Size()+otherSpace.Size() != origSize {
		t.Fatalf("Fail")
	}

	if otherSpace['a'].ubnd != 1000 {
		t.Fatalf("Fail")
	}

}

func TestCount(t *testing.T) {
	workflows := make(Lookup)

	workflows.Add([]byte("xx{s>3:R,s<2:R,A}"))

	workflows.Add([]byte("qqv{s<3:A,R}"))
	workflows.Add([]byte("qqw{s<4:A,R}"))
	workflows.Add([]byte("qqx{s<2:A,R}"))
	workflows.Add([]byte("qqr{s<1:A,R}"))

	workflows.Add([]byte("qqy{s>1:A,R}"))
	workflows.Add([]byte("qqt{s>3:A,R}"))
	workflows.Add([]byte("qqu{s>3:R,A}"))

	workflows.Add([]byte("qqz{s>1:R,A}"))

	xx := workflows["xx"]
	xxCount := xx.Count(NewSpace(4))
	if xxCount != 4*4*4*2 {
		t.Fatalf("Fail")
	}

	qqu := workflows["qqu"]
	qquCount := qqu.Count(NewSpace(4))
	if qquCount != 4*4*4*3 {
		t.Fatalf("Fail")
	}

	qqt := workflows["qqt"]
	qqtCount := qqt.Count(NewSpace(4))
	if qqtCount != 4*4*4 {
		t.Fatalf("Fail")
	}

	qqr := workflows["qqr"]
	qqrCount := qqr.Count(NewSpace(4))
	if qqrCount != 0 {
		t.Fatalf("Fail")
	}

	qqv := workflows["qqv"]
	qqvCount := qqv.Count(NewSpace(4))
	if qqvCount != 4*4*4*2 {
		t.Fatalf("Fail")
	}

	qqw := workflows["qqw"]
	qqwCount := qqw.Count(NewSpace(4))
	if qqwCount != 4*4*4*3 {
		t.Fatalf("Fail")
	}

	qqx := workflows["qqx"]
	qqxCount := qqx.Count(NewSpace(4))

	if qqxCount != 4*4*4 {
		t.Fatalf("Fail")
	}

	qqy := workflows["qqy"]
	qqyCount := qqy.Count(NewSpace(4))

	if qqyCount != 4*4*4*3 {
		t.Fatalf("Fail")
	}

	qqz := workflows["qqz"]
	qqzCount := qqz.Count(NewSpace(4))
	if qqzCount != 4*4*4 {
		t.Fatalf("Fail")
	}
}

func TestExplore(t *testing.T) {
	acceptedRatings := Explore("test01.txt")
	if acceptedRatings != 167409079868000 {
		t.Fatalf("Fail")
	}
}
