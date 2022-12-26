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

type Direction int

const (
	Right = iota
	Down
	Left
	Up
)

type FacingChanger func(Direction) Direction

func noChange(d Direction) Direction {
	return d
}

func turnLeft(d Direction) Direction {
	return Direction(utils.Mod(int(d-1), 4))
}

func turnRight(d Direction) Direction {
	return Direction(utils.Mod(int(d+1), 4))
}

func invert(d Direction) Direction {
	return Direction(utils.Mod(int(d+2), 4))
}

type Tile struct {
	x, y                                          int
	Left, Right, Up, Down                         *Tile
	FacingLeft, FacingRight, FacingUp, FacingDown FacingChanger
}

func NewTile(x, y int) *Tile {
	return &Tile{
		x: x, y: y,
		FacingLeft:  noChange,
		FacingRight: noChange,
		FacingUp:    noChange,
		FacingDown:  noChange,
	}
}

type Traveler struct {
	Tile   *Tile
	Facing Direction
}

func (t *Traveler) TurnLeft() {
	t.Facing = turnLeft(t.Facing)
}

func (t *Traveler) TurnRight() {
	t.Facing = turnRight(t.Facing)
}

func (t *Traveler) Move(n int) {
	for i := 0; i < n; i++ {
		switch t.Facing {
		case Right:
			if t.Tile.Right != nil {
				t.Facing = t.Tile.FacingRight(t.Facing)
				t.Tile = t.Tile.Right
			} else {
				break
			}
		case Down:
			if t.Tile.Down != nil {
				t.Facing = t.Tile.FacingDown(t.Facing)
				t.Tile = t.Tile.Down
			} else {
				break
			}
		case Left:
			if t.Tile.Left != nil {
				t.Facing = t.Tile.FacingLeft(t.Facing)
				t.Tile = t.Tile.Left
			} else {
				break
			}
		case Up:
			if t.Tile.Up != nil {
				t.Facing = t.Tile.FacingUp(t.Facing)
				t.Tile = t.Tile.Up
			} else {
				break
			}
		}
	}
}

type Actions []func(*Traveler)

func parseInput(input string) (grid [][]byte, actions Actions) {
	s := strings.Split(input, "\n\n")

	for _, line := range strings.Split(s[0], "\n") {
		grid = append(grid, []byte(line))
	}

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

func travel(tile *Tile, actions Actions) int {
	traveler := Traveler{tile, Right}
	for _, action := range actions {
		action(&traveler)
	}
	return 1000*(traveler.Tile.y+1) + 4*(traveler.Tile.x+1) + int(traveler.Facing)
}

func p1(grid [][]byte, actions Actions) int {
	tiles := make(map[Position]*Tile)
	var firstTile *Tile
	teleportX := make(map[int]*Tile)
	teleportY := make(map[int]*Tile)
	prevLen := 0

	for y, rangeX := range grid {
		for x, field := range rangeX {
			switch field {
			case '.':
				tile := NewTile(x, y)
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
				if y == len(grid)-1 {
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

	return travel(firstTile, actions)
}

func joinTiles(tiles map[Position]*Tile, walls utils.Set[Position], A, diffA Position, dirA Direction, B, diffB Position, dirB Direction) {

	var facingChangerA, facingChangerB FacingChanger

	switch dirA - dirB {
	case 0:
		facingChangerA = invert
		facingChangerB = invert
	case 2, -2:
		facingChangerA = noChange
		facingChangerB = noChange
	case 1, -3:
		facingChangerA = turnRight
		facingChangerB = turnLeft
	case -1, 3:
		facingChangerA = turnLeft
		facingChangerB = turnRight

	}

	tileA, okA := tiles[A]
	tileB, okB := tiles[B]
	for (okA || walls.Contains(A)) && (okB || walls.Contains(B)) {
		if okA && okB {
			switch dirA {
			case Right:
				tileA.Right = tileB
				tileA.FacingRight = facingChangerA
			case Down:
				tileA.Down = tileB
				tileA.FacingDown = facingChangerA
			case Left:
				tileA.Left = tileB
				tileA.FacingLeft = facingChangerA
			case Up:
				tileA.Up = tileB
				tileA.FacingUp = facingChangerA
			}

			switch dirB {
			case Right:
				tileB.Right = tileA
				tileB.FacingRight = facingChangerB
			case Down:
				tileB.Down = tileA
				tileB.FacingDown = facingChangerB
			case Left:
				tileB.Left = tileA
				tileB.FacingLeft = facingChangerB
			case Up:
				tileB.Up = tileA
				tileB.FacingUp = facingChangerB
			}
		}
		A, B = Position{A.x + diffA.x, A.y + diffA.y}, Position{B.x + diffB.x, B.y + diffB.y}
		tileA, okA = tiles[A]
		tileB, okB = tiles[B]
	}

	if okA || walls.Contains(A) {
		if diffB.x != 0 {
			var newDir Direction
			if diffB.x > 0 {
				newDir = Right
			} else {
				newDir = Left
			}
			newB := Position{B.x - diffB.x, B.y}
			checkB := Position{B.x - diffB.x, B.y - 1}
			if _, nbOk := tiles[checkB]; nbOk || walls.Contains(checkB) {
				joinTiles(tiles, walls, A, diffA, dirA, newB, Position{0, -1}, newDir)
			}
			checkB = Position{B.x - diffB.x, B.y + 1}
			if _, nbOk := tiles[checkB]; nbOk || walls.Contains(checkB) {
				joinTiles(tiles, walls, A, diffA, dirA, newB, Position{0, 1}, newDir)
			}
		}
		if diffB.y != 0 {
			var newDir Direction
			if diffB.y > 0 {
				newDir = Down
			} else {
				newDir = Up
			}
			newB := Position{B.x, B.y - diffB.y}
			checkB := Position{B.x - 1, B.y - diffB.y}
			if _, nbOk := tiles[checkB]; nbOk {
				joinTiles(tiles, walls, A, diffA, dirA, newB, Position{-1, 0}, newDir)
			}
			checkB = Position{B.x + 1, B.y - diffB.y}
			if _, nbOk := tiles[checkB]; nbOk {
				joinTiles(tiles, walls, A, diffA, dirA, newB, Position{1, 0}, newDir)
			}
		}
	}

	if okB || walls.Contains(B) {
		if diffA.x != 0 {
			var newDir Direction
			if diffA.x > 0 {
				newDir = Right
			} else {
				newDir = Left
			}
			newA := Position{A.x - diffA.x, A.y}
			checkA := Position{A.x - diffA.x, A.y - 1}
			if _, naOk := tiles[checkA]; naOk || walls.Contains(checkA) {
				joinTiles(tiles, walls, newA, Position{0, -1}, newDir, B, diffB, dirB)
			}
			checkA = Position{A.x - diffA.x, A.y + 1}
			if _, naOk := tiles[checkA]; naOk || walls.Contains(checkA) {
				joinTiles(tiles, walls, newA, Position{0, 1}, newDir, B, diffB, dirB)
			}
		}
		if diffA.y != 0 {
			var newDir Direction
			if diffA.y > 0 {
				newDir = Down
			} else {
				newDir = Up
			}
			newA := Position{A.x, A.y - diffA.y}
			checkA := Position{A.x - 1, A.y - diffA.y}
			if _, naOk := tiles[checkA]; naOk || walls.Contains(checkA) {
				joinTiles(tiles, walls, newA, Position{-1, 0}, newDir, B, diffB, dirB)
			}
			checkA = Position{A.x + 1, A.y - diffA.y}
			if _, naOk := tiles[checkA]; naOk || walls.Contains(checkA) {
				joinTiles(tiles, walls, newA, Position{1, 0}, newDir, B, diffB, dirB)
			}
		}
	}

}

func p2(grid [][]byte, actions Actions) int {
	tiles := make(map[Position]*Tile)
	walls := make(utils.Set[Position])
	var firstTile *Tile

	for y, rangeX := range grid {
		for x, field := range rangeX {
			switch field {
			case '.':
				tile := NewTile(x, y)
				if up, ok := tiles[Position{x, y - 1}]; ok {
					tile.Up = up
					up.Down = tile
				}
				if left, ok := tiles[Position{x - 1, y}]; ok {
					tile.Left = left
					left.Right = tile
				}
				if firstTile == nil {
					firstTile = tile
				}
				tiles[Position{x, y}] = tile
			case '#':
				walls.Add(Position{x, y})
			}
		}
	}

	tilesWalls := make(utils.Set[Position])
	for p := range tiles {
		tilesWalls.Add(p)
	}
	tilesWalls = tilesWalls.Union(walls)

	for pos := range tilesWalls {
		for _, diff := range []struct {
			x, y       int
			dirA, dirB Direction
		}{{1, 1, Down, Right}, {1, -1, Up, Right}, {-1, 1, Down, Left}, {-1, -1, Up, Left}} {
			if !tilesWalls.Contains(Position{pos.x + diff.x, pos.y + diff.y}) {
				p1, p2 := Position{pos.x + diff.x, pos.y}, Position{pos.x, pos.y + diff.y}
				if tilesWalls.Contains(p1) && tilesWalls.Contains(p2) {
					joinTiles(tiles, walls, p1, Position{diff.x, 0}, diff.dirA, p2, Position{0, diff.y}, diff.dirB)
				}
			}
		}
	}

	return travel(firstTile, actions)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	grid, actions := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(grid, actions))
	fmt.Printf("Part 2: %v\n", p2(grid, actions))
}
