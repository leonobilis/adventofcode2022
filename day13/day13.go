package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Item interface {
	Iter() []Item
	Cmp(Item) int
	Val() int
}

type Value int

func (v Value) Val() int {
	return int(v)
}

func (v Value) Iter() []Item {
	return []Item{v}
}

func (v Value) Cmp(item Item) int {
	iVal := item.Val()
	if iVal == -1 {
		return NewListFromSlice([]Item{v}).Cmp(item)
	}
	if v.Val() < iVal {
		return 1
	} else if v.Val() > iVal {
		return -1
	}
	return 0
}

type List struct {
	items []Item
}

func (l *List) Iter() []Item {
	return l.items
}

func (l *List) Cmp(item Item) int {
	left := l.Iter()
	right := item.Iter()

	length := len(left)
	if length > len(right) {
		length = len(right)
	}

	for i := 0; i < length; i++ {
		switch left[i].Cmp(right[i]) {
		case -1:
			return -1
		case 1:
			return 1
		}
	}

	if len(left) < len(right) {
		return 1
	} else if len(left) > len(right) {
		return -1
	}
	return 0
}

func (l *List) Val() int {
	return -1
}

func (l *List) Add(item Item) {
	l.items = append(l.items, item)
}

func NewList() *List {
	return &List{make([]Item, 0)}
}

func NewListFromSlice(items []Item) *List {
	return &List{items}
}

func parseInput(input string) (output []Item) {
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}
		itemStack := []*List{NewList()}
		var intBuffer strings.Builder

		checkBuffer := func() {
			if intBuffer.Len() > 0 {
				itemStack[len(itemStack)-1].Add(Value(utils.Atoi(intBuffer.String())))
				intBuffer.Reset()
			}
		}

		for _, s2 := range s[1 : len(s)-1] {
			switch s2 {
			case '[':
				itemStack = append(itemStack, NewList())
			case ']':
				checkBuffer()
				l := len(itemStack)
				itemStack[l-2].Add(itemStack[l-1])
				itemStack = itemStack[:l-1]
			case ',':
				checkBuffer()
			default:
				intBuffer.WriteRune(s2)
			}
		}
		checkBuffer()
		output = append(output, itemStack[0])
	}
	return
}

func p1(inp []Item) (sum int) {
	for i := 0; i < len(inp); i += 2 {
		if inp[i].Cmp(inp[i+1]) == 1 {
			sum += i/2 + 1
		}
	}
	return
}

func p2(inp []Item) (sum int) {
	divider1 := NewListFromSlice([]Item{NewListFromSlice([]Item{Value(2)})})
	divider2 := NewListFromSlice([]Item{NewListFromSlice([]Item{Value(6)})})
	inp = append(inp, divider1, divider2)

	sort.Slice(inp, func(i, j int) bool {
		return inp[i].Cmp(inp[j]) == 1
	})

	divider1Iindex, divider2Iindex := 0, 0

	for i, item := range inp {
		if item.Cmp(divider1) == 0 {
			divider1Iindex = i + 1
		} else if item.Cmp(divider2) == 0 {
			divider2Iindex = i + 1
		}
	}

	return divider1Iindex * divider2Iindex
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
