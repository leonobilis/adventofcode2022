package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Monkey struct {
	items             []int
	operation         func(int) int
	monkeyT, monkeyF  *Monkey
	div, inspectCount int
}

func (m *Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) Inspect(item int) int {
	m.inspectCount++
	return m.operation(item)
}

func (m *Monkey) Throw(item int) {
	if item%m.div == 0 {
		m.monkeyT.AddItem(item)
	} else {
		m.monkeyF.AddItem(item)
	}
}

func (m *Monkey) ClearItems() {
	m.items = make([]int, 0)
}

func parseInput(input string) []Monkey {
	re := regexp.MustCompile(`[\d]+`)
	entries := strings.Split(input, "\n\n")
	monkeys := make([]Monkey, len(entries))
	for i, s := range entries {
		monkey := &monkeys[i]
		s2 := strings.Split(s, "\n")
		for _, m := range re.FindAllString(s2[1], -1) {
			monkey.AddItem(utils.Atoi(m))
		}

		var oper func(int, int) int
		switch s2[2][23] {
		case '+':
			oper = func(a, b int) int { return a + b }
		case '*':
			oper = func(a, b int) int { return a * b }
		}

		switch strings.Count(s2[2], "old") {
		case 1:
			number := utils.Atoi(re.FindString(s2[2]))
			monkey.operation = func(old int) int { return oper(old, number) }
		case 2:
			monkey.operation = func(old int) int { return oper(old, old) }
		}

		monkey.div = utils.Atoi(re.FindString(s2[3]))
		monkey.monkeyT = &monkeys[utils.Atoi(re.FindString(s2[4]))]
		monkey.monkeyF = &monkeys[utils.Atoi(re.FindString(s2[5]))]
	}
	return monkeys
}

func process(monkeys []Monkey, rounds int, worryLevelMutator func(int) int) int {
	for i := 0; i < rounds; i++ {
		for mi := range monkeys {
			monkey := &monkeys[mi]
			for _, item := range monkey.items {
				worryLevel := worryLevelMutator(monkey.Inspect(item))
				monkey.Throw(worryLevel)
			}
			monkey.ClearItems()
		}
	}
	var activity []int
	for _, monkey := range monkeys {
		activity = append(activity, monkey.inspectCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(activity)))
	return activity[0] * activity[1]
}

func p1(monkeys []Monkey) int {
	return process(monkeys, 20, func(x int) int { return x / 3 })

}

func p2(monkeys []Monkey) int {
	var mod int = 1
	for _, monkey := range monkeys {
		mod *= monkey.div
	}
	return process(monkeys, 10000, func(x int) int { return x % mod })
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	inp2 := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp2))
}
