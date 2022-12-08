package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput(input string) (output [][]byte) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, []byte(s))
	}
	return
}

func checkVisibility(grid [][]byte, y, x int) bool {
	hideCount := 0
	for i := 0; i < y; i++ {
		if grid[i][x] >= grid[y][x] {
			hideCount++
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= grid[y][x] {
			hideCount++
			break
		}
	}
	for i := 0; i < x; i++ {
		if grid[y][i] >= grid[y][x] {
			hideCount++
			break
		}
	}
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= grid[y][x] {
			hideCount++
			break
		}
	}
	return hideCount < 4
}

func p1(inp [][]byte) (sum int) {
	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp[i]); j++ {
			if checkVisibility(inp, i, j) {
				sum++
			}
		}
	}
	return
}

func scenicScore(grid [][]byte, y, x int) int {
	up, down, left, right := 0, 0, 0, 0
	for i := y - 1; i >= 0; i-- {
		up++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		down++
		if grid[i][x] >= grid[y][x] {
			break
		}
	}
	for i := x - 1; i >= 0; i-- {
		left++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}
	for i := x + 1; i < len(grid[y]); i++ {
		right++
		if grid[y][i] >= grid[y][x] {
			break
		}
	}
	return up * down * left * right
}

func p2(inp [][]byte) (max int) {
	for i := 1; i < len(inp); i++ {
		for j := 1; j < len(inp[i]); j++ {
			score := scenicScore(inp, i, j)
			if score > max {
				max = score
			}
		}
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
