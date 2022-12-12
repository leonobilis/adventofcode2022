package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/yourbasic/graph"
)

func parseInput(input string) (output [][]byte, start, end int) {
	for _, line := range strings.Split(input, "\n") {
		output = append(output, []byte(line))
	}
	for y, line := range output {
		for x, val := range line {
			if val == 'S' {
				output[y][x] = 'a'
				start = y*len(output[0]) + x
			} else if val == 'E' {
				output[y][x] = 'z'
				end = y*len(output[0]) + x
			}
		}
	}
	return
}

func addEdge(g *graph.Mutable, first, second int, reversed bool) {
	if reversed {
		g.AddCost(second, first, 1)
	} else {
		g.AddCost(first, second, 1)
	}
}

func getGraph(grid [][]byte, reversed bool) *graph.Mutable {
	rows, cols := len(grid), len(grid[0])
	g := graph.New(rows * cols)
	for y, line := range grid {
		for x, val := range line {
			if x > 0 && int(grid[y][x-1])-int(val) < 2 {
				addEdge(g, y*cols+x, y*cols+x-1, reversed)
			}
			if x < cols-1 && int(grid[y][x+1])-int(val) < 2 {
				addEdge(g, y*cols+x, y*cols+x+1, reversed)
			}
			if y > 0 && int(grid[y-1][x])-int(val) < 2 {
				addEdge(g, y*cols+x, (y-1)*cols+x, reversed)
			}
			if y < rows-1 && int(grid[y+1][x])-int(val) < 2 {
				addEdge(g, y*cols+x, (y+1)*cols+x, reversed)
			}
		}
	}
	return g
}

func p1(grid [][]byte, start, end int) int64 {
	g := getGraph(grid, false)
	_, dist := graph.ShortestPath(g, start, end)
	return dist
}

func p2(grid [][]byte, end int) int64 {
	g := getGraph(grid, true)
	_, dist := graph.ShortestPaths(g, end)
	min := int64(^uint64(0) >> 1)
	for i, d := range dist {
		if grid[i/len(grid[0])][i%len(grid[0])] == 'a' && d != -1 && d < min {
			min = d
		}
	}
	return min
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	grid, start, end := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(grid, start, end))
	fmt.Printf("Part 2: %v\n", p2(grid, end))
}
