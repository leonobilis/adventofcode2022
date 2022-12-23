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

type Moves []func(Position, utils.Set[Position]) (Position, bool)

func (moves Moves) Shift() Moves {
	lastMove := moves[1]
	return append(append(moves[:1], moves[2:]...), lastMove)
}

func NewMoves() Moves {
	return []func(Position, utils.Set[Position]) (Position, bool){
		func(p Position, inp utils.Set[Position]) (Position, bool) {
			for xdiff := -1; xdiff <= 1; xdiff++ {
				for ydiff := -1; ydiff <= 1; ydiff++ {
					if (xdiff != 0 || ydiff != 0) && inp.Contains(Position{p.x + xdiff, p.y + ydiff}) {
						return p, false
					}
				}
			}
			return p, true
		},
		func(p Position, inp utils.Set[Position]) (Position, bool) {
			if !inp.Contains(Position{p.x, p.y - 1}) && !inp.Contains(Position{p.x - 1, p.y - 1}) && !inp.Contains(Position{p.x + 1, p.y - 1}) {
				return Position{p.x, p.y - 1}, true
			}
			return p, false
		},
		func(p Position, inp utils.Set[Position]) (Position, bool) {
			if !inp.Contains(Position{p.x, p.y + 1}) && !inp.Contains(Position{p.x - 1, p.y + 1}) && !inp.Contains(Position{p.x + 1, p.y + 1}) {
				return Position{p.x, p.y + 1}, true
			}
			return p, false
		},
		func(p Position, inp utils.Set[Position]) (Position, bool) {
			if !inp.Contains(Position{p.x - 1, p.y}) && !inp.Contains(Position{p.x - 1, p.y - 1}) && !inp.Contains(Position{p.x - 1, p.y + 1}) {
				return Position{p.x - 1, p.y}, true
			}
			return p, false
		},
		func(p Position, inp utils.Set[Position]) (Position, bool) {
			if !inp.Contains(Position{p.x + 1, p.y}) && !inp.Contains(Position{p.x + 1, p.y - 1}) && !inp.Contains(Position{p.x + 1, p.y + 1}) {
				return Position{p.x + 1, p.y}, true
			}
			return p, false
		},
	}
}

func parseInput(input string) utils.Set[Position] {
	output := make(utils.Set[Position])
	for y, s := range strings.Split(input, "\n") {
		for x, s2 := range s {
			if s2 == '#' {
				output.Add(Position{x, y})
			}
		}
	}
	return output
}

func p1(inp utils.Set[Position]) (sum int) {
	moves := NewMoves()

	for i := 0; i < 10; i++ {
		proposals := make(map[Position][]Position)
		for p := range inp {
			for mi, move := range moves {
				if proposal, ok := move(p, inp); ok {
					proposals[proposal] = append(proposals[proposal], p)
					break
				} else if mi == len(moves)-1 {
					proposals[p] = []Position{p}
				}
			}
		}
		new := make(utils.Set[Position])

		for proposal, elves := range proposals {
			if len(elves) == 1 {
				new.Add(proposal)
			} else {
				for _, e := range elves {
					new.Add(e)
				}
			}
		}

		inp = new
		moves = moves.Shift()
	}

	minX, minY, maxX, maxY := int(^uint(0)>>1), int(^uint(0)>>1), 0, 0
	for p := range inp {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if !inp.Contains(Position{x, y}) {
				sum++
			}
		}
	}

	return
}

func p2(inp utils.Set[Position]) (i int) {
	moves := NewMoves()

	for count := 0; count < len(inp); i++ {
		count = 0
		proposals := make(map[Position][]Position)
		for p := range inp {
			for mi, move := range moves {
				if proposal, ok := move(p, inp); ok {
					proposals[proposal] = append(proposals[proposal], p)
					if mi == 0 {
						count++
					}
					break
				} else if mi == len(moves)-1 {
					count++
					proposals[p] = []Position{p}
				}
			}
		}
		new := make(utils.Set[Position])

		for proposal, elves := range proposals {
			if len(elves) == 1 {
				new.Add(proposal)
			} else {
				count += len(elves)
				for _, e := range elves {
					new.Add(e)
				}
			}
		}

		inp = new
		moves = moves.Shift()
	}

	return i
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
