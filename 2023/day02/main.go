package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var data, _ = os.ReadFile("02.txt")

var lineLeadExpr = regexp.MustCompile(`Game (\d+):`)
var pickExpr = regexp.MustCompile(` (\d+) (\w+)`)
var maxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func PartOne() {
	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		if match := lineLeadExpr.FindStringSubmatch(line); match != nil {
			for _, pickset := range strings.Split(line[len(match[0]):], ";") {
				for _, pick := range strings.Split(pickset, ",") {
					pickMatch := pickExpr.FindStringSubmatch(pick)
					pickCount, _ := strconv.Atoi(pickMatch[1])
					if maxColors[pickMatch[2]] < pickCount {
						goto NextLine
					}
				}
			}
			gameNum, _ := strconv.Atoi(match[1])
			total += gameNum
		}
	NextLine:
	}

	println("Part One:", total)
}
func PartTwo() {
	total := 0

	for _, line := range strings.Split(string(data), "\n") {
		if match := lineLeadExpr.FindStringSubmatch(line); match != nil {
			var colors = map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, pickset := range strings.Split(line[len(match[0]):], ";") {
				for _, pick := range strings.Split(pickset, ",") {
					pickMatch := pickExpr.FindStringSubmatch(pick)
					pickCount, _ := strconv.Atoi(pickMatch[1])
					if colors[pickMatch[2]] < pickCount {
						colors[pickMatch[2]] = pickCount
					}
				}
			}
			gamePower := colors["red"] * colors["green"] * colors["blue"]
			total += gamePower
		}
	}

	println("Part Two:", total)
}

func main() {
	PartOne()
	PartTwo()
}
