package main

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var data, _ = os.ReadFile("05.txt")
var rangeExpr = regexp.MustCompile(`^(\d+) (\d+) (\d+)$`)
var rangeStartExpr = regexp.MustCompile(`[\w-]+ map:`)

type Range struct {
	src, dst, len int
}

type RangeMap struct {
	ranges *[]Range
}

func (rmap *RangeMap) sourceIdx(val int) int {
	idx := sort.Search(len(*rmap.ranges), func(i int) bool {
		return (*rmap.ranges)[i].src > val
	})
	return idx - 1
}

func (rmap *RangeMap) translate(v int) (*Range, int) {
	ridx := rmap.sourceIdx(v)
	if ridx < 0 {
		return nil, v
	}
	rng := (*rmap.ranges)[ridx]
	if v < rng.src+rng.len {
		return &rng, rng.dst + (v - rng.src)
	} else {
		return &rng, v
	}
}

func initRangeMaps(lines []string) (rangeMaps []RangeMap) {
	var ranges *[]Range

	for _, line := range lines {
		if len(line) == 0 {
			if (ranges) != nil {
				sort.Slice(*ranges, func(i, j int) bool {
					return (*ranges)[i].src < (*ranges)[j].src
				})
				rangeMaps = append(rangeMaps, RangeMap{
					ranges: ranges,
				})
			}
			continue
		}

		if match := rangeStartExpr.FindStringSubmatch(line); match != nil {
			ranges = new([]Range)
		}

		if match := rangeExpr.FindStringSubmatch(line); match != nil {
			dst, _ := strconv.Atoi(match[1])
			src, _ := strconv.Atoi(match[2])
			len, _ := strconv.Atoi(match[3])
			*ranges = append(*ranges, Range{
				src: src,
				dst: dst,
				len: len,
			})
		}

	}
	return
}

func initSeeds(seedLine string) (seeds []int) {

	seedStrs := strings.Split(seedLine, " ")[1:]
	for _, seedStr := range seedStrs {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	return
}

func traverseMap(seed int, rangeMaps *[]RangeMap) (value int) {
	value = seed
	for _, rangeMap := range *rangeMaps {
		_, value = rangeMap.translate(value)
	}
	return

}

func traverseMaps(seeds []int, rangeMaps *[]RangeMap) (final []int) {
	for idx, seed := range seeds {
		if idx%1000000 == 0 {
			println("", idx)
		}
		final = append(final, traverseMap(seed, rangeMaps))
	}
	return
}

func minDest(dests []int) (int, int) {
	idx := 0
	min := dests[0]
	for i, dest := range dests[1:] {
		if dest < min {
			idx = i
			min = dest
		}
	}
	return idx, min
}

func initSeedsFromRanges(seedRanges []int, precision int) (seeds []int) {
	for i := 1; i < len(seedRanges); i += 2 {
		first := seedRanges[i-1]
		last := seedRanges[i-1] + seedRanges[i]
		println("Range", i, last)
		for j := first; j < last; j += precision {
			seeds = append(seeds, j)
		}
	}

	return
}

func PartOne() {
	lines := strings.Split(string(data), "\n")
	seeds := initSeeds(lines[0])
	rangeMaps := initRangeMaps(lines)
	dests := traverseMaps(seeds, &rangeMaps)
	_, min := minDest(dests)
	println("Part One:", min)
}
func PartTwo() {
	lines := strings.Split(string(data), "\n")
	rangeMaps := initRangeMaps(lines)
	//	seedRanges := initSeeds(lines[0])
	//	seeds := initSeedsFromRanges(seedRanges, 100)
	//	dests := traverseMaps(seeds, &rangeMaps)
	//	idx, min := minDest(dests)
	//	println(seeds)
	//	println(seedRanges)
	//	println("Part Two:", idx, min)
	///
	//692213638 ;; the 3 range
	//682397438 + 30365957
	//712763395

	min := traverseMap(682397438, &rangeMaps)

	for i := 682397438; i < 682397438+30365957; i++ {
		if val := traverseMap(i, &rangeMaps); val < min {
			min = val
		}
	}
	println("Part Two", min)
}

func main() {
	PartOne()
	PartTwo()
}
