package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Instruction struct {
	op  string
	arg int
}

type CRT struct {
	strings.Builder
}

func (crt *CRT) drawPixel(cycle, x int) {
	row := cycle % 40
	if row == 0 {
		crt.WriteByte('\n')
	}
	if row >= x && row < x+3 {
		crt.WriteByte('#')
	} else {
		crt.WriteByte(' ')
	}
}

func parseInput(input string) (output []Instruction) {
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, " ")
		switch s2[0] {
		case "addx":
			output = append(output, Instruction{s2[0], utils.Atoi(s2[1])})
		case "noop":
			output = append(output, Instruction{s2[0], 0})
		}
	}
	return
}

func checkSignalStrength(cycle, x int) int {
	if cycle == 20 || (cycle-20)%40 == 0 {
		return cycle * x
	}
	return 0
}

func p1(inp []Instruction) (sum int) {
	x := 1
	cycle := 0
	for _, instruction := range inp {
		cycle++
		sum += checkSignalStrength(cycle, x)
		if instruction.op == "addx" {
			cycle++
			sum += checkSignalStrength(cycle, x)
			x += instruction.arg
		}
	}
	return
}

func p2(inp []Instruction) string {
	x := 0
	cycle := 0
	var crt CRT
	for _, instruction := range inp {
		crt.drawPixel(cycle, x)
		if instruction.op == "addx" {
			cycle++
			crt.drawPixel(cycle, x)
			x += instruction.arg
		}
		cycle++
	}
	return crt.String()
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
