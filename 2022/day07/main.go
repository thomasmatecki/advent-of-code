package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	name     string
	visited  bool
	parent   *Directory
	children []*Directory
	files    []int
	size     int
}

type InputIter struct {
	cursor int
	lines  []string
}

var cdExpr = regexp.MustCompile(`\$ cd (.+)`)
var lsExpr = regexp.MustCompile(`\$ ls`)
var fileExpr = regexp.MustCompile(`(\d+) ([\w.]+)`)

var data, _ = os.ReadFile("07.txt")
var input = strings.Split(string(data), "\n")

func (iter *InputIter) HasNext() bool {
	return iter.cursor < len(iter.lines)
}

func (iter *InputIter) HasOutput() bool {
	return iter.HasNext() && !strings.HasPrefix(iter.lines[iter.cursor], "$ ")
}

func (iter *InputIter) Next() *string {
	line := &iter.lines[iter.cursor]
	iter.cursor++
	return line
}

func initDirs(iter InputIter) *Directory {

	iter.Next()
	var curr *Directory
	var root *Directory
	curr = new(Directory)
	root = curr
	root.name = "/"

	for iter.HasNext() {
		next := iter.Next()
		if matchCd := cdExpr.FindStringSubmatch(*next); len(matchCd) > 0 {
			if matchCd[1] == ".." {
				curr = curr.parent
			} else {
				newDir := new(Directory)
				newDir.parent = curr
				newDir.name = fmt.Sprintf("%s%s/", curr.name, matchCd[1])
				//
				curr.children = append(curr.children, newDir)
				curr = newDir
			}
		}

		if matchLs := lsExpr.FindStringSubmatch(*next); len(matchLs) > 0 && !curr.visited {
			dirTotal := 0
			for iter.HasOutput() {
				next = iter.Next()
				if matchFile := fileExpr.FindStringSubmatch(*next); len(matchFile) > 0 {
					size, _ := strconv.Atoi(matchFile[1])
					curr.files = append(curr.files, size)
					dirTotal += size
				}
			}
			curr.size += dirTotal
			curr.visited = true
			parent := curr.parent
			for parent != nil {
				parent.size += dirTotal
				parent = parent.parent
			}
		}
	}
	return root

}

func dirSlice(rootDir *Directory) []*Directory {

	dirs := []*Directory{rootDir}
	for i := 0; i < len(dirs); i++ {
		curr := dirs[i]
		dirs = append(dirs, curr.children...)
	}

	return dirs

}

func SumTotals(lines []string) int {
	var lineIter = InputIter{
		cursor: 0,
		lines:  lines,
	}
	rootDir := initDirs(lineIter)
	total := 0
	for _, dir := range dirSlice(rootDir) {
		if dir.size <= 100000 {
			total += dir.size
		}
	}

	return total

}

func PartOne() {
	println("Part One:", SumTotals(input))
}

func PartTwo() {

	var lineIter = InputIter{
		cursor: 0,
		lines:  input,
	}
	rootDir := initDirs(lineIter)
	dirs := dirSlice(rootDir)
	needed := -70000000 + rootDir.size + 30000000

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	for _, dir := range dirs {
		if dir.size > needed {
			println("Part Two:", dir.size)
			return
		}
	}

}

func main() {
	PartOne()
	PartTwo()
}
