package main

var times = []int{
	40, 82, 84, 92,
}

var distances = []int{
	233, 1011, 1110, 1487,
}

func PartOne() {
	total := 1
	for i := 0; i < 4; i++ {
		time := times[i]
		count := 0
		for j := 1; j < time; j++ {
			if distances[i] < (time-j)*j {
				count++
			}
		}
		total *= count
	}
	println("Part One:", total)
}
func PartTwo() {
	time := 40828492
	count := 0

	for j := 1; j < time; j++ {
		if 233101111101487 < (time-j)*j {
			count++
		}
	}
	println("Part Two", count)
}

func main() {
	PartOne()
	PartTwo()
}
