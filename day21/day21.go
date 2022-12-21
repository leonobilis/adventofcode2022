package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Monkeys map[string]func(Monkeys) int

func parseInput(input string) (Monkeys, string, string) {
	monkeys := make(Monkeys)
	var rootArg1, rootArg2 string
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, ": ")
		var operation func(Monkeys) int
		switch {
		case s2[0] == "root":
			rootArg1, rootArg2 = s2[1][:4], s2[1][7:]
			fallthrough
		case strings.Contains(s2[1], "+"):
			operation = func(m Monkeys) int { return m[s2[1][:4]](m) + m[s2[1][7:]](m) }
		case strings.Contains(s2[1], "-"):
			operation = func(m Monkeys) int { return m[s2[1][:4]](m) - m[s2[1][7:]](m) }
		case strings.Contains(s2[1], "*"):
			operation = func(m Monkeys) int { return m[s2[1][:4]](m) * m[s2[1][7:]](m) }
		case strings.Contains(s2[1], "/"):
			operation = func(m Monkeys) int { return m[s2[1][:4]](m) / m[s2[1][7:]](m) }
		default:
			number := utils.Atoi(s2[1])
			operation = func(m Monkeys) int { return number }
		}
		monkeys[s2[0]] = operation
	}
	return monkeys, rootArg1, rootArg2
}

func p1(inp Monkeys) int {
	return inp["root"](inp)
}

func p2(inp Monkeys, rootArg1, rootArg2 string) int {
	start := 0
	for i := start; ; i += 1000000000 {
		inp["humn"] = func(m Monkeys) int { return i }
		a, b := inp[rootArg1](inp), inp[rootArg2](inp)
		if a-b < 0 {
			start = i
			break
		}
	}

	for i := start; ; i -= 10000 {
		inp["humn"] = func(m Monkeys) int { return i }
		a, b := inp[rootArg1](inp), inp[rootArg2](inp)
		if a-b > 0 {
			start = i
			break
		}
	}

	for i := start; ; i++ {
		inp["humn"] = func(m Monkeys) int { return i }
		a, b := inp[rootArg1](inp), inp[rootArg2](inp)
		if a == b {
			return i
		}
	}
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp, a1, a2 := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp, a1, a2))
}
