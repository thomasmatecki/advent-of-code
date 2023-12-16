package main

import (
	"testing"
)

//func TestRecurCount00(t *testing.T) {
//	if RecurCount("???.###", []int{1, 1, 3}, 5) != 1 {
//		t.Fatalf("Fail")
//	}
//
//	if RecurCount(".#.#", []int{2, 3}, 5) != 0 {
//		t.Fatalf("Fail")
//	}
//
//	if RecurCount("#.#.#", []int{1, 2, 3}, 6) != 0 {
//		t.Fatalf("Fail")
//	}
//	//if RecurCount("?##.#", []int{1, 2, 3}, 6) != 0 {
//	//	t.Fatalf("Fail")
//	//}
//}
//func TestRecurCount01(t *testing.T) {
//	if RecurCount("#.#.###", []int{1, 1, 3}, 5) != 1 {
//		t.Fatalf("Fail")
//	}
//}
//
//func TestRecurCount02(t *testing.T) {
//	if RecurCount(".#...#....###", []int{1, 1, 3}, 5) != 1 {
//		t.Fatalf("Fail")
//	}
//}
//
//func TestRecurCount03(t *testing.T) {
//	count := RecurCount(".??..??...?##.", []int{1, 1, 3}, 5)
//	if count != 4 {
//		t.Fatalf("Fail")
//	}
//}
//func TestRecurCount04(t *testing.T) {
//	count := RecurCount(".??..??...", []int{1, 1}, 2)
//	if count != 4 {
//		t.Fatalf("Fail")
//	}
//}
//
//func TestRecurCount05(t *testing.T) {
//	count := RecurCount("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 11)
//	if count != 1 {
//		t.Fatalf("Fail")
//	}
//
//}
//func TestRecurCount06(t *testing.T) {
//	count := RecurCount("????.#...#...", []int{4, 1, 1}, 6)
//	if count != 1 {
//		t.Fatalf("Fail")
//	}
//
//}
//
//func TestRecurCount07(t *testing.T) {
//	count := RecurCount("????.######..#####.", []int{1, 6, 5}, 12)
//	if count != 4 {
//		t.Fatalf("Fail")
//	}
//
//}
//func TestRecurCount08(t *testing.T) {
//	var count = RecurCount(".???????", []int{2, 1}, 6)
//	if count != 10 {
//		t.Fatalf("Fail")
//	}
//
//	count = RecurCount(".???????", []int{2, 1}, 6)
//	if count != 10 {
//		t.Fatalf("Fail")
//	}
//
//	count = RecurCount(".###.???????", []int{3, 2, 1}, 6)
//	if count != 10 {
//		t.Fatalf("Fail")
//	}
//
//	count = RecurCount(".###????????", []int{3, 2, 1}, 6)
//	if count != 10 {
//		t.Fatalf("Fail")
//	}
//
//}
//
//func TestRecurCount09(t *testing.T) {
//	var count = RecurCount("??????????#????????", []int{1, 3, 2, 1, 1, 2}, 0)
//	if count != 95 {
//		t.Fatalf("Fail")
//	}
//
//}
//
//func TestRecurCount10(t *testing.T) {
//
//	var count = RecurCount(".?.#", []int{1}, 0)
//	if count != 1 {
//		t.Fatalf("Fail")
//	}
//	count = RecurCount(".???.#", []int{1, 1}, 0)
//	if count != 3 {
//		t.Fatalf("Fail")
//	}
//
//	count = RecurCount("#.???.#", []int{1, 1, 1}, 0)
//	if count != 3 {
//		t.Fatalf("Fail")
//	}
//}
//
//func TestRecurCount11(t *testing.T) {
//	var count = RecurCount("?????#????#?????.??", []int{8, 1, 1, 1, 1}, 0)
//	if count != 12 {
//		t.Fatalf("Fail")
//	}
//}

func TestSample01(t *testing.T) {
	count := CountFile("test01.txt", 5)
	if count != 525152 {
		t.Fatalf("Fail")
	}
}
