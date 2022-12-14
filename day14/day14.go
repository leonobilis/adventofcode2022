package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Position struct {
	x, y int
}

func parseInput(input string) (utils.Set[Position], int) {
	output := make(utils.Set[Position])
	maxY := 0
	for _, s := range strings.Split(input, "\n") {
		var prev Position
		for i, s2 := range strings.Split(s, " -> ") {
			s3 := strings.Split(s2, ",")
			newX, newY := utils.Atoi(s3[0]), utils.Atoi(s3[1])
			if i > 0 {
				fromX, toX := utils.MinMax(prev.x, newX)
				fromY, toY := utils.MinMax(prev.y, newY)
				for x := fromX; x <= toX; x++ {
					for y := fromY; y <= toY; y++ {
						output.Add(Position{x, y})
					}
				}
			}
			prev = Position{newX, newY}
			if newY > maxY {
				maxY = newY
			}
		}
	}
	return output, maxY
}

func p1(inp utils.Set[Position], maxY int) int {
	sand := make(utils.Set[Position])
	for {
		unit := Position{500, 0}
		for {
			if unit.y == maxY {
				return len(sand)
			}
			posToCheck := Position{unit.x, unit.y + 1}
			if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
				unit = posToCheck
				continue
			}
			posToCheck = Position{unit.x - 1, unit.y + 1}
			if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
				unit = posToCheck
				continue
			}
			posToCheck = Position{unit.x + 1, unit.y + 1}
			if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
				unit = posToCheck
				continue
			}
			sand.Add(unit)
			break
		}
	}
}

func p2(inp utils.Set[Position], maxY int) int {
	sand := make(utils.Set[Position])
	for {
		unit := Position{500, 0}
		for {
			if sand.Contains(unit) {
				return len(sand)
			}
			if unit.y < maxY+1 {
				posToCheck := Position{unit.x, unit.y + 1}
				if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
					unit = posToCheck
					continue
				}
				posToCheck = Position{unit.x - 1, unit.y + 1}
				if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
					unit = posToCheck
					continue
				}
				posToCheck = Position{unit.x + 1, unit.y + 1}
				if !inp.Contains(posToCheck) && !sand.Contains(posToCheck) {
					unit = posToCheck
					continue
				}
			}
			sand.Add(unit)
			break
		}
	}
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp, maxY := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp, maxY))
	fmt.Printf("Part 2: %v\n", p2(inp, maxY))
}
