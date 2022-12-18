package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Position struct {
	x, y, z int
}

func parseInput(input string) utils.Set[Position] {
	output := make(utils.Set[Position])
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, ",")
		output.Add(Position{
			x: utils.Atoi(s2[0]),
			y: utils.Atoi(s2[1]),
			z: utils.Atoi(s2[2]),
		})
	}
	return output
}

func p1(cubes utils.Set[Position]) int {
	sum := len(cubes) * 6
	for c1 := range cubes {
		for _, diff := range [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
			if cubes.Contains(Position{c1.x + diff[0], c1.y + diff[1], c1.z + diff[2]}) {
				sum--
			}
		}
	}
	return sum
}

func p2(cubes utils.Set[Position]) int {
	minX, minY, minZ, maxX, maxY, maxZ := int(^uint(0)>>1), int(^uint(0)>>1), int(^uint(0)>>1), 0, 0, 0
	for c := range cubes {
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.z < minZ {
			minZ = c.z
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		}
	}

	getStartAndCondition := func(diff, min, max int) (int, func(int) bool) {
		if diff == 1 {
			return max, func(a int) bool { return a >= min }
		}
		return min, func(a int) bool { return a <= max }
	}

	diffRange := []int{-1, 1}
	newCubesZ := []utils.Set[Position]{make(utils.Set[Position]), make(utils.Set[Position])}
	for iZ, diffZ := range diffRange {
		z, conditionZ := getStartAndCondition(diffZ, minZ, maxZ)
		for ; conditionZ(z); z -= diffZ {
			newCubesY := []utils.Set[Position]{make(utils.Set[Position]), make(utils.Set[Position])}
			for iY, diffY := range diffRange {
				y, conditionY := getStartAndCondition(diffY, minY, maxY)
				for ; conditionY(y); y -= diffY {
					newCubesX := []utils.Set[Position]{make(utils.Set[Position]), make(utils.Set[Position])}
					for iX, diffX := range diffRange {
						x, conditionX := getStartAndCondition(diffX, minX, maxX)
						for ; conditionX(x); x -= diffX {
							if !cubes.Contains(Position{x, y, z}) &&
								(cubes.Contains(Position{x + diffX, y, z}) || newCubesX[iX].Contains(Position{x + diffX, y, z})) &&
								(cubes.Contains(Position{x, y + diffY, z}) || newCubesY[iY].Contains(Position{x, y + diffY, z})) &&
								(cubes.Contains(Position{x, y, z + diffZ}) || newCubesZ[iZ].Contains(Position{x, y, z + diffZ})) {
								newCubesX[iX].Add(Position{x, y, z})
							}
						}
					}
					newCubesY[iY] = newCubesY[iY].Union(newCubesX[0].Intersect(newCubesX[1]))
				}
			}
			newCubesZ[iZ] = newCubesZ[iZ].Union(newCubesY[0].Intersect(newCubesY[1]))
		}
	}

	return p1(cubes.Union(newCubesZ[0].Intersect(newCubesZ[1])))
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
