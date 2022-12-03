package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Rucksack []rune

func parseInput(input string) (output []Rucksack) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, Rucksack(s))
	}
	return
}

func priority(v rune) int {
	if v >= 97 && v <= 122 {
		return int(v) - 96
	} else if v >= 65 && v <= 90 {
		return int(v) - 38
	}
	return 0
}

func p1(inp []Rucksack) (sum int) {
	for _, rucksack := range inp {
		half := len(rucksack) / 2
		intersection := utils.NewSet(rucksack[:half]).Intersect(utils.NewSet(rucksack[half:]))
		for v := range intersection {
			sum += priority(v)
		}
	}
	return
}

func p2(inp []Rucksack) (sum int) {
	for i := 0; i < len(inp)-2; i += 3 {
		intersection := utils.NewSet(inp[i]).
			Intersect(utils.NewSet(inp[i+1])).
			Intersect(utils.NewSet(inp[i+2]))
		for v := range intersection {
			sum += priority(v)
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
