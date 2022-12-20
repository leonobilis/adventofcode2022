package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

func parseInput(input string) (output []int) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, utils.Atoi(s))
	}
	return
}

func mix(inp []int, iter int) (sum int) {
	numbers := list.New()
	elements := make([]*list.Element, 0, len(inp))
	var zero *list.Element
	for _, number := range inp {
		element := numbers.PushBack(number)
		elements = append(elements, element)
		if number == 0 {
			zero = element
		}
	}

	for i := 0; i < iter; i++ {
		for _, element := range elements {
			number := element.Value.(int)
			if number > 0 {
				for ii := 0; ii < number%(len(inp)-1); ii++ {
					next := element.Next()
					if next == nil {
						next = numbers.Front()
					}
					numbers.MoveAfter(element, next)
				}
			} else if number < 0 {
				for ii := 0; ii < -number%(len(inp)-1); ii++ {
					prev := element.Prev()
					if prev == nil {
						prev = numbers.Back()
					}
					numbers.MoveBefore(element, prev)
				}
			}
		}
	}

	e := zero
	for i := 1; i <= 3000; i++ {
		e = e.Next()
		if e == nil {
			e = numbers.Front()
		}
		if i%1000 == 0 {
			sum += e.Value.(int)
		}
	}

	return
}

func p1(inp []int) int {
	return mix(inp, 1)
}

func p2(inp []int) int {
	for i, number := range inp {
		inp[i] = number * 811589153
	}
	return mix(inp, 10)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
