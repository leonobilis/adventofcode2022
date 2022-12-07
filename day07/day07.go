package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Dir struct {
	name  string
	files map[string]int
	dirs  map[string]Dir
}

func NewDir(name string) Dir {
	return Dir{
		name:  name,
		files: make(map[string]int),
		dirs:  make(map[string]Dir),
	}
}

func (d *Dir) Size() (size int) {
	for _, fileSize := range d.files {
		size += fileSize
	}
	for _, dir := range d.dirs {
		size += dir.Size()
	}
	return
}

func parseInput(input string) (output Dir) {
	var dirStack []Dir
	for _, s := range strings.Split(input, "\n") {
		switch s[:4] {
		case "$ cd":
			switch s[5:] {
			case "/":
				output = NewDir("/")
				dirStack = []Dir{output}
			case "..":
				dirStack = dirStack[:len(dirStack)-1]
			default:
				dirStack = append(dirStack, dirStack[len(dirStack)-1].dirs[s[5:]])
			}
		case "$ ls":
			// pass
		case "dir ":
			dirStack[len(dirStack)-1].dirs[s[4:]] = NewDir(s[4:])
		default: // file
			file := strings.Split(s, " ")
			dirStack[len(dirStack)-1].files[file[1]] = utils.Atoi(file[0])
		}
	}
	return
}

func findDirs(dir Dir, condition func(Dir) bool) (dirs []Dir) {
	for _, d := range dir.dirs {
		if condition(d) {
			dirs = append(dirs, d)
		}
		dirs = append(dirs, findDirs(d, condition)...)
	}
	return
}

func p1(inp Dir) (sum int) {
	dirs := append(make([]Dir, 0), findDirs(inp, func(d Dir) bool { return d.Size() <= 100000 })...)
	for _, dir := range dirs {
		sum += dir.Size()
	}
	return
}

func p2(inp Dir) int {
	neededSpace := 30000000 - (70000000 - inp.Size())
	dirs := append(make([]Dir, 0), findDirs(inp, func(d Dir) bool { return d.Size() >= neededSpace })...)
	min := int(^uint(0) >> 1)
	for _, dir := range dirs {
		size := dir.Size()
		if min > size {
			min = size
		}
	}
	return min
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
