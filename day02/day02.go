package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Round struct {
	Opponent, You int
}

func parseInput(input string) (output []Round) {
	for _, s := range strings.Split(input, "\n") {
		var round Round
		s2 := strings.Split(s, " ")
		switch s2[0] {
		case "A":
			round.Opponent = 1
		case "B":
			round.Opponent = 2
		case "C":
			round.Opponent = 3
		}
		switch s2[1] {
		case "X":
			round.You = 1
		case "Y":
			round.You = 2
		case "Z":
			round.You = 3
		}
		output = append(output, round)
	}
	return
}

func p1(inp []Round) int {
	score := 0
	for _, round := range inp {
		score += round.You
		switch round.You - round.Opponent {
		case 0:
			score += 3
		case 1, -2:
			score += 6
		}
	}
	return score
}

func p2(inp []Round) int {
	score := 0
	for _, round := range inp {
		switch round.You {
		case 1:
			score += (round.Opponent+1)%3 + 1
		case 2:
			score += 3 + round.Opponent
		case 3:
			score += 6 + round.Opponent%3 + 1
		}
	}
	return score
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
