package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type ElvesPair struct {
	start1, end1, start2, end2 int
}

func parseInput(input string) (output []ElvesPair) {
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, ",")
		sections1 := strings.Split(s2[0], "-")
		sections2 := strings.Split(s2[1], "-")
		output = append(output, ElvesPair{
			start1: utils.Atoi(sections1[0]),
			end1:   utils.Atoi(sections1[1]),
			start2: utils.Atoi(sections2[0]),
			end2:   utils.Atoi(sections2[1]),
		})
	}
	return
}

func p1(inp []ElvesPair) (sum int) {
	for _, pair := range inp {
		if (pair.start1 <= pair.start2 && pair.end1 >= pair.end2) || (pair.start2 <= pair.start1 && pair.end2 >= pair.end1) {
			sum += 1
		}
	}
	return
}

func p2(inp []ElvesPair) (sum int) {
	for _, pair := range inp {
		if (pair.start1 <= pair.start2 && pair.end1 >= pair.start2) || (pair.start2 <= pair.start1 && pair.end2 >= pair.start1) {
			sum += 1
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
