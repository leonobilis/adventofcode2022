package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

func parseInput(input string) (output [][]int) {
	for _, s := range strings.Split(input, "\n\n") {
		var temp []int
		for _, s2 := range strings.Split(s, "\n") {
			temp = append(temp, utils.Atoi(s2))
		}
		output = append(output, temp)
	}
	return
}

func p1(calories [][]int) int {
	max := 0
	for _, c := range calories {
		sum := 0
		for _, c2 := range c {
			sum += c2
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func p2(calories [][]int) int {
	max := make([]int, 3)
	for _, c := range calories {
		sum := 0
		for _, c2 := range c {
			sum += c2
		}
		for i := 0; i < 3; i++ {
			if sum > max[i] {
				max[i] = sum
				break
			}
		}
	}

	return max[0] + max[1] + max[2]
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	calories := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(calories))
	fmt.Printf("Part 2: %v\n", p2(calories))
}
