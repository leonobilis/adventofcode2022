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

type Direction int

const (
	Right = iota
	Down
	Left
	Up
)

type Blizzards map[Position][]Direction

func (blizzards Blizzards) Move(max Position) Blizzards {
	newBlizzards := make(Blizzards)
	for p, b := range blizzards {
		for _, b2 := range b {
			var p2 Position
			switch b2 {
			case Right:
				p2 = Position{mod(p.x+1, max.x+1), p.y}
			case Down:
				p2 = Position{p.x, mod(p.y+1, max.y+1)}
			case Left:
				p2 = Position{mod(p.x-1, max.x+1), p.y}
			case Up:
				p2 = Position{p.x, mod(p.y-1, max.y+1)}
			}
			newBlizzards[p2] = append(newBlizzards[p2], b2)
		}
	}
	return newBlizzards
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func parseInput(input string) (Blizzards, Position) {
	blizzards := make(Blizzards)
	var p Position
	for y, s := range strings.Split(input, "\n") {
		for x, s2 := range s {
			p = Position{x - 1, y - 1}
			switch s2 {
			case '>':
				blizzards[p] = append(blizzards[p], Right)
			case 'v':
				blizzards[p] = append(blizzards[p], Down)
			case '<':
				blizzards[p] = append(blizzards[p], Left)
			case '^':
				blizzards[p] = append(blizzards[p], Up)
			}
		}
	}
	return blizzards, Position{p.x - 2, p.y - 1}
}

func traverse(blizzards Blizzards, start, end, max Position) (int, Blizzards) {
	possiblePositions := make(utils.Set[Position])
	possiblePositions.Add(start)

	for minute := 1; ; minute++ {
		blizzards = blizzards.Move(max)
		newPossiblePositions := make(utils.Set[Position])

		for p := range possiblePositions {
			if p == end {
				return minute, blizzards
			}
			if _, ok := blizzards[p]; !ok {
				newPossiblePositions.Add(p)
			}
			if p.x < max.x && p.y != -1 && len(blizzards[Position{p.x + 1, p.y}]) == 0 {
				newPossiblePositions.Add(Position{p.x + 1, p.y})
			}
			if p.y < max.y && len(blizzards[Position{p.x, p.y + 1}]) == 0 {
				newPossiblePositions.Add(Position{p.x, p.y + 1})
			}
			if p.x > 0 && p.y != max.y+1 && len(blizzards[Position{p.x - 1, p.y}]) == 0 {
				newPossiblePositions.Add(Position{p.x - 1, p.y})
			}
			if p.y > 0 && len(blizzards[Position{p.x, p.y - 1}]) == 0 {
				newPossiblePositions.Add(Position{p.x, p.y - 1})
			}
		}
		possiblePositions = newPossiblePositions
	}
}

func p1(blizzards Blizzards, max Position) int {
	minutes, _ := traverse(blizzards, Position{0, -1}, max, max)
	return minutes
}

func p2(blizzards Blizzards, max Position) int {
	minutes1, blizzards := traverse(blizzards, Position{0, -1}, max, max)
	minutes2, blizzards := traverse(blizzards, Position{max.x, max.y + 1}, Position{0, 0}, max)
	minutes3, _ := traverse(blizzards, Position{0, -1}, max, max)
	return minutes1 + minutes2 + minutes3
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	blizzards, max := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(blizzards, max))
	fmt.Printf("Part 2: %v\n", p2(blizzards, max))
}
