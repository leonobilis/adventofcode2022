package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Motion struct {
	direction byte
	steps     int
}

type Position struct {
	x, y int
}

type Rope struct {
	knots []Position
}

func (rope *Rope) adjustKnots() {
	for i := 0; i < len(rope.knots)-1; i++ {
		moveX, moveY := false, false
		if rope.knots[i].x-rope.knots[i+1].x > 1 {
			moveX = true
			rope.knots[i+1].x += rope.knots[i].x - rope.knots[i+1].x - 1
		} else if rope.knots[i].x-rope.knots[i+1].x < -1 {
			moveX = true
			rope.knots[i+1].x += rope.knots[i].x - rope.knots[i+1].x + 1
		}
		if rope.knots[i].y-rope.knots[i+1].y > 1 {
			moveY = true
			rope.knots[i+1].y += rope.knots[i].y - rope.knots[i+1].y - 1
		} else if rope.knots[i].y-rope.knots[i+1].y < -1 {
			moveY = true
			rope.knots[i+1].y += rope.knots[i].y - rope.knots[i+1].y + 1
		}

		if moveX && !moveY && rope.knots[i].y-rope.knots[i+1].y != 0 {
			rope.knots[i+1].y = rope.knots[i].y
		} else if moveY && !moveX && rope.knots[i].x-rope.knots[i+1].x != 0 {
			rope.knots[i+1].x = rope.knots[i].x
		}
	}
}

func simulate(rope *Rope, motions []Motion) int {
	tail := len(rope.knots) - 1
	visited := make(utils.Set[Position])
	for _, m := range motions {
		mX, mY := 0, 0
		switch m.direction {
		case 'R':
			mX = 1
		case 'L':
			mX = -1
		case 'U':
			mY = 1
		case 'D':
			mY = -1
		}

		for i := 0; i < m.steps; i++ {
			rope.knots[0].x += mX
			rope.knots[0].y += mY
			rope.adjustKnots()
			visited.Add(rope.knots[tail])
		}
	}
	return len(visited)
}

func parseInput(input string) (output []Motion) {
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, " ")
		output = append(output, Motion{
			direction: s2[0][0],
			steps:     utils.Atoi(s2[1]),
		})
	}
	return
}

func p1(inp []Motion) int {
	rope := Rope{[]Position{{0, 0}, {0, 0}}}
	return simulate(&rope, inp)
}

func p2(inp []Motion) int {
	rope := Rope{}
	for i := 0; i < 10; i++ {
		rope.knots = append(rope.knots, Position{0, 0})
	}
	return simulate(&rope, inp)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
