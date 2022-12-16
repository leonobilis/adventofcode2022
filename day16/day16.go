package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
	"golang.org/x/exp/slices"
)

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

type CacheItem struct {
	valve1, valve2 string
	openedValves   string
	busy1, busy2   bool
	minute         int
}

type Cache map[CacheItem]int

func (c Cache) Add(valve, valve2 string, openedValves []string, busy1, busy2 bool, minute, pressure int) {
	c[CacheItem{valve, valve2, strings.Join(openedValves, ""), busy1, busy2, minute}] = pressure
}

func (c Cache) Get(valve, valve2 string, openedValves []string, busy1, busy2 bool, minute int) (int, bool) {
	v, ok := cache[CacheItem{valve, valve2, strings.Join(openedValves, ""), busy1, busy2, minute}]
	return v, ok
}

var cache Cache

func parseInput(input string) map[string]Valve {
	output := make(map[string]Valve)
	for _, s := range strings.Split(input, "\n") {
		s2 := strings.Split(s, "; ")
		valve := s2[0][6:8]
		flowRate := utils.Atoi(s2[0][23:])
		var tunnels []string
		if len(s2[1]) == 24 {
			tunnels = []string{s2[1][22:]}
		} else {
			tunnels = strings.Split(s2[1][23:], ", ")
		}
		output[valve] = Valve{valve, flowRate, tunnels}
	}
	return output
}

func process1(inp map[string]Valve, currentValve Valve, busy bool, _openedValves []string, toOpenLen, minute int) int {
	if cached, ok := cache.Get(currentValve.name, "", _openedValves, busy, false, minute); ok {
		return cached
	}
	openedValves := append(make([]string, 0, len(_openedValves)), _openedValves...)
	pressure := 0
	releasingPressure := 0
	for _, valve := range openedValves {
		releasingPressure += inp[valve].flowRate
	}
	if len(openedValves) == toOpenLen {
		return releasingPressure * (31 - minute)
	}
	if busy {
		openedValves = append(openedValves, currentValve.name)
		sort.Strings(openedValves)
		busy = false
	} else if currentValve.flowRate > 0 && minute < 30 && !slices.Contains(openedValves, currentValve.name) {
		busy = true
	}
	pressure += releasingPressure
	if minute == 30 {
		return pressure
	}
	iter := currentValve.tunnels
	if busy {
		iter = []string{currentValve.name}
	}
	max := 0
	for _, valve := range iter {
		v := process1(inp, inp[valve], busy, openedValves, toOpenLen, minute+1)
		cache.Add(valve, "", openedValves, busy, false, minute+1, v)
		if v > max {
			max = v
		}
	}
	return pressure + max
}

func process2(inp map[string]Valve, currentValve1, currentValve2 Valve, busy1, busy2 bool, _openedValves []string, toOpenLen, minute int) int {
	if cached, ok := cache.Get(currentValve1.name, currentValve2.name, _openedValves, busy1, busy2, minute); ok {
		return cached
	}
	openedValves := append(make([]string, 0, len(_openedValves)), _openedValves...)
	pressure := 0
	releasingPressure := 0
	for _, valve := range openedValves {
		releasingPressure += inp[valve].flowRate
	}
	if len(openedValves) == toOpenLen {
		return releasingPressure * (27 - minute)
	}
	if busy1 {
		openedValves = append(openedValves, currentValve1.name)
		sort.Strings(openedValves)
		busy1 = false
	}
	if busy2 {
		openedValves = append(openedValves, currentValve2.name)
		sort.Strings(openedValves)
		busy2 = false
	}
	if currentValve1.flowRate > 0 && minute < 26 && !slices.Contains(openedValves, currentValve1.name) {
		busy1 = true
	}
	if currentValve2.name != currentValve1.name && currentValve2.flowRate > 0 && minute < 26 && !slices.Contains(openedValves, currentValve2.name) {
		busy2 = true
	}
	pressure += releasingPressure
	if minute == 26 {
		return pressure
	}
	iter1 := currentValve1.tunnels
	if busy1 {
		iter1 = []string{currentValve1.name}
	}
	iter2 := currentValve2.tunnels
	if busy2 {
		iter2 = []string{currentValve2.name}
	}
	max := 0
	for _, valve1 := range iter1 {
		for _, valve2 := range iter2 {
			var v int
			if valve1 > valve2 {
				v = process2(inp, inp[valve2], inp[valve1], busy2, busy1, openedValves, toOpenLen, minute+1)
				cache.Add(valve2, valve1, openedValves, busy2, busy1, minute+1, v)
			} else {
				v = process2(inp, inp[valve1], inp[valve2], busy1, busy2, openedValves, toOpenLen, minute+1)
				cache.Add(valve1, valve2, openedValves, busy1, busy2, minute+1, v)
			}
			if v > max {
				max = v
			}
		}
	}
	return pressure + max
}

func getValvesToOpenLen(inp map[string]Valve) (toOpenLen int) {
	for _, valve := range inp {
		if valve.flowRate > 0 {
			toOpenLen++
		}
	}
	return
}

func p1(inp map[string]Valve) int {
	cache = make(Cache)
	return process1(inp, inp["AA"], false, make([]string, 0), getValvesToOpenLen(inp), 0)
}

func p2(inp map[string]Valve) int {
	cache = make(Cache)
	return process2(inp, inp["AA"], inp["AA"], false, false, make([]string, 0), getValvesToOpenLen(inp), 0)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
