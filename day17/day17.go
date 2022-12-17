package main

import (
	"fmt"
	"io/ioutil"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Dir int

const (
	Right Dir = iota
	Left
)

type Pattern struct {
	i    int
	jets []Dir
}

func (p *Pattern) Next() Dir {
	p.i = (p.i + 1) % len(p.jets)
	return p.jets[p.i]
}

func NewPattern(s string) *Pattern {
	jets := make([]Dir, 0)
	for _, r := range s {
		switch r {
		case '>':
			jets = append(jets, Right)
		case '<':
			jets = append(jets, Left)
		}
	}
	return &Pattern{-1, jets}
}

type Position struct {
	x, y int
}

type Shape []Position

func (s Shape) move(moveX, moveY int, rocks utils.Set[Position]) (Shape, bool) {
	newShape := make(Shape, 0, len(s))
	for _, s2 := range s {
		posToCheck := Position{s2.x + moveX, s2.y + moveY}
		if posToCheck.y < 0 || posToCheck.x < 0 || posToCheck.x > 6 || rocks.Contains(posToCheck) {
			return s, false
		}
		newShape = append(newShape, posToCheck)
	}
	return newShape, true
}

func (s Shape) MoveDown(rocks utils.Set[Position]) (Shape, bool) {
	return s.move(0, -1, rocks)
}

func (s Shape) MoveRight(rocks utils.Set[Position]) (Shape, bool) {
	return s.move(+1, 0, rocks)
}

func (s Shape) MoveLeft(rocks utils.Set[Position]) (Shape, bool) {
	return s.move(-1, 0, rocks)
}

func (s Shape) Push(dir Dir, rocks utils.Set[Position]) Shape {
	var shape Shape
	switch dir {
	case Right:
		shape, _ = s.MoveRight(rocks)
	case Left:
		shape, _ = s.MoveLeft(rocks)
	}
	return shape
}

func (s Shape) MaxY() (maxY int) {
	for _, s2 := range s {
		if maxY < s2.y {
			maxY = s2.y
		}
	}
	return
}

func ShapesGenerator() func(Position) Shape {
	i := -1
	next := func(pos Position) (shape Shape) {
		i = (i + 1) % 5
		switch i {
		case 0:
			shape = Shape{
				Position{pos.x, pos.y},
				Position{pos.x + 1, pos.y},
				Position{pos.x + 2, pos.y},
				Position{pos.x + 3, pos.y},
			}
		case 1:
			shape = Shape{
				Position{pos.x + 1, pos.y + 2},
				Position{pos.x, pos.y + 1},
				Position{pos.x + 1, pos.y + 1},
				Position{pos.x + 2, pos.y + 1},
				Position{pos.x + 1, pos.y},
			}
		case 2:
			shape = Shape{
				Position{pos.x + 2, pos.y + 2},
				Position{pos.x + 2, pos.y + 1},
				Position{pos.x, pos.y},
				Position{pos.x + 1, pos.y},
				Position{pos.x + 2, pos.y},
			}
		case 3:
			shape = Shape{
				Position{pos.x, pos.y + 3},
				Position{pos.x, pos.y + 2},
				Position{pos.x, pos.y + 1},
				Position{pos.x, pos.y},
			}
		case 4:
			shape = Shape{
				Position{pos.x, pos.y + 1},
				Position{pos.x + 1, pos.y + 1},
				Position{pos.x, pos.y},
				Position{pos.x + 1, pos.y},
			}
		}
		return
	}
	return next
}

func p1(pattern *Pattern) (height int) {
	rocks := make(utils.Set[Position])
	nextShape := ShapesGenerator()
	maxY := -1
	for i := 0; i < 2022; i++ {
		shape := nextShape(Position{2, maxY + 3})
		shape = shape.Push(pattern.Next(), rocks)
		for ok := true; ok; shape, ok = shape.MoveDown(rocks) {
			shape = shape.Push(pattern.Next(), rocks)
		}
		for _, p := range shape {
			rocks.Add(p)
			if p.y > maxY {
				maxY = p.y
			}
		}
	}
	return maxY + 1
}

func p2(pattern *Pattern) int {
	rocks := make(utils.Set[Position])
	nextShape := ShapesGenerator()
	maxY := -1
	repI, repMaxY, diffI, diffMaxY := -1, -1, -1, -1
	for i := 0; ; i++ {
		shape := nextShape(Position{2, maxY + 3})
		shape = shape.Push(pattern.Next(), rocks)
		for ok := true; ok; shape, ok = shape.MoveDown(rocks) {
			shape = shape.Push(pattern.Next(), rocks)
		}
		for _, p := range shape {
			rocks.Add(p)
			if p.y > maxY {
				maxY = p.y
			}
		}

		count := 0
		for x := 0; x < 7; x++ {
			if !rocks.Contains(Position{x, maxY}) {
				break
			}
			count++
		}
		if count == 7 {
			if repMaxY == -1 {
				repI = i
				repMaxY = maxY
			} else {
				diffI = i - repI
				diffMaxY = maxY - repMaxY
				repI = i
				repMaxY = maxY
				break
			}
		}
	}

	iter := 1000000000000 - repI - 1

	for i := 0; i < iter%diffI; i++ {
		shape := nextShape(Position{2, maxY + 3})
		shape = shape.Push(pattern.Next(), rocks)
		for ok := true; ok; shape, ok = shape.MoveDown(rocks) {
			shape = shape.Push(pattern.Next(), rocks)
		}
		for _, p := range shape {
			rocks.Add(p)
			if p.y > maxY {
				maxY = p.y
			}
		}
	}

	return maxY + 1 + (iter/diffI)*diffMaxY
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	fmt.Printf("Part 1: %v\n", p1(NewPattern(string(input))))
	fmt.Printf("Part 2: %v\n", p2(NewPattern(string(input))))
}
