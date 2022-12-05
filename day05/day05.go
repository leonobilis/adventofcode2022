package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Stack []byte

type Procedure struct {
	quantity, from, to int
}

func parseInput(input string) ([]Stack, []Procedure) {
	stacks := make([]Stack, 9)
	s := strings.Split(input, "\n\n")
	s2 := strings.Split(s[0], "\n")

	for i := len(s2) - 2; i >= 0; i-- {
		for j := 1; j < len(s2[i]); j += 4 {
			if s2[i][j] >= 'A' && s2[i][j] <= 'Z' {
				stacks[(j-1)/4] = append(stacks[(j-1)/4], s2[i][j])
			}
		}
	}

	var procedures []Procedure
	re := regexp.MustCompile(`[\d]+`)
	for _, s3 := range strings.Split(s[1], "\n") {
		match := re.FindAllString(s3, 3)
		procedures = append(procedures, Procedure{
			quantity: utils.Atoi(match[0]),
			from:     utils.Atoi(match[1]) - 1,
			to:       utils.Atoi(match[2]) - 1,
		})
	}

	return stacks, procedures
}

func topCrates(stacks []Stack) string {
	var b strings.Builder
	for _, s := range stacks {
		b.WriteByte(s[len(s)-1])
	}
	return b.String()
}

func p1(stacks []Stack, procedures []Procedure) string {
	for _, p := range procedures {
		for i := 1; i <= p.quantity; i++ {
			stacks[p.to] = append(stacks[p.to], stacks[p.from][len(stacks[p.from])-i])
		}
		stacks[p.from] = stacks[p.from][:len(stacks[p.from])-p.quantity]
	}
	return topCrates(stacks)
}

func p2(stacks []Stack, procedures []Procedure) string {
	for _, p := range procedures {
		for i := p.quantity; i > 0; i-- {
			stacks[p.to] = append(stacks[p.to], stacks[p.from][len(stacks[p.from])-i])
		}
		stacks[p.from] = stacks[p.from][:len(stacks[p.from])-p.quantity]
	}
	return topCrates(stacks)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	stacks, procedures := parseInput(string(input))
	stacks2 := make([]Stack, len(stacks))
	for i := range stacks {
		stacks2[i] = append(make([]byte, 0, len(stacks[i])), stacks[i]...)
	}
	fmt.Printf("Part 1: %v\n", p1(stacks, procedures))
	fmt.Printf("Part 2: %v\n", p2(stacks2, procedures))
}
