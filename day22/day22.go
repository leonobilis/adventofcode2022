package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Position struct {
	x, y int
}

type Tile struct {
	x, y                  int
	Left, Right, Up, Down *Tile
}

type Direction int

const (
	Right = iota
	Down
	Left
	Up
)

func mod(a, b int) int {
	return (a%b + b) % b
}

type Traveler struct {
	Tile   *Tile
	Facing Direction
}

func (t *Traveler) TurnLeft() {
	t.Facing = Direction(mod(int(t.Facing-1), 4))
}

func (t *Traveler) TurnRight() {
	t.Facing = Direction(mod(int(t.Facing+1), 4))
}

func (t *Traveler) Move(n int) {
	// fmt.Printf("[%v, %v] --(%v)-> ", t.Tile.x, t.Tile.y, n)
	for i := 0; i < n; i++ {
		switch t.Facing {
		case Right:
			if t.Tile.Right != nil {
				t.Tile = t.Tile.Right
			} else {
				break
			}
		case Down:
			if t.Tile.Down != nil {
				t.Tile = t.Tile.Down
			} else {
				break
			}
		case Left:
			if t.Tile.Left != nil {
				t.Tile = t.Tile.Left
			} else {
				break
			}
		case Up:
			if t.Tile.Up != nil {
				t.Tile = t.Tile.Up
			} else {
				break
			}

		}
	}
	// fmt.Printf("[%v, %v]\n", t.Tile.x, t.Tile.y)
}

func parseInput(input string) (firstTile *Tile, actions []func(*Traveler)) {
	s := strings.Split(input, "\n\n")
	tiles := make(map[Position]*Tile)
	teleportX := make(map[int]*Tile)
	teleportY := make(map[int]*Tile)
	rangeY := strings.Split(s[0], "\n")
	prevLen := 0

	wals := make(utils.Set[Position])

	for y, rangeX := range rangeY {
		for x, s2 := range rangeX {
			switch s2 {
			case '.':
				tile := &Tile{x: x, y: y}
				if _, ok := teleportY[x]; !ok {
					teleportY[x] = tile
				}
				if _, ok := teleportX[y]; !ok {
					teleportX[y] = tile
				}
				if up, ok := tiles[Position{x, y - 1}]; ok {
					tile.Up = up
					up.Down = tile
				}
				if left, ok := tiles[Position{x - 1, y}]; ok {
					tile.Left = left
					left.Right = tile
				}
				if x == len(rangeX)-1 {
					if tx := teleportX[y]; tx != nil {
						tile.Right = tx
						tx.Left = tile
					}
				}
				if y == len(rangeY)-1 {
					if ty := teleportY[x]; ty != nil {
						tile.Down = ty
						ty.Up = tile
					}
				}
				if firstTile == nil {
					firstTile = tile
				}
				tiles[Position{x, y}] = tile
			case ' ':
				if up, ok := tiles[Position{x, y - 1}]; ok {
					if ty := teleportY[x]; ty != nil {
						up.Down = ty
						ty.Up = up
					}
				}
				if left, ok := tiles[Position{x - 1, y}]; ok {
					if tx := teleportX[y]; tx != nil {
						left.Right = tx
						tx.Left = left
					}
				}
			case '#':
				wals.Add(Position{x, y})
				if _, ok := teleportY[x]; !ok {
					teleportY[x] = nil
				}
				if _, ok := teleportX[y]; !ok {
					teleportX[y] = nil
				}
			}
		}

		for x := len(rangeX); x < prevLen; x++ {
			if up, ok := tiles[Position{x, y - 1}]; ok {
				if ty := teleportY[x]; ty != nil {
					up.Down = ty
					ty.Up = up
				}
			}
		}
		prevLen = len(rangeX)
	}

	// prev := 0
	// for i := 0; i < len(rangeY[0]); i++ {
	// 	if v, ok := teleportY[i]; ok {
	// 		if v != nil {
	// 			if v.Up != nil {
	// 				if prev != v.Up.y {
	// 					fmt.Printf("\n %v --> %v \n", prev, v.Up.y)
	// 				}
	// 				prev = v.Up.y
	// 				fmt.Print(".")
	// 			} else {
	// 				fmt.Print("#")
	// 			}
	// 		} else {
	// 			fmt.Print("N")
	// 		}
	// 	} else {
	// 		fmt.Print("E")
	// 	}
	// }

	// fmt.Println()

	re := regexp.MustCompile(`\d+|[RL]`)
	for _, action := range re.FindAllString(s[1], -1) {
		switch action {
		case "L":
			actions = append(actions, func(t *Traveler) {
				t.TurnLeft()
			})
		case "R":
			actions = append(actions, func(t *Traveler) {
				t.TurnRight()
			})
		default:
			number := utils.Atoi(action)
			actions = append(actions, func(t *Traveler) {
				t.Move(number)
			})
		}
	}

	return
}

func p1(tile *Tile, actions []func(*Traveler)) int {
	traveler := Traveler{tile, Right}
	for _, action := range actions {
		action(&traveler)
	}

	return 1000*(traveler.Tile.y+1) + 4*(traveler.Tile.x+1) + int(traveler.Facing)

}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	// parseInput(string(input))
	tile, actions := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(tile, actions))
	// fmt.Printf("Part 2: %v\n", p2(inp))
}
