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

func (p Position) Dist(p2 Position) int {
	return utils.Abs(p.x-p2.x) + utils.Abs(p.y-p2.y)
}

func parseInput(input string) map[Position]Position {
	re := regexp.MustCompile(`[-\d]+`)
	output := make(map[Position]Position)
	for _, s := range strings.Split(input, "\n") {
		s2 := re.FindAllString(s, 4)
		output[Position{utils.Atoi(s2[0]), utils.Atoi(s2[1])}] = Position{utils.Atoi(s2[2]), utils.Atoi(s2[3])}
	}
	return output
}

func p1(inp map[Position]Position) int {
	noBeacons := make(utils.Set[Position])
	for sensor, beacon := range inp {
		dist := sensor.Dist(beacon)
		for y := sensor.y - dist; y <= sensor.y+dist; y++ {
			if y == 2000000 {
				ydiff := utils.Abs(y - sensor.y)
				for x := sensor.x - dist + ydiff; x <= sensor.x+dist-ydiff; x++ {
					pos := Position{x, y}
					if pos != beacon {
						noBeacons.Add(pos)
					}
				}
			}
		}
	}

	return len(noBeacons)
}

type Chunk struct {
	a, b int
}

type Row struct {
	chunks []Chunk
}

func (r *Row) Cut(c Chunk) {
	newChunks := make([]Chunk, 0)
	for _, chunk := range r.chunks {
		switch {
		case c.a <= chunk.a && c.b >= chunk.b:
			continue
		case c.b < chunk.a || c.a > chunk.b:
			newChunks = append(newChunks, chunk)
		case c.a <= chunk.a && c.b < chunk.b:
			newChunks = append(newChunks, Chunk{c.b, chunk.b})
		case c.a > chunk.a && c.b >= chunk.b:
			newChunks = append(newChunks, Chunk{chunk.a, c.a})
		case c.a > chunk.a && c.b < chunk.b:
			newChunks = append(newChunks, Chunk{chunk.a, c.a - 1}, Chunk{c.b + 1, chunk.b})
		}
	}
	r.chunks = newChunks
}

func p2(inp map[Position]Position) int {
	sensorRange := make(map[Position]int)
	for sensor, beacon := range inp {
		sensorRange[sensor] = sensor.Dist(beacon)
	}

	for y := 0; y <= 4000000; y++ {
		row := Row{[]Chunk{{0, 4000000}}}
		for sensor, r := range sensorRange {
			limitedRange := r - utils.Abs(sensor.y-y)
			if limitedRange > 0 {
				row.Cut(Chunk{sensor.x - limitedRange, sensor.x + limitedRange})
			}
		}
		if len(row.chunks) > 0 {
			return row.chunks[0].a*4000000 + y
		}
	}
	return 0
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
